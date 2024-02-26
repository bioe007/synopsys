package cpu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tklauser/go-sysconf"
)

type CpuInfo struct {
	Cores    int
	Mhz      float64
	Siblings int // this is threads
	Stats    []*CpuTime
}

func (cpu *CpuInfo) InfoPrint() {
	clktck, err := sysconf.Sysconf(sysconf.SC_CLK_TCK)
	if err != nil {
		log.Fatal("error getting system clock", err)
	}
	s := fmt.Sprintf(
		"vc:%d f: %.2f usr:%d sys: %d: idle: %d",
		cpu.Siblings+1,
		cpu.Mhz/1000,
		int64(cpu.Stats[0].user)/clktck,
		int64(cpu.Stats[0].sys)/clktck,
		int64(cpu.Stats[0].idle)/clktck,
	)
	fmt.Println(s)
}

type cpuinfoidx int

const (
	cinf_processor = iota
	cinf_vendor_id
	cinf_cpu_family
	cinf_model
	cinf_model_name
	cinf_stepping
	cinf_microcode
	cinf_cpu_mhz
	cinf_cache_size
	cinf_physical_id
	cinf_siblings
	cinf_core_id
	cinf_cpu_cores
	cinf_apicid
	cinf_initial_apicid
	cinf_fpu
	cinf_fpu_exception
	cinf_cpuid_level
	cinf_wp
	cinf_flags
	cinf_bugs
	cinf_bogomips
	cinf_tlb_size
	cinf_clflush_size
	cinf_cache_alignment
	cinf_address_sizes
	cinf_power_management
	cinfcinf_power_management
)

type CpuTime struct {
	nr string // cpu number
	// times are in USER_HZ which is defined by sysconf(_SC_CLK_TCK)
	user       int // time in user mode
	nice       int
	sys        int
	idle       int
	iowait     int
	irq        int
	softirq    int
	steal      int
	guest      int
	guest_nice int
}

type cputimeidx int

const (
	cputNr   cputimeidx = iota
	cputUser            // time in user mode
	cputNice
	cputSys
	cputIdle
	cputIowait
	cputIrq
	cputSoftirq
	cputSteal
	cputGuest
	cputGuest_nice
)

func get_cpuinfo() (*CpuInfo, error) {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// f2, err := mmap.Open("/proc/cpuinfo")
	// if err != nil {
	// 	return nil, err
	// }
	// defer f2.Close()

	scanner := bufio.NewScanner(f)

	cpuinfo := new(CpuInfo)
	var line cpuinfoidx
	line = 0

	for scanner.Scan() {
		sp := strings.Split(scanner.Text(), ":")
		for i := range sp {
			sp[i] = strings.TrimSpace(sp[i])
		}
		switch line {
		case cinf_siblings:
			cpuinfo.Siblings, err = strconv.Atoi(sp[1])
			if err != nil {
				return nil, err
			}
		case cinf_cpu_mhz:
			cpuinfo.Mhz, err = strconv.ParseFloat(sp[1], 64)
			if err != nil {
				return nil, err
			}
		case cinf_cpu_cores:
			cpuinfo.Cores, err = strconv.Atoi(sp[1])
			if err != nil {
				return nil, err
			}
		}
		line++

		// unless there is a compelling reason, don't bother with rest of cpuinfo
		if line > cinf_cpu_cores {
			break
		}
	}

	return cpuinfo, nil
}

// Get cpunums stats. The -1 value is special and gets the overall stats
// For every cpunum add an entry to the returned slice of cputimes
func getCpuTime(numcpu int) ([]*CpuTime, error) {
	// The first line in stat is the overall CPU stats. We should make sure that's always in cpunums
	pathCpuTime := "/proc/stat"
	f, err := os.Open(pathCpuTime)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// t := regexp.MustCompile(`[ \t]`)
	var times []*CpuTime
	scanner := bufio.NewScanner(f)

	// For cputime struct only the first numcpu+1 lines have right content
	line := 0
	for scanner.Scan() {
		tsrc := strings.Fields(scanner.Text())
		times = append(times, new(CpuTime))
		var i cputimeidx
		// omg there has to be a better way
		for i = cputNr; i < cputGuest_nice; i++ {
			switch i {
			case cputNr:
				times[line].nr = tsrc[i]
			case cputUser:
				times[line].user, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			case cputNice:
				times[line].nice, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			case cputSys:
				times[line].sys, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			case cputIdle:
				times[line].idle, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			case cputIowait:
				times[line].iowait, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			case cputIrq:
				times[line].irq, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			case cputSoftirq:
				times[line].softirq, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			case cputSteal:
				times[line].steal, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			case cputGuest:
				times[line].guest, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			case cputGuest_nice:
				times[line].guest_nice, err = strconv.Atoi(tsrc[i])
				if err != nil {
					return nil, err
				}
			}
		}
		line++
		if line > numcpu {
			break
		}
	}

	return times, nil
}

func CPUStats() (*CpuInfo, error) {
	info, err := get_cpuinfo()
	if err != nil {
		return nil, err
	}

	info.Stats, err = getCpuTime(info.Siblings)
	if err != nil {
		return nil, err
	}
	return info, nil
}
