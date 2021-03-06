/*
 * Copyright 2015 Fabrício Godoy
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package slice

import "strings"

// A String represents a slice of string.
type String []string

// A PredicateStringFunc represents a function that defines a criteria and
// determines whether specified string meets that criteria.
type PredicateStringFunc func(string) bool

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (s String) Equal(other []string) bool {
	if s == nil && other == nil {
		return true
	}
	if s == nil || other == nil {
		return false
	}
	if len(s) != len(other) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (s String) Except(str []string, ignoreCase bool) String {
	var result String
	strBox := String(str)
	for _, v := range s {
		if !strBox.Exists(v, ignoreCase) {
			result = append(result, v)
		}
	}

	return result
}

// IndexOf looks for specified string into current slice, and optionally ignores
// letter casing.
func (s String) IndexOf(str string, ignoreCase bool) int {
	if ignoreCase {
		str = strings.ToLower(str)
		for i, v := range s {
			if str == strings.ToLower(v) {
				return i
			}
		}

		return -1
	}

	for i, v := range s {
		if str == v {
			return i
		}
	}

	return -1
}

// Exists determines whether specified string exists into current slice.
func (s String) Exists(str string, ignoreCase bool) bool {
	return s.IndexOf(str, ignoreCase) != -1
}

// ExistsAll determine whether all specified strings exists into
// current slice.
func (s String) ExistsAll(str []string, ignoreCase bool) bool {
	for _, v := range str {
		if s.IndexOf(v, ignoreCase) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified strings exists into current
// slice.
func (s String) ExistsAny(str []string, ignoreCase bool) bool {
	for _, v := range str {
		if s.IndexOf(v, ignoreCase) != -1 {
			return true
		}
	}

	return false
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (s String) TrueForAll(pred PredicateStringFunc) bool {
	for _, v := range s {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (s String) TrueForAny(pred PredicateStringFunc) bool {
	for _, v := range s {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (s String) Where(pred PredicateStringFunc) String {
	var result String
	for _, v := range s {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}
