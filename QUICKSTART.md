# 🚀 Guia Rápido de Início

## Pré-requisitos

- **Go 1.21+** - [Download](https://go.dev/dl/)
- **Node.js 18+** - [Download](https://nodejs.org/)
- **npm** ou **yarn**

## 🎯 Início Rápido (2 minutos)

### Opção 1: Scripts Automatizados (Recomendado)

1. **Iniciar o Backend:**
```bash
./start-backend.sh
```

2. **Em outro terminal, iniciar o Frontend:**
```bash
./start-frontend.sh
```

### Opção 2: Manual

#### Backend
```bash
cd backend
go mod download
go run cmd/server/main.go
```

#### Frontend
```bash
cd frontend
npm install
npm run dev
```

### Opção 3: Docker (Produção)

```bash
docker-compose up
```

## 🎮 Como Testar

### 1. Abrir como Host

1. Abra o navegador em `http://localhost:5173`
2. Clique em **"Host"**
3. Entre com:
   - **Usuário:** `admin`
   - **Senha:** `admin123`
4. Um código será gerado (ex: `a3b4c5`)
5. Copie o código

### 2. Abrir como Visitante

1. Abra outra aba ou janela anônima em `http://localhost:5173`
2. Clique em **"Visitante"**
3. Entre com:
   - **Seu nome:** `João` (ou qualquer nome)
   - **Código da sala:** Cole o código do Host
4. Clique em **"Entrar"**

### 3. Conversar

Agora você pode enviar mensagens entre as duas janelas em tempo real! 🎉

## 🔧 Resolução de Problemas

### Go não instalado

```bash
# Ubuntu/Debian
sudo apt install golang-go

# ou via snap
sudo snap install go
```

### Porta 8080 já em uso

```bash
# Encontrar processo usando a porta
lsof -i :8080

# Matar o processo (substitua PID)
kill -9 PID
```

### Porta 5173 já em uso

Edite `frontend/vite.config.ts`:

```typescript
export default defineConfig({
  server: {
    port: 3000  // Mude para outra porta
  }
})
```

### Erro de CORS

Verifique se o backend está rodando e se as origens permitidas estão configuradas corretamente em `backend/cmd/server/main.go`:

```go
AllowedOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
```

### WebSocket não conecta

1. Verifique se o backend está rodando
2. Verifique o console do navegador (F12) para erros
3. Certifique-se de que o token JWT foi gerado corretamente
4. Tente fazer logout e login novamente

## 📊 Endpoints da API

### Autenticação

**Login Admin:**
```bash
curl -X POST http://localhost:8080/api/admin/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

**Login Visitante:**
```bash
curl -X POST http://localhost:8080/api/guest/login \
  -H "Content-Type: application/json" \
  -d '{"code":"abc123","guest_name":"João"}'
```

### Sessões

**Criar Sessão:**
```bash
curl -X POST http://localhost:8080/api/session/create \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Health Check

```bash
curl http://localhost:8080/health
```

## 🔑 Credenciais Padrão

**Admin:**
- Usuário: `admin`
- Senha: `admin123`

**Visitante:**
- Código: Gerado pelo host após login
- Nome: Qualquer nome que você escolher

## 📝 Estrutura de URLs

| Serviço | URL Local | URL Produção |
|---------|-----------|--------------|
| Frontend | http://localhost:5173 | - |
| Backend API | http://localhost:8080/api | - |
| WebSocket | ws://localhost:8080/ws | - |
| Health Check | http://localhost:8080/health | - |

## 🎨 Customização Rápida

### Mudar credenciais do admin

Edite `backend/internal/auth/store.go`:

```go
hashedPassword, _ := HashPassword("sua_senha_aqui")
s.admins["seu_usuario"] = &models.Admin{
    ID:        uuid.New().String(),
    Username:  "seu_usuario",
    Password:  hashedPassword,
    CreatedAt: time.Now(),
}
```

### Mudar tempo de expiração da sessão

Edite `backend/internal/auth/store.go`:

```go
session := &models.Session{
    // ...
    ExpiresAt: time.Now().Add(48 * time.Hour), // Era 24h, agora 48h
}
```

### Mudar cores do tema

Edite `frontend/src/App.css` e procure por:

```css
background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
```

Substitua os códigos de cor por suas preferências!

## 📚 Próximos Passos

- Leia [README.md](./README.md) para documentação completa
- Leia [ARCHITECTURE.md](./ARCHITECTURE.md) para entender a arquitetura
- Explore o código e faça suas modificações!

## 🆘 Precisa de Ajuda?

1. Verifique os logs do backend no terminal
2. Verifique o console do navegador (F12)
3. Leia a documentação completa no README.md
4. Verifique se todas as dependências estão instaladas

---

**Dica:** Use duas janelas do navegador lado a lado para ver as mensagens em tempo real! 💬

