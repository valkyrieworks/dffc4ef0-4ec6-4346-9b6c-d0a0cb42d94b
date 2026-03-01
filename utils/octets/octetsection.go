package octets

//
//
//
func Identifier(section []byte) []byte {
	signature := make([]byte, 6)
	copy(signature, section)
	return signature
}
