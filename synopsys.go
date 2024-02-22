package main

import (
	"fmt"
	"log"

	"github.com/bioe007/synopsys/cpu"
	"github.com/bioe007/synopsys/memory"
)

func main() {
	m, err := memory.Getmeminfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mem struct ${##v}", m)

	cpuinf, err := cpu.GetCPUStats()
	fmt.Println("cpu struct ${##v}", cpuinf)
}
