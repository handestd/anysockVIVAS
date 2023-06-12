clean:
	go clean
#gen code api
gen:
	goctl api go -api anysock-api.api -dir . --style goZero
gomodtidy:
	go mod tidy
run:
