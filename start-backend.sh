#!/bin/bash

echo "🚀 Starting Chat App Backend..."
echo ""

# Verificar se Go está instalado
if ! command -v go &> /dev/null
then
    echo "❌ Go não está instalado!"
    echo ""
    echo "Por favor, instale Go usando um dos seguintes comandos:"
    echo "  sudo snap install go         # version 1.25.4, or"
    echo "  sudo apt install golang-go   # version 2:1.24~2"
    echo ""
    exit 1
fi

cd backend

# Baixar dependências
echo "📦 Baixando dependências..."
go mod download
go mod tidy

echo ""
echo "✅ Backend iniciando na porta 8080..."
echo "📝 Credenciais padrão: admin / admin123"
echo ""

# Iniciar servidor
go run cmd/server/main.go

