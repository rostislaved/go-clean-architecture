## Work in progress

v0.0.2 (27.10.2024)
* added graceful lib for graceful shutdown
* http-adapter is ready
* main is refactored. Initialization and runtime parts are separated
* configs now are parts of their packages
* consistent package names (all snake_case)
* general refactoring. Refactored names of: packages, variables, functions etc.



v0.0.1
* initial version

#### TODO
1. Get rid of config in domain (+)
2. Rethink the structure of http-adapter (+)
3. Implement graceful shutdown (+)
4. Consider making a branch with DI
5. Separate interface adapters layer and infrastructure layer
6. Think about fatals in adapters constructors (+)
7. Validate how context is propagated in adapters
8. Consider changing (or adding) Config.toml to yaml or hcl


Notes:
1. Opinionated: snake_case in package names 
2. Opinionated: I put the struct and constructor in init.go and the methods in methods.go
3. adapters implementations are not production ready. They are just examples (except http-adapter)

# Project layout
```
.
├── cmd
│   └── service
└── internal
    ├── app
    │   ├── adapters
    │   │   ├── primary
    │   │   │   ├── grpc-adapter
    │   │   │   │   ├── generated
    │   │   │   │   └── handlers
    │   │   │   ├── http-adapter
    │   │   │   │   ├── handlers
    │   │   │   │   └── router
    │   │   │   ├── kafka-adapter-subscriber
    │   │   │   │   ├── kafka-handlers
    │   │   │   │   └── kafka-queue
    │   │   │   ├── nats-adapter-subscriber
    │   │   │   │   └── nats-handlers
    │   │   │   ├── os-signal-adapter
    │   │   │   └── pprof-adapter
    │   │   └── secondary
    │   │       ├── gateways
    │   │       │   └── books-gateway
    │   │       ├── grpc-adapter
    │   │       │   └── generated
    │   │       ├── kafka-adapter-publisher
    │   │       ├── nats-adapter-publisher
    │   │       └── repositories
    │   │           ├── books-repository-clickhouse
    │   │           ├── books-repository-mongo
    │   │           └── books-repository-postgres
    │   ├── application
    │   │   └── usecases
    │   ├── config
    │   └── domain
    │       └── book
    └── libs
        ├── graceful
        ├── helpers
        ├── http-server
        ├── middleware-helpers
        ├── provider-helpers
        └── repo-helpers


```
