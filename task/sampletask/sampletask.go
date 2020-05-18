package sampletask

import (
	"github.com/harshitandro/go-task-executor/interfaces"
	"sync"
	"time"
)

type sampleTask struct {
	doneChan           chan int
	result             interfaces.Result
	err                error
	done               bool
	submissionTimeNano int64
	completionTimeNano int64
	creationTimeNano   int64
	waitGroup          *sync.WaitGroup
	data               int64
}

func (s *sampleTask) WaitFor() {
	<-s.doneChan
}

func (s *sampleTask) SetError(error error) {
	s.err = error
}

func (s sampleTask) Done() bool {
	return s.done
}

func (s sampleTask) Result() (interfaces.Result, error) {
	return s.result, s.err
}

func (s sampleTask) SetSubmissionTimeNano(currentTimeNano int64) {
	s.submissionTimeNano = currentTimeNano
}

func (s sampleTask) CreationTimeNano() int64 {
	return s.creationTimeNano
}

func (s sampleTask) SubmissionTimeNano() int64 {
	return s.submissionTimeNano
}

func (s sampleTask) CompletionTimeNano() int64 {
	return s.completionTimeNano
}

func (s *sampleTask) setCompletionTimeNano(currentTimeNano int64) {
	s.completionTimeNano = currentTimeNano
}

func (s *sampleTask) Execute() {
	defer func() {
		if s.waitGroup != nil {
			s.waitGroup.Done()
		}
	}()
	defer func() { s.done = true }()
	defer s.setCompletionTimeNano(time.Now().UnixNano())
	x := int64(0)
	cpuIntensive(&x, s.data)
	s.result = x
}

func New(waitGroup *sync.WaitGroup, data int64) interfaces.Task {
	task := &sampleTask{
		creationTimeNano: time.Now().UnixNano(),
		waitGroup:        waitGroup,
		data:             data,
	}
	return task
}

func cpuIntensive(p *int64, data int64) {
	for i := int64(1); i <= data; i++ {
		*p = i
	}
}
