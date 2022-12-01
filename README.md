# Point of Sales BackEnd System

```mermaid
flowchart TB
    CLIENT_APP<--req,res-->cmd/api
    
    storage_persistance[(Storage)]<-->pkg/cfg
    memory_cache[(Cache)]<-->pkg/cfg
    .env[(Env)]<-->pkg/cfg
    cmd/api--Log-->TMP
   
    subgraph SERVER_APP
      cmd/api<-->pkg/cfg
      cmd/api<--*en-->DEFUALT_MODULE
      cmd/api<--ctx,*cfg,*en-->ACCOUNT_MODULE
      cmd/api<--ctx,*cfg,*en-->STORE_MODULE
      cmd/api<--ctx,*cfg,*en-->CATALOG_MODULE
      cmd/api<--ctx,*cfg,*en-->TRANSACTION_MODULE
      
      subgraph DEFUALT_MODULE
        defualt_module<--*en-->default_http_handler
      end
      
      subgraph ACCOUNT_MODULE
        account_module<--*en,*cfg-->account_http_handler
        account_http_handler<--*ctx,uRepo,rRepo-->account_service
        account_service<--cfg.DBCon-->role_repository
        account_service<--cfg.DBCon-->user_repository
      end
      
      subgraph STORE_MODULE
      end
      
      subgraph CATALOG_MODULE
      end
      
      subgraph TRANSACTION_MODULE
      end
    end
```

### Create Mocks

#### Required tools:
- [Mockery](https://github.com/vektra/mockery)

#### How to use
```bash
mockery 
  --dir=internal/account/repository/mysql 
  --name=RoleSqlRepository 
  --filename=role_sql_repository.go 
  --output=domain/mocks --outpkg=mocks 
```

more info read the [docs](https://pkg.go.dev/github.com/stretchr/testify/mock).

### Database Migration

#### Required tools:
- [Golang Migrate](https://github.com/golang-migrate/migrate)

#### How to use

- Add new migration
    ```bash
    migrate create -ext sql -dir db/migrations example_table
    ```
- Run Migration

    `POSTGRESQL_URL: 'postgresql://postgres:@localhost:5432/posbe?sslmode=disable'`

  - set version (dirty state) (version: -1 before last migrate)
    ```bash
     migrate -database ${POSTGRESQL_URL} -path db/migrations force ${VERSION} 
    ```
  - up
    ```bash
    migrate -database ${POSTGRESQL_URL} -path db/migrations up
    ```
  - down
    ```bash
    migrate -database ${POSTGRESQL_URL} -path db/migrations down
    ```
    
more info read the [docs](https://pkg.go.dev/github.com/golang-migrate/migrate/v4).
