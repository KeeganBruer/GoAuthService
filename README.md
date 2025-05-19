# GoAuthService
An open-source authorization micro-service with SSO support.

## Components
| folder | component description |
| - | - |
| [auth-service](./auth-service/README.md)          | Core Service
| [router](./router/README.md)                      | a go based http router
| [sql-querybuilder](./sql-querybuilder/README.md)  | a go based query builder

## Demo
Using the service in a containerized environment is the fastest way to get the demo working.

This repo has a [docker-compose](./docker-compose.yaml) file already configured for use. 
To get started, clone this repo and run the following command in the root directory.
```bash
docker compose up
```
