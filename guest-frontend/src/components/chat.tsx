import { useState, useEffect, useRef } from "react";
import { apiService } from "../services/api";
import "./chat.css";

type ChatProps = {
  sessionId: string;
  isHost: boolean;
  guestName?: string;
};

interface ChatMessage {
  id: string;
  sender: string;
  content: string;
  timestamp: string;
}

function Chat({ sessionId, isHost, guestName }: ChatProps) {
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const [inputMessage, setInputMessage] = useState("");
  const [isConnected, setIsConnected] = useState(false);
  const wsRef = useRef<WebSocket | null>(null);
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  useEffect(() => {
    // Limpar mensagens ao mudar de sessão
    setMessages([]);
    setIsConnected(false);

    // Carregar mensagens anteriores ANTES de conectar ao WebSocket
    apiService.getMessages(sessionId).then((msgs) => {
      if (msgs && msgs.length > 0) {
        setMessages(msgs.map(msg => ({
          id: msg.id,
          sender: msg.sender,
          content: msg.content,
          timestamp: msg.timestamp,
        })));
      }
    }).catch((err) => {
      console.error("Error loading messages:", err);
    });

    // Conectar ao WebSocket
    try {
      const ws = apiService.connectWebSocket(sessionId, (data) => {
        if (data.content) {
          const newMessage: ChatMessage = {
            id: data.id || Date.now().toString(),
            sender: data.sender,
            content: data.content,
            timestamp: data.timestamp || new Date().toISOString(),
          };
          setMessages((prev) => [...prev, newMessage]);
        }
      });

      ws.onopen = () => {
        setIsConnected(true);
        console.log("WebSocket connected for session:", sessionId);
      };

      ws.onclose = () => {
        setIsConnected(false);
        console.log("WebSocket disconnected");
      };

      wsRef.current = ws;

      return () => {
        if (ws.readyState === WebSocket.OPEN) {
          ws.close();
        }
      };
    } catch (error) {
      console.error("Error connecting to WebSocket:", error);
    }
  }, [sessionId]);

  const handleSendMessage = (e: React.FormEvent) => {
    e.preventDefault();
    if (!inputMessage.trim() || !wsRef.current || wsRef.current.readyState !== WebSocket.OPEN) {
      return;
    }

    const message = {
      sender: isHost ? "host" : "guest",
      content: inputMessage.trim(),
      timestamp: new Date().toISOString(),
    };

    wsRef.current.send(JSON.stringify(message));
    setInputMessage("");
  };

  return (
    <div className="chat-wrapper">
      <div className="chat-header">
        <div className="chat-info">
          <h2>{isHost ? "Host Chat Room" : `Chat com Host`}</h2>
          <div className="connection-status">
            <span
              className={`status-indicator ${isConnected ? "connected" : "disconnected"}`}
            ></span>
            <span className="status-text">
              {isConnected ? "Conectado" : "Desconectado"}
            </span>
          </div>
        </div>
      </div>

      <div className="messages-container">
        {messages.length === 0 ? (
          <div className="no-messages">
            <p>Nenhuma mensagem ainda. Comece a conversar!</p>
          </div>
        ) : (
          messages.map((msg) => (
            <div
              key={msg.id}
              className={`message ${msg.sender === (isHost ? "host" : "guest") ? "own-message" : "other-message"}`}
            >
              <div className="message-header">
                <span className="message-sender">
                  {msg.sender === "host" ? "Host" : guestName || "Visitante"}
                </span>
                <span className="message-time">
                  {new Date(msg.timestamp).toLocaleTimeString()}
                </span>
              </div>
              <div className="message-content">{msg.content}</div>
            </div>
          ))
        )}
        <div ref={messagesEndRef} />
      </div>

      <form className="message-input-form" onSubmit={handleSendMessage}>
        <input
          type="text"
          value={inputMessage}
          onChange={(e) => setInputMessage(e.target.value)}
          placeholder="Digite sua mensagem..."
          className="message-input"
          disabled={!isConnected}
        />
        <button 
          type="submit" 
          className="send-button"
          disabled={!isConnected || !inputMessage.trim()}
        >
          Enviar
        </button>
      </form>
    </div>
  );
}

export default Chat;
