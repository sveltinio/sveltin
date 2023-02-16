/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package common contains utility functions for collections, maps and filesystem.
package common

// Contains returns true if an element is in a slice.
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// Unique removes duplicated entries from a slice of strings.
func Unique(s []string) []string {
	inResult := make(map[string]bool)
	uniqueValues := []string{}
	for _, elem := range s {
		if len(elem) != 0 {
			if _, value := inResult[elem]; !value {
				inResult[elem] = true
				uniqueValues = append(uniqueValues, elem)
			}
		}
	}
	return uniqueValues
}

// Union returns a slice containing the unique values of the input slices.
func Union(a, b []string) []string {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return Unique(a)
}

// RemoveEmpty delete an empty value in a slice of strings.
func RemoveEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
