package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	faktory "github.com/yvjessestephens/faktory/client"
	"github.com/yvjessestephens/faktory/util"
)

func TestMutateCommands(t *testing.T) {
	runServer(":7419", func() {
		cl, err := faktory.Open()
		assert.NoError(t, err)
		nfo, err := cl.Info()
		assert.NoError(t, err)
		assert.NotNil(t, nfo)

		err = cl.Clear(faktory.Retries)
		assert.NoError(t, err)

		j := faktory.NewJob("AnotherJob", "truid:67123", 3)
		j.At = util.Thens(time.Now().Add(10 * time.Second))
		err = cl.Push(j)
		assert.NoError(t, err)

		j = faktory.NewJob("SomeJob", "truid:67123", 3)
		j.At = util.Thens(time.Now().Add(10 * time.Second))
		err = cl.Push(j)
		assert.NoError(t, err)

		j = faktory.NewJob("SomeJob", "trid:67123", 5)
		j.At = util.Thens(time.Now().Add(10 * time.Second))
		err = cl.Push(j)
		assert.NoError(t, err)

		j = faktory.NewJob("FooJob", "445", 5)
		j.At = util.Thens(time.Now().Add(10 * time.Second))
		err = cl.Push(j)
		assert.NoError(t, err)
		targetJid := j.Jid

		j = faktory.NewJob("FooJob", "445", 5)
		err = cl.Push(j)
		assert.NoError(t, err)

		state, err := cl.CurrentState()
		assert.NoError(t, err)
		assert.EqualValues(t, 4, state.Data.Sets["scheduled"])

		err = cl.Discard(faktory.Scheduled, faktory.OfType("SomeJob").Matching("*uid:67123*"))
		assert.NoError(t, err)

		state, err = cl.CurrentState()
		assert.NoError(t, err)
		assert.EqualValues(t, 3, state.Data.Sets["scheduled"])
		assert.EqualValues(t, 0, state.Data.Sets["dead"])

		err = cl.Kill(faktory.Scheduled, faktory.OfType("AnotherJob"))
		assert.NoError(t, err)

		err = cl.Kill("", faktory.OfType("AnotherJob"))
		assert.Error(t, err)

		state, err = cl.CurrentState()
		assert.NoError(t, err)
		assert.EqualValues(t, 2, state.Data.Sets["scheduled"])
		assert.EqualValues(t, 1, state.Data.Sets["dead"])
		assert.EqualValues(t, 1, state.Data.Queues["default"])

		err = cl.Requeue(faktory.Scheduled, faktory.WithJids(targetJid))
		assert.NoError(t, err)

		state, err = cl.CurrentState()
		assert.NoError(t, err)
		assert.EqualValues(t, 1, state.Data.Sets["scheduled"])
		assert.EqualValues(t, 1, state.Data.Sets["dead"])
		assert.EqualValues(t, 2, state.Data.Queues["default"])

		err = cl.Clear(faktory.Dead)
		assert.NoError(t, err)

		state, err = cl.CurrentState()
		assert.NoError(t, err)
		assert.EqualValues(t, 0, state.Data.Sets["dead"])

	})
}
