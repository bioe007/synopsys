package cpu

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// For /proc/stat field order
type cputimeidx int

const (
	cputNr cputimeidx = iota
	cputUser
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

// Used to store calculated fractional values
type CpuStat struct {
	nr         string // cpu number
	user       float32
	nice       float32
	sys        float32
	idle       float32
	iowait     float32
	irq        float32
	softirq    float32
	steal      float32
	guest      float32
	guest_nice float32
}

// Lines in the /proc/cpuinfo file
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

// Holds all cpu information including previous, current counts and estimated values for percent time spent in each.
type CpuInfo struct {
	Cores     int
	Mhz       float64
	Siblings  int // this is threads
	Stats     []*CpuTime
	OldStats  []*CpuTime
	calcstats *calculatedstats
}

// This will be a heap to quickly get cpu sorted by user time
type calculatedstats []*CpuStat

func (h calculatedstats) Len() int { return len(h) }

// TODO: use configurable method
func (h calculatedstats) Less(i, j int) bool { return h[i].user > h[j].user }
func (h calculatedstats) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *calculatedstats) Push(x any)        { *h = append(*h, x.(*CpuStat)) }
func (h *calculatedstats) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (cpu *CpuTime) total() int {
	return cpu.user + cpu.nice + cpu.sys + cpu.idle + cpu.iowait + cpu.irq + cpu.softirq + cpu.steal + cpu.guest + cpu.guest_nice
}

func (cpu *CpuInfo) estimate() {
	if len(cpu.OldStats) == 0 {
		return
	}
	prev := cpu.OldStats
	cur := cpu.Stats

	cpu.calcstats = new(calculatedstats)
	heap.Init(cpu.calcstats)

	for i := range cur {
		c := new(CpuStat)
		denom := float32(cur[i].total() - prev[i].total())
		c.nr = cur[i].nr
		c.user = float32((cur[i].user - prev[i].user)) / denom
		c.sys = float32((cur[i].sys - prev[i].sys)) / denom
		c.idle = float32((cur[i].idle - prev[i].idle)) / denom
		c.iowait = float32((cur[i].iowait - prev[i].iowait)) / denom
		c.irq = float32((cur[i].irq - prev[i].irq)) / denom
		c.softirq = float32((cur[i].softirq - prev[i].softirq)) / denom
		c.steal = float32((cur[i].steal - prev[i].steal)) / denom
		c.guest = float32((cur[i].guest - prev[i].guest)) / denom
		c.guest_nice = float32((cur[i].guest_nice - prev[i].guest_nice)) / denom
		heap.Push(cpu.calcstats, c)
	}
}

func (cpu *CpuInfo) InfoPrint(num_cpus int) string {
	// clktck, err := sysconf.Sysconf(sysconf.SC_CLK_TCK)
	// if err != nil {
	// 	log.Fatal("error getting system clock", err)
	// }
	cpu.estimate()
	if len(cpu.OldStats) == 0 {
		// TODO: - be smarter than just skipping it the first time through?
		return ""
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("vc:%d\tf: %.2f\n", cpu.Siblings, cpu.Mhz/1000))

	for i := 0; i < num_cpus; i++ {
		c := heap.Pop(cpu.calcstats).(*CpuStat)
		sb.WriteString(
			fmt.Sprintf(
				"%s: usr:%.2f sys:%.2f: idle:%.2f iowait:%.2f irq:%.2f softirq:%.2f steal:%.2f guest:%.2f gnice:%.2f\n",
				c.nr,
				c.user,
				c.sys,
				c.idle,
				c.iowait,
				c.irq,
				c.softirq,
				c.steal,
				c.guest,
				c.guest_nice,
			))
	}
	return sb.String()
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

	var times []*CpuTime
	scanner := bufio.NewScanner(f)

	// For cputime struct only the first numcpu+1 lines have right content
	line := 0
	for scanner.Scan() {
		tsrc := strings.Fields(scanner.Text())
		times = append(times, new(CpuTime))
		var i cputimeidx
		// TODO: omg there has to be a better way
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

func CPUStats(ci *CpuInfo) (*CpuInfo, error) {
	info, err := get_cpuinfo()
	if err != nil {
		return nil, err
	}

	info.OldStats = ci.Stats

	info.Stats, err = getCpuTime(info.Siblings)
	if err != nil {
		return nil, err
	}
	return info, nil
}
