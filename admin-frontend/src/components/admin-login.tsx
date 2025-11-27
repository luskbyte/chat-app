import { useState } from "react";
import "./admin-login.css";

interface AdminLoginProps {
  onLogin: (username: string, password: string) => void;
  onBack: () => void;
  error?: string;
}

function AdminLogin({ onLogin, onBack, error }: AdminLoginProps) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    await onLogin(username, password);
    setIsLoading(false);
  };

  return (
    <div className="login-container">
      <form className="login-form" onSubmit={handleSubmit}>
        <h2>Login de Administrador</h2>
        <p className="login-description">Entre com suas credenciais de admin</p>
        
        {error && <div className="error-message">{error}</div>}
        
        <div className="form-group">
          <label htmlFor="username">Usuário:</label>
          <input
            type="text"
            id="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="Digite seu usuário"
            required
            disabled={isLoading}
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Senha:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Digite sua senha"
            required
            disabled={isLoading}
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
        <div className="login-hint">
          <small>Credenciais padrão: admin / admin123</small>
        </div>
      </form>
    </div>
  );
}

export default AdminLogin;

