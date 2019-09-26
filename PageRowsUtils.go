package utils

import "strconv"

//判断分页条件page rows
func judgePageRows(page, rows string) (int, string) {
	if page == "" || rows == "" {
		return PAGING_CONDITION_NIL, PAGING_CONDITION_NIL_MESSAGE
	}

	pa, e := strconv.Atoi(page)
	if e != nil {
		return PAGING_CONDITION_NIL, PAGING_CONDITION_NIL_MESSAGE
	}
	ro, e := strconv.Atoi(rows)
	if e != nil {
		return PAGING_CONDITION_NIL, PAGING_CONDITION_NIL_MESSAGE
	}

	if pa < 1 || ro < 1 {
		return PAGING_CONDITION_NIL, PAGING_CONDITION_NIL_MESSAGE
	}

	return SUCCESS, SUCCESS_MESSAGE
}
