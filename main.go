package main

import (
	"context"
	"fmt"
	"simple-go/cmd/server"
	"simple-go/config"
	"simple-go/database"
)

func main() {

	fmt.Println("INIT Config")
	config.Config()

	fmt.Println("INIT DB")
	store := database.NewStore()
	database.NewSeed(store.SqlDB.PostgreSQL).Setup(context.Background())

	fmt.Println("INIT REST-API")
	server.Init()
}
