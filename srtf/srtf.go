package srtf

type Process struct {
	id            int
	arrivalTime   int //
	burstTime     int // estimate of how long job takes
	remainingTime int
	waitingTime   int
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

	// index 0 = smallest burst time
	// 0 1 2 3        2
	// 0 2 4 5 -- 0 2 2 4 5
	// insert 3 -> 0 2 3 4 5
	// insert a process if there is a process

	if process != nil {
		//process := Process{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}

		srtf.processes[process.id] = process
		if process.burstTime < srtf.remainingTime {
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
			nextJob := srtf.queue[0]
			srtf.queue = srtf.queue[1:] //pop from queue
			srtf.processId = nextJob.id
			srtf.remainingTime = nextJob.burstTime
			//waitingTime := currentTime - nextJob.arrivalTime // nextJob.arrivalTime = start time of next job
			//
			srtf.totalProcessesExecuted += 1
		}

	}

	srtf.remainingTime--

	//

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
