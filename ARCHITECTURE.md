# Arquitetura do Sistema de Chat

## 📐 Visão Geral

Este é um sistema de chat em tempo real que separa claramente as responsabilidades entre frontend e backend, seguindo uma arquitetura cliente-servidor moderna.

## 🏛️ Estrutura do Backend (Go)

### Camadas

```
backend/
├── cmd/server/         # Ponto de entrada da aplicação
├── internal/
│   ├── api/           # Camada de apresentação (HTTP handlers)
│   ├── auth/          # Camada de autenticação e autorização
│   ├── models/        # Modelos de dados (DTOs)
│   └── websocket/     # Gerenciamento de conexões WebSocket
└── pkg/               # Pacotes reutilizáveis (vazio por enquanto)
```

### Componentes Principais

#### 1. API Layer (`internal/api`)
- **handlers.go**: Define todos os endpoints HTTP
- Responsável por validar requisições
- Converte entre JSON e modelos internos
- Gerencia autenticação via JWT

#### 2. Auth Layer (`internal/auth`)
- **auth.go**: Funções de autenticação (JWT, bcrypt)
- **store.go**: Armazenamento em memória (substituir por DB em produção)
- Gerencia tokens JWT
- Hash de senhas com bcrypt
- Geração de códigos de sessão

#### 3. WebSocket Layer (`internal/websocket`)
- **hub.go**: Gerencia todas as conexões WebSocket
- **client.go**: Representa um cliente conectado
- **register.go**: Funções de registro de clientes
- Broadcast de mensagens para sessões específicas
- Gerenciamento do ciclo de vida das conexões

#### 4. Models Layer (`internal/models`)
- Define estruturas de dados compartilhadas
- DTOs para requests e responses
- Modelos de domínio (Admin, Session, Message)

## 🎨 Estrutura do Frontend (React + TypeScript)

### Camadas

```
frontend/src/
├── components/        # Componentes React
│   ├── admin-login.tsx
│   ├── login.tsx
│   ├── chat.tsx
│   └── *.css
├── services/          # Camada de serviços
│   └── api.ts
├── App.tsx            # Componente principal
└── main.tsx           # Ponto de entrada
```

### Componentes Principais

#### 1. Services Layer (`services/api.ts`)
- Encapsula toda comunicação com backend
- Gerencia tokens JWT no localStorage
- Fornece métodos para:
  - Login de admin
  - Login de visitante
  - Criação de sessões
  - Conexão WebSocket
  - Recuperação de mensagens

#### 2. Components Layer
- **App.tsx**: Gerencia roteamento e estado global
- **admin-login.tsx**: Formulário de login do host
- **login.tsx**: Formulário de login do visitante
- **chat.tsx**: Interface de chat com WebSocket

## 🔐 Fluxo de Autenticação

### Host (Admin)

```
1. Usuário escolhe "Host" → AdminLogin
2. Envia credenciais → POST /api/admin/login
3. Backend valida credenciais
4. Backend gera JWT token (type: admin)
5. Frontend armazena token
6. Frontend cria sessão → POST /api/session/create
7. Backend gera código único de 6 caracteres
8. Frontend exibe código para compartilhar
9. Host conecta ao WebSocket com token
```

### Visitante (Guest)

```
1. Usuário escolhe "Visitante" → Login
2. Insere nome e código da sala
3. Envia dados → POST /api/guest/login
4. Backend valida código da sessão
5. Backend gera JWT token (type: guest)
6. Frontend armazena token
7. Frontend conecta ao WebSocket com token
```

## 🔌 Fluxo de Comunicação WebSocket

### Conexão

```
1. Cliente autentica via HTTP (obtém JWT)
2. Cliente conecta ao WebSocket: /ws?token=JWT&sessionID=ID
3. Backend valida token
4. Backend registra cliente no Hub
5. Hub mantém mapa de sessões → clientes
```

### Envio de Mensagem

