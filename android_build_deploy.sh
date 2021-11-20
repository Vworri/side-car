

env CGO_ENABLED=0 GOOS=android GOARCH=arm64  go build -o sidecar.exe main.go
adb push sidecar.exe /data/local/tmp

