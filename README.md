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
