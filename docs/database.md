# Database

> Documentation related to setup of the Postgres database for the project.

## Permissions

By default, we want to the user who is associated with the connection string in
our web app to have the minimum set of permissions needed. A common problem to
run into is one where a user is created but they get an error similar to:

```
Permission denied for relation <relation_name>
```

While an initial approach may be to `GRANT` privileges on the database itself,
instead we need to `GRANT` on tables directly doing something similar to the
following:

```sql
# Grant permissions to all of the existing tables in the schema
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA <schema> to <role>;

# Specify default permissions on future tables
ALTER DEFAULT PRIVILEGES IN SCHEMA <schema>
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO <role>;
```

_Note: if you're unsure which `<schema>` value to use, most likely it's
`public`!_
