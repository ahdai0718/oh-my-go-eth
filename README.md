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

  - Environment variables
    ```
    // Ethereum node
    // (default Ethereum node via docker at localhost)
    // if you have any other node
    // you can replace it

    ETH_SERVICE_ETH_DATA_SEED_URL=http://localhost:8545
    ```

    Then

    ```
    // environment variables for docker

    cp env-example .env
    ```

  - Run
    ```
    // up services
    docker compose up eth-scanner eth-service
    ```



## What you got

- Ethereum node

  - For scanning ethererum data

  - Default RPC endpoint `http://localhost:8545`

- Redis Server

  - For service data cache

  - Recommend [AnotherRedisDesktopManager](https://github.com/qishibo/AnotherRedisDesktopManager) to check redis data

- MySQL Server

  - Store ETH block、transaction、log...etc

  - Schema

    - [eth_service.sql](https://github.com/ahdai0718/oh-my-go-eth/blob/master/deployments/mysql/sql/eth_service.sql)

- ETH block scanner

  - scan N (default:20) ETH blocks per second

- API service

  - `/api/v1/eth/blocks?limit=n` - get (N) blocks

  - `/api/v1/eth/blocks/{id}` - get single block with specific id

  - `/api/v1/eth/transaction/{tx_hash}` - get single transaction with tx hash

- Swagger APIs Explorer

  - Default url [http://localhost:40001/swagger/index.html](http://localhost:40001/swagger/index.html)

## Todos

- performance tuning

- handle fork