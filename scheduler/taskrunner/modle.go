package taskrunner

const (
	READY_DISPATCH = "d"
	READY_EXECTUE  = "e"
	CLOSE          = "c"

	VIDEO_DIRECTORY = "./videos/"
)

type controlChan chan string
type dataChan chan interface{}
type fn func(dc dataChan) error
