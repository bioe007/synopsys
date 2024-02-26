package main

import (
	"fmt"
	"log"

	"github.com/bioe007/synopsys/cpu"
	"github.com/bioe007/synopsys/disk"
	"github.com/bioe007/synopsys/load"
	"github.com/bioe007/synopsys/memory"
	"github.com/bioe007/synopsys/uptime"
)

// main stuff getting done
func main() {
	c, err := cpu.CPUStats()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("CPU %+v\n", c)
	c.InfoPrint()

	m, err := memory.Getmeminfo()
	if err != nil {
		log.Fatal(err)
	}
	m.InfoPrint()

	ld, err := load.LoadAvg()
	if err != nil {
		log.Fatal("Load average failure", err)
	}
	fmt.Printf("LOAD %+v\n", ld)

	disks, err := disk.DiskStats()
	if err != nil {
		log.Fatal("disk average failure", err)
	}
	fmt.Println("Got disks: ", len(disks))

	ut, err := uptime.Read_uptime()
	if err != nil {
		log.Fatal("Unable to read uptime")
	}
	fmt.Printf("uptime: %+v\n", ut)
	// uptime.HoursMinutes(ut)
	ut.HoursMinutes()
}
