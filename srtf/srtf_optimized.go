package srtf

type Process_Opt struct {
	id            int // process identifier
	arrivalTime   int // time when process is included to run
	burstTime     int // estimate of how long job takes
	remainingTime int // estimated amount of time for job to complete
	waitingTime   int // time each process waits until it is completed
	priority      int
}

type SRTF_OPT struct {
	queue                  []Process_Opt        //priority queue sorting based on burst time
	remainingTime          int                  // time left for current running process
	processId              int                  // current process identifier
	totalWaitingTime       int                  // waiting time across all processes
	totalProcessesExecuted int                  // total number of processes executed
	processes              map[int]*Process_Opt // key: process id, value: pointer to process
	threshold              int                  // keeps track of scheduler time
	clockTime              int                  // threshold to determine next process based on priority
}

func NewSRTF_OPT(threshold int) *SRTF_OPT {
	srtf_opt := new(SRTF_OPT)
	srtf_opt.queue = make([]Process_Opt, 0)
	srtf_opt.processId = -1
	srtf_opt.processes = make(map[int]*Process_Opt)
	srtf_opt.threshold = threshold
	return srtf_opt
}

func (srtf_opt *SRTF_OPT) run(process *Process_Opt, currentTime int) {

	//increment priority
	for _, v := range srtf_opt.queue {
		srtf_opt.processes[v.id].priority += 1
		v.priority += 1
	}

	// check if there is a new process to insert
	if process != nil {

		// insert new process into hashmap
		srtf_opt.processes[process.id] = process

		// if new process has shorter remaining/burst time AND it meets the threshold for priority,
		// switch out old process for new process:
		if process.burstTime < srtf_opt.remainingTime && srtf_opt.processes[srtf_opt.processId].priority < srtf_opt.threshold {

			// represent current process with new updated burst time
			currentProcess_Opt := &Process_Opt{id: srtf_opt.processId, burstTime: srtf_opt.remainingTime}

			// run new process
			srtf_opt.processId = process.id
			srtf_opt.remainingTime = process.burstTime

			// we want to add old process back into queue:
			// update incoming process to add to be old process
			process = currentProcess_Opt

		}

		// Decide where to insert process into queue
		found := false
		for i := 0; i < len(srtf_opt.queue); i++ {
			if srtf_opt.queue[i].burstTime > process.burstTime {
				// burstime of current process is less than the ith process,
				// so add new process at i
				found = true
				if i == 0 {
					srtf_opt.queue = append([]Process_Opt{*process}, srtf_opt.queue[:]...)

				} else {
					srtf_opt.queue = append(srtf_opt.queue[:i], srtf_opt.queue[i-1:]...)
					srtf_opt.queue[i] = *process
				}
				break
			}

		}

		// in the case that queue is empty
		if !found {
			srtf_opt.queue = append(srtf_opt.queue, *process)
		}
	}

	// check if current process finished to decide which job runs next (if any)
	if srtf_opt.remainingTime == 0 {
		// calculate the current process' waiting time
		if srtf_opt.processId != -1 {
			waitingTime := currentTime - srtf_opt.processes[srtf_opt.processId].arrivalTime - srtf_opt.processes[srtf_opt.processId].burstTime
			srtf_opt.processes[srtf_opt.processId].waitingTime = waitingTime
			srtf_opt.totalWaitingTime += waitingTime
		}

		if len(srtf_opt.queue) == 0 {
			srtf_opt.processId = -1 // no more processes to execute
		} else {
			// check if there are processes that have waited too long to be executed
			// by comparing process priority with threshold
			// we want to run the process with the max priority
			maxProcess_Opt := srtf_opt.queue[0]
			maxId := -1 // position of max priority process in the queue
			for i := 0; i < len(srtf_opt.queue); i++ {
				if maxProcess_Opt.priority < srtf_opt.processes[srtf_opt.queue[i].id].priority && srtf_opt.processes[srtf_opt.queue[i].id].priority >= srtf_opt.threshold {
					maxProcess_Opt = *srtf_opt.processes[srtf_opt.queue[i].id]
					maxId = i
				}
			}

			// default is to run job with shortest burst time
			nextJob := srtf_opt.queue[0]

			// if there exists a process with high priority, run that instead
			// remove the process to be run from the queue
			if maxId != -1 {
				nextJob = srtf_opt.queue[maxId]
				srtf_opt.queue = append(srtf_opt.queue[:maxId], srtf_opt.queue[maxId+1:]...)
			} else {
				srtf_opt.queue = srtf_opt.queue[1:] //pop from queue
			}

			// update the currently running job to be the next job
			srtf_opt.processId = nextJob.id
			srtf_opt.remainingTime = nextJob.burstTime
			srtf_opt.totalProcessesExecuted += 1
		}

	}

	// only decrement remaining time of scheduler if there is a running process
	if srtf_opt.processId != -1 {
		srtf_opt.remainingTime--
	}
	srtf_opt.clockTime++
}

// get process you want to run at this time
func (srtf_opt *SRTF_OPT) getProcess() int {
	return srtf_opt.processId
}

func (srtf_opt *SRTF_OPT) getMaxWaitingTime() int {
	maxWaitingTime := (-1)
	for _, v := range srtf_opt.processes {
		if v.waitingTime > maxWaitingTime {
			maxWaitingTime = v.waitingTime
		}
	}
	return maxWaitingTime
}

func (srtf_opt *SRTF_OPT) getAvgWaitingTime() float32 {
	return float32(srtf_opt.totalWaitingTime) / float32(len(srtf_opt.processes))
}

func (srtf_opt *SRTF_OPT) getProcessWaitingTime(processId int) int {
	return srtf_opt.processes[processId].waitingTime
}
