APP_NAME=logm
DATABASE_URL?=postgres://logm:logm@localhost:5432/logm?sslmode=disable

# ---- Codegen ----

templ:
	go tool templ generate

sqlc:
	go tool sqlc generate

generate: templ sqlc

# ---- Migrations ----

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1

migrate-version:
	migrate -path migrations -database "$(DATABASE_URL)" version

# ---- Dev ----

run:
	podman compose up

build:
	podman compose up --build

down:
	podman compose down

app:
	podman compose exec -it app sh

# ---- Psql ----

psql:
	podman compose exec -it db psql logm logm -U logm

# ---- Clean ----

clean:
	rm -rf tmp
