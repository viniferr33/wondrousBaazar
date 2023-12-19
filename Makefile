include .env

clear:
	@if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
	fi
	@echo "Deleted binary ${BINARY}";

compile:
	@echo "Building binary ... ${BINARY}";
	go build -o ${BINARY} cmd/wondrous_baazar.go

run: compile
	./${BINARY}

docker_build:
	docker build . -t "${IMAGE_NAME}"

compose_up: docker_build
	docker compose up -d

compose_down:
	docker compose down