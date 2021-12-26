TAG := keyspecs/disussion-bot:latest

build:
	@echo "Building discussion-bot docker image"
	@docker build -f deploy/docker/Dockerfile -t $(TAG) .

start:
	@echo "Starting up discussion-bot..."
	@sh ./deploy/scripts/up.sh
	
		
stop:
	@echo "Stopping Backend..."
	@sh ./deploy/scripts/down.sh

	
.PHONY: build start stop