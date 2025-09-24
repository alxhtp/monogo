package parserhelper

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func SliceUUIDsStr(uuidsStr string) ([]uuid.UUID, error) {
	out := make([]uuid.UUID, 0)
	if len(uuidsStr) == 0 {
		return out, nil
	}

	items := strings.Split(uuidsStr, ",")

	for i := range items {
		var item = strings.TrimSpace(items[i])
		uid, err := uuid.Parse(item)
		if err != nil {
			return out, err
		}
		out = append(out, uid)
	}

	return out, nil
}

func SliceStringsStr(str string) []string {
	out := make([]string, 0)
	if len(str) == 0 {
		return out
	}

	items := strings.Split(str, ",")

	for i := range items {
		out = append(out, strings.TrimSpace(items[i]))
	}

	return out
}

func SliceFloat64sStr(str string) ([]float64, error) {
	out := make([]float64, 0)
	if len(str) == 0 {
		return out, nil
	}

	items := strings.Split(str, ",")

	for i := range items {
		var item = strings.TrimSpace(items[i])
		num, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return out, err
		}
		out = append(out, num)
	}

	return out, nil
}

func SliceIntsStr(str string) ([]int, error) {
	out := make([]int, 0)
	if len(str) == 0 {
		return out, nil
	}

	items := strings.Split(str, ",")

	for i := range items {
		var item = strings.TrimSpace(items[i])
		num, err := strconv.Atoi(item)
		if err != nil {
			return out, err
		}
		out = append(out, num)
	}

	return out, nil
}

func SliceBooleanStr(str string) ([]bool, error) {
	out := make([]bool, 0)
	if len(str) == 0 {
		return out, nil
	}
	items := strings.Split(str, ",")
	for i := range items {
		var item = strings.TrimSpace(items[i])
		bol, err := strconv.ParseBool(item)
		if err != nil {
			return out, err
		}
		out = append(out, bol)
	}
	return out, nil
}

func SliceStringContains(elems []string, v string) bool {
	for _, val := range elems {
		if v == val {
			return true
		}
	}
	return false
}

func SliceIntContains(elems []int, v int) bool {
	for _, val := range elems {
		if v == val {
			return true
		}
	}
	return false
}

// SliceUUIDToString converts a slice of uuid.UUID into a comma-separated string
func SliceUUIDToString(uuids []uuid.UUID, separator string) string {
	strSlice := make([]string, len(uuids))
	for i, u := range uuids {
		strSlice[i] = u.String()
	}
	return strings.Join(strSlice, separator)
}
