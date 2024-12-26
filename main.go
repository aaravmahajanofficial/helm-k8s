package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func getMessage(message string) string {
	return fmt.Sprintf("Hello, %v environment here", message)
}

func main() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	server := http.Server{
		Addr:    ":80",
		Handler: router,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, getMessage(os.Getenv("MESSAGE")))
	})

	err := server.ListenAndServe()

	if err != nil {

		fmt.Println("failed to start the server: ", err)

	}

}
