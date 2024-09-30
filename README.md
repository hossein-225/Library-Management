
# Library Management System

This project is a **Library Management System** implemented using a microservices architecture. The system comprises five core services:

1. **Auth Service (auth-service)** - Handles user authentication and token generation.
2. **User Service (user-service)** - Manages user profiles and registration.
3. **Book Service (book-service)** - Manages the library's books (CRUD operations).
4. **Borrow Service (borrow-service)** - Manages borrowing and returning books.
5. **API Gateway (api-gateway)** - Provides a unified RESTful interface for external communication and routes requests to the internal gRPC services.

Each microservice is containerized and deployed using **Docker Compose**, ensuring seamless deployment across different environments. Communication between the microservices happens via **gRPC**, while users interact with the system through RESTful APIs exposed by the API Gateway.

## Table of Contents

- [Project Overview](#project-overview)
- [Technologies](#technologies)
- [Architecture](#architecture)
- [Setup and Deployment](#setup-and-deployment)
- [Database Setup](#database-setup)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Contributing](#contributing)
- [Contact](#contact)

## Project Overview

The **Library Management System** is built as a collection of microservices, each responsible for a specific functionality. This architecture enables better scalability, maintainability, and fault tolerance. Each service has its own database, ensuring data isolation and independent scalability.

Key Features:
- Authentication service with JWT token generation.
- User management (registration, profile management).
- Book management (CRUD operations).
- Borrow and return books functionality.
- API Gateway to handle all external RESTful API calls.
- gRPC communication between microservices.
- Swagger documentation for internal APIs.
- Postman collection provided for API testing.

## Technologies

- **Golang**: The core language for all microservices.
- **gRPC**: For inter-service communication.
- **REST API**: Exposed via the API Gateway.
- **PostgreSQL**: Database used for each microservice.
- **Docker & Docker Compose**: For containerizing and orchestrating the services.
- **Swagger**: API documentation generation.
- **Makefile**: Build automation for gRPC and Swagger documentation generation.

## Architecture

### Microservices:
1. **Auth Service**: Handles user authentication and JWT token generation.
2. **User Service**: Manages user-related operations such as registration, profile management, and more.
3. **Book Service**: Manages the library's collection of books (add, update, delete, search).
4. **Borrow Service**: Handles borrowing and returning of books.
5. **API Gateway**: Acts as a single entry point for all client requests, routing them to the respective microservices.

Each service is designed to be independent with its own database and can be scaled individually.

## Setup and Deployment

To deploy the services locally, you can use **Docker Compose**. Simply navigate to the root directory and run:

```bash
docker-compose up
```

This command will spin up all the services and their respective databases. The API Gateway will be exposed on port `8080`, while the other services will only communicate internally.

### Build Instructions

Each service has its own `Makefile` which simplifies common tasks such as building protocol buffers and generating documentation. You can run the following commands within each service:

- **Build gRPC and Mock Services**:

```bash
make build-proto
```

- **Generate Swagger Documentation**:

```bash
make build-doc
```

## Database Setup

Each microservice has its own **PostgreSQL** database. There is no replication or advanced database setup currently, but the architecture is flexible enough to add such features in the future.

The database instances are also managed through Docker Compose, and no external access is allowed. All communications happen internally through the services.

## API Documentation

There are two types of API documentation:

1. **Swagger Documentation**: Each service has its Swagger docs generated inside the `/docs` folder. Swagger is used to document the gRPC services and provides an interactive way to explore the APIs. You can generate the documentation by running:

   ```bash
   make build-doc
   ```

2. **Postman Collection**: A complete Postman collection is available in the root directory as a `.json` file. This collection provides examples for all available API routes and is useful for testing the system's functionality.

## Testing

Most services come with unit tests, particularly focusing on core functionalities and gRPC communication. The tests can be executed using Go's testing tools.

To run the tests:

```bash
go test ./...
```

## Contributing

Contributions are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Submit a pull request.

Ensure that your code follows the project's coding standards and includes tests where applicable.

## Contact

**Author**: Seyed Hossein Hosseini Motlagh  
**Email**: [hossein_225@yahoo.com](mailto:hossein_225@yahoo.com)

Feel free to reach out with any questions, suggestions, or issues related to the project.
