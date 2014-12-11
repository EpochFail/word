FROM golang:latest

RUN go get 'bitbucket.org/liamstask/goose/cmd/goose'

RUN mkdir -p /opt/db
ADD ./run-migrations /opt/db/run-migrations
ADD ./dbconf.yml /opt/db/dbconf.yml
ADD ./migrations /opt/db/migrations
ADD ./seeds /opt/db/seeds

WORKDIR /opt/db

CMD ["./run-migrations"]
