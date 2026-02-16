package emailworker

import (
	"fmt"
	"log"
	"net/http"
	"os"
	

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/i5hwar-ka1m39h/order_lock/api"
	"github.com/joho/godotenv"
)

func emailworker() {
	fmt.Println("starting the email worker...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading the envs", err)
		return
	}

	service_port := os.Getenv("SERVICE_PORT")
	if service_port == "" {
		log.Fatalln("service port is empty")
		return
	}

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
		Addr:    ":" + service_port,
	}

	v1router := chi.NewRouter()
	v1router.Get("/health", handleServiceHealth)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("error while creating service worker", err)
		return
	}

	fmt.Println("service worker is listening on port ", service_port, "...")
}


func handleServiceHealth(w http.ResponseWriter, r *http.Request){
	main.JsonResponseWriter(w, 200, "service worker is healthy",)
}