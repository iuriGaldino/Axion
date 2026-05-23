# AXION - FASE 4: NÚCLEO DE NEGÓCIO

## OBJETIVO
Implementar o núcleo de negócio do AXION seguindo DDD e Clean Architecture.

## PRIORIDADE
1. Auth
2. Users
3. Categories
4. Transactions
5. Budgets
6. Goals

## ESPECIFICAÇÕES POR MÓDULO

### AUTH
- User Entity
- UserRepository
- RegisterUserUseCase
- LoginUseCase
- RefreshTokenUseCase
- ChangePasswordUseCase
- ForgotPasswordUseCase
- JWT Service
- Password Hash Service (Argon2id)
- Auth Middleware

### USERS
- User Profile
- User Preferences: Language, Theme, Accent Color, Avatar

### CATEGORIES
- Category Entity
- Category Repository
- CRUD completo
- Categorias padrão do sistema e personalizadas

### TRANSACTIONS
- Tipos: Income, Expense, Transfer
- Entity, Repository
- UseCases: Create, Update, Delete, List

### BUDGETS
- Entity, Repository
- UseCases: Create, Monitoring, Alerts

### GOALS
- Entity, Repository
- UseCases: Create, Progress Calculation

## PADRÕES OBRIGATÓRIOS
Para cada módulo:
- Entidade e Value Objects
- Repository (Interface e Implementation)
- Use Cases e DTOs
- HTTP Handlers e Presenters
- Validation e Unit Tests (Mínimo 80% cobertura)

## TECNOLOGIAS CONFIRMADAS
- Password Hashing: Argon2id
- Token: JWT
- Framework: Fiber (Go)
- Banco: PostgreSQL
