
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
