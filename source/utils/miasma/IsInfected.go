package miasma

func IsInfected(repo string) bool {

	result := false

	if IsClaudeAffected(repo) == true {
		result = true
	}

	if IsGeminiAffected(repo) == true {
		result = true
	}

	if IsNodeAffected(repo) == true {
		result = true
	}

	if IsCursorAffected(repo) == true {
		result = true
	}

	if IsVSCodeAffected(repo) == true {
		result = true
	}

	if HasImplant(repo) == true {
		result = true
	}

	return result

}

