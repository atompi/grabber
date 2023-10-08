package execute

import (
	"sync"

	"github.com/atompi/grabber/internal/options"
	"github.com/atompi/grabber/tools"
)

type executer struct {
	rowChan chan options.FileOptions
	wg      *sync.WaitGroup
}

func NewExecuter(ch chan options.FileOptions, wg *sync.WaitGroup) *executer {
	return &executer{
		rowChan: ch,
		wg:      wg,
	}
}

func (e *executer) Exec() {
	for rCh := range e.rowChan {
		e.handle(rCh)
	}
	e.wg.Done()
}

func (e *executer) handle(opts options.FileOptions) {
	err := tools.DownloadFile(opts.Dest, opts.Src)
	if err != nil {
		panic(err)
	}
}
