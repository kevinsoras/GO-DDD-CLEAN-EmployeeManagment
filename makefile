SHELL := /bin/bash

DB_URL := $(shell grep DB_URL .env | cut -d '=' -f2-)

# Por defecto employee
CONTEXT ?= employee

ifeq ($(CONTEXT),shared)
  MIGRATIONS_DIR = shared/infrastructure/persistence/migrations
else
  MIGRATIONS_DIR = contexts/$(CONTEXT)/infrastructure/persistence/migrations
endif

migrate-new:
	@if [ -z "$(name)" ]; then \
		echo "❌ Error: necesitas pasar el nombre (ej: make migrate-new name=create_employees)"; \
		exit 1; \
	fi; \
	@if [ -z "$(CONTEXT)" ] || ([ "$(CONTEXT)" != "employee" ] && [ "$(CONTEXT)" != "shared" ]); then \
		echo "❌ Error: necesitas especificar un CONTEXT válido (employee o shared) (ej: make migrate-new name=add_field CONTEXT=employee)"; \
		exit 1; \
	fi; \
	@echo "Creating migration '$(name)' for context '$(CONTEXT)' in $(MIGRATIONS_DIR)"; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) $(name);


MIGRATIONS_DIRS = shared/infrastructure/persistence/migrations \
                  contexts/*/infrastructure/persistence/migrations

migrate-up:
	@for dir in $(MIGRATIONS_DIRS); do \
		if [ -d "$$dir" ]; then \
			echo "▶️ Ejecutando migraciones en $$dir..."; \
			echo "migrate -path $$dir -database \"$(DB_URL)\" up"; \
			migrate -path $$dir -database "$(DB_URL)" up; \
		else \
			echo "⚠️  Saltando $$dir (no existe)"; \
		fi \
		;\
		done



migrate-down:
	@bash -c '\
	shopt -s nullglob; \
	all_files=(); \
	for dir in $(MIGRATIONS_DIRS); do \
		if [ -d "$$dir" ]; then \
			for file in $$dir/*.up.sql; do \
				all_files+=(\"$$file\"); \
				;\
				done; \
		fi; \
		done; \
	if [ $${#all_files[@]} -eq 0 ]; then \
		echo "⚠️  No se encontraron migraciones para revertir"; \
		exit 0; \
	fi; \
	latest_file=$$(printf "%s\n" "$${all_files[@]}" | sort | tail -n1); \
	latest_dir=$$(dirname "$$latest_file"); \
	echo "⏪ Revirtiendo última migración: $$latest_file en $$latest_dir..."; \
	migrate -path "$$latest_dir" -database "$(DB_URL)" down 1; \
	'

swagger-docs:
	@echo "Generating Swagger documentation..."
	@swag init -dir ./cmd -g main.go
