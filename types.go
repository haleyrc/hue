package hue

import "fmt"

// Light describes the properties of a single Phillips Hue lightbulb.
type Light struct {
	// ID is the simple string ID the bulbs are keyed by.
	ID string `json:"-"`
	// UniqueID is the MAC address-like ID that uniquely identifies a single
	// bulb regardless of configuration.
	UniqueID string `json:"uniqueid"`
	// Name is the user-set nickname for a bulb.
	Name string `json:"name"`
	// State is the total state of the bulb at the time of query.
	State *State `json:"state"`
}

// String implements fmt.Stringer.
func (l *Light) String() string {
	return fmt.Sprintf(
		`<id=%s,uid=%s,name=%s,state:%s>`,
		l.ID, l.UniqueID, l.Name, l.State,
	)
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
	return fmt.Sprintf(
		`<on=%t,bri=,%d,hue=%d,sat=%d,alt=%s,eff=%s,rea=%t>`,
		s.On, s.Brightness, s.Hue, s.Saturation, s.Alert, s.Effect, s.Reachable,
	)
}
