# test-go


## reference
    https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter09/09.1.md

## cover
1. terminal: `go test -cover`
2. html
```
        go test -coverprofile=cover
        go tool cover -html=cover
```

## test file or func

1. test single file `go test ./base/utils_test.go`
2. test single func `go test ./base/utils_test.go -run AddOne`