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
