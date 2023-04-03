# goFinances - Golang API

Esta é uma simples API para projeto financeiro utilizando Golang, Postgres e JWT.

## Instalação

Para instalar as dependências clone o projeto e execute o seguinte comando dentro da pasta.

```shell
yarn
```

logo após execute o sistema:

```
yarn start:dev
```

### Pré-requisitos

Para executar este container deverá ter o Docker instalado.

* [Windows](https://docs.docker.com/windows/started)
* [OS X](https://docs.docker.com/mac/started/)
* [Linux](https://docs.docker.com/linux/started/)

O projeto faz uso de banco de dados Postgres, caso o banco estiver local deverá estar instalado.

* [Postgres](https://www.postgresql.org/download/)

### Usage

#### Documentation

Generate API documentation for development environment

Gerar a documentação da API

```shell
// or npm run start:docs
yarn start:docs
```

Access documentation for browser

```shell
http://localhost:<API_PORT>/api/docs
```

#### Compose Project

O projeto já possui um docker-compose configurado, conseguirá executar o projeto somente executando o código:

```shell
docker-compose up
```

#### Environment Variables

* `NODE_ENV` - Ambiente de desenvolvimento.
* `DATABASE_NAME` - Nome da base de dados.
* `DATABASE_USERNAME` - Usuario de acesso ao banco de dados.
* `DATABASE_PASSWORD` - Senha de acesso ao banco de dados.
* `DATABASE_HOST` - Endereço da máquina vinculada ao banco de dados.
* `DATABASE_DIALECT` - Protocolo de comunicação com o banco de dados. ex: "postgres", "mysql"
* `APP_PORT` - Porta da aplicação.
* `JWT_SECRET` - Chave para criptografia.
* `TOKENTIME` - Tempo de expiração do token de acesso.
* `DATABASE_CONNECTION_STRING` - URI de conexão direta com o banco de dados
* `MONGO_URI` - URI de conexão direta ao MongoDB
* `REDIS_HOST` - Endereço da máquina vinculada ao banco de dados.
* `REDIS_PORT` - Porta habilitada da máquina vinculada ao banco de dados.

## Authors

* **Pedro Rodrigues** - *Work* - [pedrovitorrs](https://github.com/pedrovitorrs)

## Como contribuir

Esteja sempre atento à criação de novas branches, padronização de commits e comentários em código
para que possamos melhorar sua mantenabilidade.


## instalar go-Migrate
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

mv migrate $GOPATH/bin/migrate
