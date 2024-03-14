package cpu

import (
	"container/heap"
	"testing"
)

func TestEstimate(t *testing.T) {
	t.Error("not implemented")
}

func TestInfoPrint(t *testing.T) {
	t.Error("not implemented")
}

func TestGetCPUInfo(t *testing.T) {
	t.Error("not implemented")
}

func TestGetCPUTime(t *testing.T) {
	t.Error("not implemented")
}

func TestCPUStats(t *testing.T) {
	t.Error("not implemented")
}

func TestCPUHeapPopEmpty(t *testing.T) {
	t.Error("not implemented")
}

func TestCPUHeapOrder(t *testing.T) {
	c1 := new(CpuStat)
	c1.nr = "1"
	c1.irq = 0.0
	// currently the ordering is hard-coded to use user time
	c1.user = 10.0

	c2 := new(CpuStat)
	c2.nr = "2"
	c2.irq = 0.0
	c2.user = 1.0

	c3 := new(CpuStat)
	c3.nr = "3"
	c3.irq = 0.0
	c3.user = 2.0

	cstats := new(calculatedstats)
	heap.Init(cstats)
	heap.Push(cstats, c2)
	heap.Push(cstats, c1)
	heap.Push(cstats, c3)

	c := cstats.Pop().(*CpuStat)
	if c.nr != "1" {
		t.Errorf("heap order failure: expected 1, got %s", c.nr)
	}
}
