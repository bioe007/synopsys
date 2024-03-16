package load

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Load struct {
	one          float64
	five         float64
	fifteen      float64
	proc_running int
	proc_total   int
	lastpid      int
}

type LA_CONST int

const (
	LA_ONEMIN LA_CONST = iota
	LA_FIVEMIN
	LA_FIFTEENMIN
	// Remember LA_PROC_TOTAL is not a separate entry when parsing.
	LA_PROC_RUN
	LA_LASTPID
)

func (ld *Load) InfoPrint() string {
	s := fmt.Sprintf(
		"r/t: %d/%d\tla: %.2f,%.2f,%.2f",
		ld.proc_running,
		ld.proc_total,
		ld.one,
		ld.five,
		ld.fifteen,
	)
	return s
}

func LoadAvg() (*Load, error) {
	f, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		log.Fatal(err)
	}
	loadinfo := new(Load)
	sp := strings.Split(strings.TrimSuffix(string(f), "\n"), " ")

	// The running/total process entries are not space-delimited so parse those
	// a little different
	var i LA_CONST
	for i = LA_ONEMIN; i <= LA_LASTPID+1; i++ {
		switch i {
		case LA_ONEMIN:
			loadinfo.one, err = strconv.ParseFloat(sp[i], 64)
			if err != nil {
				return nil, err
			}
		case LA_FIVEMIN:
			loadinfo.five, err = strconv.ParseFloat(sp[i], 64)
			if err != nil {
				return nil, err
			}
		case LA_FIFTEENMIN:
			loadinfo.fifteen, err = strconv.ParseFloat(sp[i], 64)
			if err != nil {
				return nil, err
			}
		case LA_PROC_RUN:
			// These are not space delimited but shown like X/Y in the loadavg file
			procs := strings.Split(sp[i], "/")
			loadinfo.proc_running, err = strconv.Atoi(procs[0])
			if err != nil {
				return nil, err
			}
			loadinfo.proc_total, err = strconv.Atoi(procs[1])
			if err != nil {
				return nil, err
			}
		case LA_LASTPID:
			loadinfo.lastpid, err = strconv.Atoi(sp[i])
			if err != nil {
				return nil, err
			}
		}
	}
	return loadinfo, nil
}
