package sjf

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func scheduler(input_testcase []int, t *testing.T) float32 {
	t.Log("\n", "################### TESTCASE ", input_testcase, "###################\n")
	testcase := input_testcase
	sjf := NewSJF()
	i := 0                                              // keeps track of which element in testcase we are at
	for !(sjf.processId == -1 && i > len(testcase)-1) { // there are no more proccesses left to run
		if i > len(testcase)-1 { // run one by one until we finish completing last job
			sjf.run(nil, sjf.clockTime)
		} else if testcase[i] < 0 { // run algorithm with no new job added to queue
			num := testcase[i]
			for k := 0; k < int(math.Abs(float64(num))); k++ {
				sjf.run(nil, sjf.clockTime)
				t.Logf("%d", sjf.getProcess())
			}
		} else { // new process is to be added to queue
			sjf.run(&Process{id: i, arrivalTime: sjf.clockTime, burstTime: testcase[i]}, sjf.clockTime)
		}
		t.Logf("%d", sjf.getProcess())
		i += 1
	}

	// print waiting time for each process (waiting time saved in hashmap)
	for i := 0; i < len(testcase); i++ {
		if testcase[i] > 0 {
			t.Logf("id: %d, waitingTime: %d", i, sjf.getProcessWaitingTime(i))
			t.Logf("id: %d, completionTime: %d", i, sjf.getProcessCompletionTime(i))
		}
	}

	fmt.Println("\n###### Summary: ######")
	// print total time to run all jobs in the workload trace
	fmt.Printf("Total time to run all jobs: %d\n", sjf.clockTime)

	// print average waiting time across all processes for one workload trace
	avgWaitingTime := sjf.getAvgWaitingTime()
	fmt.Printf("Average waiting time: %f\n", avgWaitingTime)

	return avgWaitingTime

}

// func TestSJF(t *testing.T) {
func TestSJF(t *testing.T) {
	sjf := NewSJF()

	// testcase := []int{3, 2, -1, 5, 3, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1}
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

	t.Errorf("avg waiting time: %f", sjf.getAvgWaitingTime())

	//
}

func TestSJFRand(t *testing.T) {

	numTestcases := 10
	avgWaitingTime := float32(0)
	avgOptWaitingTime := float32(0)
	for i := 0; i < numTestcases; i++ {
		sjf := NewSJF()
		sjf_opt := NewSJF_OPT(120)
		// 100 jobs total
		//burst time range 100, waiting time in between jobs
		for j := 0; j < 3; j++ {
			//waiting time
			waitingTime := rand.Intn(10)
			//fmt.Println("waiting time", waitingTime)
			for k := 0; k < waitingTime; k++ {
				sjf.run(nil, sjf.clockTime)
				sjf_opt.run(nil, sjf_opt.clockTime)
			}
			//burst time

			sjf.run(&Process{id: j, arrivalTime: sjf.clockTime, burstTime: rand.Intn(100) + 1}, sjf.clockTime)
			sjf_opt.run(&Process_Opt{id: j, arrivalTime: sjf_opt.clockTime, burstTime: rand.Intn(100) + 1}, sjf_opt.clockTime)
			//fmt.Println("process", j)

		}

		for sjf.processId != -1 {
			sjf.run(nil, sjf.clockTime)
		}
		for sjf_opt.processId != -1 {
			sjf_opt.run(nil, sjf_opt.clockTime)
		}

		avgWaitingTime += sjf.getAvgWaitingTime()
		avgOptWaitingTime += sjf_opt.getAvgWaitingTime()
	}
	fmt.Println(avgOptWaitingTime)

	t.Errorf("avg waiting time: %f", avgWaitingTime/float32(numTestcases))
	t.Errorf("avg opt waiting time: %f", avgOptWaitingTime/float32(numTestcases))

}
