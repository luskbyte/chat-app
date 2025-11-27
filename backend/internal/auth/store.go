package auth

import (
	"chat-app-backend/internal/models"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Store gerencia o armazenamento em memória (para produção, use banco de dados)
type Store struct {
	admins   map[string]*models.Admin
	sessions map[string]*models.Session
	messages map[string][]*models.Message
	mu       sync.RWMutex
}

// NewStore cria uma nova instância do Store
func NewStore() *Store {
	s := &Store{
		admins:   make(map[string]*models.Admin),
		sessions: make(map[string]*models.Session),
		messages: make(map[string][]*models.Message),
	}

	// Criar admin padrão (para testes)
	hashedPassword, _ := HashPassword("admin123")
	s.admins["admin"] = &models.Admin{
		ID:        uuid.New().String(),
		Username:  "admin",
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	return s
}

// GetAdminByUsername busca um admin pelo username
func (s *Store) GetAdminByUsername(username string) (*models.Admin, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	admin, exists := s.admins[username]
	if !exists {
		return nil, errors.New("admin not found")
	}
	return admin, nil
}

// CreateSession cria uma nova sessão
func (s *Store) CreateSession(adminID string) (*models.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	code, err := GenerateSessionCode()
	if err != nil {
		return nil, err
	}

	session := &models.Session{
		ID:        uuid.New().String(),
		AdminID:   adminID,
		Code:      code,
		IsActive:  true,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	s.sessions[session.Code] = session
	return session, nil
}

// GetSessionByCode busca uma sessão pelo código
func (s *Store) GetSessionByCode(code string) (*models.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.sessions[code]
	if !exists {
		return nil, errors.New("session not found")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, errors.New("session expired")
	}

	return session, nil
}

// UpdateSessionGuest atualiza o nome do visitante na sessão
func (s *Store) UpdateSessionGuest(code, guestName string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	session, exists := s.sessions[code]
	if !exists {
		return errors.New("session not found")
	}

	session.GuestName = guestName
	return nil
}

// AddMessage adiciona uma mensagem ao chat
func (s *Store) AddMessage(sessionID, sender, content string) (*models.Message, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	message := &models.Message{
		ID:        uuid.New().String(),
		SessionID: sessionID,
		Sender:    sender,
		Content:   content,
		Timestamp: time.Now(),
	}

	s.messages[sessionID] = append(s.messages[sessionID], message)
	return message, nil
}

// GetMessages retorna todas as mensagens de uma sessão
func (s *Store) GetMessages(sessionID string) []*models.Message {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.messages[sessionID]
}

// GetSessionByID busca uma sessão pelo ID
func (s *Store) GetSessionByID(sessionID string) (*models.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, session := range s.sessions {
		if session.ID == sessionID {
			return session, nil
		}
	}

	return nil, errors.New("session not found")
}

// GetSessionsByAdminID retorna todas as sessões de um admin
func (s *Store) GetSessionsByAdminID(adminID string) []*models.Session {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var sessions []*models.Session
	for _, session := range s.sessions {
		if session.AdminID == adminID {
			sessions = append(sessions, session)
		}
	}

	return sessions
}

