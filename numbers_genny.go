// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slice

// A PredicateFloat32Func represents a function that defines a criteria and
// determines whether specified float32 meets that criteria.
type PredicateFloat32Func func(float32) bool

// A NFloat32 represents a slice of float32.
type NFloat32 []float32

// Average returns the average value from all elements of current slice.
func (n NFloat32) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NFloat32) Equal(other []float32) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NFloat32) Except(a ...float32) NFloat32 {
	var result NFloat32
	aBox := NFloat32(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified float32 exists into current slice.
func (n NFloat32) Exists(a float32) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Float32s exists into
// current slice.
func (n NFloat32) ExistsAll(a ...float32) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Float32s exists into current
// slice.
func (n NFloat32) ExistsAny(a ...float32) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified float32 into current slice.
func (n NFloat32) IndexOf(a float32) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NFloat32) Sum() float32 {
	var sum float32
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NFloat32) TrueForAll(pred PredicateFloat32Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NFloat32) TrueForAny(pred PredicateFloat32Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NFloat32) Where(pred PredicateFloat32Func) NFloat32 {
	var result NFloat32
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateFloat64Func represents a function that defines a criteria and
// determines whether specified float64 meets that criteria.
type PredicateFloat64Func func(float64) bool

// A NFloat64 represents a slice of float64.
type NFloat64 []float64

// Average returns the average value from all elements of current slice.
func (n NFloat64) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NFloat64) Equal(other []float64) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NFloat64) Except(a ...float64) NFloat64 {
	var result NFloat64
	aBox := NFloat64(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified float64 exists into current slice.
func (n NFloat64) Exists(a float64) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Float64s exists into
// current slice.
func (n NFloat64) ExistsAll(a ...float64) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Float64s exists into current
// slice.
func (n NFloat64) ExistsAny(a ...float64) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified float64 into current slice.
func (n NFloat64) IndexOf(a float64) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NFloat64) Sum() float64 {
	var sum float64
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NFloat64) TrueForAll(pred PredicateFloat64Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NFloat64) TrueForAny(pred PredicateFloat64Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NFloat64) Where(pred PredicateFloat64Func) NFloat64 {
	var result NFloat64
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateIntFunc represents a function that defines a criteria and
// determines whether specified int meets that criteria.
type PredicateIntFunc func(int) bool

// A NInt represents a slice of int.
type NInt []int

// Average returns the average value from all elements of current slice.
func (n NInt) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NInt) Equal(other []int) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NInt) Except(a ...int) NInt {
	var result NInt
	aBox := NInt(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified int exists into current slice.
func (n NInt) Exists(a int) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Ints exists into
// current slice.
func (n NInt) ExistsAll(a ...int) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Ints exists into current
// slice.
func (n NInt) ExistsAny(a ...int) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified int into current slice.
func (n NInt) IndexOf(a int) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NInt) Sum() int {
	var sum int
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NInt) TrueForAll(pred PredicateIntFunc) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NInt) TrueForAny(pred PredicateIntFunc) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NInt) Where(pred PredicateIntFunc) NInt {
	var result NInt
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateInt16Func represents a function that defines a criteria and
// determines whether specified int16 meets that criteria.
type PredicateInt16Func func(int16) bool

// A NInt16 represents a slice of int16.
type NInt16 []int16

// Average returns the average value from all elements of current slice.
func (n NInt16) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NInt16) Equal(other []int16) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NInt16) Except(a ...int16) NInt16 {
	var result NInt16
	aBox := NInt16(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified int16 exists into current slice.
func (n NInt16) Exists(a int16) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Int16s exists into
// current slice.
func (n NInt16) ExistsAll(a ...int16) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Int16s exists into current
// slice.
func (n NInt16) ExistsAny(a ...int16) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified int16 into current slice.
func (n NInt16) IndexOf(a int16) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NInt16) Sum() int16 {
	var sum int16
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NInt16) TrueForAll(pred PredicateInt16Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NInt16) TrueForAny(pred PredicateInt16Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NInt16) Where(pred PredicateInt16Func) NInt16 {
	var result NInt16
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateInt32Func represents a function that defines a criteria and
// determines whether specified int32 meets that criteria.
type PredicateInt32Func func(int32) bool

// A NInt32 represents a slice of int32.
type NInt32 []int32

// Average returns the average value from all elements of current slice.
func (n NInt32) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NInt32) Equal(other []int32) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NInt32) Except(a ...int32) NInt32 {
	var result NInt32
	aBox := NInt32(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified int32 exists into current slice.
func (n NInt32) Exists(a int32) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Int32s exists into
// current slice.
func (n NInt32) ExistsAll(a ...int32) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Int32s exists into current
// slice.
func (n NInt32) ExistsAny(a ...int32) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified int32 into current slice.
func (n NInt32) IndexOf(a int32) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NInt32) Sum() int32 {
	var sum int32
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NInt32) TrueForAll(pred PredicateInt32Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NInt32) TrueForAny(pred PredicateInt32Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NInt32) Where(pred PredicateInt32Func) NInt32 {
	var result NInt32
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateInt64Func represents a function that defines a criteria and
// determines whether specified int64 meets that criteria.
type PredicateInt64Func func(int64) bool

// A NInt64 represents a slice of int64.
type NInt64 []int64

// Average returns the average value from all elements of current slice.
func (n NInt64) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NInt64) Equal(other []int64) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NInt64) Except(a ...int64) NInt64 {
	var result NInt64
	aBox := NInt64(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified int64 exists into current slice.
func (n NInt64) Exists(a int64) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Int64s exists into
// current slice.
func (n NInt64) ExistsAll(a ...int64) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Int64s exists into current
// slice.
func (n NInt64) ExistsAny(a ...int64) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified int64 into current slice.
func (n NInt64) IndexOf(a int64) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NInt64) Sum() int64 {
	var sum int64
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NInt64) TrueForAll(pred PredicateInt64Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NInt64) TrueForAny(pred PredicateInt64Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NInt64) Where(pred PredicateInt64Func) NInt64 {
	var result NInt64
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateInt8Func represents a function that defines a criteria and
// determines whether specified int8 meets that criteria.
type PredicateInt8Func func(int8) bool

// A NInt8 represents a slice of int8.
type NInt8 []int8

// Average returns the average value from all elements of current slice.
func (n NInt8) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NInt8) Equal(other []int8) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NInt8) Except(a ...int8) NInt8 {
	var result NInt8
	aBox := NInt8(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified int8 exists into current slice.
func (n NInt8) Exists(a int8) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Int8s exists into
// current slice.
func (n NInt8) ExistsAll(a ...int8) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Int8s exists into current
// slice.
func (n NInt8) ExistsAny(a ...int8) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified int8 into current slice.
func (n NInt8) IndexOf(a int8) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NInt8) Sum() int8 {
	var sum int8
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NInt8) TrueForAll(pred PredicateInt8Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NInt8) TrueForAny(pred PredicateInt8Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NInt8) Where(pred PredicateInt8Func) NInt8 {
	var result NInt8
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateUintFunc represents a function that defines a criteria and
// determines whether specified uint meets that criteria.
type PredicateUintFunc func(uint) bool

// A NUint represents a slice of uint.
type NUint []uint

// Average returns the average value from all elements of current slice.
func (n NUint) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NUint) Equal(other []uint) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NUint) Except(a ...uint) NUint {
	var result NUint
	aBox := NUint(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified uint exists into current slice.
func (n NUint) Exists(a uint) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Uints exists into
// current slice.
func (n NUint) ExistsAll(a ...uint) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Uints exists into current
// slice.
func (n NUint) ExistsAny(a ...uint) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified uint into current slice.
func (n NUint) IndexOf(a uint) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NUint) Sum() uint {
	var sum uint
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NUint) TrueForAll(pred PredicateUintFunc) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NUint) TrueForAny(pred PredicateUintFunc) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NUint) Where(pred PredicateUintFunc) NUint {
	var result NUint
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateUint16Func represents a function that defines a criteria and
// determines whether specified uint16 meets that criteria.
type PredicateUint16Func func(uint16) bool

// A NUint16 represents a slice of uint16.
type NUint16 []uint16

// Average returns the average value from all elements of current slice.
func (n NUint16) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NUint16) Equal(other []uint16) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NUint16) Except(a ...uint16) NUint16 {
	var result NUint16
	aBox := NUint16(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified uint16 exists into current slice.
func (n NUint16) Exists(a uint16) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Uint16s exists into
// current slice.
func (n NUint16) ExistsAll(a ...uint16) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Uint16s exists into current
// slice.
func (n NUint16) ExistsAny(a ...uint16) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified uint16 into current slice.
func (n NUint16) IndexOf(a uint16) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NUint16) Sum() uint16 {
	var sum uint16
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NUint16) TrueForAll(pred PredicateUint16Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NUint16) TrueForAny(pred PredicateUint16Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NUint16) Where(pred PredicateUint16Func) NUint16 {
	var result NUint16
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateUint32Func represents a function that defines a criteria and
// determines whether specified uint32 meets that criteria.
type PredicateUint32Func func(uint32) bool

// A NUint32 represents a slice of uint32.
type NUint32 []uint32

// Average returns the average value from all elements of current slice.
func (n NUint32) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NUint32) Equal(other []uint32) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NUint32) Except(a ...uint32) NUint32 {
	var result NUint32
	aBox := NUint32(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified uint32 exists into current slice.
func (n NUint32) Exists(a uint32) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Uint32s exists into
// current slice.
func (n NUint32) ExistsAll(a ...uint32) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Uint32s exists into current
// slice.
func (n NUint32) ExistsAny(a ...uint32) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified uint32 into current slice.
func (n NUint32) IndexOf(a uint32) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NUint32) Sum() uint32 {
	var sum uint32
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NUint32) TrueForAll(pred PredicateUint32Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NUint32) TrueForAny(pred PredicateUint32Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NUint32) Where(pred PredicateUint32Func) NUint32 {
	var result NUint32
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateUint64Func represents a function that defines a criteria and
// determines whether specified uint64 meets that criteria.
type PredicateUint64Func func(uint64) bool

// A NUint64 represents a slice of uint64.
type NUint64 []uint64

// Average returns the average value from all elements of current slice.
func (n NUint64) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NUint64) Equal(other []uint64) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NUint64) Except(a ...uint64) NUint64 {
	var result NUint64
	aBox := NUint64(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified uint64 exists into current slice.
func (n NUint64) Exists(a uint64) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Uint64s exists into
// current slice.
func (n NUint64) ExistsAll(a ...uint64) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Uint64s exists into current
// slice.
func (n NUint64) ExistsAny(a ...uint64) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified uint64 into current slice.
func (n NUint64) IndexOf(a uint64) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NUint64) Sum() uint64 {
	var sum uint64
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NUint64) TrueForAll(pred PredicateUint64Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NUint64) TrueForAny(pred PredicateUint64Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NUint64) Where(pred PredicateUint64Func) NUint64 {
	var result NUint64
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}

// A PredicateUint8Func represents a function that defines a criteria and
// determines whether specified uint8 meets that criteria.
type PredicateUint8Func func(uint8) bool

// A NUint8 represents a slice of uint8.
type NUint8 []uint8

// Average returns the average value from all elements of current slice.
func (n NUint8) Average() float64 {
	if n == nil || len(n) == 0 {
		return 0
	}

	var result float64
	for _, v := range n {
		result += float64(v)
	}

	result /= float64(len(n))
	return result
}

// Equal returns true whether all elements of specified slice match the ones from current
// slice.
func (n NUint8) Equal(other []uint8) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil || other == nil {
		return false
	}
	if len(n) != len(other) {
		return false
	}

	for i := 0; i < len(n); i++ {
		if n[i] != other[i] {
			return false
		}
	}

	return true
}

