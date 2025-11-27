import { useState } from "react";
import "./login.css";

interface LoginProps {
  onLogin: (userName: string, code: string) => void;
  onBack: () => void;
  error?: string;
}

function Login({ onLogin, onBack, error }: LoginProps) {
  const [code, setCode] = useState("");
  const [userName, setUserName] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    await onLogin(userName, code);
    setIsLoading(false);
  };

  return (
    <div className="login-container">
      <form className="login-form" onSubmit={handleSubmit}>
        <h2>Login de Visitante</h2>
        <p className="login-description">Entre com o código fornecido pelo host</p>
        
        {error && <div className="error-message">{error}</div>}
        
        <div className="form-group">
          <label htmlFor="userName">Seu nome:</label>
          <input
            type="text"
            id="userName"
            value={userName}
            onChange={(e) => setUserName(e.target.value)}
            placeholder="Digite seu nome"
            required
            disabled={isLoading}
          />
        </div>
        <div className="form-group">
          <label htmlFor="code">Código da sala:</label>
          <input
            type="text"
            id="code"
            value={code}
            onChange={(e) => setCode(e.target.value.toLowerCase())}
            placeholder="Digite o código da sala"
            required
            disabled={isLoading}
            maxLength={6}
          />
        </div>
        <div className="form-actions">
          <button type="button" onClick={onBack} className="btn-secondary" disabled={isLoading}>
            Voltar
          </button>
          <button type="submit" className="btn-primary" disabled={isLoading}>
            {isLoading ? "Entrando..." : "Entrar"}
          </button>
        </div>
      </form>
    </div>
  );
}

export default Login;
