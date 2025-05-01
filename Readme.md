# Go GraphQL WebSocket API

A professional Go application that combines REST API, GraphQL, and WebSocket functionality in a clean architecture.

## Project Structure

```
internal/
├── domain/           # Domain models and business logic
│   └── model.go      # User domain model
├── graphql/          # GraphQL implementation
│   ├── schema.graphqls  # GraphQL schema
│   ├── resolver.go      # GraphQL resolvers
│   └── server.go        # GraphQL server configuration
├── websocket/        # WebSocket implementation
│   └── handler.go    # WebSocket handler
└── routes/           # Route configuration
    └── routes.go     # Main router setup
```

## Features

- **REST API**: Traditional REST endpoints for CRUD operations
- **GraphQL**: Flexible querying and mutations
- **WebSocket**: Real-time updates and communication
- **Clean Architecture**: Well-organized, maintainable code structure
- **GORM**: Database operations with GORM
- **Gin**: Fast and efficient HTTP framework

## Prerequisites

- Go 1.24 or higher
- PostgreSQL (or your preferred database)
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/mohar9h/go-graphql-ws-api.git
cd go-graphql-ws-api
```

2. Install dependencies:
```bash
go mod download
```

3. Generate GraphQL code:
```bash
go run github.com/99designs/gqlgen generate
```

4. Set up your environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

## Running the Application

```bash
go run cmd/main.go
```

## API Endpoints

### REST API
- Base URL: `/api/v1`
- Endpoints:
  - `GET /users` - List all users
  - `GET /users/:id` - Get a specific user
  - `POST /users` - Create a new user
  - `PUT /users/:id` - Update a user
  - `DELETE /users/:id` - Delete a user

### GraphQL
- Endpoint: `/graphql`
- Playground: `/graphql/playground`

Example GraphQL queries:
```graphql
# Query all users
query {
  users {
    id
    name
    email
    createdAt
    updatedAt
  }
}

# Create a new user
mutation {
  createUser(input: {
    name: "John Doe"
    email: "john@example.com"
  }) {
    id
    name
    email
    createdAt
    updatedAt
  }
}
```

### WebSocket
- Endpoint: `/ws/users`
- Protocol: WebSocket
- Events:
  - `user.created` - New user created
  - `user.updated` - User updated
  - `user.deleted` - User deleted

## Development

### Code Generation
```bash
# Generate GraphQL code
go run github.com/99designs/gqlgen generate

# Generate mocks (if using)
go generate ./...
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Architecture

The application follows clean architecture principles:

1. **Domain Layer**
   - Contains core business logic
   - Defines entities and interfaces
   - Independent of frameworks

2. **GraphQL Layer**
   - Handles GraphQL-specific concerns
   - Implements resolvers
   - Maps domain models to GraphQL types

3. **WebSocket Layer**
   - Manages real-time connections
   - Handles client communication
   - Broadcasts updates

4. **Routes Layer**
   - Orchestrates all endpoints
   - Configures middleware
   - Sets up routing

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

**🎉 Happy Coding!** 🚀