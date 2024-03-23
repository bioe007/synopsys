package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bioe007/synopsys/cpu"
	"github.com/bioe007/synopsys/disk"
	"github.com/bioe007/synopsys/load"
	"github.com/bioe007/synopsys/memory"
	"github.com/bioe007/synopsys/uptime"
)

var scaleMap = map[rune]int{
	'k': 1024,
	'K': 1000,
	'm': 1024 * 1024,
	'M': 1000 * 1000,
	'g': 1024 * 1024 * 1024,
	'G': 1000 * 1000 * 1000,
	't': 1024 * 1024 * 1024 * 1024,
	'T': 1000 * 1000 * 1000 * 1000,
}

const usage = `Usage:
    synopsys [options]

Options:
    -i, --interval  [integer]   Duration in seconds between updates, default 1.
    -c, --cpu       [integer]   Max number of CPU you want to see output.
                                Default 8.
    -d, --disks     [integer]   Max number of disks you want to see output.
                                Default 8.
    -m, --memscale  [kKmMgGtT]  Units of memory to display, in kilo/Kibi etc.
                                Default is megabytes.
    -D, --disk-only             Show only disk activity
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n", usage)
	}

	var (
		num_disks, num_cpu, num_seconds int
		mem_scale                       string
		disk_only                       bool
	)
	flag.IntVar(&num_seconds, "interval", 1,
		"The number of seconds to wait between updates.")
	flag.IntVar(&num_seconds, "i", 1,
		"The number of seconds to wait between updates.")
	flag.IntVar(&num_cpu, "cpu", 8, "How many 'hot' CPU to display")
	flag.IntVar(&num_cpu, "c", 8, "How many 'hot' CPU to display")
	flag.IntVar(&num_disks, "disks", 8, "How many 'hot' CPU to display")
	flag.IntVar(&num_disks, "d", 8, "How many 'hot' CPU to display")
	flag.StringVar(&mem_scale, "memory", "m", "Choose how to scale memory")
	flag.StringVar(&mem_scale, "m", "m", "Choose how to scale memory")
	flag.BoolVar(&disk_only, "D", false, "Only show disk activity")
	flag.Parse()

	// TODO - parse this as an arg
	ms := []rune(mem_scale)
	memory.SetScale(scaleMap[ms[0]])

	ticker := time.NewTicker(time.Duration(num_seconds) * time.Second)
	var err error
	go func() {
		c := new(cpu.CpuInfo)
		disks := new(disk.DiskInfo)
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

			disks, err := disk.DiskStats(disks)
			if err != nil {
				log.Fatal("disk average failure", err)
			}

			ut, err := uptime.Read_uptime()
			if err != nil {
				log.Fatal("Unable to read uptime")
			}
			if !disk_only {
				fmt.Printf(
					"up:%s %s cpu:%s\nmem: %s\ndisks:%s\n",
					ut.HoursMinutes(),
					ld.InfoPrint(),
					// TODO - accept as parameter
					c.InfoPrint(num_cpu),
					m.InfoPrint(),
					disks.InfoPrint(num_disks),
				)
			} else {
				fmt.Printf("disks:\n%s\n", disks.InfoPrint(num_disks))
			}
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
