# API Chat - GoLang

Este projeto é uma API simples em GoLang, com apenas um endpoint `POST /api/chat`.

## ✨ Funcionalidades

- Endpoint `POST /api/chat`
  - Body esperado:
    ```json
    {
      "content": "sua mensagem aqui"
    }
    ```

## 🚀 Pré-requisitos

- GoLang **versão 1.23** instalado
- `make` instalado na sua máquina (ou navegar ao diretório cmd/api e rodar go run main.go)
- Chave de API da Groq (https://groq.com/)

## ⚙️ Configuração

1. Clone o repositório.
2. Preencha o arquivo `.env` no diretório `cmd/api/` com a sua chave da Groq:

   **cmd/api/.env**

## 🛠️ Scripts disponíveis

Utilizando o `Makefile`, você pode rodar:

- **Desenvolvimento**

```bash
make api-dev 
ou 
go run main.go
```
- **Acessar**
http://localhost:9000/
http://localhost:9000/api/chat
