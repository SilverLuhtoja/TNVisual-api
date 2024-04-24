#!/bin/bash


# read -t 5 -p "Migrate UP or DOWN: " option
cd ./sql/schema;
goose postgres postgres://test:test@localhost:25432/test up