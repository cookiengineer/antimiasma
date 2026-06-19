package miasma

import utils_pkgbuild "antimiasma/utils/pkgbuild"
import "os"
import "path/filepath"

func IsMakePkgAffected(repo string) bool {

	pkgbuild  := filepath.Join(repo, "PKGBUILD")
	data, err := os.ReadFile(pkgbuild)

	if err != nil {
		return false
	}

	content := string(data)

	if utils_pkgbuild.HasDepends(content, "bun") {
		return true
	}

	makepkg_hooks := []string{
		"prepare",
		"build",
		"check",
		"package",
	}

	for _, hook := range makepkg_hooks {

		if utils_pkgbuild.HasBashCall(content, hook, "bun add") {
			return true
		}

		if utils_pkgbuild.HasBashCall(content, hook, "/usr/bin/bun add") {
			return true
		}

		// Obfuscated bash call
		if utils_pkgbuild.HasBashCall(content, hook, "$'\\x") {
			return true
		}

		// Obfuscated bash call
		if utils_pkgbuild.HasBashCall(content, hook, "$\"") {
			return true
		}

	}

	return false

}


