package cpu

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type CpuInfo struct {
	cores    int
	mhz      float64
	siblings int // this is threads
}

func get_cpuinfo() *CpuInfo {

	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cpuinfo := new(CpuInfo)
	for scanner.Scan() {
		sp := strings.Split(scanner.Text(), ":")
		for i := range sp {
			sp[i] = strings.TrimSpace(sp[i])
		}
		switch sp[0] {
		case "siblings":
			cpuinfo.siblings, err = strconv.Atoi(sp[1])
			if err != nil {
				log.Fatal(err)
			}
		case "cpu MHz":
			cpuinfo.mhz, err = strconv.ParseFloat(sp[1], 64)
			if err != nil {
				log.Fatal(err)
			}
		case "cpu cores":
			cpuinfo.cores, err = strconv.Atoi(sp[1])
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return cpuinfo

}

func GetCPUStats() (*CpuInfo, error) {
	c := get_cpuinfo()
	return c, nil
}
