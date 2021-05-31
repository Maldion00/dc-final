package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/pub"
	"go.nanomsg.org/mangos/protocol/pull"
)

type workload struct {
	id             string
	filter         string
	name           string
	status         string
	runningJobs    int
	FilteredImages []string
}

var (
	controllerAddress = "tcp://localhost:40899"
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
	return time.Now().Format(time.ANSIC)
}

func Start() {


	RecieveMessage(controllerAddress)

	}
}

func RecieveMessage(url string) {
	var sock mangos.Socket
	var err error
	var msg []byte
	if sock, err = pull.NewSocket(); err != nil {
		die("can't get new pull socket: %s", err)
	}
	if err = sock.Listen(controllerAddress); err != nil {
		die("can't listen on pull socket: %s", err.Error())
	}
	msg, err = sock.Recv()
	if err != nil {
		die("cannot receive from mangos Socket: %s", err.Error())
	}
	fmt.Printf("CONTROLLER: RECEIVED \"%s\"\n", msg)

	if string(msg) == "ACTIVATE" {

		fmt.Println("CONTROLLER: ACTIVATED")

	}
}
