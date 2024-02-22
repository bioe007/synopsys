package memory

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const MEMINFO_MAX = 57

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
	MEMMemTotal MemInfoFileLine = 0 << iota
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
		switch rec[0] {
		case "MemTotal":
			m.MemTotal = value
			if err != nil {
				log.Fatal("Unable to parse MemTotal", err)
			}
		case "MemFree":
			m.MemFree = value
			if err != nil {
				log.Fatal("Unable to parse MemFree", err)
			}
		case "MemAvailable":
			m.MemAvailable = value
			if err != nil {
				log.Fatal("Unable to parse MemAvailable", err)
			}
		case "Buffers":
			m.Buffers = value
			if err != nil {
				log.Fatal("Unable to parse Buffers", err)
			}
		case "Cached":
			m.Cached = value
			if err != nil {
				log.Fatal("Unable to parse Cached", err)
			}
		case "SwapCached":
			m.SwapCached = value
			if err != nil {
				log.Fatal("Unable to parse SwapCached", err)
			}
		case "Active":
			m.Active = value
			if err != nil {
				log.Fatal("Unable to parse Active", err)
			}
		case "Inactive":
			m.Inactive = value
			if err != nil {
				log.Fatal("Unable to parse Inactive", err)
			}
		case "Active(anon)":
			m.Active_anon = value
			if err != nil {
				log.Fatal("Unable to parse Active_anon", err)
			}
		case "Inactive(anon)":
			m.Inactive_anon = value
			if err != nil {
				log.Fatal("Unable to parse Inactive_anon", err)
			}
		case "Active(file)":
			m.Active_file = value
			if err != nil {
				log.Fatal("Unable to parse Active_file", err)
			}
		case "Inactive(file)":
			m.Inactive_file = value
			if err != nil {
				log.Fatal("Unable to parse Inactive_file", err)
			}
		case "Unevictable":
			m.Unevictable = value
			if err != nil {
				log.Fatal("Unable to parse Unevictable", err)
			}
		case "Mlocked":
			m.Mlocked = value
			if err != nil {
				log.Fatal("Unable to parse Mlocked", err)
			}
		case "SwapTotal":
			m.SwapTotal = value
			if err != nil {
				log.Fatal("Unable to parse SwapTotal", err)
			}
		case "SwapFree":
			m.SwapFree = value
			if err != nil {
				log.Fatal("Unable to parse SwapFree", err)
			}
		case "Zswap":
			m.Zswap = value
			if err != nil {
				log.Fatal("Unable to parse Zswap", err)
			}
		case "Zswapped":
			m.Zswapped = value
			if err != nil {
				log.Fatal("Unable to parse Zswapped", err)
			}
		case "Dirty":
			m.Dirty = value
			if err != nil {
				log.Fatal("Unable to parse Dirty", err)
			}
		case "Writeback":
			m.Writeback = value
			if err != nil {
				log.Fatal("Unable to parse Writeback", err)
			}
		case "AnonPages":
			m.AnonPages = value
			if err != nil {
				log.Fatal("Unable to parse AnonPages", err)
			}
		case "Mapped":
			m.Mapped = value
			if err != nil {
				log.Fatal("Unable to parse Mapped", err)
			}
		case "Shmem":
			m.Shmem = value
			if err != nil {
				log.Fatal("Unable to parse Shmem", err)
			}
		case "KReclaimable":
			m.KReclaimable = value
			if err != nil {
				log.Fatal("Unable to parse KReclaimable", err)
			}
		case "Slab":
			m.Slab = value
			if err != nil {
				log.Fatal("Unable to parse Slab", err)
			}
		case "SReclaimable":
			m.SReclaimable = value
			if err != nil {
				log.Fatal("Unable to parse SReclaimable", err)
			}
		case "SUnreclaim":
			m.SUnreclaim = value
			if err != nil {
				log.Fatal("Unable to parse SUnreclaim", err)
			}
		case "KernelStack":
			m.KernelStack = value
			if err != nil {
				log.Fatal("Unable to parse KernelStack", err)
			}
		case "PageTables":
			m.PageTables = value
			if err != nil {
				log.Fatal("Unable to parse PageTables", err)
			}
		case "SecPageTables":
			m.SecPageTables = value
			if err != nil {
				log.Fatal("Unable to parse SecPageTables", err)
			}
		case "NFS_Unstable":
			m.NFS_Unstable = value
			if err != nil {
				log.Fatal("Unable to parse NFS_Unstable", err)
			}
		case "Bounce":
			m.Bounce = value
			if err != nil {
				log.Fatal("Unable to parse Bounce", err)
			}
		case "WritebackTmp":
			m.WritebackTmp = value
			if err != nil {
				log.Fatal("Unable to parse WritebackTemp", err)
			}
		case "CommitLimit":
			m.CommitLimit = value
			if err != nil {
				log.Fatal("Unable to parse CommitLimit", err)
			}
		case "Committed_AS":
			m.Committed_AS = value
			if err != nil {
				log.Fatal("Unable to parse CommitLimit_AS", err)
			}
		case "VmallocTotal":
			m.VmallocTotal = value
			if err != nil {
				log.Fatal("Unable to parse VmallocTotal", err)
			}
		case "VmallocUsed":
			m.VmallocUsed = value
			if err != nil {
				log.Fatal("Unable to parse VmallocUsed", err)
			}
		case "VmallocChunk":
			m.VmallocChunk = value
			if err != nil {
				log.Fatal("Unable to parse VmallocChunk", err)
			}
		case "Percpu":
			m.Percpu = value
			if err != nil {
				log.Fatal("Unable to parse Percpu", err)
			}
		case "HardwareCorrupted":
			m.HardwareCorrupted = value
			if err != nil {
				log.Fatal("Unable to parse HardwareCorrupted", err)
			}
		case "AnonHugePages":
			m.AnonHugePages = value
			if err != nil {
				log.Fatal("Unable to parse AnonHugePages", err)
			}
		case "ShmemHugePages":
			m.ShmemHugePages = value
			if err != nil {
				log.Fatal("Unable to parse ShmemHugePages", err)
			}
		case "ShmemPmdMapped":
			m.ShmemPmdMapped = value
			if err != nil {
				log.Fatal("Unable to parse ShmemPmdMapped", err)
			}
		case "FileHugePages":
			m.FileHugePages = value
			if err != nil {
				log.Fatal("Unable to parse FileHugePages", err)
			}
		case "FilePmdMapped":
			m.FilePmdMapped = value
			if err != nil {
				log.Fatal("Unable to parse FilePmdMapped", err)
			}
		case "CmaTotal":
			m.CmaTotal = value
			if err != nil {
				log.Fatal("Unable to parse CmaTotal", err)
			}
		case "CmaFree":
			m.CmaFree = value
			if err != nil {
				log.Fatal("Unable to parse CmaFree", err)
			}
		case "Unaccepted":
			m.Unaccepted = value
			if err != nil {
				log.Fatal("Unable to parse Unaccepted", err)
			}
		case "HugePages_Total":
			m.HugePages_Total = value
			if err != nil {
				log.Fatal("Unable to parse HugePages_Total", err)
			}
		case "HugePages_Free":
			m.HugePages_Free = value
			if err != nil {
				log.Fatal("Unable to parse HugePages_Free", err)
			}
		case "HugePages_Rsvd":
			m.HugePages_Rsvd = value
			if err != nil {
				log.Fatal("Unable to parse HugePages_Rsvd", err)
			}
		case "HugePages_Surp":
			m.HugePages_Surp = value
			if err != nil {
				log.Fatal("Unable to parse HugePages_Surp", err)
			}
		case "Hugepagesize":
			m.Hugepagesize = value
			if err != nil {
				log.Fatal("Unable to parse Hugepagesize", err)
			}
		case "Hugetlb":
			m.Hugetlb = value
			if err != nil {
				log.Fatal("Unable to parse Hugetlb", err)
			}
		case "DirectMap4k":
			m.DirectMap4k = value
			if err != nil {
				log.Fatal("Unable to parse DirectMap4k", err)
			}
		case "DirectMap2M":
			m.DirectMap2M = value
			if err != nil {
				log.Fatal("Unable to parse DirectMap2M", err)
			}
		case "DirectMap1G":
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
