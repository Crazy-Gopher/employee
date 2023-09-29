package employee

import "github.com/Crazy-Gopher/department"

func Greet(msg string) string {
	return "Hello" + " " + department.MagicString(msg)
}
