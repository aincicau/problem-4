package main

import (
	"log"
	"net/http"
	"pb4/config"
	"pb4/db"
	"pb4/rest"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Not able to load environment file")

	}

	err = config.InitConfig()

	if err != nil {
		log.Fatal("Not able to create config")

	}

	db.InitDatabase()

	router := chi.NewRouter()
	router.Route("/"+config.GetConfig().APPVersion, func(r chi.Router) {
		r.Get("/student", rest.GetStudent)
		r.Post("/student", rest.PostStudent)
		r.Delete("/student", rest.DeleteStudent)
		r.Put("/student", rest.PutStudent)
		r.Get("/students", rest.ListOfStudents)
		r.Get("/class", rest.GetClass)
		r.Post("/class", rest.PostClass)
		r.Delete("/class", rest.DeleteClass)
		r.Put("/class", rest.PutClass)
		r.Get("/classes", rest.ListOfClasses)
		r.Post("/enroll", rest.Enroll)
	})
	http.ListenAndServe(":"+config.GetConfig().Port, router)
}
