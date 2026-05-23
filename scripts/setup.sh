#!/bin/bash
echo "Iniciando setup do AXION..."

# Criar .env se não existir
if [ ! -f .env ]; then
    cp .env.example .env
    echo ".env criado a partir do .env.example"
fi

# Inicializar backend
echo "Configurando backend..."
cd backend && go mod tidy
cd ..

echo "Setup concluído!"
