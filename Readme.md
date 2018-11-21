## Unbeatable Tic Tac Toe game in Go.

## Requirements
- golang version 1.11 (you can download go from [here](https://golang.org/dl/))

## How to run it
- clone this repository `git clone https://github.com/koszkota/tictactoe`
- `cd tictactoe`
- `export GOPATH=$HOME/path/to-tictactoe`
- run `echo $GOPATH` to check if the path is correct. I placed the directory on root, and so the path for me was: `Users/<my_name>/tictactoe`
- run `go get github.com/tj/go-spin` to get the dependency
- `go run main.go`

(If you want to run or test the project using an IDE, remember to change the `$GOPATH` in project's settings in the IDE.) 

## How to test it (98.8% statements covered)
In order to run all tests, on the root level run `cd src/tictactoe` and run: `$ go test ./...`
This command should run all tests in current directory and all of its subdirectories.
As a result, you'll see a list of directories marked as `ok` when they have tests or `?` in case of no test files.

## Wish List 
If I had more time, I would think more about the following design issues:  
1. Computer player needs to know both players' marks for the minimax algorithm. 
There is a design question concerning the place from which the computer should get them.
Right now it gets them from the `board`, as the `board` is passed to the `pickMove()` method anyway, but I feel that there could be a better place than a `board` to keep the marks. 
I've made an attempt at it by creating a `marksRepo` struct, but it's still accessible via `board`, so it doesn't completly solve this issue. If marks were constant (weren't picked in the initial menu) I would just write a `const` for them and import it when needed. However,
because the marks aren't known from the beginning, they need to be set and passed to the instances that need them. 
2. I've used the factory pattern to create games and players what made the code quite skinny. I think there is still a place left 
for a higher abstraction that would eliminate the `switch` statement I have in the `factory` struct in `game` package.
I was experimenting with this idea and came to a conclusion that right now it would be an overkill that would make the code much harder to read.
If I had more time, I would make another attempt at finding a golden ratio between the `switch` statement and abstractions. 

## Packages diagram
![Class Diagram](gotictactoe.jpg)
