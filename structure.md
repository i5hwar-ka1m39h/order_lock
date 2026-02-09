
job-system/
├── api/
│   └── main.go
├── worker/
│   └── main.go
├── queue/
│   ├── queue.go        // interface
│   ├── redis.go
│   ├── rabbitmq.go
│   └── kafka.go
├── internal/
│   ├── retry/
│   ├── metrics/
│   └── logging/
└── docker-compose.yml


shit to consider 
first create the db folder which should contains the schema(migration) and queries
schema changes/create the shit in table using goose command 