/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package npmc ...
package npmc

import (
	"bytes"
	"encoding/json"

	"github.com/spf13/afero"
)

// WriteToFile saves a json string to a file on the file system
func WriteToFile(appFS afero.Fs, pkg *PackageJSON, saveAs string, prefix string, indent string) error {
	file, _ := jsonMarshalIndent(pkg, prefix, indent)
	return afero.WriteFile(appFS, saveAs, file, 0644)
}

func jsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func jsonMarshalIndent(t interface{}, prefix, indent string) ([]byte, error) {
	b, err := jsonMarshal(t)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
