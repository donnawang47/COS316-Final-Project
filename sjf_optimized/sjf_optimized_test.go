package sjf_optimized

import (
	"testing"
)

// func TestSJF(t *testing.T) {
func TestSJF_Optimized(t *testing.T) {
	sjf := NewSJF()

	// testcase := []int{3, 2, -1, 5, 3, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1}
	// processes w smaller bursts keep coming in
	testcase := []int{5, 5, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	for i := 0; i < len(testcase); i++ {
		//i is arrival time, testcase[i] is bursttime
		// completion time - arrival time
		// startime - arrival time
		if testcase[i] == -1 {
			sjf.run(nil, i)
		} else {
			sjf.run(&Process{id: i, arrivalTime: i, burstTime: testcase[i]}, i)
		}
		t.Errorf("%d", sjf.getProcess())

	}

	for i := 0; i < len(testcase); i++ {
		//i is arrival time, testcase[i] is bursttime
		// completion time - arrival time
		// startime - arrival time
		if testcase[i] != -1 {
			t.Errorf("%d, %d:", i, sjf.getProcessWaitingTime(i))
		}

	}

	t.Errorf("avg waiting time: %f", sjf.getAvgWaitingTIme())
}
