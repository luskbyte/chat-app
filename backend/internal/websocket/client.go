package websocket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Tempo permitido para escrever uma mensagem ao peer
	writeWait = 10 * time.Second

	// Tempo permitido para ler a próxima mensagem pong do peer
	pongWait = 60 * time.Second

	// Enviar pings ao peer com este período. Deve ser menor que pongWait
	pingPeriod = (pongWait * 9) / 10

	// Tamanho máximo de mensagem permitido do peer
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client representa um cliente WebSocket
type Client struct {
	hub       *Hub
	conn      *websocket.Conn
	send      chan []byte
	sessionID string
	userType  string // "admin" or "guest"
}

// NewClient cria um novo cliente
func NewClient(hub *Hub, conn *websocket.Conn, sessionID, userType string) *Client {
	return &Client{
		hub:       hub,
		conn:      conn,
		send:      make(chan []byte, 256),
		sessionID: sessionID,
		userType:  userType,
	}
}

// MessageData representa os dados de uma mensagem
type MessageData struct {
	Sender    string `json:"sender"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

// ReadPump bombeia mensagens da conexão WebSocket para o hub
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		
		// Parse da mensagem JSON
		var msgData MessageData
		if err := json.Unmarshal(message, &msgData); err != nil {
			log.Printf("error parsing message: %v", err)
			continue
		}

		// Salvar mensagem no Store
		savedMsg, err := c.hub.store.AddMessage(c.sessionID, msgData.Sender, msgData.Content)
		if err != nil {
			log.Printf("error saving message: %v", err)
			continue
		}

		// Preparar mensagem para broadcast
		broadcastData, err := json.Marshal(savedMsg)
		if err != nil {
			log.Printf("error marshaling message: %v", err)
			continue
		}

		// Broadcast para todos os clientes da sessão
		c.hub.BroadcastToSession(c.sessionID, broadcastData)
	}
}

// WritePump bombeia mensagens do hub para a conexão WebSocket
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// O hub fechou o canal
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Adiciona mensagens enfileiradas ao atual WebSocket message
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

