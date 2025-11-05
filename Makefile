.PHONY: install run migrate seed

install:
	go mod download
	go mod tidy

run:
	go run main.go

migrate:
	psql -U postgres -d budgetwise -f database/migrations.sql

seed:
	@echo "Seeding admin user..."
	@psql -U postgres -d budgetwise -c "INSERT INTO users (email, password_hash, full_name, role) VALUES ('admin@budgetwise.com', '\$$2a\$$10\$$YtcSqB5h7rKYOhB5YjW8/.fKj0Z4HQlGLj6kZ.rX6YqJXJqYQJhfy', 'Administrator', 'admin') ON CONFLICT (email) DO NOTHING;"

build:
	go build -o bin/budgetwise main.go

clean:
	rm -rf bin/
