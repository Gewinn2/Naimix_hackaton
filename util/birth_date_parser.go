package util

import (
	"errors"
	"strconv"
	"strings"
)

func BirthDateParser(birthDate string) (int, int, int, error) {
	birthDateSlice := strings.Split(birthDate, "-")
	if len(birthDateSlice) != 3 {
		return 0, 0, 0, errors.New("")
	}
	year, err := strconv.Atoi(birthDateSlice[0])
	if err != nil {
		return 0, 0, 0, err
	}
	month, err := strconv.Atoi(birthDateSlice[1])
	if err != nil {
		return 0, 0, 0, err
	}
	day, err := strconv.Atoi(birthDateSlice[2])
	if err != nil {
		return 0, 0, 0, err
	}

	return year, month, day, nil
}
