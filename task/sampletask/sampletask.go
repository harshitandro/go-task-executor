/*
 * Copyright 2020 Harshit Singh Lodha
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sampletask

import (
	"crypto/sha256"
	"fmt"
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
	var x [32]byte
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

func cpuIntensive(p *[32]byte, data int64) {
	for i := int64(1); i <= data; i++ {
		*p = sha256.Sum256([]byte(fmt.Sprintf("hello : %s")))
	}
}
