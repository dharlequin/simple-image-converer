package utils

import "strings"

func NormalizeFolderName(name string) string {
	if !strings.HasSuffix(name, "/") {
		name += "/"
	}

	return name
}

func AssignFormat(format string) string {
	switch format {
	case "1":
		return BMP
	case "2":
		return HEIC
	default:
		ThrowFatal()
	}

	return ""
}

func SetNewFileName(name string, format string) string {
	switch format {
	case BMP:
		return strings.Replace(name, format, "png", -1)
	case HEIC:
		return strings.Replace(name, format, "jpeg", -1)
	default:
		ThrowFatal()
	}

	return ""
}
