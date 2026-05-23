# ADR 0001: Arquitetura Inicial do Projeto AXION

## Status
Aceito

## Contexto
O projeto AXION necessita de uma estrutura escalável, de fácil manutenção e que suporte múltiplos clientes (Android, Desktop, Web no futuro).

## Decisão
Adotamos o padrão **Monorepo** com as seguintes camadas:
- **Backend:** Go utilizando Fiber, seguindo Clean Architecture e DDD.
- **Frontend:** Kotlin Compose Multiplatform seguindo MVVM.

## Consequências
- **Positivas:** Código de domínio isolado de infraestrutura, facilidade em compartilhar tipos de dados, deploy orquestrado.
- **Negativas:** Curva de aprendizado inicial mais alta para desenvolvedores juniores.
