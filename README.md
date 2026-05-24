# Axion - Controle Financeiro Inteligente

AXION é uma plataforma financeira moderna construída com foco em simplicidade, performance e segurança.

## Stack Técnica

- **Backend:** Go (Fiber), PostgreSQL, Redis, JWT.
- **Frontend:** Kotlin, Compose Multiplatform, Koin, Ktor.
- **Arquitetura:** DDD, Clean Architecture, Monorepo.
- **Infra:** Docker, Docker Compose, Traefik.

## Estrutura do Projeto

- `backend/`: Código fonte do servidor API em Go.
- `frontend/`: Aplicativo mobile/desktop em Kotlin Compose.
- `infrastructure/`: Configurações de banco de dados, docker e monitoramento.
- `scripts/`: Scripts utilitários de automação.
- `docs/`: Documentação técnica das fases do projeto.

## Como Iniciar

1. Clone o repositório.
2. Copie o arquivo `.env.example` para `.env` e ajuste as variáveis.
3. Execute `make up` para subir o ambiente de desenvolvimento.
