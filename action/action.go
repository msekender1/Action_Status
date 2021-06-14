package action

import (
	"encoding/json"
	"sync"
)

type MyMap struct {
	Info sync.Map // sync.Map for concurrent use
}

type Value struct {
	Avg   float64 // Contains average time of an action item
	Count int     // Contains total count of an action item
}

var mymap *MyMap = &MyMap{} // mymap will contain information of all action items. Key is the action item name in this sync map.

var m sync.Mutex // Mutex is used to provide atomicity in sync Map. Atomicity is needed for sync map load & store operations.

// AddAction function takes the input Json as a string parameter. It returns error if there is any, otherwise returns nil.
func AddAction(input string) error {
	type Item struct {
		Action string  // Contains action name of input json
		Time   float64 // Contains time value of input json
	}
	var item Item
	err := json.Unmarshal([]byte(input), &item) // Json unmarshalling of input json
	if err != nil {
		return err // Error in Json unmarshalling of input json
	}
	var myval Value
	m.Lock()
	val, _ := mymap.Info.Load(item.Action)
	if val == nil {
		myval.Avg = float64(item.Time) // This action item was not in mymap earlier
		myval.Count = 1
		mymap.Info.Store(item.Action, myval) // Storing input Json value in mymap
	} else {
		var value Value = val.(Value)                                                               // This action item was in mymap earlier
		myval.Avg = ((value.Avg * (float64(value.Count))) + item.Time) / (float64(value.Count + 1)) // New average value calculation
		myval.Count += 1
		mymap.Info.Store(item.Action, myval) // Updating mymap for new input Json value
	}
	m.Unlock()
	return nil // Action adding is successful
}

// GetStats function returns average value for each action in Json format
func GetStats() string {
	type Record struct {
		Action string  `json:"action"` // Contains name of an action item
		Avg    float64 `json:"avg"`    // Contains average time of an action item
	}
	var record Record
	count := 0 // count contains total number of action items in sync map
	m.Lock()
	mymap.Info.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	slc := make([]Record, count) // Slice of Record
	count = 0
	mymap.Info.Range(func(k, v interface{}) bool {
		var value Value = v.(Value)
		record.Action = k.(string)
		record.Avg = value.Avg
		slc[count] = record // count is used to provide direct access in the slice and thus using "append()" is avoided to get faster performance
		count++
		return true
	})
	m.Unlock()
	result, err := json.MarshalIndent(slc, "", "\t") // Json marshalling to get the output Json.
	if err != nil {
		return ("Error: " + err.Error()) // Error in Json marshalling
	}
	return string(result) // Output Json is returned
}
