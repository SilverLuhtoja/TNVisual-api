#!/bin/bash

cd ./sql/schema;
goose postgres postgres://postgres:postgres@localhost:15432/postgres up