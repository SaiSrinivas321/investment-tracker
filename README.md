# Investment Tracker API

A simple backend API to track investments in various asset types (stocks, gold, mutual funds) across different accounts.

### Built with:
- **Go (Golang)**
- **PostgreSQL**

---

## Table of Contents
- [Prerequisites](#prerequisites)
- [Setup Instructions](#setup-instructions)
- [API Endpoints](#api-endpoints)
- [Future Enhancements](#future-enhancements)
- [License](#license)

---

## Prerequisites

- **Go**: Version 1.20+
- **PostgreSQL**: Version 14+
- **Git**: Version control

---

## Setup Instructions

1. **Install Go**: Follow the official [Go installation guide](https://golang.org/doc/install).
2. **Install PostgreSQL**: Follow system-specific instructions to install and configure PostgreSQL.
3. **Set Environment Variables**: Set the required database connection variables (`DB_HOST`, `DB_PORT`, etc.).
4. **Install Dependencies**: Run `go mod tidy`.
5. **Run the Server**: Navigate to `cmd/server` and run `go run main.go`.

---

## API Endpoints

- **POST** `/investments`: Add a new investment.
- **GET** `/investments`: List all investments.
- **GET** `/investments/aggregate`: Aggregate investments by asset type and account name.

---

## Future Enhancements

- Viewing and editing individual investments.
- Live price tracking for assets.
- Authentication for secure access.

---

## License

MIT License.
