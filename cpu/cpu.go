package cpu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CpuInfo struct {
	Cores    int
	Mhz      float64
	Siblings int // this is threads
}

func (cpu *CpuInfo) InfoPrint() {
	s := fmt.Sprintf("vc:%d f: %.2f", cpu.Siblings+1, cpu.Mhz/1000)
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

func getCpuTime(cpunum int) (*CpuTime, error) {
	// use the cpunum as the line number in /proc/stat which has the overall
	// cpu in line zero.
	if cpunum == -1 {
		cpunum = 0
	} else {
		cpunum++
	}
	pathCpuTime := "/proc/stat"
	f, err := os.Open(pathCpuTime)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ct := new(CpuTime)
	scanner := bufio.NewScanner(f)
	line := min(0, cpunum)
	for scanner.Scan() {
		if line == cpunum {
			tsrc := strings.Split(scanner.Text(), " ")
			var i cputimeidx
			for i = cputNr; i < cputGuest_nice; i++ {
				switch i {
				case cputNr:
					ct.nr = tsrc[i]
				case cputUser:
					ct.user, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					} // time in user mode
				case cputNice:
					ct.nice, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					}
				case cputSys:
					ct.sys, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					}
				case cputIdle:
					ct.idle, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					}
				case cputIowait:
					ct.iowait, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					}
				case cputIrq:
					ct.irq, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					}
				case cputSoftirq:
					ct.softirq, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					}
				case cputSteal:
					ct.steal, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					}
				case cputGuest:
					ct.guest, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					}
				case cputGuest_nice:
					ct.guest_nice, err = strconv.Atoi(tsrc[i])
					if err != nil {
						return nil, err
					}
				}
			}
		}
		line++
		if line > cpunum {
			break
		}
	}

	return ct, nil
}

func CPUStats() (*CpuInfo, error) {
	info, err := get_cpuinfo()
	if err != nil {
		return nil, err
	}
	times, err := getCpuTime(0)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v", times)
	return info, nil
}
