## Unbeatable Tic Tac Toe game in Go.

## How to run in 

## How to test it
In order to run all tests, go to the root directory and run: `$ go test ./...`
This command should run all tests in current directory and all of its subdirectories.
As a result, you'll see a list of directories marked as `ok` when they have tests or `?` in case of no test files.

## Issues 
In the `clui` package, there is a test missing for `ShowBoard()`. The reason for it is that the method which is widely used in Go for testing console output - `Example` notation and checking `// Output: ` (see `clui_tests` file for example) - doesn't strip per-line whitespace.
It makes testing multi-line output (exactly what `ShowBoard()` is doing) impossible. This issue has been raised on Go's github (see [here](https://github.com/golang/go/issues/6416)). To better understand the problem, check it out on The Go Playground [here](https://play.golang.org/p/51D2DYVHTy). It's also worth mentioning, that
the `buildBoardString` method, which is responsible for how the board looks like when is shown in `ShowBoard()` is fully tested, what makes this issue less relevant.
