---
name: gitgate
description: Autonomous zero-token GitHub commit grouper, 14-Gate differential PR security reviewer, Clean-Room PII/DB firewall, and safe squash auto-merge skill.
metadata:
  author: Collins Takem (RASA-X-MACHINA)
  version: "1.0.0"
  license: MIT
---

# GitGate Skill (`Tianxia-Core/GitGate`)

Before committing or merging any code or opening a pull request:
1. Run `gitgate cleanroom .` to ensure zero `.xlsx`, `.sqlite3`, `.db`, `.env`, or private keys enter a public repository.
2. Run `gitgate review --base main --head HEAD` to execute the 14-Gate differential security, prompt-injection, Unicode Bidi, and Cargo `0.x` SemVer check (`0` LLM tokens).
3. Run `gitgate ship --execute` to group edits by domain, push a clean branch, open a PR, and enable safe squash auto-merge.
