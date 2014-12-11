# WORD

Web service for doin' things with words. Yup.

## Use it

Create a new database user, and give it the password of `krampus` when prompted

    createuser -P word

Create a new database
    
    createdb -O word word

Install goose to run the database migrations
    
    go get bitbucket.org/liamstask/goose/cmd/goose

Run it (this will also populate the initial word list, which can take some time, so be patient)
    
    ./lazy-run

Try it from your console with
    
    curl http://localhost:8001/api/word

Result looks like 
    
    {"id":0,"word":"glittery","rating":0}

## Docker stuff 
    cd db && docker build -t word-db . && cd -
    cd migrations && docker build -t word-migrations . && cd -
    cd cmd/word && docker build -t word-app . && cd -

    docker run -d --name word-db word-db
    docker run --rm --link word-db:pg word-migrations
    docker run -d -p 80:80 -v /var/run/docker.sock:/tmp/docker.sock jwilder/nginx-proxy
    docker run -d --link word-db:pg --name word-app -e VIRTUAL_HOST=word.ralreegorganon.com word-app
