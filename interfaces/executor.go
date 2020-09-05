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

package interfaces

/*
	Main executor interface which have all the required functions.
*/
type Executor interface {

	/*
		Submits the given task in non - blocking mode. If the executor if full & task is unable to
		get submitted, the function returns the boolean as false with non null error.
	*/
	Submit(task *Task) (bool, error)

	/*
		Submits the given task in blocking mode. The function blocks until the executor has enough
		space to allocate the given task.
	*/
	SubmitBlocking(task *Task) error

	/*
		Shut down the executor.
	*/
	Close() (justClosed bool)

	/*
		Starts the executor & all the supporting routines.
	*/
	Start() error

	/*
		Set the level of parallelism for this executor. This eventually translates to the number of
		go-routines being used to execute submitted tasks.
	*/
	SetParallelism(maxProc int)

	/*
		Returns the level of parallelism for this executor ie. the number of go-routines being used
		to execute submitted tasks.
	*/
	Parallelism() int

	/*
		Sets the capacity of this executor to hold the tasks submitted to it.
	*/
	SetTaskBufferSize(bufferSize int)

	/*
		Returns the capacity of this executor to hold the tasks submitted to it.
	*/
	TaskBufferSize() int
}
