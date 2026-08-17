package antimiasma

import "os/exec"
import "strings"

func gitCommitAndPush(repo string, implant_paths []string) bool {

	for _, implant_path := range implant_paths {

		git_add := exec.Command("git", "add", implant_path)
		git_add.Dir = repo

		if err := git_add.Run(); err != nil {
			return false
		}

	}

	git_commit := exec.Command("git", "commit", "-m", "☢️ 💉 Antimiasma Vaccination 💉 ☢️")
	git_commit.Dir = repo

	if err := git_commit.Run(); err != nil {
		return false
	}

	git_push := exec.Command("git", "push")
	git_push.Dir = repo

	if err := git_push.Run(); err == nil {
		return true // pushed to default remote, same behavior as current miasma samples
	}

	branch_bytes, err0 := exec.Command("git", "-C", repo, "branch", "--show-current").Output()

	if err0 != nil {
		return false
	}

    remote_bytes, err1 := exec.Command("git", "-C", repo, "remote").Output()

    if err1 != nil {
        return false
    }

    branch  := strings.TrimSpace(string(branch_bytes))
    remotes := strings.Fields(string(remote_bytes))

	if len(remotes) > 0 {

		for _, remote := range remotes {

			git_push_update := exec.Command("git", "push", "-u", remote, branch)
			git_push_update.Dir = repo
			git_push_update.Run()

		}

		return true

	} else {
		return false
	}

}
