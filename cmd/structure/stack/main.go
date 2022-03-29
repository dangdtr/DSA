package main

import (
	"fmt"
	"github.com/dawnpanpan/go-dsa/algs"
	"regexp"
)

//  If operand -> output.
//  If "(" -> stack.
//  If ")" -> pop stack util encounter ")" and pop ")".
// 	If opreator (cur)
// 		- while stack.peek is operator
//			and priority(peek) >= priority(cur) then pop stack -> output
// 		- push (cur)
func InfixToPostfix(infix string) string {
	FormatExpression(infix)

	var postfix string
	stack := algs.NewStack()

	for _, char := range infix {
		s := string(char)

		if IsOperand(s) {
			postfix = postfix + s
		} else if char == '(' {
			stack.Push(s)
		} else if char == ')' {
			for !stack.IsEmpty() && fmt.Sprint(stack.Peek()) != "(" {
				postfix = postfix + fmt.Sprint(stack.Pop())
			}
			stack.Pop()
		} else {
			for !stack.IsEmpty() && GetPriority(s) <= GetPriority(fmt.Sprint(stack.Peek())) {
				postfix = postfix + fmt.Sprint(stack.Peek())
				stack.Pop()
			}
			stack.Push(s)
		}
	}

	for !stack.IsEmpty() {
		if fmt.Sprint(stack.Peek()) == "(" {
			return "Invalid Expression"
		}
		postfix = postfix + fmt.Sprint(stack.Peek())
		stack.Pop()
	}
	return postfix
}

func GetPriority(c string) int {
	if c == "^" {
		return 3
	}
	if c == "*" || c == "/" || c == "%" {
		return 2
	}

	if c == "+" || c == "-" {
		return 1
	}
	return 0
}

func FormatExpression(s string) {
	rp := regexp.MustCompile("\\s*")
	rp.ReplaceAllString(s, "")
}

// [0-9|a-z|A-Z]
func IsOperand(s string) bool {
	match, _ := regexp.Match("[0-9|a-z|A-Z]", []byte(s))
	return match
}

// +-*/%
func IsOperator(s string) bool {
	match, _ := regexp.Match("\\^\\+|\\-|\\*|\\/|\\%", []byte(s))
	return match
}

func main() {
	s := "2+3*(2^3-5)^(2+1*2)-4"
	fmt.Println(InfixToPostfix(s))
}
