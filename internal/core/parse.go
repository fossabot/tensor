package core

import (
	"strconv"
)

// strAsInt converts provided string to integer type.
func strAsInt(s string) int64 {
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return n
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return int64(f)
	}

	if b, err := strconv.ParseBool(s); err == nil {
		if b {
			return 1
		}
		return 0
	}

	panic(NewError("cannot convert %q to an integer value", s))
}

// strAsFloat converts provided string to floating point number type.
func strAsFloat(s string) float64 {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	if b, err := strconv.ParseBool(s); err == nil {
		if b {
			return 1.0
		}
		return 0.0
	}

	panic(NewError("cannot convert %q to a float value", s))
}

// strAsBool converts provided string to boolean type.
func strAsBool(s string) bool {
	if b, err := strconv.ParseBool(s); err == nil {
		return b
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f != 0.0
	}

	panic(NewError("cannot convert %q to a boolean value", s))
}
