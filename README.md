# GitGate (`Tianxia-Core/GitGate`)

[![CI](https://github.com/Tianxia-Core/GitGate/actions/workflows/ci.yml/badge.svg)](https://github.com/Tianxia-Core/GitGate/actions/workflows/ci.yml)
[![Release](https://img.shields.io/badge/release-v1.0.0-00f59b)](https://github.com/Tianxia-Core/GitGate/releases/tag/v1.0.0)
[![License: MIT](https://img.shields.io/badge/License-MIT-22d3ee)](LICENSE)
[![Governed by Build Law](https://img.shields.io/badge/Governed%20by-Build%20Law%20v5.2-00f59b)](https://github.com/Tianxia-Core/Build-Law)

**`GitGate` (`v1.0.0`)** is the freestanding open-source (`MIT`) **Autonomous GitHub Commit Grouper, Zero-Token 14-Gate Differential Pull-Request Security Reviewer, Clean-Room PII/DB Firewall, Safe Auto-Merge CLI, Composite GitHub Action, Claude Skill & MCP Server** created by **[Collins Takem](https://rasaxmachina.com/collins)** (**RASA-X-MACHINA** · `tianxiaOS`).

> **Composable by Design:** Use `GitGate` completely standalone in any repository (`uses: Tianxia-Core/GitGate@v1` or `gitgate ship`), OR run it seamlessly inside **[`nimbus²`](https://github.com/Tianxia-Core/nimbus2)** (`nimbus2 ship` / `nimbus2 review-diff`) and **[`Build Law`](https://github.com/Tianxia-Core/Build-Law)** (`build-law ship`).

<!-- BUILD_LAW_SHOWCASE_START -->

---

## 🌍 In Plain English — What `GitGate` Is & Why It Matters (Layman's Guide)

- **The Everyday Problem:** Reviewing pull requests by hand misses subtle security regressions or breaking dependency bumps, while AI PR review bots burn expensive API tokens and can be tricked by prompt injection inside code comments.
- **How It Works (Without the Jargon):** `GitGate` groups workspace changes cleanly (`gitgate ship`), enforces Clean-Room public isolation (`gitgate cleanroom`), and runs a deterministic 14-Gate differential security, Bidi-Unicode, prompt-injection, and SemVer `0.x` audit (`gitgate review`) at **zero LLM token cost**.
- **Why It Is Different:** Available freestanding under `MIT` (`uses: Tianxia-Core/GitGate@main` or static binary) AND integrated natively into **[`nimbus²`](https://github.com/Tianxia-Core/nimbus2)** and **[`Build Law`](https://github.com/Tianxia-Core/Build-Law)**.
- **How It Benefits You:** Engineering teams merge safe PRs faster, block accidental `.env`/`.xlsx`/`.sqlite3` leaks before push, and pay `€0.00` in review bot fees.
- **How to Install & Use It in 60 Seconds:** Add `uses: Tianxia-Core/GitGate@main` to GitHub Actions or install the signed binary from **[GitGate Releases](https://github.com/Tianxia-Core/GitGate/releases/latest)**.

---

## 🏛️ Governed by [`Build Law`](https://github.com/Tianxia-Core/Build-Law) — Clean-Room Public Edition & Stakeholder Pathways

This repository ships as a **`PUBLIC_CLEANROOM_SEED` Edition** under **[`Build Law`](https://github.com/Tianxia-Core/Build-Law)**—containing zero pre-populated personal data, zero embedded secrets, and deterministic `init-clean` bootstrapping (`nimbus2 init-clean` / `build-law init-clean`) so every adopter starts with a fresh, sovereign workspace of their own:
- **Developers & AI Builders:** Combine **[`FC-Estimator`](https://github.com/Tianxia-Core/FC-Estimator)** (`95%` accurate pre-build sizing), **[`Build Law`](https://github.com/Tianxia-Core/Build-Law)** (`5.0×–8.8×` token efficiency across 20+ IDEs/agents), and **[`nimbus²`](https://github.com/Tianxia-Core/nimbus2)** (`nimbus2 ship` + `$0.00` 14-Gate PR review).
- **CTOs & B2B Engineering Leaders:** Need similar `Rust`/`Go`/`Nix`/Security/FinOps rigor applied to your infrastructure? Scope a fixed-price **72-Hour `VRTTI` Architecture Sprint** at **[vrtti.rasaxmachina.com](https://vrtti.rasaxmachina.com)**.
- **Investors, Researchers, Press & Executive Hiring Committees:** Inspect cryptographic release proofs across **[`@Tianxia-Core`](https://github.com/Tianxia-Core)**, press assets in **[`@Tianxia-Presence/brand-kit`](https://github.com/Tianxia-Presence/brand-kit)**, commercial suites in **[`@Tianxia-Revenues`](https://github.com/Tianxia-Revenues)**, and founder dossier at **[rasaxmachina.com/collins](https://rasaxmachina.com/collins)**.

### 🔗 Composable Sovereign Stack — Works Natively With (`@Tianxia-Core` · `@Tianxia-Presence` · `@Tianxia-Revenues`)

Every tool in the **RASA-X-MACHINA** ecosystem is standalone (`0` lock-in bloat), yet designed to compose seamlessly across pre-build sizing, IDE governance, polyglot compilation, security custody, and autonomous operations:

| Lifecycle Layer | Companion Tool(s) | How They Work Together |
| :--- | :--- | :--- |
| **Pre-Build Sizing & FinOps** | **[`FC-Estimator`](https://github.com/Tianxia-Core/FC-Estimator)** (`MIT`) | Sizes `ceil(FC / 8)` microservice bounds, active build minutes (`95.0%` accuracy) & `8.8×` EUR token savings before coding. |
| **Constitutional Governance** | **[`Build Law`](https://github.com/Tianxia-Core/Build-Law)** (`MIT`) | Embeds the 4-Pillar, 14-Gate rules & `< 5 ms` `$0.00` static verifier across Claude, Cursor, VS Code, OpenCode & MCP. |
| **Executive Build, PR & Proxy Manager** | **[`nimbus²`](https://github.com/Tianxia-Core/nimbus2)** (`MPL-2.0`) & **[`GitGate`](https://github.com/Tianxia-Core/GitGate)** (`MIT`) | Runs `gitgate` / `nimbus2 ship`, zero-token 14-Gate `review-diff`, `init-clean` zero-PII seeding & high-concurrency TLS 1.3 proxy routing. |
| **Polyglot Wire Contracts** | **[`Jivascriptocle`](https://github.com/Tianxia-Core/Jivascriptocle)** (`Apache-2.0`) | Compiles single `.jiva` contracts into synchronized `Rust`, `Zig`, `OCaml`, `TypeScript`, `Nix` & compact `TOON` schemas. |
| **DevSecOps & Secret Enclave** | **[`AegisScan`](https://github.com/Tianxia-Core/AegisScan)** & **[`AetherVault`](https://github.com/Tianxia-Core/AetherVault)** (`MIT`) | 8-domain static supply-chain/prompt-injection scanner paired with `AES-256-GCM` `chmod 0600` runtime secret custody. |
| **Sovereign OS, Firewall & Mesh** | **[`TianxiaOS`](https://github.com/Tianxia-Core/TianxiaOS)** · **[`Faery-Fortress`](https://github.com/Tianxia-Core/Faery-Fortress)** · **[`tianxia-mobile`](https://github.com/Tianxia-Core/tianxia-mobile)** | Declarative NixOS operating substrate, `< 4 µs` `AF_XDP` + Rosenpass post-quantum firewall & signed Android ARM64 mesh node. |
| **Intelligence, Clearing, Ledger & Logistics** | **[`Ndimmorwyn`](https://github.com/Tianxia-Core/Ndimmorwyn)** · **[`AlatyrAtal`](https://github.com/Tianxia-Core/AlatyrAtal)** · **[`Tetraledger`](https://github.com/Tianxia-Core/Tetraledger)** · **[`MURUA`](https://github.com/Tianxia-Core/MURUA)** · **[`Scour`](https://github.com/Tianxia-Core/Scour)** | Constructal `-79.6%` router + `EjAwen` CRM, 4-Stratum `EFEMERA→ETAMANGA` clearinghouse, 4-way `x402` ledger, postal telemetry & signal radar. |
| **Press Assets & B2B Sprints** | **[`@Tianxia-Presence/brand-kit`](https://github.com/Tianxia-Presence/brand-kit)** & **[`@Tianxia-Revenues`](https://github.com/Tianxia-Revenues)** | Official vector identity/design tokens (`WCAG 2.2 AA`) and fixed-scope **72h `VRTTI` Architecture Sprints** at [rasaxmachina.com](https://rasaxmachina.com). |

<!-- BUILD_LAW_SHOWCASE_END -->

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
