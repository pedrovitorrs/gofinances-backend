# goFinances - Golang API

Esta é uma simples API para projeto financeiro utilizando Golang, Postgres e JWT.

## Instalação

Para instalar as dependências clone o projeto e execute o seguinte comando dentro da pasta.

```
make server
```

### Pré-requisitos

Para executar este container deverá ter o Docker instalado.

* [Windows](https://docs.docker.com/windows/started)
* [OS X](https://docs.docker.com/mac/started/)
* [Linux](https://docs.docker.com/linux/started/)

O projeto faz uso de banco de dados Postgres, caso o banco estiver local deverá estar instalado.

* [Postgres](https://www.postgresql.org/download/)

## Usage

#### Documentação

Gerar a documentação da API para ambiente de desenvolvimento:

```shell
make doc-gen
```

Acesse a documentação do navegador:

```shell
http://localhost:<API_PORT>/swagger/index.html
```

### Estrutura de Diretorio

.
├── cmd
│   └── server
│       └── main.go
├── config.json
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   └── api
│       └── v1
│           ├── dto
│           │   ├── request
│           │   │   └── createuserrequest.go
│           │   └── response
│           │       └── responseerror.go
│           ├── handlers
│           │   └── user_handler.go
│           ├── helpers
│           │   └── random.go
│           ├── repository
│           │   ├── sqlc
│           │   │   ├── account.sql.go
│           │   │   ├── category.sql.go
│           │   │   ├── db.go
│           │   │   ├── models.go
│           │   │   ├── querier.go
│           │   │   ├── store.go
│           │   │   └── user.sql.go
│           │   └── test
│           │       ├── account_test.go
│           │       ├── category_test.go
│           │       ├── main_test.go
│           │       └── user_test.go
│           └── usecase
│               └── user_usecase.go
├── Makefile
├── pkg
│   ├── config
│   │   └── config.go
│   ├── database
│   │   ├── migrations
│   │   │   ├── 000001_create_initials_table.down.sql
│   │   │   └── 000001_create_initials_table.up.sql
│   │   └── queries
│   │       ├── account.sql
│   │       ├── category.sql
│   │       └── user.sql
│   └── web
│       └── middlewares
│           └── HttpMiddleware.go
├── README.md
└── sqlc.yaml

#### Compose Project

O projeto já possui um docker-compose configurado, conseguirá executar o projeto somente executando o código:

```shell
docker-compose up -d
```

#### Environment Variables

* `DB_DRIVER` - Protocolo de comunicação com o banco de dados. ex: "postgres", "mysql"
* `DB_SOURCE` - URI de conexão direta ao banco de dados.
* `SERVER_ADDRESS` - Endereço da máquina vinculada a aplicação.

## Authors

* **Pedro Rodrigues** - *Work* - [pedrovitorrs](https://github.com/pedrovitorrs)

## Como contribuir

Esteja sempre atento à criação de novas branches, padronização de commits e comentários em código para que possamos melhorar sua mantenabilidade.
