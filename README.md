# postgres-backup

postgres-backup is a tool to backup daily postgres database and transfer the dump file in Minio server.

The tool saves dump file as archive \*.dump.

## Configuration

| Environment variable |   Type   | Value       |
| :------------------: | :------: | ----------- | --- |
|    POSTGRES_HOST     |  string  |             |
|    POSTGRES_USER     |  string  |             |     |
|    POSTGRES_PORT     |  string  |             |
|     POSTGRES_DB      | []string | db1,db2,db3 |
|      PGPASSWORD      |  string  |             |
|      TIMER_MIN       |  string  | 360         |

## Restore

1. Download the dump file archive in Minio server.

2. Restore the database with pg_restore

```sh
# Create db if not exist
createdb -h HOSTNAME -p PORT -U USERNAME -T template0 yourDb
# Restore db
psql -h HOSTNAME -p PORT -U USERNAME yourDb < file.dump
```
