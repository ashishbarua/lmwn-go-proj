### LMWN Golang assignment
Following are some steps to get started with running and testing the project. All commands are to be run from the root directory of the project, unless specified otherwise

#### Pre-requisite
- Run the following command for installing all dependencies
```
$ go get .
```

#### Running the code
- Run the following command to start the project
```
$ go run .
```
- Call the path `http://localhost:3000/covid/summary` (in your browser or any REST client e.g. Postman) to get the results in JSON format

#### Run the tests
- Run the following command to run all tests
```
$ go test -v ./...
```
Ideally, all the tests should pass
