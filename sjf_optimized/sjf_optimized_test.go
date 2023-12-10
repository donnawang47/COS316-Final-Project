package sjf_optimized

import (
	"fmt"
	"math"
	"testing"
)

func scheduler(input_testcase []int, t *testing.T) float32 {
	t.Log("\n", "################### TESTCASE ", input_testcase, "###################\n")
	testcase := input_testcase
	sjf := NewSJF()
	i := 0                                              // keeps track of which element in testcase we are at
	for !(sjf.processId == -2 && i > len(testcase)-1) { // there are no more proccesses left to run
		skip := false
		if i > len(testcase)-1 { // run one by one until we finish completing last job
			sjf.run(nil, sjf.clockTime)
		} else if testcase[i] < 0 { // run algorithm with no new job added to queue
			skip = true
			num := testcase[i]
			for k := 0; k < int(math.Abs(float64(num))); k++ {
				sjf.run(nil, sjf.clockTime)
				t.Logf("process id = %d", sjf.getProcess())
			}
		} else { // new process is to be added to queue
			sjf.run(&Process{id: i, arrivalTime: sjf.clockTime, burstTime: testcase[i]}, sjf.clockTime)
		}
		if !skip {
			t.Logf("process id = %d", sjf.getProcess())
		}
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

	sjf.printSummary() // prints total waiting time and jobs run

	// print average waiting time across all processes for one workload trace
	avgWaitingTime := sjf.getAvgWaitingTime()
	fmt.Printf("Average waiting time: %f\n", avgWaitingTime)

	return avgWaitingTime

}

// func TestSJF(t *testing.T) {
func TestSJF_Optimized(t *testing.T) {

	// testcase := []int{3, 2, -1, 5, 3, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1}
	// testcase := []int{5, 5, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	// testcase := []int{5, -1, 3, 4, 7, -1}
	var testcase []int

	testcase = []int{16, 10, -1, 1} // avg waiting = 9.67
	scheduler(testcase, t)

	testcase = []int{16, 10, -1, 1, -13, 23, -10, 4} // avg waiting = 12.20
	scheduler(testcase, t)
}
