package fs

import "io/fs"
import "os"
import "path/filepath"

func ScanPackages(root string) []string {

	found := make(map[string]bool)

	filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {

		if err == nil {

			if entry.IsDir() {

				switch entry.Name() {

				case "node_modules":

					entries, err1 := os.ReadDir(path)

					if err1 == nil {

						for _, pkgentry := range entries {

							if pkgentry.IsDir() {
								found[filepath.Join(path, pkgentry.Name())] = true
							}

						}

						// Continue to scan for nested node_modules
						return nil

					} else {
						return filepath.SkipDir
					}

				case "site-packages":

					python_versions, err1 := os.ReadDir(path)

					if err1 == nil {

						for _, python_version := range python_versions {

							if python_version.IsDir() {

								version_path := filepath.Join(path, python_version.Name())

								pkgentries, err2 := os.ReadDir(version_path)

								if err2 == nil {

									for _, pkgentry := range pkgentries {

										if pkgentry.IsDir() {
											found[filepath.Join(version_path, pkgentry.Name())] = true
										}

									}

								} else {
									continue
								}

							}

						}

						// Continue to scan for nested site-packages
						return nil

					} else {
						return filepath.SkipDir
					}

				}

				return nil

			} else {

				switch entry.Name() {
				case "__init__.py":
					found[filepath.Dir(path)] = true
				case ".gemspec":
					found[filepath.Dir(path)] = true
				case "Gemfile":
					found[filepath.Dir(path)] = true
				case "composer.json":
					found[filepath.Dir(path)] = true
				case "package.json":
					found[filepath.Dir(path)] = true
				case "go.mod":
					found[filepath.Dir(path)] = true
				}

				return nil

			}

		} else {
			return nil
		}

	})

	packages := make([]string, 0)

	for path, _ := range found {
		packages = append(packages, path)
	}

	return packages

}
