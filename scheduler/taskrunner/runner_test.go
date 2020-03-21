package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {

	dispatch := func(data dataChan) error {
		for i := 0; i < 10; i++ {
			log.Println("dispatch send data:", i)
			data <- i
		}
		return nil
	}

	executor := func(data dataChan) error {
	forLoop:
		for {
			select {
			case element := <-data:
				{
					log.Println("executer received:", element)
				}
			default:
				break forLoop
			}

		}
		return errors.New("received all data")
	}

	r := NewRunner(10, false, executor, dispatch)
	r.StartAll()
	time.Sleep(1 * time.Second)
}
