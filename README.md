# GoPython

GoPython is a small Python-inspired interpreter and runner written in Go. It is a work in progress and is currently in **beta**.

## Features

- Assign string and numeric values to variables.
- Define functions with positional parameters.
- Return values from functions.
- Call built-in `print()` and `abs()`.
- Call user-defined functions, including nested calls.
- Ignore blank lines and comments beginning with `#`.

## Getting started

The Go module is located in `src/` and targets Go 1.22.

To run the example program:

```sh
cd src
go run example.go GoPython.go
```

The example defines a `magnitude(value)` function and prints the absolute value of `-42`.

## Python-like syntax supported

```python
# Define a function
def magnitude(value):
    return abs(value)

# Call it and print the result
print(magnitude(-42))
```

Assignments use `name = expression`. Expressions currently support quoted strings, numeric literals, variable references, and function calls. Function definitions use `def name(parameter, ...):` with an indented body. A function can return a value with `return expression` or return an empty value with `return`.

### Built-in functions

| Function | Description |
| --- | --- |
| `print(...)` | Writes its arguments to the interpreter output, separated by spaces. |
| `abs(number)` | Returns the absolute value of a numeric argument. |

## Using GoPython from Go

The interpreter exposes `RunCode(source string) string`, which executes source text and returns captured output:

```go
package main

import (
    "fmt"
    gopython "gopython"
)

func main() {
    source := `print(abs(-42))`
    fmt.Println(gopython.RunCode(source))
}
```

Run code from the module directory (`src/`) so Go can resolve the local module.

## Current limitations

GoPython is experimental and does not implement all of Python. Its current interpreter is intentionally small: it has no general control-flow statements, imports, classes, or full Python expression/operator support. Invalid or unsupported expressions are not reported as rich user-facing errors yet. Recursion is bounded by an execution-depth limit.

## Contributing

Issues and pull requests are welcome. Please keep changes focused and include a small example when adding or changing language behavior.

## License

This project is distributed under the MIT License. See [LICENSE](LICENSE).
