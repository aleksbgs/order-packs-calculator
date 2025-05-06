# Order Packs Calculator

The Order Packs Calculator is a Go-based web application designed to calculate the minimum number of packs needed to fulfill a given order amount. It follows a clean architecture pattern, separating concerns into distinct layers: domain (core business logic), service (application logic), presentation (HTTP handlers), and infrastructure (external concerns like configuration and logging). The application provides RESTful endpoints for calculating packs, retrieving pack sizes, and updating pack sizes, along with a simple web UI for user interaction.

## Technologies Used

- **Go**: Backend programming language (version 1.23).
- **Fiber**: Web framework for handling HTTP requests (`v2.52.6`).
- **Viper**: Configuration management (`v1.20.1`).
- **GoMock**: Mocking library for unit tests (`v1.6.0`).
- **Testify**: Testing framework for assertions and suites (`v1.10.0`).
- **Docker**: Containerization for running the application.

## Project Structure
order-packs-calculator/
├── cmd/
│   └── api/
│       └── main.go           # Application entry point
├── internal/
│   ├── domain/              # Core business logic
│   │   ├── pack.go
│   │   └── pack_test.go
│   ├── service/             # Application logic
│   │   ├── calculate_packs.go
│   │   ├── calculate_packs_test.go
│   │   └── mocks/
│   │       └── calculate_packs_mock.go
│   ├── presentation/        # HTTP handlers
│   │   ├── pack_controller.go
│   │   └── pack_controller_test.go
│   ├── infrastructure/      # External concerns (config, logging, repository)
│   │   ├── config/
│   │   │   ├── config.go
│   │   │   └── config_test.go
│   │   ├── logging/
│   │   │   ├── logger.go
│   │   │   └── logger_test.go
│   │   └── repository/
│   │       ├── pack_repository.go
│   │       ├── pack_repository_test.go
│   │       └── mocks/
│   │           └── pack_repository_mock.go
├── tests/                   # Integration tests
│   └── integration_test.go
├── web/                     # Static files for UI
│   ├── index.html
│   └── script.js
├── Dockerfile               # Docker configuration for building the app
├── docker-compose.yaml      # Docker Compose for running the app
├── config.yaml              # Configuration file for Viper
├── Makefile                 # Makefile for common tasks
├── go.mod                   # Go module dependencies
├── go.sum                   # Go dependency checksums
└── README.md                # Project documentation


## Features

- **Pack Calculation**: Determines the minimum number of packs needed for a given order amount.
- **Pack Size Management**: Retrieve and update available pack sizes.
- **Web UI**: Simple interface to interact with the application.
- **Configuration**: Supports configuration via `config.yaml` or environment variables.
- **Testing**: Comprehensive unit and integration tests using GoMock and Testify.

## Prerequisites

- **Go**: Version 1.23 or later ([Download](https://golang.org/dl/)).
- **Docker** and **Docker Compose**: For containerized deployment ([Download](https://www.docker.com/get-started)).
- **make**: For running Makefile commands.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd order-packs-calculator

# Usage

## Using Makefile (Recommended)

The project includes a Makefile to automate common tasks:

### Build the Application
```bash
make build

# Order Packs Calculator – Usage Guide

## 🛠 Using Makefile (Recommended)

Automate tasks with the provided Makefile:

### 🔧 Build the Application
```bash
make build
```

### ▶️ Run Locally
```bash
make run
```

### 🐳 Run with Docker
```bash
make docker-up
```

### 🛑 Stop Docker Containers
```bash
make docker-down
```

---

## 🐳 Running with Docker (Alternative)

Build and run the application:
```bash
docker-compose up --build
```

Access the application at: [http://localhost:3000](http://localhost:3000)

---

## 💻 Running Locally (Alternative)

### Build:
```bash
go build -o order-packs-calculator ./cmd/api
```

### Run:
```bash
./order-packs-calculator
```

Access the application at: [http://localhost:3000](http://localhost:3000)

---

## ⚙️ Configuration

Configure using `config.yaml` or environment variables:

### `config.yaml`
```yaml
port: ":3000"
pack_sizes: "250,500,1000,2000,5000"
```

### Environment Variables
- `PORT`: Override port
  ```bash
  export PORT=3000
  ```
- `PACK_SIZES`: Override pack sizes
  ```bash
  export PACK_SIZES=100,200,300
  ```

---

## 📡 API Endpoints

### `POST /api/calculate`
Calculate minimum packs for an order.

**Request Body:**
```json
{ "orderAmount": 263 }
```

**Response:**
```json
{ "packs": { "500": 1 }, "totalItems": 500 }
```

---

### `GET /api/pack-sizes`
Fetch current pack sizes.

**Response:**
```json
{ "packSizes": [250, 500, 1000, 2000, 5000] }
```

---

### `POST /api/pack-sizes`
Update pack sizes.

**Request Body:**
```json
{ "packSizes": [100, 200, 300] }
```

**Response:**
```json
{ "message": "Pack sizes updated successfully" }
```

---

## 🧪 Testing

### Run All Tests
```bash
make test
```
---

## 🤝 Contributing

1. Fork the repository.
2. Create your branch:
    ```bash
    git checkout -b feature/your-feature
    ```
3. Commit changes:
    ```bash
    git commit -m "Add your feature"
    ```
4. Push to your fork:
    ```bash
    git push origin feature/your-feature
    ```
5. Open a pull request.
6. Ensure all tests pass using `make test`.

---

## 📄 License

Licensed under the MIT License. See `LICENSE` for details.

---

## 🙏 Acknowledgments

- Built with [Go](https://golang.org) and [Fiber](https://gofiber.io)
- Configuration powered by [Viper](https://github.com/spf13/viper)
- Tests via [GoMock](https://github.com/golang/mock) and [Testify](https://github.com/stretchr/testify)



