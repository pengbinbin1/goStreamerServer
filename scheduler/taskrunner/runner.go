package taskrunner

import (
	"log"
)

type Runner struct {
	Controller controlChan
	Error      controlChan
	Size       int
	Data       dataChan
	LongLived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longLived bool, execute fn, dispatcher fn) *Runner {
	runner := &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		Size:       size,
		LongLived:  longLived,
		Dispatcher: dispatcher,
		Executor:   execute,
	}
	return runner
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.LongLived {
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()
	for {
		select {
		case c := <-r.Controller:
			if c == READY_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					log.Println("dispathcer failed:", err)
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_EXECTUE
				}

			}
			if c == READY_EXECTUE {
				err := r.Executor(r.Data)
				if err != nil {
					log.Println("exectue failed:", err)

					r.Error <- CLOSE
				} else {
					r.Controller <- READY_DISPATCH

				}

			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:
		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_DISPATCH
	r.startDispatch()

}

create table coments(
	id int(64) not null auto increment,
	authID varchar (64),
	videoID varchar,
	time datetime default current_timestamp,
	primary key(id)
)