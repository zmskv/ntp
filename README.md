# **ntp**

CLI-программа на Go для получения точного времени с NTP-сервера.

---

## **Функции**

* Получение точного времени через [NTP (Network Time Protocol)](https://en.wikipedia.org/wiki/Network_Time_Protocol)
* Вывод времени в формате **RFC1123**
* Настраиваемый NTP-сервер (по умолчанию `pool.ntp.org`)
* Разделение стандартного вывода (`STDOUT`) и ошибок (`STDERR`)

---

## **Использование**

### Получение времени с сервера по умолчанию:

```bash
go run ./cmd/ntp/main.go
```

Пример вывода:

```
Sun, 24 Aug 2025 20:00:00 UTC
```

### Указание другого сервера:

```bash
go run cmd/ntp/main.go time.google.com
```

## **Проверка качества кода**

Проверка с `vet`:

```bash
go vet ./...
```

Проверка с `golint` (если установлен):

```bash
golint ./...
```

---

## **Структура проекта**

```
ntp/
├── go.mod
├── cmd/
│   └── ntp/
│       └── main.go          
└── internal/
    └── fetcher/
        └── fetch.go         
```

---

## **Примеры публичных NTP-серверов**

| Сервер                | Описание                |
| --------------------- | ----------------------- |
| `pool.ntp.org`        | Пул глобальных серверов |
| `time.google.com`     | Серверы Google NTP      |
| `time.cloudflare.com` | Серверы Cloudflare      |

---
