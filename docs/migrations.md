# Migrations

We use the [`mattes/migrate`](https://github.com/mattes/migrate) tool to handle
migrations for Keep Up. You can find installation instructions for your device
from the link for that project. Below are some helpful commands when working
with this tool locally.

- Creating a migration

```bash
migrate create -ext sql -dir ./graphql/internal/database/migrations <migration-name>
```

- Running migrations

```bash
migrate -path ./graphql/internal/database/migrations/ -database postgres://postgres:postgres@0.0.0.0:5432/keep_up_dev?sslmode=disable up
```

- Rolling back migrations

```bash
migrate -path ./graphql/internal/database/migrations/ -database postgres://postgres:postgres@0.0.0.0:5432/keep_up_dev?sslmode=disable down
# Or, to remove completely
migrate -path ./graphql/internal/database/migrations/ -database postgres://postgres:postgres@0.0.0.0:5432/keep_up_dev?sslmode=disable drop
```
