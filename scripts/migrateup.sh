#!/bin/bash

# read -t 5 -p "Migrate UP or DOWN: " option
cd ./sql/schema;
goose postgres postgres://postgres:postgres@localhost:15432/postgres up