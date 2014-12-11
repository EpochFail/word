#!/bin/bash

export PGUSER="${PGUSER:-word}"
export PGPASSWORD="${PGPASSWORD:-krampus}"
export PGDB="${PGDB:-word}"

gosu postgres postgres --single <<- EOSQL
create database $PGDB;
create user $PGUSER password '$PGPASSWORD';
grant all privileges on database $PGDB to $PGUSER;
EOSQL
