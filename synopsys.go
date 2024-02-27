package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bioe007/synopsys/cpu"
	"github.com/bioe007/synopsys/disk"
	"github.com/bioe007/synopsys/load"
	"github.com/bioe007/synopsys/memory"
	"github.com/bioe007/synopsys/uptime"
)

// main stuff getting done
func main() {
	// Argument(s) is only an int to agjust frequency
	var num_seconds int
	var err error
	if len(os.Args) < 2 {
		num_seconds = 1
	} else {
		num_seconds, err = strconv.Atoi(os.Args[1])
		if err != nil {
			usage()
		}
	}

	ticker := time.NewTicker(time.Duration(num_seconds) * time.Second)
	go func() {
		c := new(cpu.CpuInfo)
		for ; ; <-ticker.C {
			c, err = cpu.CPUStats(c)
			if err != nil {
				log.Fatal(err)
			}

			m, err := memory.Getmeminfo()
			if err != nil {
				log.Fatal(err)
			}

			ld, err := load.LoadAvg()
			if err != nil {
				log.Fatal("Load average failure", err)
			}

			disks, err := disk.DiskStats()
			if err != nil {
				log.Fatal("disk average failure", err)
			}
			fmt.Println("Got disks: ", len(disks))

			ut, err := uptime.Read_uptime()
			if err != nil {
				log.Fatal("Unable to read uptime")
			}
			fmt.Printf(
				"up:%s %s cpu:%s\nmem: %s\n",
				ut.HoursMinutes(),
				ld.InfoPrint(),
				c.InfoPrint(),
				m.InfoPrint(),
			)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	// Wait until getting SIGINT or SIGTERM
	go func() {
		_ = <-sigs
		done <- true
	}()
	<-done
}

func usage() {
	fmt.Printf("Usage: %s [<interval>]", os.Args[0])
	os.Exit(0)
}
