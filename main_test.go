package main

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	fmt.Println("Password: " + generatePassword(20))
}
