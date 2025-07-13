package main

import (
	// "fmt"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/BhanukarB/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // _ means include this code in my program even though I am not calling it directly
)

type apiConfig struct{
	DB *database.Queries;
}

func main() {

	// fmt.Println("Hello, World!")

	godotenv.Load(".env")
	portStr := os.Getenv("PORT")

	if( portStr == "" ) {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL :=os.Getenv("DB_URL")
	if( dbURL == "" ) {
		log.Fatal("DB_URL environment variable is not set")
	}	
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}


	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + portStr,
		Handler: router,
	}

	log.Printf("starting server on port %s\n", portStr)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal( err)
	}

}