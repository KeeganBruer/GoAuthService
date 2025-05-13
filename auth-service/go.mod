module go-auth-service

go 1.24.3

replace kbrouter => ../router

require kbrouter v0.0.0-00010101000000-000000000000

require github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
