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

package executor

import (
	"fmt"
	"github.com/harshitandro/go-task-executor/interfaces"
	"runtime"
	"sync"
	"time"
)

type executor struct {
	taskBuffer     chan *interfaces.Task
	parallelism    int
	taskBufferSize int
	started        bool
	wg             sync.WaitGroup
	name           string
}

func (e *executor) Submit(task *interfaces.Task) (bool, error) {
	(*task).SetSubmissionTimeNano(time.Now().UnixNano())
	select {
	case e.taskBuffer <- task:
		return true, nil
	default:
		return false, fmt.Errorf("unable to submit due to buffer full in executor '%s'", e.name)
	}
}

func (e *executor) SubmitBlocking(task *interfaces.Task) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("'%s' Executor : Unable to submit task due to error : %s\n", e.name, r)
			err = fmt.Errorf("unable to submit to executor '%s'", e.name)
		}
	}()
	(*task).SetSubmissionTimeNano(time.Now().UnixNano())
	e.taskBuffer <- task
	return
}

func (e *executor) Close() (justClosed bool) {

	defer func() {
		if recover() != nil {
			justClosed = false
		} else {
			fmt.Printf("'%s' Executor closed.\n", e.name)
		}
	}()

	close(e.taskBuffer) // panic if ch is closed
	return true
}

func (e *executor) Start() (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error while spawning workers for executor '%s' : %s", e.name, r)
		} else {
			fmt.Printf("'%s' Executor started succesfully.\n", e.name)
		}
	}()

	if e.started == true {
		fmt.Printf("'%s' Executor already started", e.name)
		return nil
	}

	for i := 0; i < e.parallelism; i++ {
		e.wg.Add(1)
		go e.worker(i, &e.wg)
	}
	e.started = true
	e.wg.Wait()
	return
}

func (e *executor) SetParallelism(maxProc int) {
	e.parallelism = maxProc
}

func (e executor) Parallelism() int {
	return e.parallelism
}

func (e *executor) SetTaskBufferSize(bufferSize int) {
	e.taskBufferSize = bufferSize
}

func (e executor) TaskBufferSize() int {
	return e.taskBufferSize
}

func (e *executor) worker(id int, group *sync.WaitGroup) {
	var task *interfaces.Task
	defer func() {
		if r := recover(); r != nil && task != nil {
			fmt.Printf("'%s' Executor : Error while executing task in worker %d : %s\n", e.name, id, r)
			(*task).SetError(fmt.Errorf("Error while executing task in worker %d : %s\n", id, r))
		}
	}()
	fmt.Printf("'%s' Executor : Started worker id %d\n", e.name, id)
	group.Done()
	for task = range e.taskBuffer {
		time.Sleep(0)
		(*task).Execute()
	}
	fmt.Printf("'%s' Executor : Worker id %d exiting\n", e.name, id)
	return
}

func New(executorName string) interfaces.Executor {
	executor := executor{
		taskBufferSize: 10000,
		taskBuffer:     make(chan *interfaces.Task, 10000),
		parallelism:    runtime.GOMAXPROCS(0),
		name:           executorName,
	}
	return &executor
}
