package server

import (
	handlers "chat/internal/server/handlers"
	mux "github.com/gorilla/mux"
)

func NewRouter(v string) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api/" + v).Subrouter()
	handlers.RegisterConversationRoutes(router)
	return router
}
