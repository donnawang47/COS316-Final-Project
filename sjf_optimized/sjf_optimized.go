package sjf_optimized

type Process struct {
	id          int
	arrivalTime int //
	burstTime   int // estimate of how long job takes
	// execution time = arrivaltime + bursttime
	waitingTime int
	priority    int
}

type SJF struct {
	queue                  []Process //priority queue sorting based on burst time
	remainingTime          int       // time left for current running process
	processId              int       // current process
	totalWaitingTime       int
	totalProcessesExecuted int
	processes              map[int]*Process
}

func NewSJF() *SJF {
	sjf := new(SJF)
	sjf.queue = make([]Process, 0)
	sjf.processId = -1
	sjf.processes = make(map[int]*Process)
	return sjf
}

func (sjf *SJF) run(process *Process, currentTime int) {
	// sort all processes according to arrival time
	// arrival time is end of scheduler len(scheduler) - 1
	// select process that has min arrival time and min burst time
	//add to queue

	//queue = append(queue, Process{id = processId, burstTime = burstTime, arrivalTime = arrivalTime})

	for _, v := range sjf.queue { //increment priority
		sjf.processes[v.id].priority += 1
		v.priority += 1
	} //process := Process{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}

	if process != nil {
		//process := Process{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}

		sjf.processes[process.id] = process

		found := false

		for i := 0; i < len(sjf.queue); i++ {
			if sjf.queue[i].burstTime > process.burstTime {
				found = true
				if i == 0 {
					sjf.queue = append([]Process{*process}, sjf.queue[:]...)

				} else {
					sjf.queue = append(sjf.queue[:i], sjf.queue[i-1:]...)
					sjf.queue[i] = *process
				}
				break
			}

		}
		// in the case that queue is empty
		if !found {
			sjf.queue = append(sjf.queue, *process)
		}
	}

	// (0,1), (1,1), (1,2)

	if sjf.remainingTime == 0 {

		if len(sjf.queue) == 0 {
			sjf.processId = -1
		} else {
			THRESHOLD := 2
			maxProcess := sjf.queue[0]
			maxId := -1
			for i := 0; i < len(sjf.queue); i++ {
				if maxProcess.priority < sjf.processes[sjf.queue[i].id].priority && sjf.processes[sjf.queue[i].id].priority >= THRESHOLD {
					maxProcess = *sjf.processes[sjf.queue[i].id]
					maxId = i
				}
			}
			nextJob := sjf.queue[0]
			if maxId != -1 {
				nextJob = sjf.queue[maxId]
				sjf.queue = append(sjf.queue[:maxId], sjf.queue[maxId+1:]...)
			} else {
				sjf.queue = sjf.queue[1:] //pop from queue
			}

			sjf.processId = nextJob.id
			sjf.remainingTime = nextJob.burstTime
			waitingTime := currentTime - nextJob.arrivalTime // nextJob.arrivalTime = start time of next job
			sjf.processes[sjf.processId].waitingTime = waitingTime

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
	return float32(sjf.totalWaitingTime) / float32(sjf.totalProcessesExecuted)
}

func (sjf *SJF) getProcessWaitingTime(processId int) int {
	return sjf.processes[processId].waitingTime
}
