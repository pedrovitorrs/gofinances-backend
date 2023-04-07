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

#### Estrutura de Diretorio

```shell
├── cmd                           # Diretório para os arquivos principais da aplicação
├── config.json                   # Arquivo de configuração da aplicação
├── docker-compose.yml            # Arquivo de configuração do Docker Compose
├── docs                          # Documentação da aplicação
├── go.mod                        # Arquivo de definição de dependências do Go
├── go.sum                        # Arquivo de somas de verificação de dependências do Go
├── internal                      # Diretório principal para os arquivos internos da aplicação
│   └── api
│       └── v1
│           ├── dto               # Data Transfer Objects
│           ├── handlers          # HTTP handlers para a API
│           ├── helpers           # Funções auxiliares
│           ├── repository        # Repositórios da aplicação
│           │   ├── sqlc          # Arquivos gerados pelo SQLC
│           │   └── test          # Testes para os repositórios
│           └── usecase           # Casos de uso da aplicação
├── Makefile                      # Arquivo Make para automatizar as tarefas comuns
├── pkg                           # Diretório para os arquivos de pacotes compartilhados
│   ├── config                    # Configurações gerais da aplicação
│   ├── database                  # Pacotes para manipulação de bancos de dados
│   │   ├── migrations            # Diretório para as migrações do banco de dados
│   │   └── queries               # Arquivos de consulta SQL
│   └── web                       # Pacotes para manipulação de HTTP
├── README.md                     # Arquivo com informações sobre a aplicação
└── sqlc.yaml                     # Arquivo de configuração do SQLC
```

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
