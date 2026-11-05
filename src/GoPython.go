package gopython

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

type userFunction struct {
	parameters []string
	body       []string
}

type interpreter struct {
	variables map[string]string
	functions map[string]userFunction
	output    []string
}

func RunCode(source string) string {
	vm := &interpreter{
		variables: make(map[string]string),
		functions: make(map[string]userFunction),
	}
	vm.execute(strings.Split(source, "\n"), vm.variables, 0)
	return strings.Join(vm.output, "\n")
}

func (vm *interpreter) execute(lines []string, variables map[string]string, depth int) (string, bool) {
	if depth > 100 {
		return "", false
	}

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "def ") {
			name, parameters, ok := parseFunctionHeader(line)
			if !ok {
				continue
			}
			body := make([]string, 0)
			for i+1 < len(lines) {
				next := lines[i+1]
				if strings.TrimSpace(next) == "" {
					i++
					continue
				}
				if !unicode.IsSpace(rune(next[0])) {
					break
				}
				i++
				body = append(body, strings.TrimSpace(next))
			}
			vm.functions[name] = userFunction{parameters: parameters, body: body}
			continue
		}

		if strings.HasPrefix(line, "return ") {
			value, ok := vm.evaluate(strings.TrimSpace(strings.TrimPrefix(line, "return ")), variables, depth)
			return value, ok
		}
		if line == "return" {
			return "", true
		}

		if name, expression, ok := splitAssignment(line); ok {
			if value, valid := vm.evaluate(expression, variables, depth); valid {
				variables[name] = value
			}
			continue
		}

		vm.evaluate(line, variables, depth)
	}
	return "", false
}

func (vm *interpreter) evaluate(expression string, variables map[string]string, depth int) (string, bool) {
	expression = strings.TrimSpace(expression)
	if expression == "" {
		return "", true
	}

	if (strings.HasPrefix(expression, "\"") && strings.HasSuffix(expression, "\"")) ||
		(strings.HasPrefix(expression, "'") && strings.HasSuffix(expression, "'")) {
		if expression[0] == '\'' {
			return expression[1 : len(expression)-1], true
		}
		value, err := strconv.Unquote(expression)
		return value, err == nil
	}
	if value, ok := variables[expression]; ok {
		return value, true
	}
	if _, err := strconv.ParseFloat(expression, 64); err == nil {
		return expression, true
	}

	name, arguments, ok := parseCall(expression)
	if !ok {
		return "", false
	}
	values := make([]string, 0, len(arguments))
	for _, argument := range arguments {
		value, valid := vm.evaluate(argument, variables, depth)
		if !valid {
			return "", false
		}
		values = append(values, value)
	}

	switch name {
	case "abs":
		if len(values) != 1 {
			return "", false
		}
		number, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			return "", false
		}
		return strconv.FormatFloat(math.Abs(number), 'f', -1, 64), true
	case "print":
		vm.output = append(vm.output, strings.Join(values, " "))
		return "", true
	}

	function, exists := vm.functions[name]
	if !exists || len(function.parameters) != len(values) {
		return "", false
	}
	localVariables := make(map[string]string, len(values))
	for i, parameter := range function.parameters {
		localVariables[parameter] = values[i]
	}
	return vm.execute(function.body, localVariables, depth+1)
}

func parseFunctionHeader(line string) (string, []string, bool) {
	line = strings.TrimSpace(strings.TrimPrefix(line, "def "))
	name, arguments, ok := parseCall(strings.TrimSuffix(line, ":"))
	if !ok || name == "" {
		return "", nil, false
	}
	for i := range arguments {
		arguments[i] = strings.TrimSpace(arguments[i])
	}
	return name, arguments, true
}

func parseCall(expression string) (string, []string, bool) {
	open := strings.IndexByte(expression, '(')
	if open <= 0 || !strings.HasSuffix(strings.TrimSpace(expression), ")") {
		return "", nil, false
	}
	name := strings.TrimSpace(expression[:open])
	if strings.ContainsAny(name, " ()") {
		return "", nil, false
	}
	inside := strings.TrimSpace(expression[open+1 : strings.LastIndex(expression, ")")])
	if inside == "" {
		return name, nil, true
	}
	return name, splitArguments(inside), true
}

func splitArguments(text string) []string {
	var arguments []string
	start, depth := 0, 0
	var quote rune
	for i, character := range text {
		if quote != 0 {
			if character == quote {
				quote = 0
			}
			continue
		}
		switch character {
		case '\'', '"':
			quote = character
		case '(':
			depth++
		case ')':
			depth--
		case ',':
			if depth == 0 {
				arguments = append(arguments, strings.TrimSpace(text[start:i]))
				start = i + 1
			}
		}
	}
	return append(arguments, strings.TrimSpace(text[start:]))
}

func splitAssignment(line string) (string, string, bool) {
	index := strings.Index(line, "=")
	if index <= 0 || strings.Contains(line, "==") {
		return "", "", false
	}
	name := strings.TrimSpace(line[:index])
	if strings.ContainsAny(name, " ()") {
		return "", "", false
	}
	return name, strings.TrimSpace(line[index+1:]), true
}



