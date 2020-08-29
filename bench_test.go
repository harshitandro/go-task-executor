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
	"github.com/harshitandro/go-task-executor/interfaces"
	"github.com/harshitandro/go-task-executor/task/sampletask"
	"os"
	"sync"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	setup()
	defer taskExecutor.Close()
	b.ResetTimer()
	b.StartTimer()

	var wg sync.WaitGroup
	wg.Add(1)
	go createTasks(taskExecutor, &wg, 1000)
	wg.Wait()

	b.StopTimer()
}

func BenchmarkRunSerial(b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	b.ResetTimer()
	b.StartTimer()

	var task interfaces.Task

	for i := 0; i < 1000; i++ {
		task = sampletask.New(nil, int64(i))
		task.Execute()
	}

	b.StopTimer()
}
