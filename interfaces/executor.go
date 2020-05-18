package interfaces

type Executor interface {
	Submit(task *Task) (bool, error)
	SubmitBlocking(task *Task) error
	Close() (justClosed bool)
	Start() error
	SetParallelism(maxProc int)
	Parallelism() int
	SetTaskBufferSize(bufferSize int)
	TaskBufferSize() int
}
