import { useEffect, useState } from "react";
import { apiService, type Session } from "../services/api";
import "./session-list.css";

interface SessionListProps {
  onSelectSession: (session: Session) => void;
  currentSessionId?: string;
}

function SessionList({ onSelectSession, currentSessionId }: SessionListProps) {
  const [sessions, setSessions] = useState<Session[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    loadSessions();
  }, []);

  const loadSessions = async () => {
    try {
      setLoading(true);
      setError("");
      const data = await apiService.getAdminSessions();
      // Ordenar por data de criação (mais recentes primeiro)
      const sorted = data.sort((a, b) => 
        new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
      );
      setSessions(sorted);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Erro ao carregar sessões");
    } finally {
      setLoading(false);
    }
  };

  const handleRefresh = () => {
    loadSessions();
  };

  if (loading) {
    return (
      <div className="session-list">
        <h3>Histórico de Conversas</h3>
        <div className="session-list-loading">Carregando...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="session-list">
        <h3>Histórico de Conversas</h3>
        <div className="session-list-error">{error}</div>
      </div>
    );
  }

  if (sessions.length === 0) {
    return (
      <div className="session-list">
        <h3>Histórico de Conversas</h3>
        <div className="session-list-empty">
          <p>Nenhuma conversa ainda.</p>
          <small>Compartilhe o código da sala para iniciar uma conversa.</small>
        </div>
      </div>
    );
  }

  return (
    <div className="session-list">
      <div className="session-list-header">
        <h3>Conversas</h3>
        <button className="refresh-button" onClick={handleRefresh} title="Atualizar lista">
          ↻
        </button>
      </div>
      <div className="session-list-items">
        {sessions.map((session) => (
          <div
            key={session.id}
            className={`session-item ${currentSessionId === session.id ? "active" : ""}`}
            onClick={() => onSelectSession(session)}
          >
            <div className="session-avatar">
              {session.guest_name ? session.guest_name.charAt(0).toUpperCase() : "?"}
            </div>
            <div className="session-content">
              <div className="session-top">
                <span className="session-guest-name">
                  {session.guest_name || "Aguardando visitante..."}
                </span>
                <span className="session-date">
                  {new Date(session.created_at).toLocaleDateString('pt-BR', {
                    day: '2-digit',
                    month: '2-digit'
                  })}
                </span>
              </div>
              <div className="session-bottom">
                <span className="session-preview">
                  Código: {session.code.toUpperCase()}
                </span>
                <span className={`session-status ${session.is_active ? "active" : "inactive"}`}>
                  {session.is_active ? "●" : ""}
                </span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default SessionList;

