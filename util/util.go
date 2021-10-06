package util

func utf82Hex(r byte) byte {
	if r > '9' {
		return byte(r - '0' - 39)
	} else {
		return byte(r - '0')
	}
}

func String2ByteArray(s string) []byte {
	bytes := make([]byte, 0)
	i := 0
	length := len(s)
	for i < length {
		b := utf82Hex(s[i])<<4 + utf82Hex(s[i+1])
		// log.Printf("%x\n", cha)
		bytes = append(bytes, b)
		i += 2
	}
	return bytes
}
