package main

import (
	"fmt"
	"go-task-executor/executor"
	"go-task-executor/interfaces"
	"go-task-executor/task/sampletask"
	"math"
	"sync"
	"time"
)

var taskExecutor interfaces.Executor

func setup() {
	taskExecutor = executor.New("Sample")
	taskExecutor.SetParallelism(4)
	taskExecutor.Start()
}

func main() {
	setup()
	defer taskExecutor.Close()

	run(50000)
	run_serial(50000)
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

func run_serial(totalTaskCount int) {
	var task interfaces.Task
	startTime := time.Now().UnixNano()
	for i := 0; i < totalTaskCount; i++ {
		task = sampletask.New(nil, int64(i))
		task.Execute()
	}
	endTime := time.Now().UnixNano()
	fmt.Println("Total time taken serially : ", float64(endTime-startTime)*math.Pow(10, -6), "ms")
}
