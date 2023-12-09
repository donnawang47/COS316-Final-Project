package srtf

import (
	"testing"
)

// func Testsrtf(t *testing.T) {
func TestSRTF(t *testing.T) {
	srtf := NewSRTF()

	// testcase := []int{3, 2, -1, 5, 3, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1}
	testcase := []int{1, 1, 1, 1, 3, 1, -1, -1, -1}
	// testcase := []int{5, 5, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	for i := 0; i < len(testcase); i++ {
		//i is arrival time, testcase[i] is bursttime
		// completion time - arrival time
		// startime - arrival time
		if testcase[i] == -1 {
			srtf.run(nil, i)
		} else {
			srtf.run(&Process{id: i, arrivalTime: i, burstTime: testcase[i], waitingTime: -1}, i)
		}
		t.Errorf("%d", srtf.getProcess())

	}

	for i := 0; i < len(testcase); i++ {
		//i is arrival time, testcase[i] is bursttime
		// completion time - arrival time
		// startime - arrival time
		if testcase[i] != -1 {
			t.Errorf("%d, %d:", i, srtf.getProcessWaitingTime(i))
		}

	}

	t.Errorf("avg waiting time: %f", srtf.getAvgWaitingTIme())
}
