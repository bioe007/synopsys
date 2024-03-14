package disk

import (
	"bufio"
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

type Stat struct {
	old    diskStat
	new    diskStat
	values statValues
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

type Disk struct {
	sick int
}

func diskparse(s string) (*DiskStat, error) {
	ds := new(DiskStat)

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
		}
	}

	return ds, nil
}

// Create an array od diskStat and return it
func DiskStats() ([]*diskStat, error) {
	var ds []*diskStat

	f, err := os.Open("/proc/diskstats")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for linenum := 0; scanner.Scan(); linenum++ {
		line := scanner.Text()
		curdisk, err := diskparse(line)
		ds = append(ds, curdisk)
		if err != nil {
			return nil, err
		}
	}
	return ds, nil
}
