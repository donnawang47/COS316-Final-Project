package sjf

import (
	"fmt"
)

type Process_Opt struct {
	id          int // process identifier
	arrivalTime int // time when process is included to run
	burstTime   int // estimate of how long job takes
	waitingTime int // time each process waits until it is completed
	priority    int // priority of process in queue
}

type SJF_OPT struct {
	queue                  []Process_Opt        //priority queue sorting based on burst time
	remainingTime          int                  // time left for current running process
	processId              int                  // current process identifier
	totalWaitingTime       int                  // waiting time across all processes
	totalProcessesExecuted int                  // total number of processes executed
	processes              map[int]*Process_Opt // key: process id, value: pointer to process
	clockTime              int                  // keeps track of scheduler time
	threshold              int                  // threshold to determine next process based on priority
}

func NewSJF_OPT(threshold int) *SJF_OPT {
	sjf_opt := new(SJF_OPT)
	sjf_opt.queue = make([]Process_Opt, 0)
	sjf_opt.processId = -1
	sjf_opt.processes = make(map[int]*Process_Opt)
	sjf_opt.threshold = threshold
	return sjf_opt
}

func (sjf_opt *SJF_OPT) run(process *Process_Opt, currentTime int) {

	// increment priority of each process from queue (waiting) in hashmap
	for _, v := range sjf_opt.queue {
		sjf_opt.processes[v.id].priority += 1
		v.priority += 1
	}

	// check if there is a new process to insert
	if process != nil {

		// insert new process into hashmap
		sjf_opt.processes[process.id] = process

		// Decide where to place new process in the queue (which is the "waiting line")
		found := false

		index := -1
		for i := 0; i < len(sjf_opt.queue); i++ {
			// burstime of current process is less than the ith process,
			// so add new process at i
			if sjf_opt.queue[i].burstTime > process.burstTime {
				found = true
				index = i
				break
			}

		}

		if found {
			if index == 0 {
				sjf_opt.queue = append([]Process_Opt{*process}, sjf_opt.queue[:]...)

			} else {
				sjf_opt.queue = append(sjf_opt.queue[:index], sjf_opt.queue[index-1:]...)
				sjf_opt.queue[index] = *process
			}
		} else {
			// in the case that queue is empty
			sjf_opt.queue = append(sjf_opt.queue, *process)
		}
	}

	// check if current process finished to decide which job runs next (if any)
	if sjf_opt.remainingTime == 0 {

		if len(sjf_opt.queue) == 0 {
			sjf_opt.processId = -1 // no more processes to execute
		} else {
			// check if there are processes that have waited too long to be executed
			// by comparing process priority with threshold
			// we want to run the process with the max priority
			maxProcess_Opt := sjf_opt.queue[0]
			maxId := -1 // position of max priority process in the queue
			for i := 0; i < len(sjf_opt.queue); i++ {
				if maxProcess_Opt.priority < sjf_opt.processes[sjf_opt.queue[i].id].priority && sjf_opt.processes[sjf_opt.queue[i].id].priority >= sjf_opt.threshold {
					maxProcess_Opt = *sjf_opt.processes[sjf_opt.queue[i].id]
					maxId = i
				}
			}
			// default is to run job with shortest burst time
			nextJob := sjf_opt.queue[0]

			// if there exists a process with high priority, run that instead
			// remove the process to be run from the queue
			if maxId != -1 {
				nextJob = sjf_opt.queue[maxId]
				sjf_opt.queue = append(sjf_opt.queue[:maxId], sjf_opt.queue[maxId+1:]...) // exclude nextJob = pop it!
			} else {
				sjf_opt.queue = sjf_opt.queue[1:] // pop from queue
			}

			// update the currently running job to be the next job
			sjf_opt.processId = nextJob.id
			sjf_opt.remainingTime = nextJob.burstTime

			// update waiting time of this new job
			waitingTime := currentTime - nextJob.arrivalTime
			sjf_opt.processes[sjf_opt.processId].waitingTime = waitingTime
			sjf_opt.totalWaitingTime += waitingTime
			sjf_opt.totalProcessesExecuted += 1
		}

	}

	// only decrement remaining time of scheduler if there is a running process
	if sjf_opt.processId != -1 {
		sjf_opt.remainingTime--
	}
	sjf_opt.clockTime++

}

// get process you want to run at this time
func (sjf_opt *SJF_OPT) getProcess() int {
	return sjf_opt.processId
}

func (sjf_opt *SJF_OPT) getAvgWaitingTime() float32 {
	return float32(sjf_opt.totalWaitingTime) / float32(len(sjf_opt.processes))
}

func (sjf_opt *SJF_OPT) getMaxWaitingTime() int {
	maxWaitingTime := (-1)
	for _, v := range sjf_opt.processes {
		if v.waitingTime > maxWaitingTime {
			maxWaitingTime = v.waitingTime
		}
	}
	return maxWaitingTime
}

func (sjf_opt *SJF_OPT) getProcessWaitingTime(processId int) int {
	return sjf_opt.processes[processId].waitingTime
}

func (sjf_opt *SJF_OPT) getProcessCompletionTime(processId int) int {
	return sjf_opt.processes[processId].waitingTime + sjf_opt.processes[processId].burstTime
}

func (sjf_opt *SJF_OPT) printSummary() {
	fmt.Println("Total waiting time = ", sjf_opt.totalWaitingTime)
	fmt.Println("Total jobs = ", sjf_opt.totalProcessesExecuted)
}
