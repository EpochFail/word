# WORD

Web service for doin' things with words. Yup.

## Use it

Create a new database user, and give it the password of ````test```` when prompted

    createuser -P jsname

Create a new database
    
    createdb -O jsname jsname_dev

Install goose to run the database migrations
    
    go get bitbucket.org/liamstask/goose/cmd/goose

Run the migrations. The migrations will also populate the initial word list, which can take some time, so be patient
    
    goose up

Run it locally

    go build
    PORT=5555 ./word

Try it from your console with
    
    curl http://localhost:5555/api/word

Result looks like 
    
    {"Word":"glittery","Rating":0}
