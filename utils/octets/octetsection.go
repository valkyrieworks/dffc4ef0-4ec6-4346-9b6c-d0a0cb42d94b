package octets

//
//
//
func Footprint(section []byte) []byte {
	footprint := make([]byte, 6)
	copy(footprint, section)
	return footprint
}
