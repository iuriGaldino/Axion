# Guia de Onboarding AXION

Seja bem-vindo ao time!

## Requisitos de Ambiente
- Docker & Docker Compose
- Go 1.22+
- JDK 17+
- Android Studio ou IntelliJ IDEA

## Setup Inicial
1. Clone o repositório.
2. Execute `./scripts/setup.sh` para configurar o `.env`.
3. Execute `make up` para subir o banco de dados e Redis.
4. Para rodar o backend localmente: `cd backend && go run cmd/api/main.go`.
5. Para rodar o frontend: Abra o módulo `frontend` no IntelliJ e execute a tarefa Gradle `:composeApp:run`.

## Fluxo de Desenvolvimento
- Trabalhe sempre em branches nomeadas `feature/nome-da-task`.
- Escreva testes unitários para toda lógica de negócio em `internal/application/usecase`.
