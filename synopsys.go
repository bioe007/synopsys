package main

import (
	"fmt"
	"log"

	"github.com/bioe007/synopsys/cpu"
	"github.com/bioe007/synopsys/load"
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

	load, err := load.GetLoadAvg()
	if err != nil {
		log.Fatal("Load average failure", err)
	}
	fmt.Println("Load average stuff", load)

}
