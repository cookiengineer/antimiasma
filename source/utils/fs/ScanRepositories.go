package fs

import "io/fs"
import "path/filepath"

func ScanRepositories(root string) []string {

	found := make(map[string]bool)

	filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {

		if err == nil {

			if entry.IsDir() && entry.Name() == ".git" {

				found[filepath.Dir(path)] = true

				return filepath.SkipDir

			} else {
				return nil
			}

		} else {
			return nil
		}

	})

	repositories := make([]string, 0)

	for path, _ := range found {
		repositories = append(repositories, path)
	}

	return repositories

}
