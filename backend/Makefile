.PHONY: run build clean test

# Executar o servidor
run:
	go run cmd/server/main.go

# Compilar o servidor
build:
	go build -o chat-app-backend cmd/server/main.go

# Limpar arquivos compilados
clean:
	rm -f chat-app-backend

# Baixar dependências
deps:
	go mod download
	go mod tidy

# Executar testes
test:
	go test -v ./...

# Build para produção (Linux)
build-linux:
	GOOS=linux GOARCH=amd64 go build -o chat-app-backend-linux cmd/server/main.go

# Build para produção (Windows)
build-windows:
	GOOS=windows GOARCH=amd64 go build -o chat-app-backend.exe cmd/server/main.go

# Build para produção (macOS)
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o chat-app-backend-mac cmd/server/main.go

