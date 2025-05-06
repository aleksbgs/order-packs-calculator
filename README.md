# Order Packs Calculator

The Order Packs Calculator is a Go-based web application designed to calculate the minimum number of packs needed to fulfill a given order amount. It follows a clean architecture pattern, separating concerns into distinct layers: domain (core business logic), service (application logic), presentation (HTTP handlers), and infrastructure (external concerns like configuration and logging). The application provides RESTful endpoints for calculating packs, retrieving pack sizes, and updating pack sizes, along with a simple web UI for user interaction.

## 🧰 Technologies Used

- **Go**: Backend programming language (version 1.23)
- **Fiber**: Web framework (`v2.52.6`)
- **Viper**: Configuration management (`v1.20.1`)
- **GoMock**: Mocking for unit tests (`v1.6.0`)
- **Testify**: Testing framework (`v1.10.0`)
- **Docker**: Containerization

## 🗂 Project Structure

```
order-packs-calculator/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── domain/                    # Core business logic
│   │   ├── pack.go
│   │   └── pack_test.go
│   ├── service/                   # Application logic
│   │   ├── calculate_packs.go
│   │   ├── calculate_packs_test.go
│   │   └── mocks/
│   │       └── calculate_packs_mock.go
│   ├── presentation/              # HTTP handlers
│   │   ├── pack_controller.go
│   │   └── pack_controller_test.go
│   ├── infrastructure/            # External concerns
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
├── tests/                         # Integration tests
│   └── integration_test.go
├── web/                           # Static files
│   ├── index.html
│   └── script.js
├── Dockerfile
├── docker-compose.yaml
├── config.yaml
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

The Order Packs Calculator follows a **clean architecture** pattern, which ensures a clear separation of concerns, enhances testability, and makes the application easier to maintain and extend. The architecture is divided into four main layers: **domain**, **service**, **presentation**, and **infrastructure**. Each layer has a specific role and interacts with other layers in a controlled manner, adhering to dependency rules (inner layers are independent of outer layers).

### 🌟 Domain Layer
- **Location**: `internal/domain`
- **Role**: Contains the core business logic and rules of the application, independent of any external frameworks or systems.
- **Key Files**:
   - `pack.go`: Implements the `CalculatePacks` function, which calculates the minimum number of packs needed for a given order amount using a dynamic programming approach.
   - `pack_test.go`: Tests the `CalculatePacks` function with various scenarios (e.g., exact matches, overshooting, error cases).
- **Dependencies**: None. The domain layer is pure and does not depend on any other layers, ensuring that business logic remains isolated and reusable.

### 🛠 Service Layer
- **Location**: `internal/service`
- **Role**: Orchestrates application logic, acting as a bridge between the domain and presentation layers. It handles use cases like calculating packs and managing pack sizes.
- **Key Files**:
   - `calculate_packs.go`: Defines the `CalculatePacksUseCase`, which interacts with the repository to fetch pack sizes and calls the domain layer to perform calculations.
   - `calculate_packs_test.go`: Tests the use case with a mocked repository.
   - `mocks/calculate_packs_mock.go`: GoMock-generated mock for the `CalculatePacksService` interface, used in presentation layer tests.
- **Dependencies**: Depends on the **domain** layer (for business logic) and the **infrastructure/repository** layer (for data access). It does not depend on the presentation layer, maintaining separation.

### 🌐 Presentation Layer
- **Location**: `internal/presentation`
- **Role**: Handles HTTP requests and responses, exposing the application’s functionality via RESTful endpoints and serving the web UI.
- **Key Files**:
   - `pack_controller.go`: Implements the `PackController`, which defines handlers for the `/api/calculate`, `/api/pack-sizes` (GET), and `/api/pack-sizes` (POST) endpoints.
   - `pack_controller_test.go`: Tests the HTTP handlers using a mocked `CalculatePacksService`.
- **Dependencies**: Depends on the **service** layer (to perform use cases) and the **infrastructure/logging** layer (for logging requests and errors). It uses Fiber to handle HTTP requests.

### 🛠 Infrastructure Layer
- **Location**: `internal/infrastructure`
- **Role**: Manages external concerns such as configuration, logging, and data storage.
- **Sub-Layers**:
   - **`config`**: Handles configuration loading with Viper.
      - `config.go`: Loads settings from `config.yaml` or environment variables.
      - `config_test.go`: Tests configuration loading scenarios.
   - **`logging`**: Provides logging functionality.
      - `logger.go`: Implements a simple logger for info and error messages.
      - `logger_test.go`: Tests logging behavior.
   - **`repository`**: Manages data persistence (currently in-memory).
      - `pack_repository.go`: Implements an in-memory repository for storing pack sizes.
      - `pack_repository_test.go`: Tests repository operations.
      - `mocks/pack_repository_mock.go`: GoMock-generated mock for the `PackRepository` interface, used in service layer tests.
- **Dependencies**: Depends on external libraries (e.g., Viper for configuration). It does not depend on the presentation or service layers, adhering to clean architecture principles.

### 🔄 Layer Interactions
- **Presentation → Service**: The `PackController` in the presentation layer calls methods on the `CalculatePacksService` interface (implemented by `CalculatePacksUseCase`) to perform use cases like calculating packs or updating pack sizes.
- **Service → Domain**: The `CalculatePacksUseCase` in the service layer calls the `CalculatePacks` function in the domain layer to perform the core pack calculation logic.
- **Service → Infrastructure**: The `CalculatePacksUseCase` interacts with the `PackRepository` interface (implemented in `infrastructure/repository`) to fetch or update pack sizes.
- **Presentation → Infrastructure**: The `PackController` uses the `Logger` from `infrastructure/logging` to log requests and errors.

### 🛡️ Dependency Rule
- Inner layers (`domain`, `service`) do not depend on outer layers (`presentation`, `infrastructure`).
- The **domain** layer is completely isolated, containing only pure business logic.
- The **service** layer acts as a mediator, depending only on the domain and infrastructure layers.
- The **presentation** and **infrastructure** layers are outer layers, handling external interactions (HTTP, configuration, logging, etc.).

This architecture ensures that business logic remains independent, testable, and adaptable to changes in external systems (e.g., replacing the in-memory repository with a database).

## 🚀 Features

- **Pack Calculation** – Computes the minimum number of packs for an order.
- **Pack Size Management** – Retrieve and update available pack sizes.
- **Web UI** – Basic interface for interaction.
- **Flexible Configuration** – Supports `config.yaml` and environment variables.
- **Testing** – Unit + integration tests via GoMock & Testify.

## 📦 Prerequisites

- **Go** 1.23+ → [Download](https://golang.org/dl/)
- **Docker** & **Docker Compose** → [Download](https://www.docker.com/get-started)
- **make**

## 📥 Installation

```bash
git clone <repository-url>
cd order-packs-calculator
```

## 🛠 Using Makefile (Recommended)

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

## 🐳 Docker (Alternative)

```bash
docker-compose up --build
```
Access: [http://localhost:3000](http://localhost:3000)

---

## 💻 Local Run (Alternative)

```bash
go build -o order-packs-calculator ./cmd/api
./order-packs-calculator
```
Access: [http://localhost:3000](http://localhost:3000)

---

## ⚙️ Configuration

### config.yaml
```yaml
port: ":3000"
pack_sizes: "250,500,1000,2000,5000"
```

### Env Vars
```bash
export PORT=3000
export PACK_SIZES=100,200,300
```

---

## 📡 API Endpoints

### `POST /api/calculate`
```json
Request:  { "orderAmount": 263 }
Response: { "packs": { "500": 1 }, "totalItems": 500 }
```

### `GET /api/pack-sizes`
```json
Response: { "packSizes": [250, 500, 1000, 2000, 5000] }
```

### `POST /api/pack-sizes`
```json
Request:  { "packSizes": [100, 200, 300] }
Response: { "message": "Pack sizes updated successfully" }
```

---

## 🧪 Testing

```bash
make test            # All tests
make test-unit       # Unit tests
make test-integration # Integration tests
make test-coverage   # Coverage
```

---

## 🤝 Contributing

1. Fork this repo
2. Create a branch:
   ```bash
   git checkout -b feature/my-feature
   ```
3. Commit & Push
   ```bash
   git commit -m "Add feature"
   git push origin feature/my-feature
   ```
4. Open a Pull Request

---

## 📄 License

MIT License – see `LICENSE` file.

---

## 🙏 Acknowledgments

- [Go](https://golang.org)
- [Fiber](https://gofiber.io)
- [Viper](https://github.com/spf13/viper)
- [GoMock](https://github.com/golang/mock)
- [Testify](https://github.com/stretchr/testify)
