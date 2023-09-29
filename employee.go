package employee

import (
	"strings"

	"github.com/Crazy-Gopher/department"
)

func Greet(msg string) string {
	return "Hello" + " " + department.MagicString(msg)
}

func GreetCapital(msg string) string {
	return strings.ToUpper("Hello" + " " + department.MagicString(msg))
}
