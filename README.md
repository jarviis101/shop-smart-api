Shop Smart Backend API
====

## Requirements

- Make (https://www.gnu.org/software/make/)
- Go 1.20 (https://go.dev/)
- Docker Compose (https://docs.docker.com/compose/)

## Run API
```bash
cp .env.dist .env
make run-server
```

## Environment variables
API environments:
1. `PORT` - API port
2. `APP_SECRET` - Application secret key
3. `APP_ENV` - Environment
4. `APP_DEBUG` - If you enable debug mode, OTP provider will be disabled and OTP always equals "1111"

Database environments:
1. `POSTGRES_PASSWORD` - Database user password
2. `POSTGRES_USER` - Database username
3. `POSTGRES_DB` - Database name
4. `DATABASE_URL` - Database URL for connecting with API

Other environments:
1. `SMS_API_KEY` - API key from sms provider

## Migrations
For up migrations use command:
```bash
make up-migration
```

For down migrations use command:
```bash
make down-migration
```

## Codestyle
Use command for run linter:
```bash
make run-lint
```
