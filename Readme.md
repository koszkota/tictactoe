## Unbeatable Tic Tac Toe game in Go.

## How to run in 

## How to test it
In order to run all tests, go to the root directory and run: `$ go test ./...`
This command should run all tests in current directory and all of its subdirectories.
As a result, you'll see a list of directories marked as `ok` when they have tests or `?` in case of no test files.

## Issues 
In the clui package, there is a test missing for ShowBoard(). The reason for it is that the method which is widely used in go for testing console output (using `Example` notation and checking `// Output: ` (see tests in clui tests file for example) doesn't strip per-line whitespace. It makes testing multi-line output (like the board) impossible. This issue has been rised on Go's github (see [here](https://github.com/golang/go/issues/6416). To better understand the problem, check it out on The Go Playground [here](https://play.golang.org/p/51D2DYVHTy).
