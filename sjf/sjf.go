/*
Shortest job first scheduling policy: selects the waiting
process with smallest execution time to execute next,
regardless of an incoming job having a shorter estimated
execution time than the job currently running.
*/

package sjf

type Process struct {
	id          int // process identifier
	arrivalTime int // time when process is included to run
	burstTime   int // estimated amount of time for job to complete
	waitingTime int // time each process waits until it is completed
}

type SJF struct {
	queue                  []Process        // priority queue sorted based on burst time
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
	return sjf
}

func (sjf *SJF) run(process *Process, currentTime int) {

	// check if there is a new process to insert
	if process != nil {

		// insert new process into hashmap
		sjf.processes[process.id] = process

		// Decide where to place new process in the queue (which is the "waiting line")
		found := false

		index := -1
		for i := 0; i < len(sjf.queue); i++ {
			// burstime of current process is less than the ith process,
			// so add new process at i
			// push all remaining processes to after i
			if sjf.queue[i].burstTime > process.burstTime {
				found = true
				index = i
				break
			}
		}

		if found {
			// check if append to front of queue
			if index == 0 {
				sjf.queue = append([]Process{*process}, sjf.queue[:]...)

			} else {
				sjf.queue = append(sjf.queue[:index], sjf.queue[index-1:]...)
				sjf.queue[index] = *process
			}
		} else {
			// in the case that queue is empty, append new process
			sjf.queue = append(sjf.queue, *process)
		}
	}

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

	// only decrement remaining time of scheduler if there is a running process
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
	return float32(sjf.totalWaitingTime) / float32(len(sjf.processes))
}

func (sjf *SJF) getMaxWaitingTime() int {
	maxWaitingTime := -1
	for _, v := range sjf.processes {
		if v.waitingTime > maxWaitingTime {
			maxWaitingTime = v.waitingTime
		}
	}
	return maxWaitingTime
}

func (sjf *SJF) getProcessWaitingTime(processId int) int {
	return sjf.processes[processId].waitingTime
}

func (sjf *SJF) getProcessCompletionTime(processId int) int {
	return sjf.processes[processId].waitingTime + sjf.processes[processId].burstTime
}
