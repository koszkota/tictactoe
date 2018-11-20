## Unbeatable Tic Tac Toe game in Go.

## Requirements
- golang version 1.11 (you can download go from [here](https://golang.org/dl/))

## How to run it
- clone this repository `git clone https://github.com/koszkota/tictactoe`
- `cd tictactoe`
- `export GOPATH=$HOME/path/to-tictactoe`
- run `echo $GOPATH` to check if the path is correct. I placed the directory on root, and so the path for me was: `Users/<my_name>/tictactoe`
- run `go get github.com/logrunorgru/aurora` and `go get github.com/tj/go-spin` to get the dependencies
- `go run main.go`

(If you want to run or test the project using an IDE, remember to change the `$GOPATH` in project's settings in the IDE.) 

## How to test it (98.8% coverage)
In order to run all tests, on the root level run `cd src/tictactoe` and run: `$ go test ./...`
This command should run all tests in current directory and all of its subdirectories.
As a result, you'll see a list of directories marked as `ok` when they have tests or `?` in case of no test files.

## Packages diagram
![Class Diagram](gotictactoe.jpg)
