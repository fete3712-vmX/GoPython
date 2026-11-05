package main

import (
	"fmt"
	pysnake "gopython/pysnake"
)

func main() {
	program := `def magnitude(value):
    return abs(value)
print(magnitude(-42))`
	fmt.Print(pysnake.RunCode(program))
}
