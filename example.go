package main

import (
	"fmt"
	"github.com/hakimamarullah/gocollections/golinkedlist"
	"github.com/hakimamarullah/gocollections/gostack"
)

func main() {
	stack := gostack.NewStack[string]()
	stack.Push("Halo")
	fmt.Println(stack)

	linkedList := golinkedlist.NewLinkedList[any]()
	linkedList.Add(1)
	linkedList.Add("Hello")
	linkedList.AddFirst("HAI")
	linkedList.Insert("Hei", 2)
	linkedList.Insert("Hei2", 2)
	linkedList.Remove("Hei2")
	_, err := linkedList.Insert("Hei2", 3)
	if err != nil {
		fmt.Println(err.Error())
	}
	linkedList.Remove("Hello")
	fmt.Println(linkedList.Contains("heii"))
	fmt.Println(linkedList.Size())
	fmt.Println(linkedList)
}

// Usage example
func balancedParentheses(token string) bool {
	stack := gostack.NewStack[string]()
	for i := 0; i < len(token); i++ {
		currentChar := token[i : i+1]
		if stack.IsEmpty() {
			stack.Push(currentChar)
			continue
		}
		top, _ := stack.Peek()
		if (top == "(" && currentChar == ")") ||
			(top == "{" && currentChar == "}") ||
			(top == "[" && currentChar == "]") {

			_, err := stack.Pop()
			if err != nil {
				return false
			}
		} else {
			stack.Push(currentChar)
		}
	}

	return stack.IsEmpty()
}
