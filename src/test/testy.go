package test

import (
	"fmt"

	"github.com/alcionai/clues"
)

func Do() {
	fmt.Println("test - before -", clues.Conceal("hello"))
	clues.SetProtection(clues.NoProtection())
	fmt.Println("test - after -", clues.Conceal("hello"))
}
