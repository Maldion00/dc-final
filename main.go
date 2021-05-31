package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Maldion00/dc-final/api"
	"github.com/Maldion00/dc-final/controller"
	"github.com/Maldion00/dc-final/scheduler"
)

func main() {
	log.Println("Hi!")

	// Start Controller
	go controller.Start()

	// Start Scheduler
	jobs := make(chan scheduler.Job)
	go scheduler.Start(jobs)
	// Send sample jobs
	sampleJob := scheduler.Job{Address: "localhost:50051", RPCName: "hello"}

	// API
	go api.Start()

	for {
		// add SayHelloFunction
		sampleJob.RPCName = fmt.Sprintf("hello-%v", rand.Intn(10000))
		jobs <- sampleJob
		time.Sleep(time.Second * 5)
	}
}
