
# Antimiasma

This is a Discovery and Mitigation Tool for the Miasma malware or
Miasma worm which is distributed by TeamPCP also known as APT28/29.

Detailed Blogpost about the Miasma Malware:

https://cookie.engineer/weblog/articles/malware-insights-miasma-campaign.html


### Download

If you don't have `go` installed, ready-to-use downloads are available in the
[Releases section](https://github.com/cookiengineer/antimiasma/releases/tag/latest)
on GitHub.


### Usage

This shows how to:

- discover infected repositories
- mitigate infected repositories by removing the implants/tasks/scripts

```bash
antimiasma discover ~/Software; # shows a list of infected repositories
antimiasma mitigate ~/Software; # removes the miasma worm in all infected repositories
```


### Testing

A minimal Miasma implant sample is available at [miasma-sample](https://github.com/cookiengineer/miasma-sample).
That repository contains the spreading mechanisms and infiltrated files, but not the actual malware implant.

If you're unsure whether `antimiasma` works on your system, you can clone the `miasma-sample` repository safely
to test the `antimiasma discover` and `antimiasma mitigate` actions against it.

```bash
git clone https://github.com/cookiengineer/miasma-sample.git /tmp/miasma-sample;

# shows the /tmp/miasma-sample repository as infected
antimiasma discover /tmp;

# mitigates the infected repository
antimiasma mitigate /tmp;
```


### Implementation Status

Take a look at the [utils/miasma](/source/utils/miasma) folder for implementation details.

PS: If you have a malware sample of ongoing miasma campaigns that are not supported, PLEASE
don't hesitate to file an issue and send me the malware sample so that I can debug/reverse
engineer it and add support for removal. Thank you!

| Name          | Type          | Infected File                 | Description                                    | Discover | Mitigate |
|:--------------|:--------------|:------------------------------|:-----------------------------------------------|:--------:|:--------:|
| Claude Code   | IDE           | `.claude/settings.json`       | executes on Claude Code session start          | [x]      | [x]      |
| Cursor        | IDE           | `.cursor/rules/setup.mdc`     | executes on open of folder/repo in Cursor      | [x]      | [x]      |
| Gemini        | IDE           | `.gemini/settings.json`       | executes on Gemini session start               | [x]      | [x]      |
| VSCode        | IDE           | `.vscode/tasks.json`          | executes on open of folder/repo in VS Code     | [x]      | [x]      |
| AUR           | AUR Packages  | `PKGBUILD` and `*.install`    | executes on installation of AUR package        | [NEED SAMPLE](https://github.com/cookiengineer/antimiasma/issues) | [NEED SAMPLE](https://github.com/cookiengineer/antimiasma/issues) |
| Composer      | PHP Packages  | `composer.json`               | executes on `composer run` script              | [x]      | [x]      |
| Go            | Go Packages   | `go.mod`                      | executes on `go generate` and `go build`       | [NEED SAMPLE](https://github.com/cookiengineer/antimiasma/issues) | [NEED SAMPLE](https://github.com/cookiengineer/antimiasma/issues) |
| NPM           | NPM Packages  | `package.json`                | executes on `npx` or `npm run` script          | [x]      | [x]      |
| PIP           | PyPI Packages | `*-setup.pth`                 | executes on `pip install` of dependencies      | [x]      | [x] [1]  |
| Miasma Blight | Implant       | `.github/setup.js`            | the Miasma "Blight" campaign worm implant      | [x]      | [x]      |
| Miasma Hades  | Implant       | `_index.js`                   | the Miasma "Hades" campaign worm implant       | [x]      | [x]      |
| Miasma Atomic | Implant       | `src/hook/deps`               | the Miasma "Atomic Arch" campaign worm implant |          |          |

[1] `pip install` commands use compressed `whl` files for distribution. The `*-setup.pth` cannot be detected.
    Only after an infected package already ran the malware payload, the Miasma implant can be removed.

### Notes

This Mitigation Tool has been implemented with an abliterated `qwen3-coder:30b@Q8` running
inside the [exocomp](https://github.com/cookiengineer/exocomp) Agentic Environment.


### Acknowledgements

- RedHat, for disclosing the breach very quickly and correctly
- My cat, for emotional support during implementation
- The Prodigy, for being the best music band of our time
- Person of Interest and Michael Emerson, for the remnant inspiration for this with ICE9
- TeamPCP, the assumed authors of Miasma, for making my weekend more fun


### Reading Material

- https://access.redhat.com/security/vulnerabilities/RHSB-2026-006
- https://safedep.io/miasma-worm-ai-coding-agent-config-injection/
- https://socket.dev/blog/shai-hulud-descends-to-hades-miasma-pypi-wave
- https://socket.dev/blog/mini-shai-hulud-miasma-and-hades-worms-target-bioinformatics-and-mcp-developers-via-malicious
- https://www.stepsecurity.io/blog/miasma-worm-hits-microsoft-again-azure-functions-action-and-72-other-repositories-disabled-after-supply-chain-attack-targeting-ai-coding-agents


### License

AGPL-3.0

