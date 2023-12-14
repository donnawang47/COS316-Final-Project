package srtf

type Process struct {
	id            int // process identifier
	arrivalTime   int // time when process is included to run
	burstTime     int // estimate of how long job takes
	remainingTime int // estimated amount of time for job to complete
	waitingTime   int // time each process waits until it is completed
}

type SRTF struct {
	queue                  []Process        //priority queue sorting based on burst time
	remainingTime          int              // time left for current running process
	processId              int              // current process identifier
	totalWaitingTime       int              // waiting time across all processes
	totalProcessesExecuted int              // total number of processes executed
	processes              map[int]*Process // key: process id, value: pointer to process
	clockTime              int              // keeps track of scheduler time
}

func NewSRTF() *SRTF {
	srtf := new(SRTF)
	srtf.queue = make([]Process, 0)
	srtf.processId = -1
	srtf.processes = make(map[int]*Process)
	return srtf
}

func (srtf *SRTF) run(process *Process, currentTime int) {

	// check if there is a new process to insert
	if process != nil {
		// insert new process into hashmap
		srtf.processes[process.id] = process

		// to ensure we are always running process with shortest remaining time
		// check if new process has a shorter burst time
		// than remaining time of current running process
		// (which is already shortest burst time seen thus far)
		if process.burstTime < srtf.remainingTime {
			// since new process has shorter remaining/burst time,
			// switch out old process for new process:

			// represent current process with new updated burst time
			currentProcess := &Process{id: srtf.processId, burstTime: srtf.remainingTime}

			// run new process
			srtf.processId = process.id
			srtf.remainingTime = process.burstTime

			// we want to add old process back into queue:
			// update incoming process to add to be old process
			process = currentProcess

		}

		// Decide where to insert process into queue
		found := false
		for i := 0; i < len(srtf.queue); i++ {
			if srtf.queue[i].burstTime > process.burstTime {
				// burstime of current process is less than the ith process,
				// so add new process at i
				found = true
				if i == 0 {
					srtf.queue = append([]Process{*process}, srtf.queue[:]...)

				} else {
					srtf.queue = append(srtf.queue[:i], srtf.queue[i-1:]...)
					srtf.queue[i] = *process
				}
				break
			}

		}
		// in the case that queue is empty
		if !found {
			srtf.queue = append(srtf.queue, *process)
		}
	}

	// check if current process finished to decide which job runs next (if any)
	if srtf.remainingTime == 0 {

		// calculate the current process' waiting time
		if srtf.processId != -1 {
			waitingTime := currentTime - srtf.processes[srtf.processId].arrivalTime - srtf.processes[srtf.processId].burstTime
			srtf.processes[srtf.processId].waitingTime = waitingTime
			srtf.totalWaitingTime += waitingTime
		}

		if len(srtf.queue) == 0 {
			srtf.processId = -1 // no more processes to execute
		} else {
			// take next job from the front of queue
			nextJob := srtf.queue[0]
			srtf.queue = srtf.queue[1:] //pop from queue
			srtf.processId = nextJob.id
			srtf.remainingTime = nextJob.burstTime
			srtf.totalProcessesExecuted += 1
		}

	}

	// only decrement remaining time of scheduler if there is a running process
	if srtf.processId != -1 {
		srtf.remainingTime--
	}

	srtf.clockTime++

}

// get process you want to run at this time
func (srtf *SRTF) getProcess() int {
	return srtf.processId
}

func (srtf *SRTF) getMaxWaitingTime() int {
	maxWaitingTime := (-1)
	for _, v := range srtf.processes {
		if v.waitingTime > maxWaitingTime {
			maxWaitingTime = v.waitingTime
		}
	}
	return maxWaitingTime
}

func (srtf *SRTF) getAvgWaitingTime() float32 {
	return float32(srtf.totalWaitingTime) / float32(len(srtf.processes))
}

func (srtf *SRTF) getProcessWaitingTime(processId int) int {
	return srtf.processes[processId].waitingTime
}
