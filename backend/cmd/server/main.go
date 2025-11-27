package main

import (
	"chat-app-backend/internal/api"
	"chat-app-backend/internal/auth"
	"chat-app-backend/internal/websocket"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Inicializar store e hub
	store := auth.NewStore()
	hub := websocket.NewHub(store)
	go hub.Run()

	// Inicializar handler
	handler := api.NewHandler(store, hub)

	// Configurar rotas
	router := mux.NewRouter()

	// Rotas de autenticação
	router.HandleFunc("/api/admin/login", handler.AdminLogin).Methods("POST")
	router.HandleFunc("/api/guest/login", handler.GuestLogin).Methods("POST")

	// Rotas de sessão
	router.HandleFunc("/api/session/create", handler.CreateSession).Methods("POST")
	router.HandleFunc("/api/sessions", handler.GetAdminSessions).Methods("GET")

	// Rotas de mensagens
	router.HandleFunc("/api/messages", handler.GetMessages).Methods("GET")

	// Rota WebSocket
	router.HandleFunc("/ws", handler.WebSocketHandler)

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	handler2 := c.Handler(router)

	// Iniciar servidor
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Printf("Default admin credentials - username: admin, password: admin123")
	if err := http.ListenAndServe(port, handler2); err != nil {
		log.Fatal(err)
	}
}

