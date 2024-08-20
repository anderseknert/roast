package encoding

import (
	jsoniter "github.com/json-iterator/go"

	_ "github.com/anderseknert/roast/internal/encoding"
)

// JSON returns the fastest jsoniter configuration
// It is preferred using this function instead of jsoniter.ConfigFastest directly
// as there as the init function needs to be called to register the custom types,
// which will happen automatically on import.
func JSON() jsoniter.API {
	return jsoniter.ConfigFastest
}
