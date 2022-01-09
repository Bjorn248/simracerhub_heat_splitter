build:
	go build -o simracerhub_heat_splitter main.go

release:
	export GOARCH=amd64
	GOOS=linux go build -o simracerhub_heat_splitter_linux main.go
	GOOS=darwin go build -o simracerhub_heat_splitter_mac main.go
	GOOS=windows go build -o simracerhub_heat_splitter_windows.exe main.go
