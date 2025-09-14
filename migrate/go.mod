module github.com/KaranMali2001/mini-ride-booking/migrate

go 1.24.4

replace github.com/KaranMali2001/mini-ride-booking/common => ../common

require (
	github.com/KaranMali2001/mini-ride-booking/common v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
	github.com/pressly/goose/v3 v3.25.0
)

require (
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
)
