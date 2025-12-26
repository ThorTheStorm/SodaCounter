package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"sodaCounter_rest/handlers"

	"sodaCounter_rest/pkg/db"
	"sodaCounter_rest/pkg/framework/response"
)

// TODO!: Rebuild this to be a middleware chain function that takes in multiple middlewares and applies them in order
func middleware(next http.Handler) http.Handler {

	// Chain multiple middlewares here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mws := []func(http.Handler) http.Handler{
			loggingMiddleware,
		}
		for i := len(mws) - 1; i >= 0; i-- {
			next = mws[i](next)
		}
		next.ServeHTTP(w, r)
	})
}

// Loggin middleware that logs request method, path and time taken to process the request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[LOG]", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

var database *sql.DB

// func init() {
// }

const tableName = "sodas"

func main() {
	// 1. Connect to the database here and setup any required tables
	database, err := db.DBConnect(database, "devuser", "devpasss", "localhost", "5432", "thedb")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.DBClose(database) // Close the database connection when main() exits

	// 2. Check if the required table exists, if not create it
	exists, err := db.TableExists(database, tableName)
	if err != nil {
		log.Fatalf("Failed to check if table exists: %v", err)
	}
	if !exists {
		err = db.TableInit(database, tableName)
		if err != nil {
			log.Fatalf("Failed to initialize table: %v", err)
		}
	}

	// 3. Setup HTTP server and routes
	api := http.NewServeMux()                     // Create a new ServeMux (router-object) for routing requests to handlers
	h := handlers.NewHandler(database, tableName) // Initialize a new handler instance

	router(api, h) // Setup routes - Contains logic to route requests to appropriate handler methods

	server := &http.Server{
		Addr:    ":8080",
		Handler: api,
	}

	// Log server start in terminal
	fmt.Println("Server started on port 8080")

	serverErr := make(chan error, 1)

	// Start the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			serverErr <- fmt.Errorf("the server failed for some reason: %s", err.Error())
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Wait for a termination signal or server error
	select {
	case err := <-serverErr:
		log.Fatalf("application crashed: %s", err.Error())
	case err := <-quit:
		log.Println("application dies unexpectedly, or was terminated", err)
	}

	// Create context for shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// If shutdown fails, log a fatal error
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("a bad day to be a REST server.")
	}
}

// router sets up the routes for the HTTP server
func router(api *http.ServeMux, h *handlers.Handler) {
	// Define how to handle everything on /soda
	api.Handle("/soda", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAllSoda(w, r)
		case http.MethodPost:
			h.CreateSoda(w, r, tableName)
		default:
			http.Error(w, "/soda: method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	// Define how to handle everything on /soda/ like /soda/{id}
	api.Handle("/soda/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Logic for extracting ID from URL path
		idStr := strings.TrimPrefix(r.URL.Path, "/soda/")
		// fmt.Println("idStr:", idStr)
		if idStr == "" {
			response.Err(w, response.ErrInvalidInput)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Err(w, response.ErrInvalidInput)
			return
		}

		// Route to appropriate handler method based on HTTP method
		switch r.Method {
		case http.MethodGet:
			h.GetSoda(w, r, tableName, id) // TODO: Implement middleware to extract and validate ID from URL path
		case http.MethodPut:
			h.UpdateSoda(w, r, tableName, id) // TODO: Implement middleware to extract and validate ID from URL path
		case http.MethodDelete:
			h.DeleteSoda(w, r, tableName, id) // TODO: Implement middleware to extract and validate ID from URL path
		default:
			http.Error(w, "/soda/{id}: method not allowed", http.StatusMethodNotAllowed)
		}
	}))
}
