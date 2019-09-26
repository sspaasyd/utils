package utils

import (
	"errors"
	"strconv"
)

//判断分页条件page rows
func JudgePageRows(page, rows string) (int, string, error) {
	err := errors.New("PAGING_CONDITION_NIL_MESSAGE")
	if page == "" || rows == "" {
		return PAGING_CONDITION_NIL, PAGING_CONDITION_NIL_MESSAGE, err
	}

	pa, e := strconv.Atoi(page)
	if e != nil {
		return PAGING_CONDITION_NIL, PAGING_CONDITION_NIL_MESSAGE, err
	}
	ro, e := strconv.Atoi(rows)
	if e != nil {
		return PAGING_CONDITION_NIL, PAGING_CONDITION_NIL_MESSAGE, err
	}

	if pa < 1 || ro < 1 {
		return PAGING_CONDITION_NIL, PAGING_CONDITION_NIL_MESSAGE, err
	}

	return SUCCESS, SUCCESS_MESSAGE, nil
}
