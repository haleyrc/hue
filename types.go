package hue

import (
	"encoding/json"
)

type Update struct {
	State       string `json:"state"`
	LastInstall string `json:"lastinstall"`
}

type ColorTemperature struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Control struct {
	MinDimLevel      int              `json:"mindimlevel"`
	MaxLumen         int              `json:"maxlumen"`
	ColorGamutType   string           `json:"colorgamuttype"`
	ColorGamut       [][]float64      `json:"colorgamut"`
	ColorTemperature ColorTemperature `json:"ct"`
}

type Streaming struct {
	Renderer bool `json:"renderer"`
	Proxy    bool `json:"proxy"`
}

type Capabilities struct {
	Certified bool      `json:"certified"`
	Control   Control   `json:"control"`
	Streaming Streaming `json:"streaming"`
}

type Startup struct {
	Mode       string `json:"mode"`
	Configured bool   `json:"configured"`
}

type Config struct {
	Archetype string  `json:"archetype"`
	Function  string  `json:"function"`
	Direction string  `json:"direction"`
	Startup   Startup `json:"startup"`
}

// Light describes the properties of a single Phillips Hue lightbulb.
type Light struct {
	// ID is the simple string ID the bulbs are keyed by.
	ID string `json:"-"`

	// State is the total state of the bulb at the time of query.
	State State `json:"state"`

	Update Update `json:"swupdate"`

	Type string `json:"type"`

	// Name is the user-set nickname for a bulb.
	Name string `json:"name"`

	ModelID string `json:"modelid"`

	ManufacturerName string `json:"manufacturername"`

	ProductName string `json:"productname"`

	Capabilities Capabilities `json:"capabilities"`

	Config Config `json:"config"`

	// UniqueID is the MAC address-like ID that uniquely identifies a single
	// bulb regardless of configuration.
	UniqueID string `json:"uniqueid"`

	SoftwareVersion string `json:"swversion"`

	SoftwareConfigID string `json:"swconfigid"`

	ProductID string `json:"productid"`
}

// String implements fmt.Stringer.
func (l *Light) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

// State describes the accumulated state of a single Phillips Hue bulb. We only
// work with hue and saturation at the moment, since they are the easiest
// without doing a lot of conversion.
type State struct {
	// On specifies whether the bulb is on or off.
	On bool `json:"on"`
	// Brightness is a value between 1 and 254. 1 is the lowest the bulb can
	// produce, but is not off.
	Brightness int `json:"bri"`
	// Hue is a value between 0 and 65535. Both 0 and 65535 are red, 25500 is
	// green, and 46920 is blue.
	Hue int `json:"hue"`
	// Saturation is the color saturation of the light. 254 is the most
	// saturated and 0 is the least (white).
	Saturation int `json:"sat"`
	// Alert is the last alert sent to the light. It is either "none", "select",
	// or "lselect".
	Alert string `json:"alert"`
	// Effect is currently either "none" or "colorloop".
	Effect string `json:"effect"`
	// Reachable indicates the bulb is reachable from the bridge (and can thus
	// be controlled).
	Reachable bool `json:"reachable"`
}

// String implements fmt.Stringer.
func (s *State) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// Group represents a group of Phillips Hue lights. The group may either be
// user- or system-defined. The group with ID 0 is the master group and can be
// used to control all lights regardless of their other group or room
// associations.
type Group struct {
	// ID is the string ID of the group.
	ID string `json:"-"`
	// Name is the user- or system-defined name of the group.
	Name string `json:"name"`
	// Lights is a list of light IDs that belong to the group.
	Lights []string `json:"lights"`
	// GroupType is the type of group.
	GroupType string `json:"type"`
	// Action is the last state command issued to the group.
	Action State `json:"action"`
}

// String implements fmt.Stringer.
func (g *Group) String() string {
	b, _ := json.Marshal(g)
	return string(b)
}
