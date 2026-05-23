# Especificação da API - AXION

## Base URL: `/api/v1`

### Auth
- `POST /auth/register`: Registro de usuário.
- `POST /auth/login`: Autenticação.

### Finance
- `GET /transactions`: Lista transações do usuário.
- `POST /transactions`: Nova transação.
- `GET /categories`: Lista categorias.

### AI & OCR
- `GET /insights`: Gera insights com IA.
- `POST /ocr`: Processa comprovante.
