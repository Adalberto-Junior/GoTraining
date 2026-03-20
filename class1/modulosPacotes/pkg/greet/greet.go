package greet

import (
	"fmt"
)

func Hello(name string) string {
	if name == "" {
		name = "mundo"
	}

	return fmt.Sprintf("Olá, %s!", name)
}
