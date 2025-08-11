build:
	GOOS=darwin GOARCH=arm64 go build -o bigary execute/main.go


run:
	go run execute/main.go command name="ali reza" --age 25 family=mohammadi -foo "barbique"