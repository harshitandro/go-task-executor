package interfaces

type Task interface {
	WaitFor()
	Done() bool
	Result() (Result, error)
	CreationTimeNano() int64
	SubmissionTimeNano() int64
	CompletionTimeNano() int64

	Execute()
	SetSubmissionTimeNano(currentTimeNano int64)
	SetError(error error)
}
