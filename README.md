# GitGate (`Tianxia-Core/GitGate`)

[![CI](https://github.com/Tianxia-Core/GitGate/actions/workflows/ci.yml/badge.svg)](https://github.com/Tianxia-Core/GitGate/actions/workflows/ci.yml)
[![Release](https://img.shields.io/badge/release-v1.0.0-00f59b)](https://github.com/Tianxia-Core/GitGate/releases/tag/v1.0.0)
[![License: MIT](https://img.shields.io/badge/License-MIT-22d3ee)](LICENSE)
[![Governed by Build Law](https://img.shields.io/badge/Governed%20by-Build%20Law%20v5.2-00f59b)](https://github.com/Tianxia-Core/Build-Law)

**`GitGate` (`v1.0.0`)** is the freestanding open-source (`MIT`) **Autonomous GitHub Commit Grouper, Zero-Token 14-Gate Differential Pull-Request Security Reviewer, Clean-Room PII/DB Firewall, Safe Auto-Merge CLI, Composite GitHub Action, Claude Skill & MCP Server** created by **[Collins Takem](https://rasaxmachina.com/collins)** (**RASA-X-MACHINA** · `tianxiaOS`).

> **Composable by Design:** Use `GitGate` completely standalone in any repository (`uses: Tianxia-Core/GitGate@v1` or `gitgate ship`), OR run it seamlessly inside **[`nimbus²`](https://github.com/Tianxia-Core/nimbus2)** (`nimbus2 ship` / `nimbus2 review-diff`) and **[`Build Law`](https://github.com/Tianxia-Core/Build-Law)** (`build-law ship`).

---

## ⚡ 60-Second Quick Start

```bash
# 1. Audit directory for Clean-Room public purity (blocks .xlsx, .sqlite3, .env, private PEMs)
gitgate cleanroom .

# 2. Run 14-Gate differential PR security, prompt-injection & Cargo 0.x SemVer review (0 LLM tokens)
gitgate review --base main --head HEAD

# 3. Group dirty edits by domain, commit explicit NUL-safe pathspecs, push, open PR & enable squash auto-merge
gitgate ship --execute

# 4. Run as a Model Context Protocol (MCP) JSON-RPC 2.0 server
gitgate mcp-stdio
```

### Use as a 2-Line GitHub Action
```yaml
- uses: actions/checkout@v4
  with: { fetch-depth: 0 }
- uses: Tianxia-Core/GitGate@main
```
