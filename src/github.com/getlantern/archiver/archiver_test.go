package archiver

import (
	"testing"
	"time"

	"github.com/getlantern/testify/assert"
)

var total int

type Counter struct {
	i int
}

func (c *Counter) Append(interface{}) error {
	c.i = c.i + 1
	return nil
}

func (c *Counter) SaveTemp() (d interface{}, err error) {
	return c.i, nil
}

func (c *Counter) LoadTemp(d interface{}) error {
	c.i = d.(int)
	return nil
}

func (c *Counter) Archive() error {
	total = total + c.i
	return nil
}

func TestAdder(t *testing.T) {
	c := Counter{}
	New(&c, 5*time.Millisecond)
	c.Append(nil)
	c.Append(nil)

	assert.Equal(t, total, 0, "should not archive before the time")
	time.Sleep(19 * time.Millisecond)
	assert.Equal(t, c.i, 2, "should keep value")
	assert.Equal(t, total, 6, "should archived 3 times")

	/*SaveTemp()
	c.Append(nil)
	LoadTemp()
	assert.Equal(t, c.i, 2, "should keep value")
	assert.Equal(t, total, 6, "should archived 3 times")*/
}
