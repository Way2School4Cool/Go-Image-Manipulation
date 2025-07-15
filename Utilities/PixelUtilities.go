package utilities

func OverflowCheck(value uint32) uint32 {
	if value > 255 {
		return 255
	}
	return value
}
