package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	lower := "abcdefghijklmnopqrstvuwxyz"
	upper := strings.ToUpper(lower)
	symbols := "!@#$%^&*(){}_"
	number := "1234567890"
	all := lower + upper + symbols + number

	length := 12

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	password := ""

	for i := 0; i < length; i++ {
		password += string(all[random.Intn(len(all))])
	}
	fmt.Println(password)
}
