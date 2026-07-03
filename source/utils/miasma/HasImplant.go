package miasma

import "antimiasma/utils/log"
import "os"
import "path/filepath"

func HasImplant(repo string) bool {

	log.Printf("  checking worm implants: %s\n", repo)

	result := false

	implantPath := filepath.Join(repo, ".github", "setup.js")
	info1, err1 := os.Stat(implantPath)

	if err1 == nil && info1.IsDir() == false {
		result = true
	}

	indexPath   := filepath.Join(repo, "_index.js")
	info2, err2 := os.Stat(indexPath)

	if err2 == nil && info2.IsDir() == false {
		result = true
	}

	depsPath    := filepath.Join(repo, "src", "hooks", "deps")
	info3, err3 := os.Stat(depsPath)

	if err3 == nil && info3.IsDir() == false {

		data, err := os.ReadFile(depsPath)

		if err == nil {

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

		}

	}

	return result

}
