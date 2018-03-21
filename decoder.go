// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"encoding/json"
	"github.com/json-iterator/go"
)

// Decoder is used to decode responses from Elasticsearch.
// Users of elastic can implement their own marshaler for advanced purposes
// and set them per Client (see SetDecoder). If none is specified,
// DefaultDecoder is used.
type Decoder interface {
	Decode(data []byte, v interface{}) error
}

// DefaultDecoder uses json.Unmarshal from the Go standard library
// to decode JSON data.
type DefaultDecoder struct{}

// Decode decodes with json.Unmarshal from the Go standard library.
func (u *DefaultDecoder) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// DefaultDecoder uses jsoniter.ConfigDefault.Unmarshal from json-iterator
// to decode JSON data.
type IteratorDefaultDecoder struct{}

// Decode decodes with jsoniter.ConfigDefault.Unmarshal from the Go standard library.
func (u *IteratorDefaultDecoder) Decode(data []byte, v interface{}) error {
	return jsoniter.ConfigDefault.Unmarshal(data, v)
}

// DefaultDecoder uses jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal from json-iterator
// to decode JSON data.
type IteratorStandardDecoder struct{}

// Decode decodes with jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal from the Go standard library.
func (u *IteratorStandardDecoder) Decode(data []byte, v interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, v)
}

// DefaultDecoder uses jsoniter.ConfigFastest.Unmarshal from json-iterator
// to decode JSON data.
type IteratorFastestDecoder struct{}

// Decode decodes with jsoniter.ConfigFastest.Unmarshal from the Go standard library.
func (u *IteratorFastestDecoder) Decode(data []byte, v interface{}) error {
	return jsoniter.ConfigFastest.Unmarshal(data, v)
}
