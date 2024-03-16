package memory

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const MEMINFO_MAX = 57

// something in a comment
type Meminfo struct {
	MemTotal          int
	MemFree           int
	MemAvailable      int
	Buffers           int
	Cached            int
	SwapCached        int
	Active            int
	Inactive          int
	Active_anon       int
	Inactive_anon     int
	Active_file       int
	Inactive_file     int
	Unevictable       int
	Mlocked           int
	SwapTotal         int
	SwapFree          int
	Zswap             int
	Zswapped          int
	Dirty             int
	Writeback         int
	AnonPages         int
	Mapped            int
	Shmem             int
	KReclaimable      int
	Slab              int
	SReclaimable      int
	SUnreclaim        int
	KernelStack       int
	PageTables        int
	SecPageTables     int
	NFS_Unstable      int
	Bounce            int
	WritebackTmp      int
	CommitLimit       int
	Committed_AS      int
	VmallocTotal      int
	VmallocUsed       int
	VmallocChunk      int
	Percpu            int
	HardwareCorrupted int
	AnonHugePages     int
	ShmemHugePages    int
	ShmemPmdMapped    int
	FileHugePages     int
	FilePmdMapped     int
	CmaTotal          int
	CmaFree           int
	Unaccepted        int
	HugePages_Total   int
	HugePages_Free    int
	HugePages_Rsvd    int
	HugePages_Surp    int
	Hugepagesize      int
	Hugetlb           int
	DirectMap4k       int
	DirectMap2M       int
	DirectMap1G       int
}

type MemInfoFileLine int

const (
	MEMMemTotal MemInfoFileLine = iota
	MEMMemFree
	MEMMemAvailable
	MEMBuffers
	MEMCached
	MEMSwapCached
	MEMActive
	MEMInactive
	MEMActive_anon
	MEMInactive_anon
	MEMActive_file
	MEMInactive_file
	MEMUnevictable
	MEMMlocked
	MEMSwapTotal
	MEMSwapFree
	MEMZswap
	MEMZswapped
	MEMDirty
	MEMWriteback
	MEMAnonPages
	MEMMapped
	MEMShmem
	MEMKReclaimable
	MEMSlab
	MEMSReclaimable
	MEMSUnreclaim
	MEMKernelStack
	MEMPageTables
	MEMSecPageTables
	MEMNFS_Unstable
	MEMBounce
	MEMWritebackTmp
	MEMCommitLimit
	MEMCommitted_AS
	MEMVmallocTotal
	MEMVmallocUsed
	MEMVmallocChunk
	MEMPercpu
	MEMHardwareCorrupted
	MEMAnonHugePages
	MEMShmemHugePages
	MEMShmemPmdMapped
	MEMFileHugePages
	MEMFilePmdMapped
	MEMCmaTotal
	MEMCmaFree
	MEMUnaccepted
	MEMHugePages_Total
	MEMHugePages_Free
	MEMHugePages_Rsvd
	MEMHugePages_Surp
	MEMHugepagesize
	MEMHugetlb
	MEMDirectMap4k
	MEMDirectMap2M
	MEMDirectMap1G
)

var scale int

func init() {
	scale = 1024
}

func SetScale(v int) {
	scale = v
}

func (m *Meminfo) InfoPrint() string {
	// free/total cache buff
	scale := 1024
	s := fmt.Sprintf(
		"free/tot: %d/%d\t\tbuff/cache:%d/%d",
		m.MemFree/scale,
		m.MemTotal/scale,
		m.Buffers/scale,
		m.Cached/scale,
	)
	return s
}

