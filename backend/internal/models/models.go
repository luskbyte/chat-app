package models

import "time"

// Admin representa um usuário administrador (host)
type Admin struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // Never send password to client
	CreatedAt time.Time `json:"created_at"`
}

// Session representa uma sessão de chat
type Session struct {
	ID         string    `json:"id"`
	AdminID    string    `json:"admin_id"`
	Code       string    `json:"code"`
	GuestName  string    `json:"guest_name,omitempty"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}

// Message representa uma mensagem do chat
type Message struct {
	ID        string    `json:"id"`
	SessionID string    `json:"session_id"`
	Sender    string    `json:"sender"` // "host" or "guest"
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// LoginRequest representa uma requisição de login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse representa a resposta de login
type LoginResponse struct {
	Token   string `json:"token"`
	AdminID string `json:"admin_id"`
	Message string `json:"message"`
}

// GuestLoginRequest representa uma requisição de login de visitante
type GuestLoginRequest struct {
	Code      string `json:"code"`
	GuestName string `json:"guest_name"`
}

// GuestLoginResponse representa a resposta de login de visitante
type GuestLoginResponse struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Message   string `json:"message"`
}

// CreateSessionRequest representa uma requisição para criar sessão
type CreateSessionRequest struct {
	AdminID string `json:"admin_id"`
}

// CreateSessionResponse representa a resposta de criação de sessão
type CreateSessionResponse struct {
	Session *Session `json:"session"`
	Message string   `json:"message"`
}

// ErrorResponse representa uma resposta de erro
type ErrorResponse struct {
	Error string `json:"error"`
}

