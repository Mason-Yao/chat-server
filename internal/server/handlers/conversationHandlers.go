package handlers

import (
	http "net/http"
	mux "github.com/gorilla/mux"
)

func RegisterConversationRoutes(router *mux.Router) *mux.Router {
	conversationsRouter := router.PathPrefix("/conversations").Subrouter()
	conversationsRouter.HandleFunc("/", getConversations);
	conversationsRouter.HandleFunc("/", addConversation).Methods("POST")
	conversationsRouter.HandleFunc("/{id:[0-9]+}", getConversation).Methods("GET")
	conversationsRouter.HandleFunc("/{id:[0-9]+}", deleteConversation).Methods("DELETE")
	conversationsRouter.HandleFunc("/{id:[0-9]+}/messages", getMessages).Methods("GET")
	conversationsRouter.HandleFunc("/{id:[0-9]+}/messages", addMessage).Methods("POST")
	return router
}

func getConversations(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Return all conversations"))
}

func addConversation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add a conversation"))
}

func getConversation(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Write([]byte("Return conversation by id: " + id))
}

func deleteConversation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete conversation by id"))
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Return all messages by conversation id"))
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add a message to conversation"))
}
