package utils

func ByteToString(endMsg []byte, length int) string {
	return string(endMsg)[:length]
}