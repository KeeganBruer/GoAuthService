module go-auth-service

go 1.24.3

replace kbrouter => ../router

require kbrouter v0.0.0-00010101000000-000000000000

require filippo.io/edwards25519 v1.1.0 // indirect

require (
	github.com/go-sql-driver/mysql v1.9.2
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
)
