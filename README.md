# Description:
GraphQL API for an audio shorts directory.

# System:
Linux Mint Kernel

# Requirements:
1. GraphQL
2. Golang
3. PostgreSQL

# Setup:
$ go run main.go

# Usage:
1. Please type this command in command terminal: $ go run main.go

2. Server will run at: localhost:8080

3. Please do run this in postgres CLI after selecting database and table 'audio': INSERT INTO audio(id, title, description, category, audio_file, creator_name, creator_email) VALUES (2, 'Aurora', 'Song from a movie', 'Soundtrack', 'aurora_mp3', 'Hans Zimmer', 'hans@gmail.com');

4. Please do add more data records by these commands: INSERT INTO audio(2, 'Posterity', 'Song from a movie', 'Soundtrack', 'posterity_mp3', 'Ludwig', 'ludwig@gmail.com');

5. Redis server will will be running on: 127.0.0.1:6379

6. Please open http://localhost:8080/graphql?query={audio(title:"Aurora"){id,title,description,category,audio_file,creator_name,creator_email}} to get audio by title.

7. Please open http://localhost:8080/graphql?query={audios{id,title,description,category,audio_file,creator_name,creator_email}} to get all audios.

8. Please open http://localhost:8080/graphql?query=mutation {createAudio(title:"Time", description:"Song from a movie", category:"Soundtrack", audio_file:"time_mp3", creator_name:"Zimmer", creator_email:"zimmer@gmail.com"){title,description,category,audio_file,creator_name,creator_email}} to post or create audio.
  
9. Please open http://localhost:8080/graphql?query=mutation {updateAudio(title:"Posterity", description:"Song from a movie",category:"Soundtrack",audio_file:"posterity_mp3",creator_name:"John",creator_email:"john@gmail.com"){title,description,category,audio_file,creator_name,creator_email}} to update audio by title.

10. Please open http://localhost:8080/graphql?query=mutation {deleteAudio(title:"Aurora"){id,title,description,category,audio_file,creator_name,creator_email}} to delete audio by title.