```
1. Cliente envia mensagem via WebSocket
2. Hub recebe mensagem
3. Hub identifica sessionID do cliente
4. Hub faz broadcast para todos clientes da mesma sessão
5. Todos clientes recebem mensagem em tempo real
```

### Desconexão

```
1. Cliente desconecta (intencional ou erro)
2. Hub detecta desconexão
3. Hub remove cliente do registro
4. Hub fecha canal de envio do cliente
```

## 💾 Armazenamento de Dados

### Atual (Em Memória)

```go
type Store struct {
    admins   map[string]*Admin      // username → Admin
    sessions map[string]*Session    // code → Session
    messages map[string][]*Message  // sessionID → []Message
}
```

**Vantagens:**
- Simples de implementar
- Rápido para desenvolvimento/testes
- Zero configuração

**Desvantagens:**
- Dados perdidos ao reiniciar servidor
- Não escala horizontalmente
- Limitado pela memória RAM

### Recomendado para Produção

**PostgreSQL + Redis:**

```
PostgreSQL:
- Tabela users (admins)
- Tabela sessions
- Tabela messages (histórico)

Redis:
- Cache de sessões ativas
- Lista de usuários online
- Pub/Sub para múltiplas instâncias
```

## 🔒 Segurança

### Autenticação
- JWT com expiração de 24h
- Senhas com bcrypt (cost 10)
- Validação de token em todos endpoints protegidos

### CORS
- Configurado para origens específicas
- Credentials habilitados
- Headers permitidos controlados

### WebSocket
- Token validado na conexão
- SessionID validado
- Clientes isolados por sessão

## 🚀 Escalabilidade

### Pontos de Melhoria

1. **Banco de Dados Persistente**
   - Implementar repository pattern
   - Usar PostgreSQL para dados permanentes
   - Usar Redis para cache e sessões

2. **Múltiplas Instâncias**
   - Redis Pub/Sub para broadcast entre instâncias
   - Load balancer com sticky sessions
   - Shared storage para sessões

3. **Microserviços**
   - Separar serviço de autenticação
   - Separar serviço de mensagens
   - Usar message broker (RabbitMQ/Kafka)

4. **Observabilidade**
   - Logging estruturado (logrus/zap)
   - Métricas (Prometheus)
   - Tracing distribuído (Jaeger)

## 📊 Diagrama de Sequência - Fluxo Completo

```
Host                Frontend           Backend              WebSocket Hub
 |                      |                  |                      |
 |--[1] Choose Host---->|                  |                      |
 |                      |                  |                      |
 |--[2] Login--------->|--POST /admin/login->|                   |
 |                      |<---JWT Token-------|                   |
 |                      |                  |                      |
 |                      |--POST /session/create->|               |
 |                      |<---Session+Code----|                   |
 |                      |                  |                      |
 |<-[3] Display Code----|                  |                      |
 |                      |                  |                      |
 |--[4] Connect WS----->|--WS /ws?token--->|--Register Client---->|
 |                      |                  |                      |
 |                      |                  |    [Guest Joins]     |
 |                      |                  |                      |
 |--[5] Send Msg------->|----WS Message--->|----Broadcast-------->|
 |                      |                  |                      |
 |<-----Receive---------|<----WS Message---|<--To All Clients-----|
Guest                   |<----WS Message---|<--To All Clients-----|
 |<-----Receive---------|                  |                      |
```

## 🧪 Testes (A Implementar)

### Backend
- Testes unitários para auth
- Testes de integração para API
- Testes de WebSocket
- Mock do Store

### Frontend
- Testes de componentes (Vitest + Testing Library)
- Testes E2E (Playwright)
- Testes de integração com backend mockado

## 📝 Próximos Passos

1. ✅ Implementar autenticação diferenciada
2. ✅ WebSocket para chat em tempo real
3. ✅ Geração de códigos de sessão
4. ⏳ Migrar para banco de dados PostgreSQL
5. ⏳ Implementar Redis para cache
6. ⏳ Adicionar testes automatizados
7. ⏳ Implementar CI/CD
8. ⏳ Deploy em produção (AWS/GCP/Azure)

