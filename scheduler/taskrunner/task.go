package taskrunner

import (
	"errors"
	"goStreamerServer/scheduler/data"
	"log"
	"os"
	"sync"
)

func delVideo(vid string) error {
	err := os.Remove(VIDEO_DIRECTORY + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Println("delete video failed:", err)
		return err
	}
	return nil
}
func VideoClearDispatcher(dc dataChan) error {
	vids, err := data.ReadVideo(5)
	if err != nil {
		log.Println("get vide failed:", err)
		return err
	}
	if len(vids) == 0 {
		log.Println("videos length is 0")
		return errors.New("videos length is 0")
	}

	for _, v := range vids {
		dc <- v
	}
	return nil
}

func VideoClearExectuer(dc dataChan) error {
	errMap := sync.Map{}
	var errs error
mainloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				err := delVideo(id.(string))
				if err != nil {
					errMap.Store(id.(string), err)
					return

				}
				err = data.DelVideo(id.(string))
				if err != nil {
					errMap.Store(id.(string), err)
					return
				}
			}(vid)

		default:
			break mainloop
		}
	}

	errMap.Range(func(k, v interface{}) bool {
		if v.(error) != nil {
			errs = v.(error)
			return false
		}
		return true
	})

	return errs
}
