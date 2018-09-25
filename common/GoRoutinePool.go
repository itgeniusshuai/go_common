package common

const MAX_QUEUE_SIZE = 10000

var queue = make(chan Job, MAX_QUEUE_SIZE)

type Job interface {
	Do() error
}
// define job channel
type JobChan chan Job

// define worker channer
type WorkerChan chan JobChan

var (
	JobQueue          JobChan
	WorkerPool        WorkerChan
)

type Worker struct {
	JobChannel JobChan
	quit       chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			// regist current job channel to worker pool
			WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if err := job.Do(); err != nil {
					Println("excute job failed with err: %v", err)
				}
				// recieve quit event, stop worker
			case <-w.quit:
				return
			}
		}
	}()
}

type Dispatcher struct {
	Workers []*Worker
	quit    chan bool
	PoolSize int
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.PoolSize; i++ {
		worker := NewWorker()
		d.Workers = append(d.Workers, worker)
		worker.Start()
	}

	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChan := <-WorkerPool
				jobChan <- job
			}(job)
			// stop dispatcher
		case <-d.quit:
			return
		}
	}
}

func NewWorker() *Worker{
	worker := Worker{}
	return &worker
}

func MakeGoroutinePoolAndRun(poolSize int) *Dispatcher{
	dispatcher := Dispatcher{PoolSize:poolSize}
	go dispatcher.Run()
	return &dispatcher
}

func SendToPool(job Job){
	JobQueue <- job
}

func (this  *Dispatcher)Close(){
	for _,e := range this.Workers{
		e.quit <- true
	}
	this.quit <- true
}

