package sjf

import (
	"fmt";
)

/*
Shortest job first scheduling policy: selects the waiting process with smallest execution time to execute next,
regardless of an incoming job having a shorter estimated execution time than the job currently running.
*/

type Process struct {
	id          int // identifier for each process
	arrivalTime int // time when process is included to run
	burstTime   int // estimated amount of time for job to complete
	// execution time = arrivaltime + bursttime
	waitingTime int // how long each process waits from the moment they arrive until it is completed 
}

type SJF struct {
	queue                  []Process // priority queue sorting based on burst time
	remainingTime          int       // time left for current running process
	processId              int       // current process identifier
	totalWaitingTime       int 		// waiting time across all processes
	totalProcessesExecuted int 		// total number of processes executed 
	processes              map[int]*Process // key: process id, value: pointer to process 
}

func NewSJF() *SJF {
	sjf := new(SJF)
	sjf.queue = make([]Process, 0)
	sjf.processId = -1 
	sjf.processes = make(map[int]*Process)
	return sjf
}

func (sjf *SJF) run(process *Process, currentTime int) {
	// index 0 = smallest burst time
	// 0 1 2 3        2
	// 0 2 4 5 -- 0 2 2 4 5
	// insert 3 -> 0 2 3 4 5
	// insert a process if there is a process

	if process != nil {
		//process := Process{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}
		
		// insert new process into hashmap
		sjf.processes[process.id] = process

		found := false
		// Decide where to place new process in the queue (which is the "waiting line")
		for i := 0; i < len(sjf.queue); i++ {
			if sjf.queue[i].burstTime > process.burstTime {
				found = true
				// append new process to front of queue because it has smallest burstime
				if i == 0 {
					sjf.queue = append([]Process{*process}, sjf.queue[:]...)
				
				// burstime of current process is less than the ith process, 
				// so place new process before it 
				} else { 
					sjf.queue = append(sjf.queue[:i], sjf.queue[i-1:]...)
					sjf.queue[i] = *process
				}
				break
			}
		}
		// in the case that queue is empty, append new process
		if !found {
			sjf.queue = append(sjf.queue, *process)
		}
	}

	// (0,1), (1,1), (1,2)

	// check if current process finished to decide which job runs next (if any)
	if sjf.remainingTime == 0 {
		if len(sjf.queue) == 0 {
			sjf.processId = -2 // no more processes to execute
		} else {
			nextJob := sjf.queue[0] // take next job from the front of queue
			sjf.queue = sjf.queue[1:] // pop from queue
			waitingTime := currentTime - nextJob.arrivalTime 
			nextJob.waitingTime = waitingTime
			sjf.processId = nextJob.id // set nextJob to be the current process now
			sjf.remainingTime = nextJob.burstTime // set remaining time to burstTime of new process to execute

			sjf.totalWaitingTime += waitingTime
			sjf.totalProcessesExecuted += 1
		}

	}

	sjf.remainingTime--
}

// get process you want to run at this time
func (sjf *SJF) getProcess() int {
	return sjf.processId
}

func (sjf *SJF) getAvgWaitingTIme() float32 {
	fmt.Println("total waiting time = ", sjf.totalWaitingTime)
	fmt.Println("total jobs = ", sjf.totalProcessesExecuted)
	return float32(sjf.totalWaitingTime) / float32(sjf.totalProcessesExecuted)
}

func (sjf *SJF) getProcessWaitingTime(processId int) int {
	return sjf.processes[processId].waitingTime
}

