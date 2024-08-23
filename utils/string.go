package utils

import "strconv"

// Contains function to check if a string is present in a slice of strings
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// StrToInt converts a string to an integer
func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// IntToStr converts an integer to a string
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

// StrToFloat converts a string to a float64
func StrToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// FloatToStr converts a float64 to a string
func FloatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// StrToBool converts a string to a boolean
func StrToBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

// BoolToStr converts a boolean to a string
func BoolToStr(b bool) string {
	return strconv.FormatBool(b)
}

// StrToInt8 converts a string to an int8
func StrToUint(s string) uint {
	i, _ := strconv.ParseUint(s, 10, 64)
	return uint(i)
}

// Int8ToStr converts an int8 to a string
func UintToStr(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

// StrToInt8 converts a string to an int8
func StrToInt8(s string) int8 {
	i, _ := strconv.ParseInt(s, 10, 8)
	return int8(i)
}

// Int8ToStr converts an int8 to a string
func Int8ToStr(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}
