package gring

import (
	"encoding/json"
	"github.com/atedja/gmap"
)

// Returns a map[string]interface{}
func (r *Ring) marshalMap() (map[string]interface{}, error) {
	mp := map[string]interface{}{}
	mp["head"] = r.head
	mp["length"] = r.length
	nodes := make([]map[string]interface{}, len(r.nodes))
	for i, v := range r.nodes {
		nodes[i] = map[string]interface{}{}
		nodes[i]["next"] = v.next
		nodes[i]["prev"] = v.prev
		nodes[i]["value"] = v.value
	}
	mp["nodes"] = nodes
	return mp, nil
}

func (r *Ring) unmarshalMap(mp map[string]interface{}) error {
	var err error

	gmp := gmap.Map(mp)
	r.head, err = gmp.Int("head", r.head)
	r.length, err = gmp.Int("length", r.length)
	mpnodes, err := gmp.Array("nodes", nil)
	nodeLength := len(mpnodes)
	if nodeLength > 0 {
		nodes := make([]*node, nodeLength)
		for i, v := range mpnodes {
			nmp := gmap.Map(v.(map[string]interface{}))
			nodes[i] = &node{}
			nodes[i].next, err = nmp.Int("next", -1)
			nodes[i].prev, err = nmp.Int("prev", -1)
			nodes[i].value = nmp["value"]
		}
		r.nodes = nodes
	}

	return err
}

func (r *Ring) MarshalJSON() ([]byte, error) {
	mp, err := r.marshalMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(mp)
}

func (r *Ring) UnmarshalJSON(data []byte) error {
	mp := map[string]interface{}{}
	json.Unmarshal(data, &mp)
	return r.unmarshalMap(mp)
}
