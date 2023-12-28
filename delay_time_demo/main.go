package main

import "time"

type Task struct {
	ExecuteTime time.Time

	Job func()
}

func NewTask(executeTime time.Duration, job func()) *Task {
	// 返回一个Task对象
	return &Task{
		ExecuteTime: time.Now().Add(executeTime),
		Job:         job,
	}
}

func (t *Task) Start() {
    time.AfterFunc(time.Until(t.ExecuteTime), t.Job)
}

func main() {
	task := NewTask(10 * time.Second, func() {
		println(time.Now().String())
	})
	task.Start()
	time.Sleep(11 * time.Second)
}