package srtf

import (
	"fmt"
	"math/rand"
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

	t.Errorf("avg waiting time: %f", srtf.getAvgWaitingTime())
}

func TestSRTFRand(t *testing.T) {
	numTestcases := 10
	avgWaitingTime := float32(0)
	avgOptWaitingTime := float32(0)
	for i := 0; i < numTestcases; i++ {
		srtf := NewSRTF()
		srtf_opt := NewSRTF_OPT(150)
		// 100 jobs total
		//burst time range 100, waiting time in between jobs
		for j := 0; j < 3; j++ {
			//waiting time
			waitingTime := rand.Intn(100)
			//fmt.Println("waiting time", waitingTime)
			for k := 0; k < waitingTime; k++ {
				srtf.run(nil, srtf.clockTime)
				srtf_opt.run(nil, srtf_opt.clockTime)
			}
			//burst time

			srtf.run(&Process{id: j, arrivalTime: srtf.clockTime, burstTime: rand.Intn(100) + 1}, srtf.clockTime)
			srtf_opt.run(&Process_Opt{id: j, arrivalTime: srtf_opt.clockTime, burstTime: rand.Intn(100) + 1}, srtf_opt.clockTime)
			//fmt.Println("process", j)

		}

		for srtf.processId != -1 {
			srtf.run(nil, srtf.clockTime)
		}
		for srtf_opt.processId != -1 {
			srtf_opt.run(nil, srtf_opt.clockTime)
		}

		avgWaitingTime += srtf.getAvgWaitingTime()
		avgOptWaitingTime += srtf_opt.getAvgWaitingTime()
	}
	fmt.Println(avgOptWaitingTime)

	t.Errorf("avg waiting time: %f", avgWaitingTime/float32(numTestcases))
	t.Errorf("avg opt waiting time: %f", avgOptWaitingTime/float32(numTestcases))
}
