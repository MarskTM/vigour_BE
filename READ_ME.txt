Create swagger docs file to build interface API in networks:

    swag init -pd -g ***header_file_env***

    - header.go : run server with localhost
    - header_dir.go : run server with railway server

Run file server.

    go run ./main -db ***param***
    
    - 1: migrate database
    - 0: run with default database

Run Docker package:
    
    docker build -t emotional_rescues .
    heroku container:push web -a recues  
    heroku container:release web -a recues

    - a: pointer to the APP.