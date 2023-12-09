package srtf

type Process struct {
	id            int
	arrivalTime   int //
	burstTime     int // estimate of how long job takes
	remainingTime int
	waitingTime   int
	priority      int
	// execution time = arrivaltime + bursttime
}

type SRTF struct {
	queue                  []Process //priority queue sorting based on burst time
	remainingTime          int       // time left for current running process
	processId              int       // current process
	totalWaitingTime       int
	totalProcessesExecuted int
	processes              map[int]*Process
}

func NewSRTF() *SRTF {
	srtf := new(SRTF)
	srtf.queue = make([]Process, 0)
	srtf.processId = -1
	srtf.processes = make(map[int]*Process)
	return srtf
}

func (srtf *SRTF) run(process *Process, currentTime int) {
	THRESHOLD := 2

	for _, v := range srtf.queue { //increment priority
		srtf.processes[v.id].priority += 1
		v.priority += 1
	}

	if process != nil {
		//process := Process{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}

		srtf.processes[process.id] = process
		// check to switch out process
		if process.burstTime < srtf.remainingTime && srtf.processes[srtf.processId].priority < THRESHOLD {
			// get current process
			// update burst time
			currentProcess := &Process{id: srtf.processId, burstTime: srtf.remainingTime}

			// run new process
			srtf.processId = process.id
			srtf.remainingTime = process.burstTime

			// add process back into queue
			process = currentProcess

		}

		found := false
		for i := 0; i < len(srtf.queue); i++ {
			if srtf.queue[i].burstTime > process.burstTime {
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

	// (0,1), (1,1), (1,2)

	if srtf.remainingTime == 0 {

		if srtf.processId != -1 {
			waitingTime := currentTime - srtf.processes[srtf.processId].arrivalTime - srtf.processes[srtf.processId].burstTime
			srtf.processes[srtf.processId].waitingTime = waitingTime
			srtf.totalWaitingTime += waitingTime
		}

		if len(srtf.queue) == 0 {
			srtf.processId = -1
		} else {

			maxProcess := srtf.queue[0]
			maxId := -1
			for i := 0; i < len(srtf.queue); i++ {
				if maxProcess.priority < srtf.processes[srtf.queue[i].id].priority && srtf.processes[srtf.queue[i].id].priority >= THRESHOLD {
					maxProcess = *srtf.processes[srtf.queue[i].id]
					maxId = i
				}
			}

			nextJob := srtf.queue[0]

			if maxId != -1 {
				nextJob = srtf.queue[maxId]
				srtf.queue = append(srtf.queue[:maxId], srtf.queue[maxId+1:]...)
			} else {
				srtf.queue = srtf.queue[1:] //pop from queue
			}
			srtf.processId = nextJob.id
			srtf.remainingTime = nextJob.burstTime
			srtf.totalProcessesExecuted += 1
		}

	}

	srtf.remainingTime--

}

// get process you want to run at this time
func (srtf *SRTF) getProcess() int {
	return srtf.processId
}

func (srtf *SRTF) getAvgWaitingTIme() float32 {
	return float32(srtf.totalWaitingTime) / float32(srtf.totalProcessesExecuted)
}

func (srtf *SRTF) getProcessWaitingTime(processId int) int {
	return srtf.processes[processId].waitingTime
}
