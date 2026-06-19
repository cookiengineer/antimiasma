package miasma

import utils_pkgbuild "antimiasma/utils/pkgbuild"
import "os"
import "path/filepath"

func FixMakePkg(repo string) bool {

	pkgbuild  := filepath.Join(repo, "PKGBUILD")
	data, err := os.ReadFile(pkgbuild)

	if err != nil {
		return true // nothing to fix
	}

	content := string(data)

	if utils_pkgbuild.HasDepends(content, "bun") {
		content = utils_pkgbuild.RemoveDepends(content, "bun")
	}

	makepkg_hooks := []string{
		"prepare",
		"build",
		"check",
		"package",
	}

	for _, hook := range makepkg_hooks {

		if utils_pkgbuild.HasBashCall(content, hook, "bun add") {
			content = utils_pkgbuild.RemoveBashCall(content, hook, "bun add")
		}

		if utils_pkgbuild.HasBashCall(content, hook, "/usr/bin/bun add") {
			content = utils_pkgbuild.RemoveBashCall(content, hook, "/usr/bin/bun add")
		}

		// Obfuscated bash call
		if utils_pkgbuild.HasBashCall(content, hook, "$'\\x") {
			content = utils_pkgbuild.RemoveBashCall(content, hook, "$'\\x")
		}

		// Obfuscated bash call
		if utils_pkgbuild.HasBashCall(content, hook, "$\"") {
			content = utils_pkgbuild.RemoveBashCall(content, hook, "$\"")
		}

	}

	if err := os.WriteFile(pkgbuild, []byte(content), 0644); err != nil {
		return false
	}

	return true

}
