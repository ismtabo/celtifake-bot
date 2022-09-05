# Phrases of the year telegram bot


This repository contains a Telegram bot to handle phrases of the year.

Please modify `config.yml` with proper data. Or set sensible configuration into environment variables or `.env` file.

## Usage

To execute:

```
make build
make run
```

To clean:

```
make clean
```

To interact with dev environment using docker-compose:
```
make develenv-up   # docker-compose up whole environment
make develenv-down # docker-compose down whole environment
make develenv-sh   # logs bash session inside develenv container
```

To run test:
```
make develenv-up
make develenv-sh
make test-acceptance # (once inside develenv)
```

# Configuration

The following code shows the yaml configuration needed by the application:

```yaml
telegram:
  token: TOKEN
log:
  level: DEBUG
database:
  host: localhost
  port: 5432
  user:
  password:
```

This configuration can also be achieved by environment variables:

| **Environment Variable** | **Description**                     |
| :----------------------- | :---------------------------------- |
| `TELEGRAM_TOKEN`         | Telegram bot token                  |
| `LOG_LEVEL`              | Log level                           |
| `DATABASE_HOST`          | Postgres host                       |
| `DATABASE_PORT`          | Postgres port                       |
| `DATABASE_USER`          | Postgres user                       |
| `DATABASE_PASS`          | Postgres password *(sensible data)* |
| `DATABASE_NAME`          | Postgres Database                   |

## Authors

- Ismael Taboada Rodero: [@ismtabo](https://github.com/ismtabo)