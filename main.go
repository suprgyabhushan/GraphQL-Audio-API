package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "thanks123"
	DB_NAME     = "database1"
)

type Audio struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	Audio_File    string `json:"audio_file"`
	Creator_Name  string `json:"creator_name"`
	Creator_Email string `json:"creator_email"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Making connection to the postgres database :
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	defer db.Close()

	audioType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Audio",
		Description: "An audio file",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "ID of the audio file",
				Resolve: func(a graphql.ResolveParams) (interface{}, error) {
					if audio, ok := a.Source.(*Audio); ok {
						return audio.ID, nil
					}
					return nil, nil
				},
			},
			"title": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "title of the audio file",
				Resolve: func(a graphql.ResolveParams) (interface{}, error) {
					if audio, ok := a.Source.(*Audio); ok {
						return audio.Title, nil
					}
					return nil, nil
				},
			},

			"description": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "description of the audio file",
				Resolve: func(a graphql.ResolveParams) (interface{}, error) {
					if audio, ok := a.Source.(*Audio); ok {
						return audio.Description, nil
					}

					return nil, nil
				},
			},
			"category": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Category of the audio file",
				Resolve: func(a graphql.ResolveParams) (interface{}, error) {
					if audio, ok := a.Source.(*Audio); ok {
						return audio.Category, nil
					}

					return nil, nil
				},
			},
			"audio_file": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "audio file",
				Resolve: func(a graphql.ResolveParams) (interface{}, error) {
					if audio, ok := a.Source.(*Audio); ok {
						return audio.Audio_File, nil
					}

					return nil, nil
				},
			},

			"creator_name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "creator name",
				Resolve: func(a graphql.ResolveParams) (interface{}, error) {
					if audio, ok := a.Source.(*Audio); ok {
						return audio.Creator_Name, nil
					}

					return nil, nil
				},
			},
			"creator_email": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "creator email",
				Resolve: func(a graphql.ResolveParams) (interface{}, error) {
					if audio, ok := a.Source.(*Audio); ok {
						return audio.Creator_Email, nil
					}

					return nil, nil
				},
			},
		},
	})

	// All GET queries (Get audio by title and Get all audios) :
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"audio": &graphql.Field{
				Type:        audioType,
				Description: "Get an audio",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					title, _ := params.Args["title"].(string)

					audio := &Audio{}
					err = db.QueryRow("select * from audio where title = $1", title).Scan(&audio.ID, &audio.Title, &audio.Description, &audio.Category, &audio.Audio_File, &audio.Creator_Name, &audio.Creator_Email)
					checkErr(err)
					return audio, nil
				},
			},

			"audios": &graphql.Field{
				Type:        graphql.NewList(audioType),
				Description: "List of audio files",
				Resolve: func(a graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT id,title,description,category,audio_file,creator_name,creator_email FROM audio")
					checkErr(err)
					var audios []*Audio

					for rows.Next() {
						audio := &Audio{}

						err = rows.Scan(&audio.ID, &audio.Title, &audio.Description, &audio.Category, &audio.Audio_File, &audio.Creator_Name, &audio.Creator_Email)
						checkErr(err)
						audios = append(audios, audio)
					}
					return audios, nil
				},
			},
		},
	})

	// All the mutations (create audio, update audio by title and delete audio by title) :
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			// Author
			"createAudio": &graphql.Field{
				Type:        audioType,
				Description: "Create new audio",
				Args: graphql.FieldConfigArgument{

					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"audio_file": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"creator_name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"creator_email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					title, _ := params.Args["title"].(string)
					description, _ := params.Args["description"].(string)
					category, _ := params.Args["category"].(string)
					audio_file, _ := params.Args["audio_file"].(string)
					creator_name, _ := params.Args["creator_name"].(string)
					creator_email, _ := params.Args["creator_email"].(string)
					var lastInsertId int

					err = db.QueryRow("INSERT INTO audio(title,description,category,audio_file,creator_name,creator_email) VALUES($1, $2, $3, $4, $5, $6) returning id ;", title, description, category, audio_file, creator_name, creator_email).Scan(&lastInsertId)
					checkErr(err)

					newAudio := &Audio{
						ID:            lastInsertId,
						Title:         title,
						Description:   description,
						Category:      category,
						Audio_File:    audio_file,
						Creator_Name:  creator_name,
						Creator_Email: creator_email,
					}

					return newAudio, nil
				},
			},
			"updateAudio": &graphql.Field{
				Type:        audioType,
				Description: "Update an audio",
				Args: graphql.FieldConfigArgument{

					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"audio_file": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"creator_name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"creator_email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					title, _ := params.Args["title"].(string)
					description, _ := params.Args["description"].(string)
					category, _ := params.Args["category"].(string)
					audio_file, _ := params.Args["audio_file"].(string)
					creator_name, _ := params.Args["creator_name"].(string)
					creator_email, _ := params.Args["creator_email"].(string)

					stmt, err := db.Prepare("UPDATE audio SET  description = $1, category = $2, audio_file = $3, creator_name= $4,creator_email = $5 WHERE title = $6")
					checkErr(err)

					_, err2 := stmt.Exec(description, category, audio_file, creator_name, creator_email, title)
					checkErr(err2)

					newAudio := &Audio{

						Title:         title,
						Description:   description,
						Category:      category,
						Audio_File:    audio_file,
						Creator_Name:  creator_name,
						Creator_Email: creator_email,
					}

					return newAudio, nil
				},
			},
			"deleteAudio": &graphql.Field{
				Type:        audioType,
				Description: "Delete an audio",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					title, _ := params.Args["title"].(string)

					stmt, err := db.Prepare("DELETE FROM audio WHERE title = $1")
					checkErr(err)

					_, err2 := stmt.Exec(title)
					checkErr(err2)

					return nil, nil
				},
			},
		},
	})
	// Schema definition :
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	// Hander
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve HTTP on port 8080 :
	http.Handle("/graphql", h)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
