# WORD

Web service for getting a word. Yup.

## Use it

There's a copy hosted on Heroku at http://tranquil-lowlands-2993.herokuapp.com/api/word

Try it from your console with
    
    curl http://tranquil-lowlands-2993.herokuapp.com/api/word

Result looks like 
    
    {"word":"spellbinding"}

Running it locally

    go get
    PORT=5555 word
    curl http://localhost:5555/api/word

Running it on Heroku

    heroku create -b https://github.com/kr/heroku-buildpack-go.git
    git push heroku master
