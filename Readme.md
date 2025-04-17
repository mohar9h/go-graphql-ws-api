# **Go API + GraphQL + WebSocket Project**

A modern Go backend with **REST API**, **GraphQL**, and **WebSocket** support. Structured for scalability, maintainability, and ease of development.

---

## **📂 Project Structure**

```
my-go-project/
├── cmd/                          # Application entry points
│   └── myapp/                    # Main application
│       └── main.go               # Server initialization
│
├── internal/                     # Private application code
│   ├── api/                      # REST API components
│   │   ├── handlers/             # HTTP handlers (e.g., user_handler.go)
│   │   ├── middleware/           # Middleware (e.g., auth_middleware.go)
│   │   └── routes.go             # API route definitions
│   │
│   ├── graphql/                  # GraphQL server
│   │   ├── schema/               # GraphQL schema definitions
│   │   ├── resolvers/            # Query/mutation resolvers
│   │   └── server.go             # GraphQL server setup
│   │
│   ├── ws/                       # WebSocket (real-time communication)
│   │   ├── handlers/             # WebSocket handlers (e.g., chat_handler.go)
│   │   └── server.go             # WebSocket server setup
│   │
│   ├── config/                   # Configuration management
│   ├── models/                   # Data models (e.g., User)
│   ├── services/                 # Business logic layer
│   └── utils/                    # Helper functions
│
├── pkg/                          # Reusable packages
│   ├── database/                 # Database connection & queries
│   └── logger/                   # Logging utilities
│
├── migrations/                   # Database migration scripts
├── scripts/                      # Helper scripts (e.g., start.sh)
│
├── .env                          # Environment variables
├── .gitignore                    # Git ignore rules
├── go.mod                        # Go module definition
├── go.sum                        # Dependency checksums
└── README.md                     # Project documentation
```

---

## **🚀 Getting Started**

### **Prerequisites**
- Go 1.20+
- PostgreSQL / MySQL (or any preferred database)
- Dependencies:
    - `github.com/gorilla/mux` (HTTP routing)
    - `github.com/graphql-go/graphql` (GraphQL support)
    - `github.com/gorilla/websocket` (WebSocket)
    - `github.com/joho/godotenv` (Environment variables)

### **Installation**
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/my-go-project.git
   cd my-go-project
   ```

2. Install dependencies:
   ```sh
   go mod download
   ```

3. Set up `.env` (example):
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=myapp
   JWT_SECRET=your-secret-key
   PORT=8080
   ```

4. Run database migrations:
   ```sh
   # Example (adjust based on your migration tool)
   psql -U postgres -d myapp -a -f migrations/001_create_users_table.sql
   ```

5. Start the server:
   ```sh
   go run cmd/myapp/main.go
   ```

---

## **🌐 API Endpoints**

### **REST API**
| Endpoint       | Method | Description                     | Auth Required |
|---------------|--------|---------------------------------|--------------|
| `/api/login`  | POST   | User login (JWT token)          | No           |
| `/api/users`  | GET    | Fetch all users                 | Yes (JWT)    |

### **GraphQL**
- **Endpoint**: `/graphql`
- **Example Query**:
  ```graphql
  query {
    users {
      id
      name
      email
    }
  }
  ```

### **WebSocket**
- **Endpoint**: `/ws`
- **Example Usage**:
  ```javascript
  const socket = new WebSocket("ws://localhost:8080/ws");
  socket.onmessage = (event) => console.log(event.data);
  ```

---

## **🔧 Configuration**

### **Environment Variables**
| Variable      | Description                | Default   |
|--------------|----------------------------|-----------|
| `DB_HOST`    | Database host              | `localhost` |
| `DB_PORT`    | Database port              | `5432` (PostgreSQL) |
| `JWT_SECRET` | JWT signing key            | (Required) |
| `PORT`       | Server port                | `8080` |

### **Database Setup**
- Modify `pkg/database/db.go` for your DB driver (PostgreSQL, MySQL, etc.).
- Run migrations from `migrations/`.

---

## **🛠️ Development**

### **Running the Server**
```sh
go run cmd/myapp/main.go
```

### **Testing**
```sh
go test ./...
```

### **Building for Production**
```sh
go build -o bin/myapp cmd/myapp/main.go
```

---

## **📜 License**
MIT

---

## **📌 Notes**
- **Security**: Always use HTTPS in production.
- **Scaling**: Consider adding Redis for caching WebSocket connections.
- **Monitoring**: Integrate Prometheus for metrics.

---

**🎉 Happy Coding!** 🚀