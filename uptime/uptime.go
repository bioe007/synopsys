package uptime

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// example
// 221671.25 3315800.64

const (
	uptimePath     = "/proc/uptime"
	secondsPerMin  = 60
	secondsPerHour = 3600
)

type Uptime struct {
	UptimeSeconds float64
	IdleSeconds   float64
}

func (ut *Uptime) HoursMinutes() {
	seconds := int64(ut.UptimeSeconds)
	hours := seconds / secondsPerHour
	seconds %= secondsPerHour
	minutes := seconds / secondsPerMin
	seconds %= secondsPerMin
	fmt.Printf("%d:%d:%d\n", hours, minutes, seconds)
}

func Read_uptime() (*Uptime, error) {
	ufile, err := os.ReadFile(uptimePath)
	if err != nil {
		return nil, err
	}
	sp := strings.Split(strings.TrimSuffix(string(ufile), "\n"), " ")

	ut := new(Uptime)
	ut.UptimeSeconds, err = strconv.ParseFloat(sp[0], 64)
	if err != nil {
		return nil, err
	}
	ut.IdleSeconds, err = strconv.ParseFloat(sp[1], 64)
	if err != nil {
		return nil, err
	}

	return ut, nil
}
