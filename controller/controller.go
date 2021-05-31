package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/Maldion00/dc-final/api"
	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/rep"
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
	var sock mangos.Socket
	var err error
	RecieveMessage(api.SendMessage())
	/*if sock, err = pub.NewSocket(); err != nil {
		die("can't get new pub socket: %s", err)
	}
	if err = sock.Listen(controllerAddress); err != nil {
		die("can't listen on pub socket: %s", err.Error())
	}
	for {
		// Could also use sock.RecvMsg to get header
		d := date()
		log.Printf("Controller: Publishing Date %s\n", d)
		if err = sock.Send([]byte(d)); err != nil {
			die("Failed publishing: %s", err.Error())
		}
		time.Sleep(time.Second * 3)
	}*/
}

func RecieveMessage(apiMessage byte) {

	var sock mangos.Socket
	var err error
	var msg []byte
	if sock, err = rep.NewSocket(); err != nil {
		die("can't get new rep socket: %s", err)
	}
	if err = sock.Listen(apiMessage); err != nil {
		die("can't listen on rep socket: %s", err.Error())
	}
	for {

		msg, err = sock.Recv()
		if err != nil {
			die("cannot receive on rep socket: %s", err.Error())
		}
		if string(msg) == "ACTIVATE" {
			fmt.Println("CONTROLLER: RECEIVED ACTIVATE REQUEST")
		}
	}
}
