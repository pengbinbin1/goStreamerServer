package taskrunner

import (
	"log"
	"time"
)

type Worker struct {
	ticker *time.Ticker
	r      *Runner
}

func NewWorker(t time.Duration, r *Runner) *Worker {
	w := &Worker{}
	w.ticker = time.NewTicker(t * time.Second)
	w.r = r
	return w
}

func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			log.Println("start a new time ticker")
			go w.r.StartAll()
		}
	}

}

func Start() {
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExectuer)
	w := NewWorker(3, r)
	w.startWorker()
}
