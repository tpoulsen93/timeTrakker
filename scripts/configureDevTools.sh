#!/bin/bash

# start postgresql
sudo service postgresql start

# add a psql user and create a db
sudo runuser -l postgres -c "createuser -lds $(whoami)"
createdb $(whoami)

# create the tables
psql -f db/schema.sql