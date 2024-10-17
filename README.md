# Pinterest API

## Setup Instructions

### Prerequisites
- Go 1.19+
- PostgreSQL
- MinIO
- Swagger for API documentation

### Installation

1. Clone the repo:
     git clone https://github.com/Bakdaulett/Pinterest.git
     cd pinterest
   
2. Install dependencies:
    go mod tidy

3. Configure environment variables: Create a .env file in the root and add the necessary variables:
    MINIO_ENDPOINT=localhost:9000
    MINIO_ACCESS_KEY=minioadmin
    MINIO_SECRET_KEY=minioadmin123
    POSTGRES_DSN=host=localhost user=postgres password=postgres dbname=pinterest port=5432 sslmode=disable
    JWT_SECRET=your_secret_key

4. Run the application:
    go run cmd/main.go

