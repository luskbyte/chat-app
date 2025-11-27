import { useState } from "react";
import "./App.css";
import AdminLogin from "./components/admin-login";
import Chat from "./components/chat";
import SessionList from "./components/session-list";
import { apiService, type Session } from "./services/api";

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [sessionId, setSessionId] = useState("");
  const [sessionCode, setSessionCode] = useState("");
  const [guestName, setGuestName] = useState("");
  const [error, setError] = useState("");

  const handleHostLogin = async (username: string, password: string) => {
    try {
      setError("");
      await apiService.adminLogin(username, password);
      
      // Criar sessão após login
      const sessionResponse = await apiService.createSession();
      
      setIsLoggedIn(true);
      setSessionId(sessionResponse.session.id);
      setSessionCode(sessionResponse.session.code);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Login failed");
    }
  };

  const handleBack = () => {
    apiService.clearToken();
    setIsLoggedIn(false);
    setSessionId("");
    setSessionCode("");
    setGuestName("");
    setError("");
  };

  const handleSelectSession = (session: Session) => {
    setSessionId(session.id);
    setSessionCode(session.code);
    setGuestName(session.guest_name || "");
  };

  const handleNewConversation = async () => {
    try {
      const sessionResponse = await apiService.createSession();
      setSessionId(sessionResponse.session.id);
      setSessionCode(sessionResponse.session.code);
      setGuestName("");
    } catch (err) {
      console.error("Error creating new session:", err);
    }
  };

  if (!isLoggedIn) {
    return (
      <AdminLogin 
        onLogin={handleHostLogin} 
        onBack={() => window.location.href = "http://localhost:3000"}
        error={error}
      />
    );
  }

  return (
    <div className="admin-page">
      <div className="admin-header">
        <h1>Chat App - Administrador</h1>
        <button className="back-button" onClick={handleBack}>Sair</button>
      </div>
      
      <div className="admin-content">
        <div className="sidebar">
          <div className="session-info">
            <p>Compartilhe este código:</p>
            <div className="session-code">{sessionCode.toUpperCase()}</div>
            <button className="new-conversation-btn" onClick={handleNewConversation}>
              + Nova Conversa
            </button>
          </div>
          
          <SessionList 
            onSelectSession={handleSelectSession}
            currentSessionId={sessionId}
          />
        </div>
        
        <div className="chat-area">
          <div className="chat-container">
            <Chat 
              sessionId={sessionId}
              isHost={true}
              guestName={guestName}
            />
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
