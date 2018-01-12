package gring

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJSON(t *testing.T) {
	r := NewFromArray([]int{2, 1, 0, 4, 3})
	r.SetValue(0, "hello")
	r.SetValue(1, "world")
	j, err := json.Marshal(r)
	assert.Nil(t, err)
	assert.Equal(t, `{"head":0,"length":5,"nodes":[{"next":4,"prev":1,"value":"hello"},{"next":0,"prev":2,"value":"world"},{"next":1,"prev":3,"value":null},{"next":2,"prev":4,"value":null},{"next":3,"prev":0,"value":null}]}`, string(j))

	r2 := New()
	err = json.Unmarshal(j, &r2)
	assert.Nil(t, err)
	assert.Equal(t, r.head, r2.head)
	assert.Equal(t, r.length, r2.length)
	assert.Equal(t, r.nodes, r2.nodes)
}
