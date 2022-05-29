package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//Handler => the struct that stores pointer to comments service
type Handler struct {
	Router *mux.Router
}

// NewHandler => returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

//SetupRoutes => sets up all the routes for the applications
func (h *Handler) SetupRoutes() {
	fmt.Println("Settings up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")

	})
}
