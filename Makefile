include .env
export

all: build

# Build the web server
build-web:
	@echo "Building..."
	@templ generate
	@npx tailwindcss -i cmd/ui/assets/css/input.css -o cmd/ui/assets/css/output.css
	@go build -o main cmd/web/main.go

# Run the web
run-web:
	@go run cmd/web/main.go

run: build webhook-proxy run-api run-web run-bot

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

# db query generate
generate:
	@sqlc generate

# db create migration, expects NAME to be set
migrate-create:
	@cd migrations && goose create $(NAME) sql

# db up migration
migrate-up:
	@cd migrations && goose up

# db down migration
migrate-down:
	@cd migrations && goose down

.PHONY: all build run test clean watch generate migrate-create migrate-up migrate-down
