package fs

import "os"
import "path/filepath"

func ParentRepository(path string) (string, error) {

	dir, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	// If path points to a file, start from its parent directory.
	if info, err := os.Stat(dir); err == nil && !info.IsDir() {
		dir = filepath.Dir(dir)
	}

	for {

		gitPath := filepath.Join(dir, ".git")

		info, err := os.Stat(gitPath)
		if err == nil {

			// Accept either a .git directory or git worktree file.
			if info.IsDir() || info.Mode().IsRegular() {
				return dir, nil
			}

		} else if !os.IsNotExist(err) {
			return "", err
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root.
			return "", nil
		}

		dir = parent

	}

}
