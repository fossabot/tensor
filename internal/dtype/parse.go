package dtype

import (
	"strconv"
	"strings"

	"github.com/ppknap/tensor/internal/errorc"
)

// strAsBool converts provided string to boolean type.
func strAsBool(s string) bool {
	if s == "" {
		return false
	}

	if b, err := strconv.ParseBool(s); err == nil {
		return b
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f != 0.0
	}

	panic(errorc.New("cannot convert %q to a boolean value", s))
}

// strAsInt converts provided string to integer type.
func strAsInt(s string) int64 {
	if s == "" {
		return 0
	}

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

	panic(errorc.New("cannot convert %q to integer value", s))
}

// strAsUint converts provided string to unsigned integer type.
func strAsUint(s string) uint64 {
	if s == "" {
		return 0
	}

	if n, err := strconv.ParseUint(s, 10, 64); err == nil {
		return n
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return uint64(f)
	}

	if b, err := strconv.ParseBool(s); err == nil {
		if b {
			return 1
		}
		return 0
	}

	panic(errorc.New("cannot convert %q to unsigned integer value", s))
}

// strAsFloat converts provided string to floating point number type.
func strAsFloat(s string) float64 {
	if s == "" {
		return 0.0
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	if b, err := strconv.ParseBool(s); err == nil {
		if b {
			return 1.0
		}
		return 0.0
	}

	panic(errorc.New("cannot convert %q to a float value", s))
}

// strAsComplex converts provided string to complex number type.
func strAsComplex(s string) complex128 {
	tok := strings.Replace(strings.Trim(s, "(,)"), "j", "i", -1)
	if tok == "" {
		return 0.0
	}

	idx := strings.LastIndexFunc(tok, func(r rune) bool {
		return r == '-' || r == '+'
	})

	var re, im string
	switch {
	case idx <= 0 && tok[len(tok)-1] == 'i':
		re, im = "0", tok[:len(tok)-1]
	case idx <= 0:
		re, im = tok, "0"
	case idx > 0 && tok[len(tok)-1] == 'i':
		re, im = tok[:idx], tok[idx:len(tok)-1]
	case idx > 0:
		re, im = tok[:idx], tok[idx:]
	}

	return complex(strAsFloat(re), strAsFloat(im))
}
