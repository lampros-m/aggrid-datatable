BIN = ./bin
CMD_SERVERHTTP = ./cmd/serverhttp
OUT_SERVERHTTP = $(BIN)/serverhttp

.PHONY: build
build:
	CGO_ENABLED=0 go build -o "${OUT_SERVERHTTP}" "${CMD_SERVERHTTP}/"


.PHONY: run
run:
	./"${OUT_SERVERHTTP}" 