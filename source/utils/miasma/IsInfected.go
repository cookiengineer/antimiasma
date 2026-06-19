package miasma

func IsInfected(repo string) bool {

	result := false

	if IsClaudeAffected(repo) == true {
		result = true
	}

	if IsGeminiAffected(repo) == true {
		result = true
	}

	if IsCursorAffected(repo) == true {
		result = true
	}

	if IsVSCodeAffected(repo) == true {
		result = true
	}

	if IsComposerAffected(repo) == true {
		result = true
	}

	if IsMakePkgAffected(repo) == true {
		result = true
	}

	if IsNodeAffected(repo) == true {
		result = true
	}

	if IsPacmanAffected(repo) == true {
		result = true
	}

	if HasImplant(repo) == true {
		result = true
	}

	return result

}

