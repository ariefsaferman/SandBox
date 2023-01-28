# GORM - Raw SQL Backend

### API Lists:
- Get all trainers data
  ```
  GET /v1/trainers
  GET /v1/gorm/trainers
  ```
- Add new trainer
  ```
  POST /v1/trainers
    Content-Type: application/json
    Body:
    {
      "email": "user@email.com",
      "password": "password"
    }
  ```
- Register
  ```
  POST /v1/register
  ```
- Login
  ```
  POST /v1/login
  POST /v1/gorm/login
    Content-Type: application/json
    Body:
    {
      "email": "user@email.com",
      "password": "password"
    }
  ```

### Benchmarking with wrk
```
wrk -c100 -t100 -d20s -s ./benchmarks/login-post.lua http://localhost:3000/v1/login
```

### Benchmarking with apache bench
```
wrk -c100 -t100 -d20s -s ./benchmarks/login-post.lua http://localhost:3000/v1/login
```

### Run using Docker
- Run `docker-compose up` to build the docker container and image.
- After the build have been completed, you can access the API but the database is still empty
- Go to `http://localhost:5050/` from your browser to access pgAdmin for the database
- Register a new server with the following connection detail:
  ```
  Name: Postgres
  Host name/address full_db_postgres
  Port: 5432
  Username: admin
  Password: admin
  ```
- Restore the sql dump in `/db/db-dump.sql` by opening the terminal of `full_db_postgres` and type:
  ```
  cat db-dump.sql| docker exec -i full_db_postgres psql -U admin trainers_db
  ```

### References
- [Review API GORM vs Database/SQL di Golang](https://www.youtube.com/watch?v=x9QWtzCzWv8&ab_channel=CloudEngineeringwithImre)