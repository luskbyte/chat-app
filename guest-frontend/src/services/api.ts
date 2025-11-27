const API_URL = 'http://localhost:8080/api';

export interface AdminLoginRequest {
  username: string;
  password: string;
}

export interface AdminLoginResponse {
  token: string;
  admin_id: string;
  message: string;
}

export interface GuestLoginRequest {
  code: string;
  guest_name: string;
}

export interface GuestLoginResponse {
  token: string;
  session_id: string;
  message: string;
}

export interface Session {
  id: string;
  admin_id: string;
  code: string;
  guest_name?: string;
  is_active: boolean;
  created_at: string;
  expires_at: string;
}

export interface CreateSessionResponse {
  session: Session;
  message: string;
}

export interface Message {
  id: string;
  session_id: string;
  sender: string;
  content: string;
  timestamp: string;
}

export class ApiService {
  private token: string | null = null;

  setToken(token: string) {
    this.token = token;
    localStorage.setItem('token', token);
  }

  getToken(): string | null {
    if (!this.token) {
      this.token = localStorage.getItem('token');
    }
    return this.token;
  }

  clearToken() {
    this.token = null;
    localStorage.removeItem('token');
  }

  async adminLogin(username: string, password: string): Promise<AdminLoginResponse> {
    const response = await fetch(`${API_URL}/admin/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Login failed');
    }

    const data = await response.json();
    this.setToken(data.token);
    return data;
  }

  async guestLogin(code: string, guestName: string): Promise<GuestLoginResponse> {
    const response = await fetch(`${API_URL}/guest/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ code, guest_name: guestName }),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Login failed');
    }

    const data = await response.json();
    this.setToken(data.token);
    return data;
  }

  async createSession(): Promise<CreateSessionResponse> {
    const token = this.getToken();
    if (!token) {
      throw new Error('No authentication token');
    }

    const response = await fetch(`${API_URL}/session/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to create session');
    }

    return await response.json();
  }

  async getMessages(sessionId: string): Promise<Message[]> {
    const token = this.getToken();
    if (!token) {
      throw new Error('No authentication token');
    }

    const response = await fetch(`${API_URL}/messages?sessionID=${sessionId}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to fetch messages');
    }

    return await response.json();
  }

  async getAdminSessions(): Promise<Session[]> {
    const token = this.getToken();
    if (!token) {
      throw new Error('No authentication token');
    }

    const response = await fetch(`${API_URL}/sessions`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to fetch sessions');
    }

    return await response.json();
  }

  connectWebSocket(sessionId: string, onMessage: (data: any) => void): WebSocket {
    const token = this.getToken();
    if (!token) {
      throw new Error('No authentication token');
    }

    const ws = new WebSocket(`ws://localhost:8080/ws?token=${token}&sessionID=${sessionId}`);

    ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        onMessage(data);
      } catch (error) {
        console.error('Error parsing message:', error);
      }
    };

    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    ws.onclose = () => {
      console.log('WebSocket connection closed');
    };

    return ws;
  }
}

export const apiService = new ApiService();

