# ⚙️ Go Microservices 🦫

Simple microservices written in Go, designed to explore and learn distributed systems architecture.

## Overview

This ongoing project showcases microservices that perform various tasks, enabling an understanding of concepts like communication, scalability, and orchestration in a distributed environment.

_Almost done! Only deployment to Docker Swarms and Kubernetes is left to go!_

### Microservices Included:
- **Front-End**: Serves the UI for user interactions (mostly just endpoint testing).
- **Authentication**: Manages user logins (simple bcrypt-hashed password checking).
- **Logging**: Tracks and stores application logs for monitoring.
- **Broker**: Acts as a message broker between services.
- **Listener**: Consumes and processes messages from the broker.
- **Mail Service**: Handles email notifications.

### Tech Stack
- **Programming Language**: [Go](https://golang.org/)
- **Communication**: gRPC
- **Message Queues**: RabbitMQ
- **Web Server**: Caddy
- **Databases**: MongoDB, PostgreSQL
- **Containerization**: Docker

### Deployment
The project is containerized and can be deployed using (both will be done):
- Docker Swarm
- Kubernetes

### Goals

- **Learn**: Deepen understanding of distributed systems.
- **Experiment**: Test microservices architecture and orchestration tools.
- **Build**: Create a foundation for future distributed systems projects.

### License

This project is licensed under the [MIT License](LICENSE).

---

Created with ❤️ by [Andy Tran](https://github.com/andynapoleon).

