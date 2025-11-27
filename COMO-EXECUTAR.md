# 🚀 Como Executar - Segurança Nível 3

## 📋 Pré-requisitos

- Docker e Docker Compose instalados
- Portas disponíveis: 3000, 3001, 8080

## 🎯 Opção 1: Docker (Recomendado)

### Iniciar todas as aplicações:

```bash
sudo docker compose up --build
```

### Acessar:

- **Visitante:** http://localhost:3000
- **Admin:** http://localhost:3001
- **API:** http://localhost:8080

### Parar:

```bash
sudo docker compose down
```

## 🛠️ Opção 2: Desenvolvimento Local

### Abrir 3 terminais:

**Terminal 1 - Backend:**
```bash
./start-backend.sh
```

**Terminal 2 - Guest Frontend:**
```bash
./start-guest.sh
```

**Terminal 3 - Admin Frontend:**
```bash
./start-admin.sh
```

### Acessar:

- **Visitante:** http://localhost:3000
- **Admin:** http://localhost:3001
- **API:** http://localhost:8080

## 🔑 Credenciais Padrão

**Admin:**
- Usuário: `admin`
- Senha: `admin123`

**Visitante:**
- Use o código gerado pelo admin após login

## 📦 Estrutura

```
chat-app/
├── backend/              # Backend Go (porta 8080)
├── guest-frontend/       # Frontend Visitante (porta 3000)
├── admin-frontend/       # Frontend Admin (porta 3001)
└── docker-compose.yml    # Orquestração dos 3 serviços
```

## 🔒 Segurança

### Nível Atual
✅ Aplicações completamente separadas
✅ Código admin NÃO existe no bundle guest
✅ Portas diferentes
✅ Deploys independentes

### Melhorias para Produção

**1. Bloquear acesso externo ao admin:**
```bash
sudo ufw allow 3000       # Guest (público)
sudo ufw deny 3001        # Admin (bloquear)
```

**2. Permitir admin apenas de IPs específicos:**
```bash
sudo ufw allow from 192.168.1.0/24 to any port 3001
```

**3. Usar subdomínios:**
- guest.seusite.com → porta 3000
- admin.seusite.com → porta 3001

## 🐛 Troubleshooting

### Porta em uso:
```bash
# Liberar portas
lsof -ti:3000,3001,8080 | xargs kill -9
```

### Rebuild completo:
```bash
sudo docker compose down -v
sudo docker compose up --build
```

### Logs:
```bash
# Ver logs em tempo real
sudo docker compose logs -f

# Logs de um serviço específico
sudo docker compose logs -f guest-frontend
sudo docker compose logs -f admin-frontend
sudo docker compose logs -f backend
```

## 📖 Mais Informações

- **Arquitetura:** LEVEL3-SECURITY.md
- **Documentação completa:** FILES_CREATED.md

