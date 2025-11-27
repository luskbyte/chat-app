#!/bin/bash

echo "🚀 Starting Admin Frontend (Administrador)..."
echo ""

cd admin-frontend

# Verificar se node_modules existe
if [ ! -d "node_modules" ]; then
    echo "📦 Instalando dependências..."
    npm install
fi

echo ""
echo "✅ Admin Frontend iniciando na porta 3001..."
echo "🌐 Acesse: http://localhost:3001"
echo ""

# Iniciar servidor de desenvolvimento
npm run dev -- --port 3001

