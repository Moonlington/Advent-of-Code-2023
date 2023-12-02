build:
	@go build .

run: build
	@./AOC23 $(day)