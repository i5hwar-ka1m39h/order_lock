package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)



func main(){
	fmt.Println("starting the email worker...")


	err :=godotenv.Load()
	if err != nil {
		log.Fatalln("error loading the envs", err)
		return
	}

	service_port := os.Getenv("SERVICE_PORT")
	if service_port == ""{
		log.Fatalln("service port is empty")
		return
	}

	r := chi.NewRouter()
	
	server := &http.Server{
		Handler: r,
		Addr: ":" + service_port,
	}


	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("error while creating service worker", err)
		return
	}

	fmt.Println("service worker is listening on port ", service_port, "...")
}




