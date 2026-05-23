.PHONY: help up down build logs backend-shell db-shell test

help:
	@echo "Comandos disponíveis:"
	@echo "  make up             - Sobe os containers (docker-compose up -d)"
	@echo "  make down           - Derruba os containers"
	@echo "  make build          - Builda as imagens"
	@echo "  make logs           - Mostra os logs"
	@echo "  make backend-shell  - Acessa o shell do backend"
	@echo "  make db-shell       - Acessa o shell do banco de dados"
	@echo "  make test           - Executa os testes do backend"

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build

logs:
	docker-compose logs -f

backend-shell:
	docker exec -it axion-backend sh

db-shell:
	docker exec -it axion-db psql -U axion -d axion_db

test:
	cd backend && go test ./...
