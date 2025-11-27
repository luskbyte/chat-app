# 💬 Chat App - Sistema de Chat em Tempo Real

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
[![React](https://img.shields.io/badge/React-18+-61DAFB?logo=react)](https://reactjs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5+-3178C6?logo=typescript)](https://www.typescriptlang.org/)

Sistema de chat em tempo real com **Segurança Nível 3** - Aplicações completamente separadas para visitantes e administradores.

## 🎯 Características

- 🔐 **Segurança Máxima**: Aplicações completamente separadas (guest/admin)
- ⚡ **Chat em Tempo Real**: WebSocket para comunicação instantânea
- 🏗️ **Arquitetura Moderna**: Backend Go + Frontend React separados
- 🔑 **Autenticação Diferenciada**:
  - Admin: Login com usuário/senha + JWT
  - Visitante: Código de acesso gerado pelo admin
- 📱 **Interface Responsiva**: Design moderno estilo Telegram/WhatsApp
- 🐳 **Docker**: Pronto para produção com Docker Compose
- 📊 **Dashboard Admin**: Histórico de conversas e gerenciamento de sessões

## 🏗️ Arquitetura - Segurança Nível 3

```
┌─────────────────────────────────────────────────────┐
│  👤 VISITANTE (localhost:3000)                      │
│  guest-frontend/ - Apenas login + chat             │
│  Bundle: ~43 KB (28% menor!)                       │
│  ❌ SEM código admin                                │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│  🔐 ADMIN (localhost:3001)                          │
│  admin-frontend/ - Dashboard completo              │
│  Bundle: ~60 KB (todas as features)                │
│  ❌ SEM código guest                                │
└─────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────┐
│  ⚙️  BACKEND (localhost:8080)                       │
│  backend/ - API REST + WebSocket                   │
│  Go + JWT + bcrypt                                 │
└─────────────────────────────────────────────────────┘
```

### Estrutura do Projeto

```
chat-app/
├── backend/              # Backend Go
│   ├── cmd/server/       # Aplicação principal
│   ├── internal/
│   │   ├── api/          # Handlers REST
│   │   ├── auth/         # Autenticação JWT
│   │   ├── models/       # Estruturas de dados
│   │   └── websocket/    # WebSocket hub
│   ├── Dockerfile
│   └── go.mod
│
├── guest-frontend/       # Frontend do Visitante
│   ├── src/
│   │   ├── components/   # Login + Chat
│   │   └── services/     # API client
│   ├── Dockerfile
│   └── package.json
│
├── admin-frontend/       # Frontend do Admin
│   ├── src/
│   │   ├── components/   # Admin Login + Chat + SessionList
│   │   └── services/     # API client
│   ├── Dockerfile
│   └── package.json
│
└── docker-compose.yml    # Orquestração
```

## 🚀 Quick Start

### Com Docker (Recomendado)

```bash
# Clone o repositório
git clone https://github.com/seu-usuario/chat-app.git
cd chat-app

# Inicie todos os serviços
docker compose up --build
```

**Pronto!** Acesse:
- 👤 **Visitante**: http://localhost:3000
- 🔐 **Admin**: http://localhost:3001
- ⚙️ **Backend**: http://localhost:8080

### Desenvolvimento Local

Abra **3 terminais**:

```bash
# Terminal 1 - Backend
./start-backend.sh

# Terminal 2 - Guest Frontend
./start-guest.sh

# Terminal 3 - Admin Frontend
./start-admin.sh
```

## 🔑 Credenciais Padrão

> ⚠️ **AVISO DE SEGURANÇA**: Estas são credenciais de desenvolvimento. **NUNCA use em produção!**
> 
> Veja [SECURITY.md](./SECURITY.md) para instruções de segurança em produção.

**Admin:**
- Usuário: `admin`
- Senha: `admin123`

**Visitante:**
- Use o código gerado pelo admin após o login

## 📖 Como Usar

### Como Administrador

1. Acesse http://localhost:3001
2. Faça login com `admin` / `admin123`
3. Um código será gerado automaticamente
4. Compartilhe o código com o visitante
5. Aguarde o visitante conectar
6. Use o botão "Nova Conversa" para gerar novos códigos

### Como Visitante

1. Acesse http://localhost:3000
2. Digite seu nome
3. Insira o código fornecido pelo admin
4. Comece a conversar!

## 🔒 Segurança

### Isolamento Completo

✅ **Zero Vazamento de Código**
- Código admin NÃO existe no bundle do visitante
- Visitante NUNCA vê componentes admin

✅ **Isolamento Físico**
- Aplicações completamente separadas
- Builds independentes
- Deploys independentes

✅ **Redes Isoladas**
- Docker networks separadas
- Admin pode ter rede interna

✅ **Portas Diferentes**
- Guest: 3000 (pública)
- Admin: 3001 (pode ser bloqueada)

### Proteções Adicionais (Produção)

**Nível 2 - Firewall:**
```bash
sudo ufw allow 3000        # Guest público
sudo ufw deny 3001         # Admin bloqueado
sudo ufw allow from 192.168.1.0/24 to any port 3001
```

**Nível 3 - VPN:**
- Admin acessível apenas via VPN
- Guest público normal

**Nível 4 - Servidores Diferentes:**
- guest.seusite.com → Servidor público
- admin.seusite.com → Servidor privado/VPN

## 🛠️ Tecnologias

### Backend
- **Go** 1.21+
- **Gorilla Mux** - HTTP router
- **Gorilla WebSocket** - WebSocket
- **JWT** - Autenticação
- **bcrypt** - Hash de senhas

### Frontend
- **React** 18
- **TypeScript** 5
- **Vite** - Build tool
- **Nginx** - Servidor (Docker)

## 📚 Documentação

- 🔒 [SECURITY.md](./SECURITY.md) - **LEIA ANTES DE FAZER DEPLOY!**
- 📖 [LEVEL3-SECURITY.md](./LEVEL3-SECURITY.md) - Detalhes de segurança
- 🚀 [COMO-EXECUTAR.md](./COMO-EXECUTAR.md) - Instruções completas
- 📝 [CHANGELOG-LEVEL3.md](./CHANGELOG-LEVEL3.md) - Changelog
- 📄 [FILES_CREATED.md](./FILES_CREATED.md) - Lista de arquivos

## 🔌 API Endpoints

### Autenticação

**Admin Login**
```http
POST /api/admin/login
Content-Type: application/json

{
  "username": "admin",
  "password": "admin123"
}
```

**Guest Login**
```http
POST /api/guest/login
Content-Type: application/json

{
  "code": "ABC123",
  "guest_name": "João"
}
```

### Sessões

**Criar Sessão** (Admin)
```http
POST /api/session/create
Authorization: Bearer <token>
```

**Listar Sessões** (Admin)
```http
GET /api/admin/sessions
Authorization: Bearer <token>
```

### Mensagens

**Obter Mensagens**
```http
GET /api/messages?sessionID=<id>
Authorization: Bearer <token>
```

### WebSocket

**Conectar ao Chat**
```
WS /ws?token=<token>&sessionID=<id>
```

## 📊 Comparação

| Aspecto | Antes (Rotas) | Agora (Separado) |
|---------|---------------|------------------|
| Bundle Guest | ~60 KB | ~43 KB (-28%) ✅ |
| Bundle Admin | ~60 KB | ~60 KB |
| Código Admin | Incluído ❌ | NÃO EXISTE! ✅ |
| Isolamento | Lógico | Físico ✅ |
| Portas | 1 (5173) | 2 (3000, 3001) ✅ |
| Deploy | Único | Independente ✅ |
| Segurança | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |

## 🐳 Docker

**Iniciar:**
```bash
docker compose up --build
```

**Parar:**
```bash
docker compose down
```

**Rebuild completo:**
```bash
docker compose down -v
docker compose up --build
```

**Ver logs:**
```bash
docker compose logs -f
docker compose logs -f guest-frontend
docker compose logs -f admin-frontend
```

## 🎯 Roadmap

- [ ] Banco de dados persistente (PostgreSQL)
- [ ] Upload de arquivos/imagens
- [ ] Notificações de digitação
- [ ] Emojis e markdown
- [ ] Histórico paginado
- [ ] Testes automatizados
- [ ] CI/CD
- [ ] Kubernetes deployment

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor:

1. Fork o projeto
2. Crie uma branch (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## ✨ Créditos

Desenvolvido com ❤️ usando:
- [Go](https://go.dev/)
- [React](https://reactjs.org/)
- [TypeScript](https://www.typescriptlang.org/)
- [Docker](https://www.docker.com/)

---

**⭐ Se você gostou deste projeto, considere dar uma estrela no GitHub!**
