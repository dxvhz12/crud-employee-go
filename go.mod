module github.com/dxvhz12/crud-employee-go

go 1.25.0

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // Library JWT (JSON Web Token).
	github.com/jackc/pgx/v5 v5.10.0 // Driver PostgreSQL modern.
	github.com/joho/godotenv v1.5.1 // Mirip Laravel .env.
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect; Membaca file PostgreSQL .pgpass
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect; Membaca konfigurasi PostgreSQL dari pg_service.conf
	github.com/jackc/puddle/v2 v2.2.2 // indirect; Connection Pool untuk PGX.
	golang.org/x/sync v0.17.0 // indirect; Utility untuk goroutine. (dependency internal pgx/v5)
	golang.org/x/text v0.29.0 // indirect; Library pengolahan teks dan Unicode. (dependency internal pgx/v5)
)
