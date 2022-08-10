package main

import "github.com/itsapep/golang-api-with-gin/delivery"

func main() {
	// DB_HOST=localhost DB_PORT=5432 DB_NAME=db_api_gin DB_USER=postgres DB_PASSWORD=12345678 API_URL=localhost:8888 ENV=migration FILE_PATH=/home/yurham/Documents/Golang/golang-intermediate/golang-api-with-gin/img go run .
	delivery.Server().Run()
}
