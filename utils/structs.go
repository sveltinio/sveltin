/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package utils

import (
	"fmt"
	"reflect"

	"github.com/fatih/structtag"
)

// GetStructTags extracts the additional meta data information attached to fields of a struct.
func GetStructTags(v any) ([]*structtag.Tags, error) {
	var tags []*structtag.Tags

	t := reflect.ValueOf(v)
	if t.Kind() == reflect.Ptr {
		t = reflect.Indirect(t)
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s is neither a struct or a pointer to struct", v)
	}

	tt := t.Type()
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)

		fieldTags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			return nil, err
		}
		tags = append(tags, fieldTags)

	}
	return tags, nil
}
