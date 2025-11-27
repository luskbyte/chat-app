package websocket

import (
	"chat-app-backend/internal/auth"
	"sync"
)

// Hub mantém o conjunto de clientes ativos e transmite mensagens
type Hub struct {
	// Clientes registrados por sessão
	clients map[string]map[*Client]bool

	// Mensagens de broadcast dos clientes
	broadcast chan *BroadcastMessage

	// Registrar requisições dos clientes
	register chan *Client

	// Desregistrar requisições dos clientes
	unregister chan *Client

	// Store para salvar mensagens
	store *auth.Store

	mu sync.RWMutex
}

// BroadcastMessage representa uma mensagem para broadcast
type BroadcastMessage struct {
	SessionID string
	Message   []byte
}

// NewHub cria uma nova instância do Hub
func NewHub(store *auth.Store) *Hub {
	return &Hub{
		broadcast:  make(chan *BroadcastMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
		store:      store,
	}
}

// Run inicia o hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if _, exists := h.clients[client.sessionID]; !exists {
				h.clients[client.sessionID] = make(map[*Client]bool)
			}
			h.clients[client.sessionID][client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if clients, exists := h.clients[client.sessionID]; exists {
				if _, ok := clients[client]; ok {
					delete(clients, client)
					close(client.send)
					if len(clients) == 0 {
						delete(h.clients, client.sessionID)
					}
				}
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.RLock()
			if clients, exists := h.clients[message.SessionID]; exists {
				for client := range clients {
					select {
					case client.send <- message.Message:
					default:
						close(client.send)
						delete(clients, client)
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

// BroadcastToSession envia uma mensagem para todos os clientes de uma sessão
func (h *Hub) BroadcastToSession(sessionID string, message []byte) {
	h.broadcast <- &BroadcastMessage{
		SessionID: sessionID,
		Message:   message,
	}
}

