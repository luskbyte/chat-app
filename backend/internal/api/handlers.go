package api

import (
	"chat-app-backend/internal/auth"
	"chat-app-backend/internal/models"
	"chat-app-backend/internal/websocket"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	ws "github.com/gorilla/websocket"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Em produção, verificar origem adequadamente
	},
}

// Handler gerencia as requisições HTTP
type Handler struct {
	store *auth.Store
	hub   *websocket.Hub
}

// NewHandler cria uma nova instância do Handler
func NewHandler(store *auth.Store, hub *websocket.Hub) *Handler {
	return &Handler{
		store: store,
		hub:   hub,
	}
}

// AdminLogin trata o login do administrador
func (h *Handler) AdminLogin(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	admin, err := h.store.GetAdminByUsername(req.Username)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if !auth.CheckPasswordHash(req.Password, admin.Password) {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := auth.GenerateToken(admin.ID, "admin")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	respondWithJSON(w, http.StatusOK, models.LoginResponse{
		Token:   token,
		AdminID: admin.ID,
		Message: "Login successful",
	})
}

// CreateSession cria uma nova sessão de chat
func (h *Handler) CreateSession(w http.ResponseWriter, r *http.Request) {
	claims, err := h.extractClaims(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if claims.UserType != "admin" {
		respondWithError(w, http.StatusForbidden, "Only admins can create sessions")
		return
	}

	session, err := h.store.CreateSession(claims.UserID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating session")
		return
	}

	respondWithJSON(w, http.StatusCreated, models.CreateSessionResponse{
		Session: session,
		Message: "Session created successfully",
	})
}

// GuestLogin trata o login do visitante
func (h *Handler) GuestLogin(w http.ResponseWriter, r *http.Request) {
	var req models.GuestLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	session, err := h.store.GetSessionByCode(req.Code)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Invalid session code")
		return
	}

	if err := h.store.UpdateSessionGuest(req.Code, req.GuestName); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating session")
		return
	}

	token, err := auth.GenerateToken(session.ID, "guest")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	respondWithJSON(w, http.StatusOK, models.GuestLoginResponse{
		Token:     token,
		SessionID: session.ID,
		Message:   "Login successful",
	})
}

// WebSocketHandler trata conexões WebSocket
func (h *Handler) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := h.extractClaimsFromQuery(r)
	if err != nil {
		log.Printf("WebSocket auth error: %v", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var sessionID string
	if claims.UserType == "admin" {
		// Para admin, pegar o sessionID da query
		sessionID = r.URL.Query().Get("sessionID")
		if sessionID == "" {
			http.Error(w, "Session ID required", http.StatusBadRequest)
			return
		}
	} else {
		// Para guest, o sessionID é o próprio UserID
		sessionID = claims.UserID
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := websocket.NewClient(h.hub, conn, sessionID, claims.UserType)
	client.Register()

	go client.WritePump()
	go client.ReadPump()
}

// GetMessages retorna todas as mensagens de uma sessão
func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	claims, err := h.extractClaims(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	sessionID := r.URL.Query().Get("sessionID")
	if sessionID == "" {
		respondWithError(w, http.StatusBadRequest, "Session ID required")
		return
	}

	// Verificar se o usuário tem acesso a esta sessão
	session, err := h.store.GetSessionByID(sessionID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Session not found")
		return
	}

	if claims.UserType == "admin" && session.AdminID != claims.UserID {
		respondWithError(w, http.StatusForbidden, "Access denied")
		return
	} else if claims.UserType == "guest" && sessionID != claims.UserID {
		respondWithError(w, http.StatusForbidden, "Access denied")
		return
	}

	messages := h.store.GetMessages(sessionID)
	respondWithJSON(w, http.StatusOK, messages)
}

// GetAdminSessions retorna todas as sessões de um admin
func (h *Handler) GetAdminSessions(w http.ResponseWriter, r *http.Request) {
	claims, err := h.extractClaims(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if claims.UserType != "admin" {
		respondWithError(w, http.StatusForbidden, "Only admins can access this endpoint")
		return
	}

	sessions := h.store.GetSessionsByAdminID(claims.UserID)
	respondWithJSON(w, http.StatusOK, sessions)
}

// extractClaims extrai os claims do JWT do header Authorization
func (h *Handler) extractClaims(r *http.Request) (*auth.Claims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, http.ErrNoCookie
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	return auth.ValidateToken(tokenString)
}

// extractClaimsFromQuery extrai os claims do JWT da query string
func (h *Handler) extractClaimsFromQuery(r *http.Request) (*auth.Claims, error) {
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		return nil, http.ErrNoCookie
	}

	return auth.ValidateToken(tokenString)
}

// respondWithError responde com uma mensagem de erro
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.ErrorResponse{Error: message})
}

// respondWithJSON responde com JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

