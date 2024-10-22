coverfile := ".coverage"


test *opts:
    go test {{ opts }} -coverprofile {{ coverfile }} .

show-coverage:
    go tool cover -html {{ coverfile }}

test-and-show-coverage: test show-coverage

run:
    go run ./cmd