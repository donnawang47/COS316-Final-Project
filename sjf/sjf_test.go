package sjf

import (
	"fmt"
	"math/rand"
	"testing"
)

// test used to verify that algorithm returns expected waiting time for each individual process and each process
// is run in the correct order
func TestSJF(t *testing.T) {
	sjf := NewSJF()
	sjf_opt := NewSJF_OPT(3) // threshold for priority (aging mechanism) = 3

	var padding_time []int
	// negative numbers represent 1 unit time of running scheduler with no new job added to it
	for i := 0; i < 28; i++ {
		padding_time = append(padding_time, -1)
	}

	burstime := []int{16, 10, -1, 1}

	testcase := append(burstime, padding_time...)
	t.Log("###########  TESTCASE: , ", testcase, " ############")

	for i := 0; i < len(testcase); i++ {
		if testcase[i] == -1 {
			sjf.run(nil, i)
			sjf_opt.run(nil, i)
		} else {
			sjf.run(&Process{id: i, arrivalTime: i, burstTime: testcase[i]}, i)
			sjf_opt.run(&Process_Opt{id: i, arrivalTime: i, burstTime: testcase[i]}, i)
		}
		t.Logf("id: %d", sjf.getProcess())
		t.Logf("id for opt: %d", sjf_opt.getProcess())
	}

	for i := 0; i < len(testcase); i++ {
		if testcase[i] != -1 {
			t.Logf("id: %d, waiting time: %d", i, sjf.getProcessWaitingTime(i))
			t.Logf("id: %d, waiting time for opt: %d", i, sjf_opt.getProcessWaitingTime(i))
		}
	}
}

// run different trials (batches) of jobs on the scheduler and compare performance.
func TestSJFRand(t *testing.T) {

	numTestcases := 10
	avgWaitingTime := float32(0)
	avgOptWaitingTime := float32(0)

	avgMaxWaitingTime := 0
	avgMaxOptWaitingTime := 0

	avgMakespan := float32(0)
	avgOptMakespan := float32(0)
	for i := 0; i < numTestcases; i++ {
		sjf := NewSJF()
		sjf_opt := NewSJF_OPT(1000)
		// 100 jobs total
		//burst time range 100, waiting time in between jobs
		for j := 0; j < 100; j++ {
			waitingTime := rand.Intn(10)
			for k := 0; k < waitingTime; k++ {
				sjf.run(nil, sjf.clockTime)
				sjf_opt.run(nil, sjf_opt.clockTime)
			}
			burstTime := rand.Intn(100) + 1
			sjf.run(&Process{id: j, arrivalTime: sjf.clockTime, burstTime: burstTime}, sjf.clockTime)
			sjf_opt.run(&Process_Opt{id: j, arrivalTime: sjf_opt.clockTime, burstTime: burstTime}, sjf_opt.clockTime)
		}

		for sjf.processId != -1 {
			sjf.run(nil, sjf.clockTime)
		}
		for sjf_opt.processId != -1 {
			sjf_opt.run(nil, sjf_opt.clockTime)
		}

		avgWaitingTime += sjf.getAvgWaitingTime()
		avgOptWaitingTime += sjf_opt.getAvgWaitingTime()

		avgMaxWaitingTime += sjf.getMaxWaitingTime()
		avgMaxOptWaitingTime += sjf_opt.getMaxWaitingTime()

		avgMakespan += float32(sjf.clockTime)
		avgOptMakespan += float32(sjf_opt.clockTime)

	}
	fmt.Println(avgOptWaitingTime)

	t.Logf("avg waiting time: %f", avgWaitingTime/float32(numTestcases))
	t.Logf("avg opt waiting time: %f", avgOptWaitingTime/float32(numTestcases))

	t.Logf("avg max waiting time: %f", float32(avgMaxWaitingTime)/float32(numTestcases))
	t.Logf("avg max opt waiting time: %f", float32(avgMaxOptWaitingTime)/float32(numTestcases))

	t.Logf("avg makespan: %f", float32(avgMakespan)/float32(numTestcases))
	t.Logf("avg opt makespan: %f", float32(avgOptMakespan)/float32(numTestcases))
}
