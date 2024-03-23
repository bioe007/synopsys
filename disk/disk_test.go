package disk

import (
	"io"
	"os"
	"testing"
	"testing/fstest"
)

// var fs filesystem = osFS{}

type filesystem interface {
	Open(name string) (file, error)
	Stat(name string) (os.FileInfo, error)
}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

//	type mockFS struct {
//		osFS
//		reportErr bool
//	}
//
//	func (osFS) Open(name string) (file, error) {
//		if osFS.reportErr {
//			return nil, os.ErrNotExist
//		}
//		return os.Open(name)
//	}
//
// func (osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
func TestInfoPrint(t *testing.T) {
	zerodisk := diskStat{
		0, 0, "dev", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	onedisk := diskStat{
		1, 1, "dev", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}
	di := new(DiskInfo)
	di.new = make([]*diskStat, 2)
	di.old = make([]*diskStat, 2)
	di.new[0] = &onedisk
	di.new[1] = &onedisk
	di.old[0] = &zerodisk
	di.old[1] = &zerodisk
	di.estimate()

	s := di.InfoPrint(1)
	// if s != "dev 1" {
	if s != "dev 1" {
		t.Errorf("infoprint failed %s != dev 1", s)
	}
}

func TestDiskEstimateUnity(t *testing.T) {
	zerodisk := diskStat{
		0, 0, "dev", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	onedisk := diskStat{
		1, 1, "dev", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}
	di := new(DiskInfo)
	di.new = make([]*diskStat, 1)
	di.old = make([]*diskStat, 1)
	di.new[0] = &onedisk
	di.old[0] = &zerodisk

	di.estimate()
	if di.values.Len() != 1 {
		t.Errorf("not right")
	}
	v := di.values.Pop().(*statValues)
	if v.devname != "dev" {
		t.Errorf("not right")
	}
	if v.num_reads_completed != float32(1.0) {
		t.Errorf(
			"v.num_readi_completed wrong value: got %.2f, expected %.2f",
			v.num_reads_completed,
			float32(1.0),
		)
	}
	if v.num_reads_merged != float32(1.0) {
		t.Errorf(
			"v.num_reav_merged wrong value: got %.2f, expected %.2f",
			v.num_reads_merged,
			float32(1.0),
		)
	}
	if v.num_sectors_read != float32(1.0) {
		t.Errorf(
			"v.num_sectors_read wrong value: got %.2f, expected %.2f",
			v.num_sectors_read,
			float32(1.0),
		)
	}
	if v.ms_reading != float32(1.0) {
		t.Errorf(
			"v.ms_reading wrong value: got %.2f, expected %.2f",
			v.ms_reading,
			float32(1.0),
		)
	}
	if v.num_writes_completed != float32(1.0) {
		t.Errorf(
			"v.num_writes_completed wrong value: got %.2f, expected %.2f",
			v.num_writes_completed,
			float32(1.0),
		)
	}
	if v.num_writes_merged != float32(1.0) {
		t.Errorf(
			"v.num_writes_merged wrong value: got %.2f, expected %.2f",
			v.num_writes_merged,
			float32(1.0),
		)
	}
	if v.num_sectors_written != float32(1.0) {
		t.Errorf(
			"v.num_sectors_written wrong value: got %.2f, expected %.2f",
			v.num_sectors_written,
			float32(1.0),
		)
	}
	if v.ms_writing != float32(1.0) {
		t.Errorf(
			"v.ms_writing wrong value: got %.2f, expected %.2f",
			v.ms_writing,
			float32(1.0),
		)
	}
	if v.num_io_in_progress != float32(1.0) {
		t.Errorf(
			"v.num_io_in_progress wrong value: got %.2f, expected %.2f",
			v.num_io_in_progress,
			float32(1.0),
		)
	}
	if v.ms_doing_io != float32(1.0) {
		t.Errorf(
			"v.ms_doing_io wrong value: got %.2f, expected %.2f",
			v.ms_doing_io,
			float32(1.0),
		)
	}
	if v.ms_doing_io_weighted != float32(1.0) {
		t.Errorf(
			"v.ms_doing_io_weighted wrong value: got %.2f, expected %.2f",
			v.ms_doing_io_weighted,
			float32(1.0),
		)
	}
	if v.num_discards_completed != float32(1.0) {
		t.Errorf(
			"v.num_discarv_completed wrong value: got %.2f, expected %.2f",
			v.num_discards_completed,
			float32(1.0),
		)
	}
	if v.num_discards_merged != float32(1.0) {
		t.Errorf(
			"v.num_discarv_merged wrong value: got %.2f, expected %.2f",
			v.num_discards_merged,
			float32(1.0),
		)
	}
	if v.num_sectors_discarded != float32(1.0) {
		t.Errorf(
			"v.num_sectors_discarded wrong value: got %.2f, expected %.2f",
			v.num_sectors_discarded,
			float32(1.0),
		)
	}
	if v.ms_spent_discarding != float32(1.0) {
		t.Errorf(
			"v.ms_spent_discarding wrong value: got %.2f, expected %.2f",
			v.ms_spent_discarding,
			float32(1.0),
		)
	}
	if v.num_flush_requests_completed != float32(1.0) {
		t.Errorf(
			"v.num_flush_requests_completed  wrong value: got %.2f, expected %.2f",
			v.num_flush_requests_completed,
			float32(1.0),
		)
	}
	if v.ms_spent_flushing != float32(1.0) {
		t.Errorf(
			"v.ms_spent_flushing wrong value: got %.2f, expected %.2f",
			v.ms_spent_flushing,
			float32(1.0),
		)
	}
}

// Tests a single disk can be parsed
func TestDiskParse(t *testing.T) {
	s := "1       2 sda 3 4 5 6 7 8 9 10 11 12  13 14 15 16 17 18 19"
	ds, _ := diskparse(s)
	if ds.major != 1 {
		t.Errorf("got %d, wanted %d", ds.major, 1)
	}
	if ds.minor != 2 {
		t.Errorf("got %d, wanted %d", ds.minor, 2)
	}
	if ds.devname != "sda" {
		t.Errorf("got %q, wanted %q", ds.devname, "sda")
	}
	if ds.num_reads_completed != 3 {
		t.Errorf("got %q, wanted %q", ds.num_reads_completed, 3)
	}

	if ds.num_reads_merged != 4 {
		t.Errorf("got %d, wanted %d", ds.num_reads_merged, 4)
	}
	if ds.num_sectors_read != 5 {
		t.Errorf("got %d, wanted %d", ds.num_sectors_read, 5)
	}
	if ds.ms_reading != 6 {
		t.Errorf("got %d, wanted %d", ds.ms_reading, 6)
	}
	if ds.num_writes_completed != 7 {
		t.Errorf("got %d, wanted %d", ds.num_writes_completed, 7)
	}
	if ds.num_writes_merged != 8 {
		t.Errorf("got %d, wanted %d", ds.num_writes_merged, 8)
	}
	if ds.num_sectors_written != 9 {
		t.Errorf("got %d, wanted %d", ds.num_sectors_written, 9)
	}
	if ds.ms_writing != 10 {
		t.Errorf("got %d, wanted %d", ds.ms_writing, 10)
	}
	if ds.num_io_in_progress != 11 {
		t.Errorf("got %d, wanted %d", ds.num_io_in_progress, 11)
	}
	if ds.ms_doing_io != 12 {
		t.Errorf("got %d, wanted %d", ds.ms_doing_io, 12)
	}
	if ds.ms_doing_io_weighted != 13 {
		t.Errorf("got %d, wanted %d", ds.ms_doing_io_weighted, 13)
	}
	if ds.num_discards_completed != 14 {
		t.Errorf("got %d, wanted %d", ds.num_discards_completed, 14)
	}
	if ds.num_discards_merged != 15 {
		t.Errorf("got %d, wanted %d", ds.num_discards_merged, 15)
	}
	if ds.num_sectors_discarded != 16 {
		t.Errorf("got %d, wanted %d", ds.num_sectors_discarded, 16)
	}
	if ds.ms_spent_discarding != 17 {
		t.Errorf("got %d, wanted %d", ds.ms_spent_discarding, 17)
	}
	if ds.num_flush_requests_completed != 18 {
		t.Errorf("got %d, wanted %d", ds.num_flush_requests_completed, 18)
	}
	if ds.ms_spent_flushing != 19 {
		t.Errorf("got %d, wanted %d", ds.ms_spent_flushing, 19)
	}
}

func TestDiskParseFailsNonNumeric(t *testing.T) {
	s := "1       a sda 3 4 5 6 7 8 9 10 11 12  13 14 15 16 17 18 19"
	ds, _ := diskparse(s)

	ds, err := diskparse(s)
	if err == nil {
		t.Errorf("should have caught nonnumeric input: %+x", ds)
	}
}

func TestGetDiskStats(t *testing.T) {
	di := new(DiskInfo)

	FILES := fstest.MapFS{
		"diskstats": {
			Data: []byte(
				`0 0 dev0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
    1 1 dev1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1`),
		},
	}

	f, _ := FILES.Open("diskstats")
	di2, err := getDiskStats(di, f)
	if err != nil {
		t.Fatal(err)
	}
	if di2.new[0].devname != "dev0" {
		t.Errorf("got di2[0].devname: %s", di2.new[0].devname) // di2.new[0].devname)
	}
	if di2.new[1].devname != "dev1" {
		t.Errorf("got di2[1].devname: %s", di2.new[1].devname) // di2.new[1].devname)
	}
}
