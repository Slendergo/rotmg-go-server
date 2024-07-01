package utils

import "strconv"

func ParseHexInt32(hex string) int32 {
	return ParseHexInt32D(hex, 0)
}

func ParseHexInt32D(hex string, defaultValue int32) int32 {
	if len(hex) < 2 || hex[0] != '0' || hex[1] != 'x' {
		return defaultValue
	}

	intType, err := strconv.ParseInt(hex[2:], 16, 32) // Parse starting from index 2 (skip "0x")
	if err != nil {
		return defaultValue
	}

	return int32(intType)
}

func ParseHexInt64(hex string) int64 {
	return ParseHexInt64D(hex, 0)
}

func ParseHexInt64D(hex string, defaultValue int64) int64 {
	if len(hex) < 2 || hex[0] != '0' || hex[1] != 'x' {
		return defaultValue
	}

	intType, err := strconv.ParseInt(hex[2:], 16, 64) // Parse starting from index 2 (skip "0x")
	if err != nil {
		return defaultValue
	}

	return int64(intType)
}
