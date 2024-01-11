test:
	go test ./response/... -coverprofile=coverage.html
	go tool cover -html=coverage.html
