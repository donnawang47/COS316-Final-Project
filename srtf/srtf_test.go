package srtf

import (
	// "fmt"
	// "math/rand"
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

	// burstime := []int{16, 10, -1, 1}
	burstime := []int{8, -1, 16, -1, -1, -1, -1, 5}

	testcase := append(burstime, padding_time...)
	t.Log("###########  TESTCASE: , ", testcase, " ############")

	for i := 0; i < len(testcase); i++ {
		//i is arrival time, testcase[i] is bursttime
		// completion time - arrival time
		// startime - arrival time
		if testcase[i] == -1 {
			srtf.run(nil, i)
			srtf_opt.run(nil, i)
		} else {
			srtf.run(&Process{id: i, arrivalTime: i, burstTime: testcase[i]}, i)
			srtf_opt.run(&Process_Opt{id: i, arrivalTime: i, burstTime: testcase[i]}, i)
		}
		t.Logf("id: %d", srtf.getProcess())
		t.Logf("id for opt: %d", srtf_opt.getProcess())
		// t.Logf("id: %d, WAITING TIME: %d", srtf_opt.getProcess(), srtf_opt.getProcessWaitingTime(srtf_opt.getProcess()))
	}

	for i := 0; i < len(testcase); i++ {
		//i is arrival time, testcase[i] is bursttime
		// completion time - arrival time
		// startime - arrival time
		if testcase[i] != -1 {
			t.Logf("id: %d, waiting time: %d", i, srtf.getProcessWaitingTime(i))
			t.Logf("id: %d, waiting time for opt: %d", i, srtf_opt.getProcessWaitingTime(i))
		}
	}
}

// func TestSRTF(t *testing.T) {
// 	srtf := NewSRTF()

// 	// negative numbers represent 1 unit time of running scheduler with no new job added to it
// 	testcase := []int{1, 1, 1, 1, 3, 1, -1, -1, -1}

// 	for i := 0; i < len(testcase); i++ {
// 		//i is arrival time, testcase[i] is bursttime
// 		// completion time - arrival time
// 		// startime - arrival time
// 		if testcase[i] == -1 {
// 			srtf.run(nil, i)
// 		} else {
// 			srtf.run(&Process{id: i, arrivalTime: i, burstTime: testcase[i], waitingTime: -1}, i)
// 		}
// 		t.Logf("%d", srtf.getProcess())

// 	}

// 	for i := 0; i < len(testcase); i++ {
// 		//i is arrival time, testcase[i] is bursttime
// 		// completion time - arrival time
// 		// startime - arrival time
// 		if testcase[i] != -1 {
// 			t.Logf("%d, %d:", i, srtf.getProcessWaitingTime(i))
// 		}

// 	}

// 	t.Logf("avg waiting time: %f", srtf.getAvgWaitingTime())
// }

// func TestSRTFRand(t *testing.T) {
// 	numTestcases := 10
// 	avgWaitingTime := float32(0)
// 	avgOptWaitingTime := float32(0)

// 	avgMaxWaitingTime := 0
// 	avgMaxOptWaitingTime := 0

// 	avgMakespan := float32(0)
// 	avgOptMakespan := float32(0)
// 	for i := 0; i < numTestcases; i++ {
// 		srtf := NewSRTF()
// 		srtf_opt := NewSRTF_OPT(4500)
// 		// 100 jobs total
// 		//burst time range 100, waiting time in between jobs
// 		for j := 0; j < 100; j++ {
// 			//waiting time
// 			waitingTime := rand.Intn(10)
// 			//fmt.Println("waiting time", waitingTime)
// 			for k := 0; k < waitingTime; k++ {
// 				srtf.run(nil, srtf.clockTime)
// 				srtf_opt.run(nil, srtf_opt.clockTime)
// 			}
// 			//burst time
// 			burstTime := rand.Intn(100) + 1
// 			srtf.run(&Process{id: j, arrivalTime: srtf.clockTime, burstTime: burstTime}, srtf.clockTime)
// 			srtf_opt.run(&Process_Opt{id: j, arrivalTime: srtf_opt.clockTime, burstTime: burstTime}, srtf_opt.clockTime)
// 			//fmt.Println("process", j)

// 		}

// 		for srtf.processId != -1 {
// 			srtf.run(nil, srtf.clockTime)
// 		}
// 		for srtf_opt.processId != -1 {
// 			srtf_opt.run(nil, srtf_opt.clockTime)
// 		}

// 		avgWaitingTime += srtf.getAvgWaitingTime()
// 		avgOptWaitingTime += srtf_opt.getAvgWaitingTime()

// 		avgMaxWaitingTime += srtf.getMaxWaitingTime()
// 		avgMaxOptWaitingTime += srtf_opt.getMaxWaitingTime()

// 		avgMakespan += float32(srtf.clockTime)
// 		avgOptMakespan += float32(srtf_opt.clockTime)
// 	}
// 	fmt.Println(avgOptWaitingTime)

// 	t.Logf("avg waiting time: %f", avgWaitingTime/float32(numTestcases))
// 	t.Logf("avg opt waiting time: %f", avgOptWaitingTime/float32(numTestcases))

// 	t.Logf("avg max waiting time: %f", float32(avgMaxWaitingTime)/float32(numTestcases))
// 	t.Logf("avg max opt waiting time: %f", float32(avgMaxOptWaitingTime)/float32(numTestcases))

// 	t.Logf("avg makespan: %f", float32(avgMakespan)/float32(numTestcases))
// 	t.Logf("avg max makespan: %f", float32(avgOptMakespan)/float32(numTestcases))
// }
