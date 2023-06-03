package main

import (
	"fmt"
	"log"

	"github.com/robertobouses/api_ejercicio1/http"
	"github.com/robertobouses/api_ejercicio1/internal/postgres"

	"github.com/robertobouses/api_ejercicio1/repository"
	"github.com/robertobouses/api_ejercicio1/user"
)

func main() {

	fmt.Println("hola api")

	db, err := postgres.NewPostgres(postgres.DBConfig{
		/*
			User:     os.Getenv("DB_USER"),
			Pass:     os.Getenv("DB_PASS"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
		*/
		User:     "username",
		Pass:     "password",
		Host:     "localhost",
		Port:     "5432",
		Database: "dbname",
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo, err := repository.NewRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	userService, err := user.NewUserAppService(repo)
	if err != nil {
		log.Fatal(err)
	}

	s := http.NewServer(usrService)
	s.Run("8080")
}
