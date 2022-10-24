package tdd_golang

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	var list_of_num []string
	if input == "" {
		return 0, nil
	}
	separated := strings.Split(input, "\n")
	if len(separated[0]) == 1 || len(separated) == 1 {
		remove_new_line := strings.ReplaceAll(input, "\n", ",")
		list_of_num = strings.Split(remove_new_line, ",")
	} else if len(separated[0]) > 1 && separated[0][:2] == "//" {
		list_of_num = strings.Split(separated[1], separated[0][3:])

	}

	result := 0
	for _, v := range list_of_num {
		v_in_num, _ := strconv.Atoi(v)
		if v_in_num < 0 {
			err := fmt.Sprintf("negatives not allowed: %s", v)
			return 0, errors.New(err)
		} else if v_in_num > 1000 {
			continue
		}
		result += v_in_num
	}

	return result, nil
}
