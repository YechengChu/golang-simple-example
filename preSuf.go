package main

import (
	"fmt"
	"strings"
)

func main() {
	var testStr = "This is a test string"
	fmt.Printf("Does the string \"%s\" has prefix \"Th\"?\n", testStr)
	fmt.Printf("%t\n", strings.HasPrefix(testStr, "Th"))
	fmt.Printf("Does the string \"%s\" has prefix \"th\"?\n", testStr)
	fmt.Printf("%t\n", strings.HasPrefix(testStr, "th"))
	fmt.Printf("Does the string \"%s\" has suffix \"ng\"?\n", testStr)
	fmt.Printf("%t\n", strings.HasSuffix(testStr, "ng"))
	fmt.Printf("Does the string \"%s\" contains \"es\"?\n", testStr)
	fmt.Printf("%t\n", strings.Contains(testStr, "es"))
} // main
