## DATABASE
to kill in use port : fuser -k 8080/tcp

CONNECTION STRING - protocol://username:password@host:port/database

sqlc -  SQLC is an amazing Go program that generates Go code from SQL queries. It's not exactly an ORM, but rather a tool that makes working with raw SQL almost as easy as using an ORM. 
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
goose - Goose is a database migration tool written in Go. It runs migrations from the same SQL files that SQLC uses, making the pair of tools a perfect fit.
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

sqlc init - creates sqlc configuration file

POSTGRES:
- start postgres server: sudo service postgresql start
- check postgres server status: sudo service postgresql status
- stop postgres server: sudo service postgresql stop

MIGRATION:
1. move to sql/schema directory
2. Migrate up : goose postgres postgres://user:pass@localhost:5432/blog_db  up  
Migrate down : goose postgres postgres://user:pass@localhost:5432/blog_db   down  
OR 
1. run bash migrate.sh to migrate

GENERATE SQL to GO (from root): sqlc generate

LINUX POSTGRES 
sudo apt update
sudo apt install postgresql postgresql-contrib
