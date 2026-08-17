package elf

func IsSuspicious(data []byte) bool {

	result := false

	// ELF binary
	if len(data) >= 4 && data[0] == 0x7f && data[1] == 'E' && data[2] == 'L' && data[3] == 'F' {
		result = true
	}

	// embedded eBPF bytecode
	if len(data) >= 20 && data[0] == 0x7f && data[1] == 'E' && data[2] == 'L' && data[3] == 'F' {

		machine := uint16(data[18]) | uint16(data[19])<<8

		if machine == 247 { // EM_BPF
			result = true
		}

	}

	return result

}
