# Contribuindo para o Chat App

Obrigado por considerar contribuir para o Chat App! 🎉

## 📋 Código de Conduta

Este projeto adere a um código de conduta. Ao participar, espera-se que você mantenha este código.

## 🚀 Como Contribuir

### Reportando Bugs

Se você encontrou um bug, por favor abra uma issue incluindo:

- Descrição clara do problema
- Passos para reproduzir
- Comportamento esperado vs atual
- Screenshots (se aplicável)
- Informações do ambiente (OS, versão do Docker, etc.)

### Sugerindo Features

Para sugerir uma nova feature:

1. Verifique se já não existe uma issue similar
2. Abra uma nova issue com o template "Feature Request"
3. Descreva claramente a feature e seu caso de uso

### Pull Requests

1. **Fork o repositório**

2. **Clone seu fork**
   ```bash
   git clone https://github.com/seu-usuario/chat-app.git
   cd chat-app
   ```

3. **Crie uma branch**
   ```bash
   git checkout -b feature/minha-feature
   # ou
   git checkout -b fix/meu-fix
   ```

4. **Faça suas alterações**
   - Mantenha o código limpo e bem documentado
   - Siga os padrões de código existentes
   - Adicione testes se aplicável

5. **Teste suas alterações**
   ```bash
   # Backend
   cd backend && go test ./...
   
   # Frontend
   cd guest-frontend && npm run build
   cd admin-frontend && npm run build
   
   # Docker
   docker compose up --build
   ```

6. **Commit suas mudanças**
   ```bash
   git add .
   git commit -m "feat: adiciona nova feature"
   ```
   
   Use o padrão de commits:
   - `feat:` para novas features
   - `fix:` para correções
   - `docs:` para documentação
   - `style:` para formatação
   - `refactor:` para refatoração
   - `test:` para testes
   - `chore:` para tarefas gerais

7. **Push para seu fork**
   ```bash
   git push origin feature/minha-feature
   ```

8. **Abra um Pull Request**
   - Vá para o repositório original
   - Clique em "New Pull Request"
   - Selecione sua branch
   - Preencha o template do PR

## 🏗️ Estrutura do Projeto

```
chat-app/
├── backend/              # Backend Go
├── guest-frontend/       # Frontend Visitante
├── admin-frontend/       # Frontend Admin
└── docs/                 # Documentação
```

## 🔧 Desenvolvimento Local

### Requisitos

- Go 1.21+
- Node.js 18+
- Docker (opcional)

### Setup

```bash
# Backend
cd backend
go mod download
go run cmd/server/main.go

# Guest Frontend
cd guest-frontend
npm install
npm run dev

# Admin Frontend
cd admin-frontend
npm install
npm run dev
```

## 📝 Padrões de Código

### Go (Backend)

- Use `gofmt` para formatação
- Siga as [Effective Go guidelines](https://golang.org/doc/effective_go.html)
- Adicione comentários em funções públicas
- Use nomes descritivos

### TypeScript/React (Frontend)

- Use TypeScript strict mode
- Componentes funcionais com hooks
- Props tipadas
- CSS modules ou styled-components
- ESLint + Prettier

### Docker

- Multi-stage builds
- Imagens mínimas (alpine)
- .dockerignore configurado

## ✅ Checklist do PR

Antes de submeter seu PR, certifique-se que:

- [ ] O código compila sem erros
- [ ] Todos os testes passam
- [ ] A documentação foi atualizada (se necessário)
- [ ] O código segue os padrões do projeto
- [ ] Commits seguem o padrão semântico
- [ ] O PR tem uma descrição clara

## 🐛 Debugging

### Backend

```bash
cd backend
go run -race cmd/server/main.go
```

### Frontend

```bash
cd guest-frontend
npm run dev
```

Abra as DevTools do navegador (F12)

### Docker

```bash
docker compose logs -f
docker compose logs -f backend
docker compose logs -f guest-frontend
```

## 📚 Recursos

- [Documentação Go](https://go.dev/doc/)
- [Documentação React](https://react.dev/)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/)
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)

## 💬 Dúvidas?

Se você tiver dúvidas:

1. Verifique a documentação
2. Procure em issues existentes
3. Abra uma nova issue com a tag "question"

## 🙏 Obrigado!

Suas contribuições tornam este projeto melhor para todos! 🎉

