package main

import (
	"fmt"
	"log"

	"github.com/bioe007/synopsys/memory"
)

func main() {
	m, err := memory.Getmeminfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mem struct ${##v}", m)
}
