package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("starting the shit")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error occured while loading env", err)
	}

	port := os.Getenv("PORT")

	r := chi.NewRouter()

	r.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"http://*", "https://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"*"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

	server := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	v1router := chi.NewRouter()

	v1router.Get("/shit", giveShit)

	r.Mount("/v1", v1router)

	log.Println("server running of the port :", port)
	err = server.ListenAndServe()
	if err != nil{
		log.Fatalln("error starting server", err)
	}
}

type Shit struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Length int32 `json:"length"`
	Amount int32 `json:"amount"`
	Unit string `json:"unit"`
}

func giveShit(w http.ResponseWriter, r *http.Request){
	data := Shit{
		Title: "the hard shit from maida",
		Description: "due to eating stuff made of maida the shit is hard as hell.",
		Length: 2,
		Amount: 50,
		Unit: "cubic-cm",
	}
	
	jsonData , err := json.Marshal(data)
	if err != nil{
		log.Fatalln("error parsing json", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)
}