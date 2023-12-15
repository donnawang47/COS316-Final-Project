package srtf

import (
	"math/rand"
	"testing"
)

// test used to verify that algorithm returns expected waiting time for each individual process and each process
// is run in the correct order
func TestSJF(t *testing.T) {
	srtf := NewSRTF()
	srtf_opt := NewSRTF_OPT(3) // threshold for priority (aging mechanism) = 3

	var padding_time []int
	// negative numbers represent 1 unit time of running scheduler with no new job added to it
	for i := 0; i < 22; i++ {
		padding_time = append(padding_time, -1)
	}

	burstime := []int{8, -1, 16, -1, -1, -1, -1, 5}

	testcase := append(burstime, padding_time...)
	t.Log("###########  TESTCASE: , ", testcase, " ############")

	for i := 0; i < len(testcase); i++ {
		if testcase[i] == -1 {
			srtf.run(nil, i)
			srtf_opt.run(nil, i)
		} else {
			srtf.run(&Process{id: i, arrivalTime: i, burstTime: testcase[i]}, i)
			srtf_opt.run(&Process_Opt{id: i, arrivalTime: i, burstTime: testcase[i]}, i)
		}
		t.Logf("id: %d", srtf.getProcess())
		t.Logf("id for opt: %d", srtf_opt.getProcess())
	}

	for i := 0; i < len(testcase); i++ {
		if testcase[i] != -1 {
			t.Logf("id: %d, waiting time: %d", i, srtf.getProcessWaitingTime(i))
			t.Logf("id: %d, waiting time for opt: %d", i, srtf_opt.getProcessWaitingTime(i))
		}
	}
}

// run different trials (batches) of jobs on the scheduler and compare performance.
func TestSRTFRand(t *testing.T) {
	numTestcases := 10
	avgWaitingTime := float32(0)
	avgOptWaitingTime := float32(0)

	avgMaxWaitingTime := 0
	avgMaxOptWaitingTime := 0

	avgMakespan := float32(0)
	avgOptMakespan := float32(0)
	for i := 0; i < numTestcases; i++ {
		srtf := NewSRTF()
		srtf_opt := NewSRTF_OPT(4500)
		// 100 jobs total
		// burst time range 100, waiting time in between jobs
		for j := 0; j < 100; j++ {
			waitingTime := rand.Intn(10)
			for k := 0; k < waitingTime; k++ {
				srtf.run(nil, srtf.clockTime)
				srtf_opt.run(nil, srtf_opt.clockTime)
			}
			burstTime := rand.Intn(100) + 1
			srtf.run(&Process{id: j, arrivalTime: srtf.clockTime, burstTime: burstTime}, srtf.clockTime)
			srtf_opt.run(&Process_Opt{id: j, arrivalTime: srtf_opt.clockTime, burstTime: burstTime}, srtf_opt.clockTime)

		}

		for srtf.processId != -1 {
			srtf.run(nil, srtf.clockTime)
		}
		for srtf_opt.processId != -1 {
			srtf_opt.run(nil, srtf_opt.clockTime)
		}

		avgWaitingTime += srtf.getAvgWaitingTime()
		avgOptWaitingTime += srtf_opt.getAvgWaitingTime()

		avgMaxWaitingTime += srtf.getMaxWaitingTime()
		avgMaxOptWaitingTime += srtf_opt.getMaxWaitingTime()

		avgMakespan += float32(srtf.clockTime)
		avgOptMakespan += float32(srtf_opt.clockTime)
	}

	t.Logf("avg waiting time: %f", avgWaitingTime/float32(numTestcases))
	t.Logf("avg opt waiting time: %f", avgOptWaitingTime/float32(numTestcases))

	t.Logf("avg max waiting time: %f", float32(avgMaxWaitingTime)/float32(numTestcases))
	t.Logf("avg max opt waiting time: %f", float32(avgMaxOptWaitingTime)/float32(numTestcases))

	t.Logf("avg makespan: %f", float32(avgMakespan)/float32(numTestcases))
	t.Logf("avg opt makespan: %f", float32(avgOptMakespan)/float32(numTestcases))
}
