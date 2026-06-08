package fs

import "io/fs"
import "path/filepath"

func ScanRepositories(root string) []string {

	repositories := make([]string, 0)

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			// Ignore unreadable directories/files and continue.
			return nil
		}

		if d.IsDir() && d.Name() == ".git" {

			repositories = append(repositories, filepath.Dir(path))
			return filepath.SkipDir

		}

		return nil

	})

	return repositories

}
