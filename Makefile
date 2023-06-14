clean:
	go clean
#gen code api
gen:
	goctl api go -api greet.api -dir . --style goZero
# get required package
de:
	go mod tidy
run:
	