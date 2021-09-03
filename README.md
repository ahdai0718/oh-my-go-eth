# oh-my-go-eth

My first Go ETH service.

## Refercences

- Go

  - [https://github.com/golang/go/wiki/Modules#how-to-use-modules](https://github.com/golang/go/wiki/Modules#how-to-use-modules)

  - [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

- ETH

    - [https://eth.wiki/](https://eth.wiki/)

    - [https://github.com/ethereum/go-ethereum](https://github.com/ethereum/go-ethereum)

## How it works

- Via `Docker`、`docker-compose`

  - [Install Docker](https://docs.docker.com/get-started/)

  - [Install docker compose](https://docs.docker.com/compose/install/)

  - Run
    ```
    // environment variables for docker
    cp env-example .env

    // up service
    docker compose up eth-service
    ```

## What you got

- Ethereum node

  - RPC endpoint `http://localhost:8545`

- Swagger APIs Explorer With 3 APIs

  - Default url [http://localhost:40001/swagger/index.html](http://localhost:40001/swagger/index.html)

- ETH block scanner on boot

  - [scanner.go](https://github.com/ahdai0718/oh-my-go-eth/blob/master/internal/app/server/eth/scanner.go)

  - scan N (default:20) ETH blocks per second

- MySQL schema

  - [eth_service.sql](https://github.com/ahdai0718/oh-my-go-eth/blob/master/deploy/mysql/sql/eth_service.sql)


## Todos

- performance tuning