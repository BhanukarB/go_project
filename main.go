package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
)

func main() {

	// fmt.Println("Hello, World!")

	godotenv.Load(".env")
	portStr := os.Getenv("PORT")

	if( portStr == "" ) {
		log.Fatal("PORT environment variable is not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // Allow all origins with http/https
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow these HTTP methods
		AllowedHeaders:   []string{"*"}, // Allow all headers
		ExposedHeaders:   []string{"Content-Length"}, // Expose Content-Length header to the browser
		AllowCredentials: true, // Allow cookies and credentials
		MaxAge:           300, // Cache preflight response for 300 seconds
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + portStr,
		Handler: router,
	}

	log.Printf("starting server on port %s\n", portStr)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal( err)
	}

}