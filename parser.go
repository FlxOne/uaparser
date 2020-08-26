package uaparser

import (
	"strings"
	"unicode"
)

type itemSpec struct {
	name             string
	mustContains     []string
	mustNotContains  []string
	versionSplitters [][]string
}

type InfoItem struct {
	Name    string
	Version string
}

type UAInfo struct {
	Browser,
	Device,
	DeviceType,
	OS *InfoItem
}

func isEmptyString(str string) bool {
	for _, char := range str {
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

func contains(ua string, tokens []string) bool {
	for _, tk := range tokens {
		if strings.Contains(ua, strings.ToLower(tk)) {
			return true
		}
	}
	return false
}

func matchSpec(ua string, spec *itemSpec) (info *InfoItem, ok bool) {
	if !contains(ua, spec.mustContains) {
		return
	}
	if contains(ua, spec.mustNotContains) {
		return
	}

	info = new(InfoItem)
	info.Name = spec.name
	ok = true

	for _, splitter := range spec.versionSplitters {
		split := strings.ToLower(splitter[0])

		if strings.Contains(ua, split) {
			if rmLeft := strings.Split(ua, split)[1]; strings.Contains(rmLeft, split) || isEmptyString(split) {
				rmRight := strings.Split(rmLeft, split)[0]
				info.Version = strings.TrimSpace(rmRight)
				break
			}
		}
	}
	return
}

func searchIn(ua string, specs []*itemSpec) (info *InfoItem) {
	for _, spec := range specs {
		if result, ok := matchSpec(ua, spec); ok {
			info = result
			break
		}
	}
	return
}

func Parse(ua string) (info *UAInfo) {
	info = new(UAInfo)

	lowerCaseUa := strings.ToLower(ua)

	info.Browser = searchIn(lowerCaseUa, _BROWSERS)
	info.Device = searchIn(lowerCaseUa, _DEVICES)
	info.DeviceType = searchIn(lowerCaseUa, _DEVICETYPES)
	info.OS = searchIn(lowerCaseUa, _OS)

	return
}
