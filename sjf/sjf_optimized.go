package sjf

import (
	"fmt"
)

type Process_Opt struct {
	id          int
	arrivalTime int //
	burstTime   int // estimate of how long job takes
	// execution time = arrivaltime + bursttime
	waitingTime int
	priority    int
}

type SJF_OPT struct {
	queue                  []Process_Opt //priority queue sorting based on burst time
	remainingTime          int           // time left for current running process
	processId              int           // current process
	totalWaitingTime       int
	totalProcessesExecuted int
	processes              map[int]*Process_Opt
	clockTime              int
	threshold              int
}

func NewSJF_OPT(threshold int) *SJF_OPT {
	sjf_opt := new(SJF_OPT)
	sjf_opt.queue = make([]Process_Opt, 0)
	sjf_opt.processId = -1
	sjf_opt.processes = make(map[int]*Process_Opt)
	sjf_opt.clockTime = 0
	sjf_opt.threshold = threshold
	return sjf_opt
}

func (sjf_opt *SJF_OPT) run(process *Process_Opt, currentTime int) {
	// sort all processes according to arrival time
	// arrival time is end of scheduler len(scheduler) - 1
	// select process that has min arrival time and min burst time
	//add to queue

	//queue = append(queue, Process_Opt{id = processId, burstTime = burstTime, arrivalTime = arrivalTime})

	for _, v := range sjf_opt.queue { // v = process
		sjf_opt.processes[v.id].priority += 1 // increment priority of each process from queue (waiting) in hashmap
		v.priority += 1                       // increment priority field of process pointer
	} //process := Process_Opt{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}

	if process != nil {
		//process := Process_Opt{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}

		sjf_opt.processes[process.id] = process

		found := false

		index := -1
		for i := 0; i < len(sjf_opt.queue); i++ {
			if sjf_opt.queue[i].burstTime > process.burstTime {
				found = true
				index = i

				break
			}

		}
		// in the case that queue is empty
		if found {
			if index == 0 {
				sjf_opt.queue = append([]Process_Opt{*process}, sjf_opt.queue[:]...)

			} else {
				sjf_opt.queue = append(sjf_opt.queue[:index], sjf_opt.queue[index-1:]...)
				sjf_opt.queue[index] = *process
			}
		} else {
			sjf_opt.queue = append(sjf_opt.queue, *process)
		}
	}

	// (0,1), (1,1), (1,2)

	if sjf_opt.remainingTime == 0 {

		if len(sjf_opt.queue) == 0 {
			sjf_opt.processId = -1
		} else {
			//waiting too long to be executed
			maxProcess_Opt := sjf_opt.queue[0]
			maxId := -1
			for i := 0; i < len(sjf_opt.queue); i++ {
				if maxProcess_Opt.priority < sjf_opt.processes[sjf_opt.queue[i].id].priority && sjf_opt.processes[sjf_opt.queue[i].id].priority >= sjf_opt.threshold {
					maxProcess_Opt = *sjf_opt.processes[sjf_opt.queue[i].id]
					maxId = i
				}
			}
			nextJob := sjf_opt.queue[0]
			if maxId != -1 {
				nextJob = sjf_opt.queue[maxId]
				sjf_opt.queue = append(sjf_opt.queue[:maxId], sjf_opt.queue[maxId+1:]...) // exclude nextJob = pop it!
			} else {
				sjf_opt.queue = sjf_opt.queue[1:] // pop from queue
			}

			sjf_opt.processId = nextJob.id
			sjf_opt.remainingTime = nextJob.burstTime
			waitingTime := currentTime - nextJob.arrivalTime               // nextJob.arrivalTime = its start time
			sjf_opt.processes[sjf_opt.processId].waitingTime = waitingTime // update waiting time field of next process/job

			sjf_opt.totalWaitingTime += waitingTime
			sjf_opt.totalProcessesExecuted += 1
		}

	}

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
