package bluecore

import (
	"encoding/json"
	"fmt"
	"time"
)

type Node map[string]any
type Nodes map[string]Node
type CustomTime time.Time

type Journey struct {
	Classification string         `json:"classification"`
	Creator        string         `json:"creator"`
	Id             int            `json:"id"`
	SegmentIds     []int          `json:"segment_ids"`
	TouchToCampMap map[string]int `json:"touch_to_campaign_map"`
	Nodes          Nodes          `json:"nodes"`
	SomeRandomProp *string        `json:"some_random_prop,omitempty"`
	Name           interface{}    `json:"name"`
	Key            Key            `json:"key"`
	Created        CustomTime     `json:"created"`
}

func (t *CustomTime) UnmarshalJSON(data []byte) (err error) {
	var tstr string
	if err = json.Unmarshal(data, &tstr); err != nil {
		return err
	}

	parsed, err := time.Parse("2006-01-02 15:04:05.000000", tstr)
	if err != nil {
		return fmt.Errorf("time parse failed: %v", err)
	}
	*t = CustomTime(parsed)
	return nil
}

type Key struct {
	Name  string
	Other string
	Len   int
}

func (k *Key) UnmarshalJSON(data []byte) (err error) {

	keyData := ""
	if err = json.Unmarshal(data, &keyData); err != nil {
		return err
	}

	k.Name = keyData[0:5]
	k.Other = keyData[5:]
	k.Len = len(k.Name) + len(k.Other)

	return nil
}

func (n *Nodes) UnmarshalJSON(data []byte) (err error) {
	m := map[string]interface{}{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	// Make sure to allocate the map.
	*n = Nodes{}

	// Loop through our nested map object and assign each Node
	for k, v := range m {
		if v == nil {
			(*n)[k] = nil
		} else {
			node := Node{}
			node.AddNode(v.(map[string]interface{}))
			(*n)[k] = node
		}
	}

	return nil
}

func (n *Node) AddNode(m map[string]any) {
	for k, v := range m {
		if v == nil {
			(*n)[k] = nil
		} else {
			node := Node{}
			node.AddNode(v.(map[string]interface{}))
			(*n)[k] = node
		}
	}
}
