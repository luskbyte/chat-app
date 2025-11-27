# 🔧 Guia de Instalação de Dependências

## 📋 Pré-requisitos

Este projeto requer:
- **Go 1.21+** (backend)
- **Node.js 18+** (frontend)
- **npm** ou **yarn** (gerenciador de pacotes)

## 🐹 Instalando Go

### Ubuntu/Debian

#### Opção 1: Via APT (Mais Simples)

```bash
# Atualizar repositórios
sudo apt update

# Instalar Go
sudo apt install golang-go

# Verificar instalação
go version
```

#### Opção 2: Via Snap (Versão Mais Recente)

```bash
# Instalar via snap
sudo snap install go --classic

# Verificar instalação
go version
```

#### Opção 3: Instalação Manual (Versão Específica)

```bash
# Baixar Go 1.21.5 (ajuste a versão conforme necessário)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz

# Remover instalação anterior (se existir)
sudo rm -rf /usr/local/go

# Extrair para /usr/local
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

# Adicionar ao PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verificar instalação
go version
```

### Fedora/RHEL/CentOS

```bash
sudo dnf install golang

# ou
sudo yum install golang
```

### Arch Linux

```bash
sudo pacman -S go
```

### macOS

```bash
# Via Homebrew
brew install go

# ou baixar do site oficial
# https://go.dev/dl/
```

### Windows

1. Baixar instalador: https://go.dev/dl/
2. Executar o instalador (.msi)
3. Reiniciar terminal
4. Verificar: `go version`

## 📦 Instalando Node.js

### Ubuntu/Debian

#### Opção 1: Via APT

```bash
# Instalar Node.js
sudo apt update
sudo apt install nodejs npm

# Verificar instalação
node --version
npm --version
```

#### Opção 2: Via NodeSource (Versão Mais Recente)

```bash
# Node.js 20.x (LTS)
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt-get install -y nodejs

# Verificar instalação
node --version
npm --version
```

#### Opção 3: Via NVM (Recomendado)

```bash
# Instalar NVM
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash

# Reiniciar terminal ou executar
source ~/.bashrc

# Instalar Node.js
nvm install 20
nvm use 20

# Verificar instalação
node --version
npm --version
```

### Fedora/RHEL/CentOS

```bash
sudo dnf install nodejs npm
```

### Arch Linux

```bash
sudo pacman -S nodejs npm
```

### macOS

```bash
# Via Homebrew
brew install node

# ou via NVM (recomendado)
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
nvm install 20
```

### Windows

1. Baixar instalador: https://nodejs.org/
2. Executar o instalador (.msi)
3. Reiniciar terminal
4. Verificar: `node --version` e `npm --version`

## 🐳 Instalando Docker (Opcional)

### Ubuntu/Debian

```bash
# Remover versões antigas
sudo apt remove docker docker-engine docker.io containerd runc

# Instalar dependências
sudo apt update
sudo apt install ca-certificates curl gnupg

# Adicionar chave GPG
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

# Adicionar repositório
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Instalar Docker
sudo apt update
sudo apt install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Adicionar usuário ao grupo docker (opcional)
sudo usermod -aG docker $USER

# Verificar instalação
docker --version
docker compose version
```

### Outras Plataformas

- **macOS:** Docker Desktop - https://docs.docker.com/desktop/install/mac-install/
- **Windows:** Docker Desktop - https://docs.docker.com/desktop/install/windows-install/

## ✅ Verificação Final

Execute os seguintes comandos para verificar se tudo está instalado:

```bash
# Go
go version
# Esperado: go version go1.21.x linux/amd64 (ou similar)

# Node.js
node --version
# Esperado: v18.x.x ou v20.x.x

# npm
npm --version
# Esperado: 9.x.x ou 10.x.x

# Docker (opcional)
docker --version
# Esperado: Docker version 24.x.x

docker compose version
# Esperado: Docker Compose version v2.x.x
```

## 🚀 Próximos Passos

Após instalar as dependências:

1. **Iniciar o Backend:**
   ```bash
   cd backend
   go mod download
   go run cmd/server/main.go
   ```

2. **Iniciar o Frontend:**
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

Ou use os scripts de inicialização:

```bash
# Terminal 1
./start-backend.sh

# Terminal 2
./start-frontend.sh
```

## 🆘 Problemas Comuns

### Go: "command not found"

**Problema:** O PATH não foi configurado corretamente.

**Solução:**
```bash
# Adicionar ao ~/.bashrc ou ~/.zshrc
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$HOME/go/bin

# Aplicar mudanças
source ~/.bashrc  # ou source ~/.zshrc
```

### Node: "EACCES: permission denied"

**Problema:** Permissões do npm.

**Solução:**
```bash
# Usar NVM (recomendado)
# ou
sudo chown -R $(whoami) ~/.npm
sudo chown -R $(whoami) /usr/local/lib/node_modules
```

### Docker: "permission denied while trying to connect"

**Problema:** Usuário não está no grupo docker.

**Solução:**
```bash
sudo usermod -aG docker $USER
# Fazer logout/login ou executar:
newgrp docker
```

### Go: "go.mod" not found

**Problema:** Executando comando na pasta errada.

**Solução:**
```bash
cd backend  # Certifique-se de estar na pasta backend
go mod download
```

## 📚 Links Úteis

- **Go:** https://go.dev/doc/install
- **Node.js:** https://nodejs.org/en/download
- **NVM:** https://github.com/nvm-sh/nvm
- **Docker:** https://docs.docker.com/get-docker/
- **Docker Compose:** https://docs.docker.com/compose/install/

## 🎯 Versões Testadas

Este projeto foi testado com:
- Go 1.21.x
- Node.js 18.x e 20.x
- npm 9.x e 10.x
- Docker 24.x
- Docker Compose 2.x

---

**Pronto para começar! 🚀**

