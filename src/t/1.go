package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	max := 100
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2; i++ {
		n := rand.Intn(max)

		tmp := []string{}
		for j := 0; j < n; j++ {
			rand.Seed(time.Now().Unix() - int64(j))
			tmp = append(tmp, string(rand.Intn(max)))
		}
		fmt.Println("n =", n)
		fmt.Println("tmp =", tmp)
	}
}
