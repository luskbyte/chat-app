#!/bin/bash

echo "🔧 Corrigindo DNS do Docker..."
echo ""

# Criar backup se o arquivo já existir
if [ -f /etc/docker/daemon.json ]; then
    echo "📋 Fazendo backup do daemon.json existente..."
    sudo cp /etc/docker/daemon.json /etc/docker/daemon.json.backup
fi

# Criar diretório se não existir
sudo mkdir -p /etc/docker

# Configurar DNS públicos
echo "📝 Configurando DNS públicos (Google e Cloudflare)..."
sudo bash -c 'cat > /etc/docker/daemon.json << "DOCKEREOF"
{
  "dns": ["8.8.8.8", "8.8.4.4", "1.1.1.1"]
}
DOCKEREOF'

echo ""
echo "✅ Configuração criada:"
cat /etc/docker/daemon.json
echo ""

# Reiniciar Docker
echo "🔄 Reiniciando Docker..."
sudo systemctl restart docker

# Aguardar o Docker iniciar
sleep 3

# Verificar status
echo ""
echo "🔍 Verificando status do Docker..."
if sudo systemctl is-active --quiet docker; then
    echo "✅ Docker está rodando"
else
    echo "❌ Docker não está rodando"
    exit 1
fi

echo ""
echo "🧪 Testando pull de imagem..."
if docker pull alpine:latest; then
    echo ""
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "✅ DNS CORRIGIDO COM SUCESSO!"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo ""
    echo "Agora você pode executar:"
    echo "  sudo docker compose up --build"
else
    echo ""
    echo "❌ Ainda há problemas. Verifique:"
    echo "  1. Conexão com internet"
    echo "  2. Firewall"
    echo "  3. Proxy (se aplicável)"
fi

