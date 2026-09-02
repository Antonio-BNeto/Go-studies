# 🐹 Go Learning Journey

![Go Version](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-em%20andamento-yellow)
![License](https://img.shields.io/badge/license-none-lightgrey)

Repositório dedicado ao estudo da linguagem **Go (Golang)** através da construção de pequenos projetos práticos semanais.

---

## 📁 Estrutura do Repositório

O repositório está organizado utilizando a estrutura padrão recomendada para múltiplos projetos Go:

| Diretório | Descrição |
|---|---|
| `cmd/` | Pontos de entrada (`main.go`) de cada mini-projeto semanal |
| `pkg/` | Módulos e pacotes auxiliares reutilizáveis |

---

## 🗺️ Cronograma de Projetos

| # | Semana | Projeto | Status |
|---|---|---|:---:|
| 01 | [`cmd/semana-01-cli-task`](cmd/semana-01-cli-task) | CLI Task Manager | ✅ |
| 02 | `cmd/semana-02-file-parser` | Log Parser CSV/JSON | ⬜ |
| 03 | `cmd/semana-03-interfaces` | Calculadora com Interfaces | ⬜ |
| 04 | `cmd/semana-04-http-api` | API REST Nativa | ⬜ |
| 05 | `cmd/semana-05-url-checker` | Verificador de URLs Concorrente | ⬜ |
| 06 | `cmd/semana-06-context-fetcher` | Fetcher com Context e Timeout | ⬜ |
| 07 | `cmd/semana-07-sql-crud` | API CRUD com SQL | ⬜ |
| 08 | `cmd/semana-08-unit-tests` | Testes Unitários | ⬜ |
| 09 | `cmd/semana-09-jwt-auth` | API com Autenticação JWT | ⬜ |
| 10 | `cmd/semana-10-worker-pool` | Worker Pool em Go | ⬜ |
| 11 | `cmd/semana-11-realtime-chat` | Chat com WebSockets / gRPC | ⬜ |
| 12 | `cmd/semana-12-dockerized-app` | Multi-Stage Docker Build | ⬜ |

---

## 🛠️ Como Executar

1. Clone o repositório:

   ```bash
   git clone https://github.com/Antonio-BNeto/Go-studies.git
   cd Go-studies
   ```

2. Execute o projeto da semana desejada (exemplo: Semana 1):

   ```bash
   go run ./cmd/semana-01-cli-task
   ```
