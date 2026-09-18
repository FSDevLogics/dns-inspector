# CLI App - DNS Lookup

Uma aplicação de linha de comando (CLI) desenvolvida em Go para realizar consultas rápidas de endereços IP e servidores de nome (DNS) de domínios na internet.

## 🚀 Recursos

* **Consulta de IP (`ip`):** Retorna os endereços IP (IPv4/IPv6) associados a um domínio.
* **Consulta de Servidores (`server`):** Retorna os servidores de nome (Name Servers - NS) responsáveis por um domínio.
* **Dockerizado:** Imagem Docker otimizada com *multi-stage build* (super leve e segura).
* Construído com as melhores práticas da linguagem Go e utilizando o pacote [urfave/cli/v2](https://github.com/urfave/cli?utm_source=gemini).

## 🛠️ Pré-requisitos

* [Go 1.26.8](https://go.dev/dl/?utm_source=gemini) ou superior (para desenvolvimento e execução local).
* [Docker](https://www.docker.com/?utm_source=gemini) (para execução via container isolado).

## 📦 Instalação e Execução Local

1. Clone o repositório:
```bash
git clone https://github.com/FSDevLogics/ip-server-cli
cd cli_app

```


2. Baixe e sincronize as dependências:
```bash
go mod tidy

```


3. Execute o aplicativo passando os comandos desejados:
```bash
# Ver o painel de ajuda e comandos disponíveis
go run main.go --help

# Buscar IPs de um domínio (o padrão é parse3d.com.br se a flag --host for omitida)
go run main.go ip --host google.com

# Buscar os servidores de nome de um domínio
go run main.go server --host github.com

```


4. (Opcional) Compilar o binário para uso direto:
```bash
go build -o cli_app .
./cli_app ip --host apple.com

```



## 🐳 Execução via Docker

O projeto inclui um `Dockerfile` que utiliza o conceito de múltiplas etapas, resultando em uma imagem final extremamente leve.

1. Construa a imagem Docker localmente:
```bash
docker build -t cli-app .

```


2. Execute o container repassando os argumentos do seu CLI:
```bash
# Ajuda
docker run --rm cli-app --help

# Consulta de IPs
docker run --rm cli-app ip --host parse3d.com.br

# Consulta de Servidores
docker run --rm cli-app server --host parse3d.com.br

```


> **Nota:** A flag `--rm` é utilizada para garantir que o container seja removido automaticamente do seu sistema assim que a consulta for concluída, não ocupando espaço desnecessário.



## 📂 Estrutura do Projeto

```text
.
├── app/
│   └── app.go         # Lógica de configuração do CLI, comandos e funções de busca DNS
├── .dockerignore      # Arquivos ignorados durante a construção da imagem Docker
├── .gitignore         # Arquivos e diretórios ignorados pelo versionamento do Git
├── Dockerfile         # Receita do container (Go 1.26.8 Builder -> Alpine)
├── go.mod             # Definição do módulo e declaração de dependências
├── go.sum             # Hashes de integridade das dependências
└── main.go            # Ponto de entrada (Entrypoint) da aplicação

```