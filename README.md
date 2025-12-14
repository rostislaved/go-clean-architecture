<p align="center">
  <img src="./logo.svg" alt="logo" width="300"/>
</p>

## Go Clean Architecture Template

Шаблон для создания приложений в соответствии с принципами чистой архитектуры.

Ключевые моменты:
- Соответствие принципам чистой архитектуры
- Используются паттерны DDD
- Учтены лучшие практики написания кода на Go
- Все названия выбраны так, чтобы, максимально соответствовать терминам, используемым в литературе, но не в ущерб корректности

Это шаблон, а не пример реального приложения. Задача шаблона быть максимально общим и показывать подходы, а не реализовывать конкретную бизнес-логику, которая будет отличаться для каждого реального приложения.

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



### Примечания
- **Opinionated**: в названиях пакетов используется `snake_case`
- **Opinionated**: структура и конструктор находятся в `init.go`, а методы — в `methods.go`



# FAQ
#### Why package names snake_case
#### Why pkg inside internal
#### It is a template of a project layout, not a example project
