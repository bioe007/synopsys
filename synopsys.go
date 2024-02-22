package main

import (
	"fmt"
	"log"

	"github.com/bioe007/synopsys/cpu"
	"github.com/bioe007/synopsys/disk"
	"github.com/bioe007/synopsys/load"
	"github.com/bioe007/synopsys/memory"
)

func main() {
	m, err := memory.Getmeminfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MEMORY %+v\n", m)

	cpuinf, err := cpu.GetCPUStats()
	fmt.Printf("CPU %+v\n", cpuinf)

	load, err := load.GetLoadAvg()
	if err != nil {
		log.Fatal("Load average failure", err)
	}
	fmt.Printf("LOAD %+v\n", load)

	disks, err := disk.GetDiskStats()
	if err != nil {
		log.Fatal("disk average failure", err)
	}
	fmt.Println("Got disks: ", len(disks))

}
