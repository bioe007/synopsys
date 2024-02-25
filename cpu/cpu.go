package cpu

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type CpuInfo struct {
	Cores    int
	Mhz      float64
	Siblings int // this is threads
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
	nr int // cpu number
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
	ct := new(CpuTime)

	return ct, nil
}

func CPUStats() (*CpuInfo, error) {
	info, err := get_cpuinfo()
	if err != nil {
		return nil, err
	}
	return info, nil
}
