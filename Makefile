build:
	CGO_ENABLED=0 GOOS=linux go build -o monobot -a -installsuffix cgo
