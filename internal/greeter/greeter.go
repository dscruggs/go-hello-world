package greeter

import (
	"fmt"

	"github.com/fatih/color"
)

func HelloWorld() {
	color.Cyan("Hello, World!") // Print in cyan color
	fmt.Println("Welcome to your first Go project!")
}
