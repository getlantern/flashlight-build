package archiver

import (
	"sync"
	"time"

	"github.com/rakyll/ticktock"
	"github.com/rakyll/ticktock/t"
)

// Archivable is can archive periodically
type Archivable interface {
	Append(item interface{}) error
	Archive() error
	SaveTemp() (d interface{}, err error)
	LoadTemp(d interface{}) error
}

type job struct {
	a Archivable
}

func (j *job) Run() error {
	return j.a.Archive()
}

var muArchivers sync.RWMutex
var archivers []Archivable
var saver Saver

func SaveTo(s Saver) {
	saver = s
}

func New(ar Archivable, d time.Duration) int {
	j := job{ar}
	ticktock.Schedule("asdfaf", &j, &t.When{
		Every: t.Every(int(d / time.Millisecond)).Milliseconds()})

	muArchivers.Lock()
	defer muArchivers.Unlock()
	archivers = append(archivers, j.a)
	return 0
}

/*func SaveTemp() {
	muArchivers.RLock()
	defer muArchivers.RUnlock()
	dataToSave := make([]interface{}, len(archivers))
	for i, a := range archivers {
		if d, err := a.SaveTemp(); err == nil {
			dataToSave[i] = d
		}
	}
	saver.Save(dataToSave)
}

func LoadTemp() {
	data := make([]interface{}, 0)
	muArchivers.Lock()
	defer muArchivers.Unlock()
	saver.Load(data)
	for _, d := range data {
		s.Save(d)
	}
}*/

func init() {
	ticktock.Start()
}
