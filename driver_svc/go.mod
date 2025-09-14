module github.com/KaranMali2001/mini-ride-booking/driver_svc

go 1.24.4

require (
	github.com/KaranMali2001/mini-ride-booking/common v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
)

require (
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
)

replace github.com/KaranMali2001/mini-ride-booking/common => ../common
