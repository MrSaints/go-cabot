// Code generated by "stringer -type=Status"; DO NOT EDIT

package cabot

import "fmt"

const _Status_name = "passingintermittentfailing"

var _Status_index = [...]uint8{0, 7, 19, 26}

func (i Status) String() string {
	if i >= Status(len(_Status_index)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}