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

package main

import (
	"fmt"
	"github.com/harshitandro/go-task-executor/executor"
	"github.com/harshitandro/go-task-executor/interfaces"
	"github.com/harshitandro/go-task-executor/task/sampletask"
	"math"
	"sync"
	"time"
)

var taskExecutor interfaces.Executor

func setup() {
	taskExecutor = executor.New("Sample")
	taskExecutor.SetParallelism(6)
	taskExecutor.Start()
}

func main() {
	setup()
	defer taskExecutor.Close()

	run(5000)
	taskExecutor.Close()
	runSerial(5000)
}

func createTasks(executor interfaces.Executor, group *sync.WaitGroup, totalTaskCount int) {
	defer group.Done()
	var task interfaces.Task
	for i := 0; i < totalTaskCount; i++ {
		group.Add(1)
		task = sampletask.New(group, int64(i))
		executor.SubmitBlocking(&task)
	}
}

func run(totalTaskCount int) {
	var wg sync.WaitGroup
	wg.Add(1)
	startTime := time.Now().UnixNano()
	go createTasks(taskExecutor, &wg, totalTaskCount)
	wg.Wait()
	endTime := time.Now().UnixNano()
	fmt.Println("Total time taken          : ", float64(endTime-startTime)*math.Pow(10, -6), "ms")
}

func runSerial(totalTaskCount int) {
	var task interfaces.Task
	startTime := time.Now().UnixNano()
	for i := 0; i < totalTaskCount; i++ {
		task = sampletask.New(nil, int64(i))
		task.Execute()
	}
	endTime := time.Now().UnixNano()
	fmt.Println("Total time taken serially : ", float64(endTime-startTime)*math.Pow(10, -6), "ms")
}
