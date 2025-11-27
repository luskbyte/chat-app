import { Navigate } from "react-router-dom";
import { apiService } from "../services/api";

interface ProtectedRouteProps {
  children: React.ReactNode;
}

function ProtectedRoute({ children }: ProtectedRouteProps) {
  const token = apiService.getToken();

  // Se não tem token, redireciona para admin login
  if (!token) {
    return <Navigate to="/admin" replace />;
  }

  // TODO: Idealmente, validar o token com o backend aqui
  // Por enquanto, apenas verifica se existe

  return <>{children}</>;
}

export default ProtectedRoute;

