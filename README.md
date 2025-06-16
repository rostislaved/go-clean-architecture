<p align="center">
  <img src="./logo.svg" alt="logo" width="300"/>
</p>

## Go Clean Architecture Template

Шаблон для создания приложений в соответствии с принципами чистой архитектуры.

Ключевые моменты:
- Соответствие принципам чистой архитектуры
- Используются паттерны DDD
- Учтены лучшие практики написания кода на Go

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

## В разработке
### v0.0.3 (16.06.2025)
- Экземпляры БД теперь создаются вне репозиториев и передаются в них как аргументы (для pg)
- Заменён `sqlx` на `pgxpool`
- Заменён `kafka-go` на `franz-go`
- Заменён `go-clickhouse` на `clickhouse-go`
- Сущность `Book` заменена на несколько сущностей с более обобщёнными именами — `Entity*`
- `controllers` переименованы в `handlers` (так как "controllers" — термин из MVC)
- `provider` переименован в `gateway` (используется для HTTP-запросов к другим микросервисам)
- `libs` переименованы в `pkg` (`pkg` — более распространённое название)
- Обновлена конфигурация линтеров (1.64.8)
- `graceful` теперь оформлен как библиотека
- Различные исправления и улучшения
- Добавлен первоначальный логотип

### v0.0.2 (27.10.2024)
- Добавлена библиотека `graceful` для корректного завершения работы
- `http-adapter` готов
- Рефакторинг `main`: разделена инициализация и выполнение
- Конфиги теперь являются частью соответствующих пакетов
- Унифицированы имена пакетов (все в `snake_case`)
- Общий рефакторинг: переименование пакетов, переменных, функций и т.д.

### v0.0.1
- Начальная версия

---

### TODO
- [x] Удалить `config` из domain
- [x] Пересмотреть структуру `http-adapter`
- [x] Реализовать корректное завершение
- [ ] Подумать над веткой с внедрением зависимостей (DI)
- [x] Разделить слой адаптеров и слой инфраструктуры
- [x] Подумать о `fatal` в конструкторах адаптеров
- [x] Проверить, как передаётся `context` в адаптерах
- [ ] Подумать о замене (или добавлении) `Config.toml` на `yaml` или `hcl`
- [ ] Добавить менеджер транзакций?
- [x] Удалить `books`
- [x] Проверить адаптеры
- [x] Логотип
- [x] Перенести `infra` в `pkg`
- [ ] Привести проект к компилируемому состоянию

---

### Примечания
- **Opinionated**: в названиях пакетов используется `snake_case`
- **Opinionated**: структура и конструктор находятся в `init.go`, а методы — в `methods.go`



# FAQ
#### Why package names snake_case
#### Why pkg inside internal
#### It is a template of a project layout, not a example project
