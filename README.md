
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
to test the discover and mitigate actions against.

```bash
git clone https://github.com/cookiengineer/miasma-sample.git /tmp/miasma-sample;

# should show the /tmp/miasma-sample repository as infected
antimiasma discover /tmp;

# should mitigate the infected repository
antimiasma mitigate /tmp;
```


### Implementation Status

Take a look at the [utils/miasma](/utils/miasma) folder for implementation details.

| Infected File             | Description                           | Discover | Mitigate |
|:--------------------------|:--------------------------------------|:--------:|:--------:|
| `.claude/settings.json`   | triggers on Claude Code session start | [x]      | [x]      |
| `.cursor/rules/setup.mdc` | triggers when opening repo in Cursor  | [x]      | [x]      |
| `.gemini/settings.json`   | triggers on Gemini session start      | [x]      | [x]      |
| `.vscode/tasks.json`      | triggers when opening repo in VS Code | [x]      | [x]      |
| `package.json`            | hijacks `npm test` script             | [x]      | [x]      |
| `.github/setup.js`        | the Miasma worm implant               | [x]      | [x]      |


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


### License

AGPL-3.0

