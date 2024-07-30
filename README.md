# Golang - Microservice Template for scalable project

## Overview

This repository provides a template for building microservices using Node.js and TypeScript, with Docker support for easy containerization and deployment. It is designed to help you quickly set up a scalable, maintainable microservice architecture.

## Features

- **Golang**: Asynchronous, event-driven JavaScript runtime.
- **Docker**: Containerization for consistent environments across different stages of development and deployment.
- **Echo**: Lightweight web framework for building RESTful APIs.
- **Redis**: In-memory data structure store, used as a database, cache, and message broker.
- **PostgreSQL**: Reliable and powerful open-source relational database.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or later)
- [Docker](https://www.docker.com/products/docker-desktop) (for containerization)
- [Docker Compose](https://docs.docker.com/compose/) (optional, for managing multi-container Docker applications)
- [PostgreSQL](https://www.postgresql.org/) (for local development)
- [Redis](https://redis.io/) (for local development)

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/Teyz/go-svc-template.git
```

### Update environments variables

Update environments variables in **docker-compose.yml**

### Run the project

```bash
docker compose up
```