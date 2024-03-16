package disk

import (
	"bufio"
	"container/heap"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

type diskStat struct {
	major               int
	minor               int
	devname             string
	num_reads_completed int // This is the total number of reads completed successfully.
	num_reads_merged    int // , field 6 -- # of writes merged (unsigned long)
	// Reads and writes which are adjacent to each other may be merged for efficiency. Thus two 4K reads may become one 8K read before it is ultimately handed to the disk, and so it will be counted (and queued) as only one I/O. This field lets you know how often this was done.
	num_sectors_read int // This is the total number of sectors read successfully.

	ms_reading           int // This is the total number of milliseconds spent by all reads (as measured from blk_mq_alloc_request() to __blk_mq_end_request()).
	num_writes_completed int // This is the total number of writes completed successfully.
	num_writes_merged    int // See the description of field 2.
	num_sectors_written  int // This is the total number of sectors written successfully.
	ms_writing           int // This is the total number of milliseconds spent by all writes (as measured from blk_mq_alloc_request() to __blk_mq_end_request()).
	num_io_in_progress   int // The only field that should go to zero. Incremented as requests are given to appropriate struct request_queue and decremented as they finish.
	ms_doing_io          int // This field increases so long as field 9 is nonzero.
	// Since 5.0 this field counts jiffies when at least one request was started or completed. If request runs more than 2 jiffies then some I/O time might be not accounted in case of concurrent requests.
	ms_doing_io_weighted   int // This field is incremented at each I/O start, I/O completion, I/O merge, or read of these stats by the number of I/Os in progress (field 9) times the number of milliseconds spent doing I/O since the last update of this field. This can provide an easy measure of both I/O completion time and the backlog that may be accumulating.
	num_discards_completed int // This is the total number of discards completed successfully.
	num_discards_merged    int // See the description of field 2
	num_sectors_discarded  int // This is the total number of sectors discarded successfully.
	ms_spent_discarding    int // This is the total number of milliseconds spent by
	// all discards (as measured from blk_mq_alloc_request() to
	// __blk_mq_end_request()).
	num_flush_requests_completed int // This is the total number of flush requests completed successfully.
	// Block layer combines flush requests and
	// executes at most one at a time. This
	// counts flush requests executed by disk.
	// Not tracked for partitions.

	ms_spent_flushing int // This is the total number of milliseconds spent by all flush requests.
}

type statValues struct {
	major                        float32
	minor                        float32
	devname                      string
	num_reads_completed          float32
	num_reads_merged             float32
	num_sectors_read             float32
	ms_reading                   float32
	num_writes_completed         float32
	num_writes_merged            float32
	num_sectors_written          float32
	ms_writing                   float32
	num_io_in_progress           float32
	ms_doing_io                  float32
	ms_doing_io_weighted         float32
	num_discards_completed       float32
	num_discards_merged          float32
	num_sectors_discarded        float32
	ms_spent_discarding          float32
	num_flush_requests_completed float32
	ms_spent_flushing            float32
}

type diskHeap []*statValues

func (h diskHeap) Len() int { return len(h) }
func (h diskHeap) Less(
	i, j int,
) bool {
	return h[i].num_writes_completed > h[j].num_writes_completed
}
func (h diskHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *diskHeap) Push(x any)   { *h = append(*h, x.(*statValues)) }
func (h *diskHeap) Pop() any {
	old := *h
	n := len(old)
	if n == 0 {
		return nil
	}
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type DiskInfo struct {
	old    []*diskStat
	new    []*diskStat
	values *diskHeap
}

type dsfields int

const (
	DSFMAJOR dsfields = iota
	DSFMINOR
	DSFNAME
	DSFNUM_READS_COMPLETED
	DSFNUM_READS_MERGED
	DSFNUM_SECTORS_READ
	DSFMS_READING
	DSFNUM_WRITES_COMPLETED
	DSFNUM_WRITES_MERGED
	DSFNUM_SECTORS_WRITTEN
	DSFMS_WRITING
	DSFNUM_IO_IN_PROGRESS
	DSFMS_DOING_IO
	DSFMS_DOING_IO_WEIGHTED
	DSFNUM_DISCARDS_COMPLETED
	DSFNUM_DISCARDS_MERGED
	DSFNUM_SECTORS_DISCARDED
	DSFMS_SPENT_DISCARDING
	DSFNUM_FLUSH_REQUESTS_COMPLETED
	DSFMS_SPENT_FLUSHING
)

const diskstats = "/proc/diskstats"

func getDiskStatPath() string {
	return diskstats
}

func (disks *DiskInfo) estimate() {
	if len(disks.old) == 0 {
		return
	}

	prev := disks.old
	cur := disks.new

	disks.values = new(diskHeap)
	heap.Init(disks.values)

	for i := range cur {
		d := new(statValues)
		d.devname = cur[i].devname
		d.major = float32(cur[i].major)
		d.minor = float32(cur[i].minor)
		d.num_reads_completed = float32(cur[i].num_reads_completed - prev[i].num_reads_completed)
		d.num_reads_merged = float32(cur[i].num_reads_merged - prev[i].num_reads_merged)
		d.num_sectors_read = float32(cur[i].num_sectors_read - prev[i].num_sectors_read)
		d.ms_reading = float32(cur[i].ms_reading - prev[i].ms_reading)
		d.num_writes_completed = float32(cur[i].num_writes_completed - prev[i].num_writes_completed)
		d.num_writes_merged = float32(cur[i].num_writes_merged - prev[i].num_writes_merged)
		d.num_sectors_written = float32(cur[i].num_sectors_written - prev[i].num_sectors_written)
		d.ms_writing = float32(cur[i].ms_writing - prev[i].ms_writing)
		d.num_io_in_progress = float32(cur[i].num_io_in_progress - prev[i].num_io_in_progress)
		d.ms_doing_io = float32(cur[i].ms_doing_io - prev[i].ms_doing_io)
		d.ms_doing_io_weighted = float32(cur[i].ms_doing_io_weighted - prev[i].ms_doing_io_weighted)
		d.num_discards_completed = float32(
			cur[i].num_discards_completed - prev[i].num_discards_completed,
		)
		d.num_discards_merged = float32(cur[i].num_discards_merged - prev[i].num_discards_merged)
		d.num_sectors_discarded = float32(
			cur[i].num_sectors_discarded - prev[i].num_sectors_discarded,
		)
		d.ms_spent_discarding = float32(cur[i].ms_spent_discarding - prev[i].ms_spent_discarding)
		d.num_flush_requests_completed = float32(
			cur[i].num_flush_requests_completed - prev[i].num_flush_requests_completed,
		)
		d.ms_spent_flushing = float32(cur[i].ms_spent_flushing - prev[i].ms_spent_flushing)
		heap.Push(disks.values, d)
	}
}

func diskparse(s string) (*diskStat, error) {
	ds := new(diskStat)

	fields := strings.Fields(s)

	var fieldnum dsfields
	var err error
	for fieldnum = DSFMAJOR; fieldnum < DSFMS_SPENT_FLUSHING+1; fieldnum++ {
		switch fieldnum {
		case DSFMAJOR:
			ds.major, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFMINOR:
			ds.minor, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNAME:
			ds.devname = fields[fieldnum]
		case DSFNUM_READS_COMPLETED:
			ds.num_reads_completed, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_READS_MERGED:
			ds.num_reads_merged, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_SECTORS_READ:
			ds.num_sectors_read, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFMS_READING:
			ds.ms_reading, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_WRITES_COMPLETED:
			ds.num_writes_completed, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_WRITES_MERGED:
			ds.num_writes_merged, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_SECTORS_WRITTEN:
			ds.num_sectors_written, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFMS_WRITING:
			ds.ms_writing, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_IO_IN_PROGRESS:
			ds.num_io_in_progress, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFMS_DOING_IO:
			ds.ms_doing_io, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFMS_DOING_IO_WEIGHTED:
			ds.ms_doing_io_weighted, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_DISCARDS_COMPLETED:
			ds.num_discards_completed, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_DISCARDS_MERGED:
			ds.num_discards_merged, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_SECTORS_DISCARDED:
			ds.num_sectors_discarded, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFMS_SPENT_DISCARDING:
			ds.ms_spent_discarding, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFNUM_FLUSH_REQUESTS_COMPLETED:
			ds.num_flush_requests_completed, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		case DSFMS_SPENT_FLUSHING:
			ds.ms_spent_flushing, err = strconv.Atoi(fields[fieldnum])
			if err != nil {
				return nil, err
			}
		}
	}

	return ds, nil
}

// Get a diskinfo and update it with new stats
func DiskStats(di *DiskInfo) (*DiskInfo, error) {
	f, err := os.Open(getDiskStatPath())
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return getDiskStats(di, f)
}

func getDiskStats(di *DiskInfo, f fs.File) (*DiskInfo, error) {
	var ds []*diskStat

	di.old = di.new
	scanner := bufio.NewScanner(f)
	for linenum := 0; scanner.Scan(); linenum++ {
		line := scanner.Text()
		curdisk, err := diskparse(line)
		ds = append(ds, curdisk)
		if err != nil {
			return nil, err
		}
	}
	di.new = ds
	di.estimate()
	return di, nil
}

func (disks *DiskInfo) InfoPrint(num_disks int) string {
	if len(disks.old) == 0 {
		return ""
	}
	total_disks := disks.values.Len()
	disk_limit := max(min(total_disks, num_disks), 0)

	if num_disks == 0 {
		return "zero"
	}

	var sb strings.Builder
	for i := 0; i < disk_limit; i++ {
		disk := heap.Pop(disks.values).(*statValues)
		sb.WriteString(
			fmt.Sprintf("%s %.0f\t",
				disk.devname,
				disk.num_writes_completed,
			))
	}

	return sb.String()
}
