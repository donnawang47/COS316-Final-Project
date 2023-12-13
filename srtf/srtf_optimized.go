package srtf

type Process_Opt struct {
	id            int
	arrivalTime   int //
	burstTime     int // estimate of how long job takes
	remainingTime int
	waitingTime   int
	priority      int
	// execution time = arrivaltime + bursttime
}

type SRTF_OPT struct {
	queue                  []Process_Opt //priority queue sorting based on burst time
	remainingTime          int           // time left for current running process
	processId              int           // current process
	totalWaitingTime       int
	totalProcessesExecuted int
	processes              map[int]*Process_Opt
	threshold              int
	clockTime              int
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

	for _, v := range srtf_opt.queue { //increment priority
		srtf_opt.processes[v.id].priority += 1
		v.priority += 1
	}

	if process != nil {
		//process := Process{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}

		srtf_opt.processes[process.id] = process
		// check to switch out process
		if process.burstTime < srtf_opt.remainingTime && srtf_opt.processes[srtf_opt.processId].priority < srtf_opt.threshold {
			// get current process
			// update burst time
			currentProcess_Opt := &Process_Opt{id: srtf_opt.processId, burstTime: srtf_opt.remainingTime}

			// run new process
			srtf_opt.processId = process.id
			srtf_opt.remainingTime = process.burstTime

			// add process back into queue
			process = currentProcess_Opt

		}

		found := false
		for i := 0; i < len(srtf_opt.queue); i++ {
			if srtf_opt.queue[i].burstTime > process.burstTime {
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

	// (0,1), (1,1), (1,2)

	if srtf_opt.remainingTime == 0 {

		if srtf_opt.processId != -1 {
			waitingTime := currentTime - srtf_opt.processes[srtf_opt.processId].arrivalTime - srtf_opt.processes[srtf_opt.processId].burstTime
			srtf_opt.processes[srtf_opt.processId].waitingTime = waitingTime
			srtf_opt.totalWaitingTime += waitingTime
		}

		if len(srtf_opt.queue) == 0 {
			srtf_opt.processId = -1
		} else {

			maxProcess_Opt := srtf_opt.queue[0]
			maxId := -1
			for i := 0; i < len(srtf_opt.queue); i++ {
				if maxProcess_Opt.priority < srtf_opt.processes[srtf_opt.queue[i].id].priority && srtf_opt.processes[srtf_opt.queue[i].id].priority >= srtf_opt.threshold {
					maxProcess_Opt = *srtf_opt.processes[srtf_opt.queue[i].id]
					maxId = i
				}
			}

			nextJob := srtf_opt.queue[0]

			if maxId != -1 {
				nextJob = srtf_opt.queue[maxId]
				srtf_opt.queue = append(srtf_opt.queue[:maxId], srtf_opt.queue[maxId+1:]...)
			} else {
				srtf_opt.queue = srtf_opt.queue[1:] //pop from queue
			}
			srtf_opt.processId = nextJob.id
			srtf_opt.remainingTime = nextJob.burstTime
			srtf_opt.totalProcessesExecuted += 1
		}

	}
	if srtf_opt.processId != -1 {
		srtf_opt.remainingTime--
	}
	srtf_opt.clockTime++

}

// get process you want to run at this time
func (srtf_opt *SRTF_OPT) getProcess() int {
	return srtf_opt.processId
}

func (srtf_opt *SRTF_OPT) getAvgWaitingTime() float32 {
	return float32(srtf_opt.totalWaitingTime) / float32(srtf_opt.totalProcessesExecuted)
}

func (srtf_opt *SRTF_OPT) getProcessWaitingTime(processId int) int {
	return srtf_opt.processes[processId].waitingTime
}