func Getmeminfo() (*Meminfo, error) {
	memfile, err := os.Open("/proc/meminfo")
	// memfile, err := os.Open("./meminfo_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer memfile.Close()

	c := csv.NewReader(memfile)
	c.Comma = ':'
	m := new(Meminfo)
	var i MemInfoFileLine
	for i = MEMMemTotal; i < MEMINFO_MAX; i++ {
		rec, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error from csv read", err)
		}

		vw := strings.TrimSpace(strings.Fields(rec[1])[0])
		value, err := strconv.Atoi(vw)
		if err != nil {
			log.Fatal("can't convert value", err)
		}
		switch i {
		case MEMMemTotal:
			m.MemTotal = value
			if err != nil {
				log.Fatal("Unable to parse MemTotal", err)
			}
		case MEMMemFree:
			m.MemFree = value
			if err != nil {
				log.Fatal("Unable to parse MemFree", err)
			}
		case MEMMemAvailable:
			m.MemAvailable = value
			if err != nil {
				log.Fatal("Unable to parse MemAvailable", err)
			}
		case MEMBuffers:
			m.Buffers = value
			if err != nil {
				log.Fatal("Unable to parse Buffers", err)
			}
		case MEMCached:
			m.Cached = value
			if err != nil {
				log.Fatal("Unable to parse Cached", err)
			}
		case MEMSwapCached:
			m.SwapCached = value
			if err != nil {
				log.Fatal("Unable to parse SwapCached", err)
			}
		case MEMActive:
			m.Active = value
			if err != nil {
				log.Fatal("Unable to parse Active", err)
			}
		case MEMInactive:
			m.Inactive = value
			if err != nil {
				log.Fatal("Unable to parse Inactive", err)
			}
		case MEMActive_anon:
			m.Active_anon = value
			if err != nil {
				log.Fatal("Unable to parse Active_anon", err)
			}
		case MEMInactive_anon:
			m.Inactive_anon = value
			if err != nil {
				log.Fatal("Unable to parse Inactive_anon", err)
			}
		case MEMActive_file:
			m.Active_file = value
			if err != nil {
				log.Fatal("Unable to parse Active_file", err)
			}
		case MEMInactive_file:
			m.Inactive_file = value
			if err != nil {
				log.Fatal("Unable to parse Inactive_file", err)
			}
		case MEMUnevictable:
			m.Unevictable = value
			if err != nil {
				log.Fatal("Unable to parse Unevictable", err)
			}
		case MEMMlocked:
			m.Mlocked = value
			if err != nil {
				log.Fatal("Unable to parse Mlocked", err)
			}
		case MEMSwapTotal:
			m.SwapTotal = value
			if err != nil {
				log.Fatal("Unable to parse SwapTotal", err)
			}
		case MEMSwapFree:
			m.SwapFree = value
			if err != nil {
				log.Fatal("Unable to parse SwapFree", err)
			}
		case MEMZswap:
			m.Zswap = value
			if err != nil {
				log.Fatal("Unable to parse Zswap", err)
			}
		case MEMZswapped:
			m.Zswapped = value
			if err != nil {
				log.Fatal("Unable to parse Zswapped", err)
			}
		case MEMDirty:
			m.Dirty = value
			if err != nil {
				log.Fatal("Unable to parse Dirty", err)
			}
		case MEMWriteback:
			m.Writeback = value
			if err != nil {
				log.Fatal("Unable to parse Writeback", err)
			}
		case MEMAnonPages:
			m.AnonPages = value
			if err != nil {
				log.Fatal("Unable to parse AnonPages", err)
			}
		case MEMMapped:
			m.Mapped = value
			if err != nil {
				log.Fatal("Unable to parse Mapped", err)
			}
		case MEMShmem:
			m.Shmem = value
			if err != nil {
				log.Fatal("Unable to parse Shmem", err)
			}
		case MEMKReclaimable:
			m.KReclaimable = value
			if err != nil {
				log.Fatal("Unable to parse KReclaimable", err)
			}
		case MEMSlab:
			m.Slab = value
			if err != nil {
				log.Fatal("Unable to parse Slab", err)
			}
		case MEMSReclaimable:
			m.SReclaimable = value
			if err != nil {
				log.Fatal("Unable to parse SReclaimable", err)
			}
		case MEMSUnreclaim:
			m.SUnreclaim = value
			if err != nil {
				log.Fatal("Unable to parse SUnreclaim", err)
			}
		case MEMKernelStack:
			m.KernelStack = value
			if err != nil {
				log.Fatal("Unable to parse KernelStack", err)
			}
		case MEMPageTables:
			m.PageTables = value
			if err != nil {
				log.Fatal("Unable to parse PageTables", err)
			}
		case MEMSecPageTables:
			m.SecPageTables = value
			if err != nil {
				log.Fatal("Unable to parse SecPageTables", err)
			}
		case MEMNFS_Unstable:
			m.NFS_Unstable = value
			if err != nil {
				log.Fatal("Unable to parse NFS_Unstable", err)
			}
		case MEMBounce:
			m.Bounce = value
			if err != nil {
				log.Fatal("Unable to parse Bounce", err)
			}
		case MEMWritebackTmp:
			m.WritebackTmp = value
			if err != nil {
				log.Fatal("Unable to parse WritebackTemp", err)
			}
		case MEMCommitLimit:
			m.CommitLimit = value
			if err != nil {
				log.Fatal("Unable to parse CommitLimit", err)
			}
		case MEMCommitted_AS:
			m.Committed_AS = value
			if err != nil {
				log.Fatal("Unable to parse CommitLimit_AS", err)
			}
		case MEMVmallocTotal:
			m.VmallocTotal = value
			if err != nil {
				log.Fatal("Unable to parse VmallocTotal", err)
			}
		case MEMVmallocUsed:
			m.VmallocUsed = value
			if err != nil {
				log.Fatal("Unable to parse VmallocUsed", err)
			}
		case MEMVmallocChunk:
			m.VmallocChunk = value
			if err != nil {
				log.Fatal("Unable to parse VmallocChunk", err)
			}
		case MEMPercpu:
			m.Percpu = value
			if err != nil {
				log.Fatal("Unable to parse Percpu", err)
			}
		case MEMHardwareCorrupted:
			m.HardwareCorrupted = value
			if err != nil {
				log.Fatal("Unable to parse HardwareCorrupted", err)
			}
		case MEMAnonHugePages:
			m.AnonHugePages = value
			if err != nil {
				log.Fatal("Unable to parse AnonHugePages", err)
			}
		case MEMShmemHugePages:
			m.ShmemHugePages = value
			if err != nil {
				log.Fatal("Unable to parse ShmemHugePages", err)
			}
		case MEMShmemPmdMapped:
			m.ShmemPmdMapped = value
			if err != nil {
				log.Fatal("Unable to parse ShmemPmdMapped", err)
			}
		case MEMFileHugePages:
			m.FileHugePages = value
			if err != nil {
				log.Fatal("Unable to parse FileHugePages", err)
			}
		case MEMFilePmdMapped:
			m.FilePmdMapped = value
			if err != nil {
				log.Fatal("Unable to parse FilePmdMapped", err)
			}
		case MEMCmaTotal:
			m.CmaTotal = value
			if err != nil {
				log.Fatal("Unable to parse CmaTotal", err)
			}
		case MEMCmaFree:
			m.CmaFree = value
			if err != nil {
				log.Fatal("Unable to parse CmaFree", err)
			}
		case MEMUnaccepted:
			m.Unaccepted = value
			if err != nil {
				log.Fatal("Unable to parse Unaccepted", err)
			}
		case MEMHugePages_Total:
			m.HugePages_Total = value
			if err != nil {
				log.Fatal("Unable to parse HugePages_Total", err)
			}
		case MEMHugePages_Free:
			m.HugePages_Free = value
			if err != nil {
				log.Fatal("Unable to parse HugePages_Free", err)
			}
		case MEMHugePages_Rsvd:
			m.HugePages_Rsvd = value
			if err != nil {
				log.Fatal("Unable to parse HugePages_Rsvd", err)
			}
		case MEMHugePages_Surp:
			m.HugePages_Surp = value
			if err != nil {
				log.Fatal("Unable to parse HugePages_Surp", err)
			}
		case MEMHugepagesize:
			m.Hugepagesize = value
			if err != nil {
				log.Fatal("Unable to parse Hugepagesize", err)
			}
		case MEMHugetlb:
			m.Hugetlb = value
			if err != nil {
				log.Fatal("Unable to parse Hugetlb", err)
			}
		case MEMDirectMap4k:
			m.DirectMap4k = value
			if err != nil {
				log.Fatal("Unable to parse DirectMap4k", err)
			}
		case MEMDirectMap2M:
			m.DirectMap2M = value
			if err != nil {
				log.Fatal("Unable to parse DirectMap2M", err)
			}
		case MEMDirectMap1G:
			m.DirectMap1G = value
			if err != nil {
				log.Fatal("Unable to parse DirectMap1G", err)
			}
		default:
			log.Println("Unkown field, not parsing ", rec)
		}
	}

	return m, nil
}
