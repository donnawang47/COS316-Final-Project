/*
Shortest job first scheduling policy: selects the waiting process with smallest execution time to execute next,
regardless of an incoming job having a shorter estimated execution time than the job currently running.
*/

package sjf

import (
	"fmt"
)

type Process struct {
	id          int // process identifier
	arrivalTime int // time when process is included to run
	burstTime   int // estimated amount of time for job to complete
	// execution time = arrivaltime + bursttime
	waitingTime int // how long each process waits from the moment they arrive until it is completed
}

type SJF struct {
	queue                  []Process        // priority queue sorting based on burst time
	remainingTime          int              // time left for current running process
	processId              int              // current process identifier
	totalWaitingTime       int              // waiting time across all processes
	totalProcessesExecuted int              // total number of processes executed
	processes              map[int]*Process // key: process id, value: pointer to process
	clockTime              int              // keeps track of scheduler time
}

func NewSJF() *SJF {
	sjf := new(SJF)
	sjf.queue = make([]Process, 0)
	sjf.processId = -1
	sjf.processes = make(map[int]*Process)
	sjf.clockTime = 0
	sjf.remainingTime = 0
	sjf.totalWaitingTime = 0
	sjf.totalProcessesExecuted = 0
	return sjf
}

func (sjf *SJF) run(process *Process, currentTime int) {
	// index 0 = smallest burst time
	// 0 1 2 3        2
	// 0 2 4 5 -- 0 2 2 4 5
	// insert 3 -> 0 2 3 4 5
	// insert a process if there is a process
	// fmt.Println(sjf.clockTime)
	// fmt.Println(sjf.getProcess())

	if process != nil {
		//process := Process{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}

		// insert new process into hashmap
		sjf.processes[process.id] = process

		found := false
		// Decide where to place new process in the queue (which is the "waiting line")
		index := -1
		for i := 0; i < len(sjf.queue); i++ {
			if sjf.queue[i].burstTime > process.burstTime {
				found = true
				// append new process to front of queue because it has smallest burstime
				index = i

				break
			}
		}
		if found {
			if index == 0 {
				sjf.queue = append([]Process{*process}, sjf.queue[:]...)

				// burstime of current process is less than the ith process,
				// so place new process before it
			} else {
				sjf.queue = append(sjf.queue[:index], sjf.queue[index-1:]...)
				sjf.queue[index] = *process
			}
		} else {
			// in the case that queue is empty, append new process
			sjf.queue = append(sjf.queue, *process)
		}
	}

	// (0,1), (1,1), (1,2)

	// check if current process finished to decide which job runs next (if any)
	if sjf.remainingTime == 0 {
		if len(sjf.queue) == 0 {
			sjf.processId = -1 // no more processes to execute
		} else {
			nextJob := sjf.queue[0]               // take next job from the front of queue
			sjf.queue = sjf.queue[1:]             // pop from queue
			sjf.processId = nextJob.id            // set nextJob to be the current process now
			sjf.remainingTime = nextJob.burstTime // set remaining time to burstTime of new process to execute
			waitingTime := sjf.clockTime - nextJob.arrivalTime
			sjf.processes[sjf.processId].waitingTime = waitingTime // update waiting time of pointers of each process

			sjf.totalWaitingTime += waitingTime
			sjf.totalProcessesExecuted += 1
		}

	}

	if sjf.processId != -1 {
		sjf.remainingTime--
	}
	sjf.clockTime++
}

// get process you want to run at this time
func (sjf *SJF) getProcess() int {
	return sjf.processId
}

func (sjf *SJF) getAvgWaitingTime() float32 {
	fmt.Println("Total waiting time = ", sjf.totalWaitingTime)
	fmt.Println("Total jobs = ", sjf.totalProcessesExecuted)
	return float32(sjf.totalWaitingTime) / float32(sjf.totalProcessesExecuted)
}

func (sjf *SJF) getProcessWaitingTime(processId int) int {
	// debugging (remove later)
	fmt.Println("Arrivaltime = ", sjf.processes[processId].arrivalTime)
	return sjf.processes[processId].waitingTime
}

func (sjf *SJF) getProcessCompletionTime(processId int) int {
	return sjf.processes[processId].waitingTime + sjf.processes[processId].burstTime
}
