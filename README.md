# HTTP Servers in Go


## http client
curl -sS https://webi.sh/curlie | sh; \
source ~/.config/envman/PATH.env

## DB setup
go install github.com/pressly/goose/v3/cmd/goose@latest
- create migration in sql/schema folder
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
- create sqlc.yaml file
- create queries in sql/queries folder
go get github.com/google/uuid
go get github.com/lib/pq
- use it for its side effects *(import _ ...)*. It is the postgresql driver 
go get github.com/joho/godotenv
- create the .env and put things such as
  - DB_URL="YOUR_CONNECTION_STRING_HERE?sslmode=disable"

## authentication
go get golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5
