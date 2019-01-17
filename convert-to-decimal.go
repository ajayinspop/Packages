package converttodecimal

import (
"strconv"
"strings"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func ConvertToDecimal(numText, devText string) (*float64, *string) {
	var val, num, dev float64
	var err string

	numText = strings.TrimSuffix(numText, "\r\n")
	devText = strings.TrimSuffix(devText, "\r\n")

	if IsNumeric(numText) {
		val, _ := strconv.ParseFloat(numText, 64)
		num = val
	} else {
		err =  "Entered numerator is not correct."
		return nil, &err
	}

	if IsNumeric(devText) {
		val, _ := strconv.ParseFloat(devText, 64)
		dev = val
	} else {
		err =  "Entered denominator is not correct . change"
		return nil, &err
	}

	if dev == 0 {
		err =  "Divide by 0 error."
		return nil, &err
	}

	val = num / dev

	return &val, nil
}