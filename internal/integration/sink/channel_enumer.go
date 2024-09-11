// Code generated by "enumer -type Channel -trimprefix Channel -transform snake"; DO NOT EDIT.

package sink

import (
	"fmt"
	"strings"
)

const _ChannelName = "milestonequota"

var _ChannelIndex = [...]uint8{0, 9, 14}

const _ChannelLowerName = "milestonequota"

func (i Channel) String() string {
	if i < 0 || i >= Channel(len(_ChannelIndex)-1) {
		return fmt.Sprintf("Channel(%d)", i)
	}
	return _ChannelName[_ChannelIndex[i]:_ChannelIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ChannelNoOp() {
	var x [1]struct{}
	_ = x[ChannelMilestone-(0)]
	_ = x[ChannelQuota-(1)]
}

var _ChannelValues = []Channel{ChannelMilestone, ChannelQuota}

var _ChannelNameToValueMap = map[string]Channel{
	_ChannelName[0:9]:       ChannelMilestone,
	_ChannelLowerName[0:9]:  ChannelMilestone,
	_ChannelName[9:14]:      ChannelQuota,
	_ChannelLowerName[9:14]: ChannelQuota,
}

var _ChannelNames = []string{
	_ChannelName[0:9],
	_ChannelName[9:14],
}

// ChannelString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ChannelString(s string) (Channel, error) {
	if val, ok := _ChannelNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ChannelNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Channel values", s)
}

// ChannelValues returns all values of the enum
func ChannelValues() []Channel {
	return _ChannelValues
}

// ChannelStrings returns a slice of all String values of the enum
func ChannelStrings() []string {
	strs := make([]string, len(_ChannelNames))
	copy(strs, _ChannelNames)
	return strs
}

// IsAChannel returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Channel) IsAChannel() bool {
	for _, v := range _ChannelValues {
		if i == v {
			return true
		}
	}
	return false
}