import { useState } from "react";
import "./App.css";
import Login from "./components/login";
import Chat from "./components/chat";
import { apiService } from "./services/api";

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [sessionId, setSessionId] = useState("");
  const [guestName, setGuestName] = useState("");
  const [error, setError] = useState("");

  const handleGuestLogin = async (name: string, code: string) => {
    try {
      setError("");
      const response = await apiService.guestLogin(code, name);
      
      setIsLoggedIn(true);
      setSessionId(response.session_id);
      setGuestName(name);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Login failed");
    }
  };

  const handleBack = () => {
    apiService.clearToken();
    setIsLoggedIn(false);
    setSessionId("");
    setGuestName("");
    setError("");
  };

  if (!isLoggedIn) {
    return (
      <>
        <Login 
          onLogin={handleGuestLogin}
          onBack={() => {}}
          error={error}
        />
        <div className="admin-link-container">
          <a 
            href="http://localhost:3001" 
            target="_blank" 
            rel="noopener noreferrer"
            className="admin-link-button"
          >
            🔑 Acesso Administrativo
          </a>
        </div>
      </>
    );
  }

  return (
    <div className="guest-page">
      <div className="guest-header">
        <h1>Chat App - Visitante</h1>
        <button className="back-button" onClick={handleBack}>Sair</button>
      </div>
      
      <div className="guest-chat-area">
        <div className="chat-container">
          <Chat 
            sessionId={sessionId}
            isHost={false}
            guestName={guestName}
          />
        </div>
      </div>
    </div>
  );
}

export default App;
