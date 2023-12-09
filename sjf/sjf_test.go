package sjf

import (
	"testing";
	"math"
)

// func TestSJF(t *testing.T) {
func TestSJF(t *testing.T) {
	sjf := NewSJF()

	// testcase := []int{3, 2, -1, 5, 3, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1}
	// testcase := []int{5, 5, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	// testcase := []int{5, -1, 3, 4, 7, -1}
	
	testcase := []int{16, 10, -1, 1} // avg waiting = 9.67
	// testcase := []int{16, 10, -1, 1, -12, 23, -13, 4} // avg waiting = 12.6

	i := 0
	for sjf.processId != -2 { // -2 indicates that there are no more proccesses left to run in queue
		if i > len(testcase)-1 { // run one by one until we finish all remaining jobs
			sjf.run(nil, i)
		} else if testcase[i] < 0 { // run algorithm with no new job added to queue
			num := testcase[i] 
			for k := 0; k < int(math.Abs(float64(num))); k++ {
				sjf.run(nil, i)
				i += 1
			}
			i -= 1 // since there is an increment at the end
		} else { // new process is to be added to queue
			sjf.run(&Process{id: i, arrivalTime: i, burstTime: testcase[i]}, i)
		}
		t.Logf("%d", sjf.getProcess())
		i += 1
	}

	// for i := 0; i < len(testcase); i++ {
	// 	//i is arrival time, testcase[i] is bursttime
	// 	// completion time - arrival time
	// 	// startime - arrival time
	// 	if testcase[i] == -1 {
	// 		sjf.run(nil, i)
	// 	} else {
	// 		sjf.run(&Process{id: i, arrivalTime: i, burstTime: testcase[i]}, i)
	// 	}
	// 	t.Logf("%d", sjf.getProcess())
	// }
	
	// print waiting time for each process (waiting time saved in hashmap)
	for i := 0; i < len(testcase); i++ { 
		if testcase[i] > 0 {
			t.Logf("id: %d, waitingT: %d", i, sjf.getProcessWaitingTime(i))
		}

	}

	t.Logf("avg waiting time: %f", sjf.getAvgWaitingTIme())
}
