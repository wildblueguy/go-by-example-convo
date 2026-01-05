package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Can customize tokenization with Scanner.Split (defaults to line-based)
	for scanner.Scan() {
	  text := scanner.Text()
	  fmt.Println(strings.ToUpper(text))
	}
	// Best practice is to call Scanner.Err after
}
