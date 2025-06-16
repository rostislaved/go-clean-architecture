## Work in progress
<p style="text-align: center;">
  <img src="./logo.svg" alt="logo" width="300"/>
</p>

v0.0.3
* Db instances now created outside of repositories and then passed into them as argument (for all db types)
* Replace sqlx with Pgxpool
* Replace kafka-go with franz-go
* Replace go-clickhouse with clickhouse-go
* Book entity replaced with several entities with more generic names - "Entity*"
* Rename controllers->handlers (controllers is a mvc term)
* Rename provider->gateway (where you make https requests to other microservices)
* Rename libs -> pkg. pkg is more widespread name
* Update linters config (1.64.8)
* Graceful is a lib now
* Various fixes and refinements
* Added logo

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
5. Separate interface adapters layer and infrastructure layer (+)
6. Think about fatals in adapters constructors (+)
7. Validate how context is propagated in adapters (+)
8. Consider changing (or adding) Config.toml to yaml or hcl
9. Add Transaction Manager?
10. get rid of "books" (+)
11. Check adapters (+)
12. Logo (+)
13. Infra to pkg (+)
14. make compile

Notes:
1. Opinionated: snake_case in package names 
2. Opinionated: I put the struct and constructor in init.go and the methods in methods.go


# Project layout
```
> tree -d
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
    │   │       │   └── entity5-gateway
    │   │       ├── grpc-adapter
    │   │       │   └── generated
    │   │       ├── kafka-adapter-publisher
    │   │       ├── kafka-adapter-publisher2
    │   │       │   └── kafka-client
    │   │       ├── nats-adapter-publisher
    │   │       └── repositories
    │   │           ├── entity1-repository
    │   │           ├── entity2-repository
    │   │           ├── entity3-repository
    │   │           └── entity4-repository
    │   ├── application
    │   │   └── usecases
    │   ├── config
    │   └── domain
    │       ├── entity1
    │       ├── entity2
    │       ├── entity3
    │       ├── entity4
    │       └── entity5
    └── pkg
        ├── clickhouse
        ├── helpers
        ├── http-server
        ├── middleware-helpers
        ├── mongo
        ├── postgres
        └── provider-helpers



```

# FAQ
#### Why package names snake_case
#### Why pkg inside internal
#### It is a template of a project layout, not a example project