// Except returns the difference of two slices. Any element specified will be ignored
// from current slice to the returned one.
func (n NUint8) Except(a ...uint8) NUint8 {
	var result NUint8
	aBox := NUint8(a)
	for _, v := range n {
		if !aBox.Exists(v) {
			result = append(result, v)
		}
	}

	return result
}

// Exists determines whether specified uint8 exists into current slice.
func (n NUint8) Exists(a uint8) bool {
	return n.IndexOf(a) != -1
}

// ExistsAll determine whether all specified Uint8s exists into
// current slice.
func (n NUint8) ExistsAll(a ...uint8) bool {
	for _, v := range a {
		if n.IndexOf(v) == -1 {
			return false
		}
	}

	return true
}

// ExistsAny determine whether any of specified Uint8s exists into current
// slice.
func (n NUint8) ExistsAny(a ...uint8) bool {
	for _, v := range a {
		if n.IndexOf(v) != -1 {
			return true
		}
	}

	return false
}

// IndexOf looks for specified uint8 into current slice.
func (n NUint8) IndexOf(a uint8) int {
	for i, v := range n {
		if v == a {
			return i
		}
	}

	return -1
}

// Sum all elements and returns the result.
func (n NUint8) Sum() uint8 {
	var sum uint8
	for _, v := range n {
		sum += v
	}

	return sum
}

// TrueForAll tests whether every element of current slice matches the
// conditions specified by predicate.
func (n NUint8) TrueForAll(pred PredicateUint8Func) bool {
	for _, v := range n {
		if !pred(v) {
			return false
		}
	}

	return true
}

// TrueForAny tests whether any element of current slice matches the conditions
// specified by predicate.
func (n NUint8) TrueForAny(pred PredicateUint8Func) bool {
	for _, v := range n {
		if pred(v) {
			return true
		}
	}

	return false
}

// Where filters the elements from current slice and return them on new slice.
func (n NUint8) Where(pred PredicateUint8Func) NUint8 {
	var result NUint8
	for _, v := range n {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}
