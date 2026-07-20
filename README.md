# 🐹 Go Learning Journey

Repositório dedicado ao estudo da linguagem **Go (Golang)** através da construção de pequenos projetos práticos semanais.

## 📁 Estrutura do Repositório

O repositório está organizado utilizando a estrutura padrão recomendada para múltiplos projetos Go:

- `cmd/`: Contém os pontos de entrada (`main.go`) de cada mini-projeto semanal.
- `pkg/`: Módulos e pacotes auxiliares reutilizáveis.

## 🗺️ Cronograma de Projetos

- [ ] **Semana 01**: CLI Task Manager (`cmd/semana-01-cli-task`)
- [ ] **Semana 02**: Log Parser CSV/JSON (`cmd/semana-02-file-parser`)
- [ ] **Semana 03**: Calculadora com Interfaces (`cmd/semana-03-interfaces`)
- [ ] **Semana 04**: API REST Nativa (`cmd/semana-04-http-api`)
- [ ] **Semana 05**: Verificador de URLs Concorrente (`cmd/semana-05-url-checker`)
- [ ] **Semana 06**: Fetcher com Context e Timeout (`cmd/semana-06-context-fetcher`)
- [ ] **Semana 07**: API CRUD com SQL (`cmd/semana-07-sql-crud`)
- [ ] **Semana 08**: Testes Unitários (`cmd/semana-08-unit-tests`)
- [ ] **Semana 09**: API com Autenticação JWT (`cmd/semana-09-jwt-auth`)
- [ ] **Semana 10**: Worker Pool em Go (`cmd/semana-10-worker-pool`)
- [ ] **Semana 11**: Chat com WebSockets / gRPC (`cmd/semana-11-realtime-chat`)
- [ ] **Semana 12**: Multi-Stage Docker Build (`cmd/semana-12-dockerized-app`)

## 🛠️ Como Executar

1. Clone o repositório:
   ```bash
   git clone [https://github.com/Antonio-BNeto/Go-studies.git](https://github.com/Antonio-BNeto/Go-studies.git)
   cd Go-studies
   ```
2. Execute o projeto da semana desejada (exemplo: Semana 1):
    ```bash
    go run ./cmd/semana-01-cli-task
    ```