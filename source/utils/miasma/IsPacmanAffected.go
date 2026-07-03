package miasma

import "antimiasma/utils/log"
import utils_pkgbuild "antimiasma/utils/pkgbuild"
import "os"
import "path/filepath"

func IsPacmanAffected(repo string) bool {

	log.Printf("  checking Pacman: %s\n", repo)

	install   := filepath.Join(repo, "install")
	data, err := os.ReadFile(install)

	if err != nil {
		return false
	}

	content := string(data)

	pacman_hooks := []string{
		"pre_install",
		"post_install",
		"pre_upgrade",
		"post_upgrade",
		"pre_remove",
		"post_remove",
	}

	for _, hook := range pacman_hooks {

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
