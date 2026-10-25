# include .env

# export

# MIGRATION_PATH=./migrations
# DATABASE_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)

include .env
export

MIGRATION_PATH=./migrations

DATABASE_URL=postgres://$(DB_USER):$(DB_PASS)@$(strip $(DB_HOST)):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)


# ============================================================
# MIGRATION UP
# ============================================================

migrate-up:
	migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" up


# Apply ONLY ONE pending migration
migrate-up-one:
	migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" up 1


# ============================================================
# MIGRATION DOWN
# ============================================================

# Rollback ONE migration
migrate-down-one:
	migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" down 1


# Rollback ALL migrations
migrate-down:
	migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" down


# ============================================================
# MIGRATION GOTO
# ============================================================

migrate-goto:
	migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" goto $(version)


# ============================================================
# FORCE MIGRATION VERSION
# ============================================================

migrate-force:
	migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" force $(version)


# ============================================================
# SHOW CURRENT VERSION
# ============================================================

migrate-version:
	migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" version


# ============================================================
# CREATE MIGRATION
# ============================================================

create-migration:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)


# ============================================================
# SWAGGER
# ============================================================

swagger:
	swag init -g cmd/server/main.go --parseDependency --parseInternal


# ============================================================
# DEBUG DATABASE VARIABLES
# ============================================================

show-db:
	@echo "DB_HOST=[$(DB_HOST)]"
	@echo "DB_PORT=[$(DB_PORT)]"
	@echo "DB_USER=[$(DB_USER)]"
	@echo "DB_NAME=[$(DB_NAME)]"
	@echo "DB_SSL=[$(DB_SSL)]"
	@echo "DATABASE_URL=[$(DATABASE_URL)]"