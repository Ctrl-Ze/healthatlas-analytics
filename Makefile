# Makefile for HealthAtlas Themis (Analytics)
# Usage:
#   make db-up            → start only Themis DB
#   make themis-up        → start Themis + DB
#   make down             → stop and remove all
#   make logs             → tail logs
#   make db               → psql into DB
#   make clean            → remove containers + volumes

DOCKER_COMPOSE := docker-compose -f docker-compose.yml

# Start only the TimescaleDB instance
db-up:
	@echo "💾 Starting Themis DB (Timescale)..."
	$(DOCKER_COMPOSE) up -d timescaledb

# Start Themis service + DB
themis-up:
	@echo "📊 Starting Themis (Analytics) and DB..."
	$(DOCKER_COMPOSE) up -d themis timescaledb

# Stop all containers
down:
	@echo "🛑 Stopping Themis containers..."
	$(DOCKER_COMPOSE) down

# Connect to TimescaleDB
db:
	@echo "🗄️ Connecting to Themis DB..."
	docker exec -it themis-timescale psql -U postgres -d themis

# Tail logs
logs:
	@echo "📜 Tailing Themis logs..."
	$(DOCKER_COMPOSE) logs -f

# Cleanup containers + volumes
clean:
	@echo "🧹 Removing Themis containers and volumes..."
	$(DOCKER_COMPOSE) down -v

# Start optional pgAdmin
pgadmin:
	@echo "🖥️ Starting pgAdmin..."
	$(DOCKER_COMPOSE) up -d pgadmin

pgadmin-logs:
	@echo "📜 Tailing pgAdmin logs..."
	$(DOCKER_COMPOSE) logs -f pgadmin
