# GoPython Documentation

GoPython is a small Python-inspired interpreter and runner written in Go. It is experimental and currently in **beta**.

## Features

- String and numeric variable assignments
- Functions with positional parameters
- Return values
- Built-in `print()` and `abs()`
- Calls to user-defined functions, including nested calls
- Blank lines and comments beginning with `#`

## Getting started

The Go module is in `src/` and targets Go 1.22.

Run the example:

```sh
cd src
go run example.go GoPython.go
```

The sample defines `magnitude(value)` and prints the absolute value of `-42`.

## Language syntax

### Comments

Lines whose trimmed contents begin with `#` are ignored.

```python
# This is a comment
```

### Variables and values

Assignments use `name = expression`. Supported expressions include quoted strings, numeric literals, variable references, and function calls.

```python
message = "Hello"
number = 42
```

### Functions

Define a function with `def`, parentheses, and a colon. Its body is indented. Parameters are positional.

```python
def magnitude(value):
    return abs(value)

print(magnitude(-42))
```

Use `return expression` to return a value or `return` to return an empty value.

### Built-ins

| Function | Behavior |
| --- | --- |
| `print(...)` | Appends arguments to interpreter output, separated by spaces. |
| `abs(number)` | Returns the absolute value of a numeric argument. |

## Go API

The interpreter exposes `RunCode(source string) string`. It executes source code and returns captured output.

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

Run Go commands from the module directory (`src/`) to resolve the module correctly.

## Limitations

GoPython is not a full Python implementation. At present, it does not support general control-flow statements, imports, classes, or full Python expression/operator semantics. Invalid or unsupported expressions do not yet produce detailed user-facing errors. Recursive execution is limited by a depth guard.

## Contributing

Issues and pull requests are welcome. Keep changes focused and include a small example when adding or changing language behavior.

## License

GoPython is distributed under the MIT License. See [LICENSE](LICENSE).
