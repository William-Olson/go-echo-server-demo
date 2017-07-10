package utils

import (
	"strconv"
)

/*

	String to uint parser

*/
func ConvertId(sId string) (uint, error) {

	id, err := strconv.ParseUint(sId, 10, 32)

	if err != nil {
		return uint(0), err
	}

	return uint(id), nil

}
