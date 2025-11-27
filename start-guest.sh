#!/bin/bash

echo "🚀 Starting Guest Frontend (Visitante)..."
echo ""

cd guest-frontend

# Verificar se node_modules existe
if [ ! -d "node_modules" ]; then
    echo "📦 Instalando dependências..."
    npm install
fi

echo ""
echo "✅ Guest Frontend iniciando na porta 3000..."
echo "🌐 Acesse: http://localhost:3000"
echo ""

# Iniciar servidor de desenvolvimento
npm run dev -- --port 3000

