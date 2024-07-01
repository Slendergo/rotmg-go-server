package assets

import (
	"strconv"
	"strings"

	"github.com/beevik/etree"
)

func StringElementValue(elem *etree.Element, name string, defaultValue string) string {
	element := elem.SelectElement(name)
	if element != nil {
		return element.Text()
	}
	return defaultValue
}

func IntElementValue(elem *etree.Element, name string, defaultValue int32) int32 {
	element := elem.SelectElement(name)
	if element != nil {
		intValue, err := strconv.ParseInt(element.Text(), 10, 32)
		if err == nil {
			return int32(intValue)
		}
	}
	return defaultValue
}

func IntArrayElementValue(elem *etree.Element, name string) []int32 {
	element := elem.SelectElement(name)
	if element == nil {
		return nil
	}

	values := strings.Split(strings.TrimSpace(element.Text()), ",")

	var intArray []int32

	for _, valStr := range values {
		hexValue, err := strconv.ParseInt(valStr, 0, 32)
		if err != nil {
			return nil
		}
		intArray = append(intArray, int32(hexValue))
	}
	return intArray
}

func HexElementValue(elem *etree.Element, name string, defaultValue int32) int32 {
	element := elem.SelectElement(name)
	if element != nil {
		value, err := strconv.ParseInt(element.Text(), 16, 32)
		if err == nil {
			return int32(value)
		}
	}
	return defaultValue
}

func FloatElementValue(elem *etree.Element, name string, defaultValue float32) float32 {
	element := elem.SelectElement(name)
	if element != nil {
		value, err := strconv.ParseFloat(element.Text(), 32)
		if err == nil {
			return float32(value)
		}
	}
	return defaultValue
}

func HasElement(elem *etree.Element, name string) bool {
	element := elem.SelectElement(name)
	return element != nil
}
