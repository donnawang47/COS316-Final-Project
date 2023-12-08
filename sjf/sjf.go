package sjf

type Process struct {
	id          int
	arrivalTime int //
	burstTime   int // estimate of how long job takes
	// execution time = arrivaltime + bursttime
}

type SJF struct {
	queue                  []Process //priority queue sorting based on burst time
	remainingTime          int       // time left for current running process
	processId              int       // current process
	totalWaitingTime       int
	totalProcessesExecuted int
}

func NewSJF() *SJF {
	sjf := new(SJF)
	sjf.queue = make([]Process, 0)
	sjf.processId = -1
	return sjf
}

func (sjf *SJF) run(process *Process, currentTime int) {
	// sort all processes according to arrival time
	// arrival time is end of scheduler len(scheduler) - 1
	// select process that has min arrival time and min burst time
	//add to queue

	//queue = append(queue, Process{id = processId, burstTime = burstTime, arrivalTime = arrivalTime})

	// index 0 = smallest burst time
	// 0 1 2 3        2
	// 0 2 4 5 -- 0 2 2 4 5
	// insert 3 -> 0 2 3 4 5
	// insert a process if there is a process

	if process != nil {
		//process := Process{id = processId, burstTime = burstTime, arrivalTime=arrivalTime}

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
			nextJob := sjf.queue[0]
			sjf.queue = sjf.queue[1:] //pop from queue
			sjf.processId = nextJob.id
			sjf.remainingTime = nextJob.burstTime
			waitingTime := currentTime - nextJob.arrivalTime // nextJob.arrivalTime = start time of next job
			sjf.totalWaitingTime += waitingTime
			sjf.totalProcessesExecuted += 1
		}

	}

	sjf.remainingTime--

	//

}

// get process you want to run at this time
func (sjf *SJF) getProcess() int {
	return sjf.processId
}

func (sjf *SJF) getAvgWaitingTIme() float32 {
	return float32(sjf.totalWaitingTime) / float32(sjf.totalProcessesExecuted)
}

// func Scheduling(stream){

// }
