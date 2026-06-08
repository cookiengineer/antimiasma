package miasma

func Fix(repo string) bool {

	result := true

	if tmp := FixClaude(repo); tmp == false {
		result = false
	}

	if tmp := FixGemini(repo); tmp == false {
		result = false
	}

	if tmp := FixCursor(repo); tmp == false {
		result = false
	}

	if tmp := FixVSCode(repo); tmp == false {
		result = false
	}

	if tmp := FixComposer(repo); tmp == false {
		result = false
	}

	if tmp := FixNode(repo); tmp == false {
		result = false
	}

	if tmp := RemoveImplant(repo); tmp == false {
		result = false
	}

	return result

}
