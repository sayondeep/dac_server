.PHONY: dac-db dac-server all create-network

# Define paths to each service
SERVICE1_DIR=./db
SERVICE2_DIR=./server

# Define the shared network name
NETWORK_NAME=server_default

# Target to create the network if it doesn't exist
create-network:
	@if [ -z "$$(docker network ls --filter name=^$(NETWORK_NAME)$$ --format '{{.Name}}')" ]; then \
		docker network create $(NETWORK_NAME); \
		echo "Network $(NETWORK_NAME) created."; \
	else \
		echo "Network $(NETWORK_NAME) already exists."; \
	fi

# Targets to start each service (ensure the network is created first)
dac-db: create-network
	cd $(SERVICE1_DIR) && docker compose up -d

dac-server: create-network
	cd $(SERVICE2_DIR) && docker compose up -d

# Combined target to start all services
all: dac-db dac-server
