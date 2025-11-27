# 🔒 Política de Segurança

## ⚠️ Aviso Importante

Este projeto contém credenciais padrão para facilitar o desenvolvimento. **NUNCA use essas credenciais em produção!**

## 🚨 Antes de Fazer Deploy em Produção

### 1. Alterar Credenciais de Admin

**Arquivo:** `backend/internal/auth/store.go`

```go
// ALTERE ESTAS CREDENCIAIS!
hashedPassword, _ := HashPassword("SUA_SENHA_FORTE_AQUI")
adminStore["admin"] = &models.Admin{
    ID:       "admin-123",
    Username: "seu_usuario",  // Altere o username
    Password: hashedPassword,
}
```

### 2. Usar Variável de Ambiente para JWT_SECRET

**Arquivo:** `backend/internal/auth/auth.go`

Substitua:
```go
var jwtSecret = []byte("your-secret-key-change-in-production")
```

Por:
```go
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
```

E configure a variável de ambiente:
```bash
export JWT_SECRET="sua-chave-secreta-muito-forte-e-aleatoria-minimo-32-caracteres"
```

### 3. Configurar Docker Compose para Produção

**Arquivo:** `docker-compose.yml`

```yaml
services:
  backend:
    environment:
      - JWT_SECRET=${JWT_SECRET}  # Use variável de ambiente
      - ADMIN_USERNAME=${ADMIN_USERNAME}
      - ADMIN_PASSWORD=${ADMIN_PASSWORD}
```

### 4. Bloquear Acesso Admin Externamente

```bash
# Permitir apenas visitante (porta 3000) publicamente
sudo ufw allow 3000

# Bloquear admin (porta 3001) externamente
sudo ufw deny 3001

# Permitir admin apenas de IPs específicos
sudo ufw allow from 192.168.1.0/24 to any port 3001
```

### 5. Usar HTTPS em Produção

Configure certificados SSL/TLS:
- Let's Encrypt (recomendado)
- Certbot
- Nginx com SSL

### 6. Banco de Dados Persistente

O projeto atual usa armazenamento em memória. Para produção:
- PostgreSQL
- MySQL
- MongoDB

## 🔐 Checklist de Segurança para Produção

- [ ] Alterar username e senha do admin
- [ ] JWT_SECRET forte e aleatório (mínimo 32 caracteres)
- [ ] JWT_SECRET como variável de ambiente
- [ ] Firewall configurado (bloquear porta 3001 externamente)
- [ ] HTTPS configurado com certificado válido
- [ ] Banco de dados persistente configurado
- [ ] Backup automático configurado
- [ ] Rate limiting implementado
- [ ] Logs de segurança configurados
- [ ] CORS configurado adequadamente
- [ ] Headers de segurança (HSTS, CSP, etc.)

## 🛡️ Níveis de Proteção

### Desenvolvimento (Atual)
✅ Credenciais padrão
✅ HTTP
✅ Armazenamento em memória

### Produção (Recomendado)
✅ Credenciais fortes únicas
✅ HTTPS obrigatório
✅ Banco de dados persistente
✅ Firewall ativo
✅ VPN para admin (opcional)
✅ IPs whitelisted para admin

## 📝 Gerando Senhas Fortes

### JWT_SECRET
```bash
# Linux/Mac
openssl rand -base64 48

# Node.js
node -e "console.log(require('crypto').randomBytes(48).toString('base64'))"
```

### Senha Admin
```bash
# Gerar senha aleatória forte
openssl rand -base64 24
```

## 🔍 Verificando Vulnerabilidades

### Backend (Go)
```bash
cd backend
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

### Frontend
```bash
cd guest-frontend
npm audit

cd ../admin-frontend
npm audit
```

## 📧 Reportar Vulnerabilidades

Se você encontrou uma vulnerabilidade de segurança:

1. **NÃO** abra uma issue pública
2. Envie um email para: [seu-email@example.com]
3. Inclua:
   - Descrição da vulnerabilidade
   - Passos para reproduzir
   - Impacto potencial
   - Sugestão de correção (se possível)

Responderemos em até 48 horas.

## 🆕 Atualizações de Segurança

### Como Atualizar Dependências

**Backend:**
```bash
cd backend
go get -u ./...
go mod tidy
```

**Frontend:**
```bash
cd guest-frontend
npm update
npm audit fix

cd ../admin-frontend
npm update
npm audit fix
```

## 📚 Recursos de Segurança

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [Go Security](https://go.dev/doc/security/)
- [React Security](https://reactjs.org/docs/dom-elements.html#dangerouslysetinnerhtml)
- [JWT Best Practices](https://tools.ietf.org/html/rfc8725)

---

**Última atualização:** 27 de Novembro de 2025

**⚠️ A segurança é responsabilidade de todos. Sempre revise o código antes de fazer deploy em produção!**
