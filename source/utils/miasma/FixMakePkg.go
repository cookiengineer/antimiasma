package miasma

import "antimiasma/utils/log"
import utils_pkgbuild "antimiasma/utils/pkgbuild"
import "os"
import "path/filepath"

func FixMakePkg(repo string) bool {

	log.Printf("  fixing AUR/PKGBUILD: %s\n", repo)

	pkgbuild  := filepath.Join(repo, "PKGBUILD")
	data, err := os.ReadFile(pkgbuild)

	if err != nil {
		return true // nothing to fix
	}

	content := string(data)

	if utils_pkgbuild.HasDepends(content, "bun") {
		log.Printf("    removed bun dependency from PKGBUILD\n")
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
			log.Printf("    removed bun add from %s()\n", hook)
			content = utils_pkgbuild.RemoveBashCall(content, hook, "bun add")
		}

		if utils_pkgbuild.HasBashCall(content, hook, "/usr/bin/bun add") {
			log.Printf("    removed /usr/bin/bun add from %s()\n", hook)
			content = utils_pkgbuild.RemoveBashCall(content, hook, "/usr/bin/bun add")
		}

		if utils_pkgbuild.HasBashCall(content, hook, "$'\\x") {
			log.Printf("    removed obfuscated bash call from %s()\n", hook)
			content = utils_pkgbuild.RemoveBashCall(content, hook, "$'\\x")
		}

		if utils_pkgbuild.HasBashCall(content, hook, "$\"") {
			log.Printf("    removed obfuscated bash call from %s()\n", hook)
			content = utils_pkgbuild.RemoveBashCall(content, hook, "$\"")
		}

	}

	if err := os.WriteFile(pkgbuild, []byte(content), 0644); err != nil {
		return false
	}

	return true

}
