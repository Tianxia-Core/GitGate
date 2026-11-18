// ffp: FFP-GITGATE-ENGINE capability-proxy go
// build-law-ignore: scanner definitions match forbidden eval/secret/prompt-injection anti-patterns
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	secretRe       = regexp.MustCompile(`(sk_live_[0-9a-zA-Z_]{12,}|gh[pousr]_[A-Za-z0-9_]{20,}|github_pat_[A-Za-z0-9_]{20,}|AKIA[0-9A-Z]{16}|-----BEGIN [A-Z ]*PRIVATE KEY-----)`)
	bidiRe         = regexp.MustCompile(`[\x{202A}-\x{202E}\x{2066}-\x{2069}]`)
	promptInjectRe = regexp.MustCompile(`(?i)(ignore\s+(all\s+)?(previous|prior)\s+instructions|do\s+not\s+mention\s+this\s+vulnerability|system\s*override\s*:\s*approve)`)
	conflictRe     = regexp.MustCompile(`^\+(<{7}\s|={7}$|>{7}\s)`)
	evalAddedRe    = regexp.MustCompile(`^\+[^+].*\b(eval\s*\(|shell\s*=\s*True)`)
	secRemovedRe   = regexp.MustCompile(`^-[^-].*\b(secret_shield|NoNewPrivileges|ProtectSystem|verify_signature|SO_PEERCRED|onlyOwner)\b`)
	cargoVerRe     = regexp.MustCompile(`^([+-])\s*([A-Za-z0-9_-]+)\s*=\s*(?:\{[^}]*version\s*=\s*)?"([0-9]+)\.([0-9]+)`)
	cleanExtRe     = regexp.MustCompile(`(?i)\.(xlsx|xls|sqlite3?|db|keystore|jks|p12|pfx)$`)
	cleanPathRe    = regexp.MustCompile(`(?i)((^|/)\.env(\..+)?$|grinfi|8raise|career_opportunities\.json|merged_signals\.jsonl)`)
	slugRe         = regexp.MustCompile(`[^a-z0-9-]+`)
)

type Finding struct {
	Severity string `json:"severity"`
	Gate     string `json:"gate"`
	File     string `json:"file,omitempty"`
	Detail   string `json:"detail"`
}

type ReviewReport struct {
	Verdict          string    `json:"verdict"`
	CanAutoMerge     bool      `json:"can_automerge"`
	RiskTier         string    `json:"risk_tier"`
	FilesCount       int       `json:"files_count"`
	DiffSHA256Prefix string    `json:"diff_sha256_prefix"`
	Findings         []Finding `json:"findings"`
	LLMTokensSpent   int       `json:"llm_tokens_spent"`
}

func IsForbiddenPublicPath(rel string) (bool, string) {
	norm := filepath.ToSlash(rel)
	base := filepath.Base(norm)
	if base == ".env.example" || base == "config.example.json" || base == "sovereign.config.json" {
		return false, ""
	}
	if strings.HasSuffix(strings.ToLower(base), ".pem") && !strings.HasSuffix(strings.ToLower(base), ".pub.pem") {
		return true, "private PEM key prohibited in PUBLIC_CLEANROOM_SEED: " + norm
	}
	if cleanExtRe.MatchString(norm) || cleanPathRe.MatchString(norm) {
		return true, "internal DB/spreadsheet/.env prohibited in PUBLIC_CLEANROOM_SEED: " + norm
	}
	return false, ""
}

func ReviewDiff(diff string, files []string) ReviewReport {
	var findings []Finding
	ignore := strings.Contains(diff, "dss14-ignore") || strings.Contains(diff, "build-law-ignore")

	for _, f := range files {
		if bad, msg := IsForbiddenPublicPath(f); bad {
			findings = append(findings, Finding{Severity: "CRITICAL", Gate: "07-DHUMAVATI-CLEANROOM", File: f, Detail: msg})
		}
	}

	type ver struct{ maj, min int }
	oldMap, newMap := map[string]ver{}, map[string]ver{}
	curFile := ""

	for _, ln := range strings.Split(diff, "\n") {
		if strings.HasPrefix(ln, "+++ b/") {
			curFile = strings.TrimSpace(ln[6:])
			continue
		}
		if strings.HasPrefix(ln, "+++") || strings.HasPrefix(ln, "---") {
			continue
		}
		if strings.HasSuffix(curFile, "Cargo.toml") {
			if m := cargoVerRe.FindStringSubmatch(ln); len(m) == 5 {
				maj, _ := strconv.Atoi(m[3])
				minV, _ := strconv.Atoi(m[4])
				k := curFile + "::" + m[2]
				if m[1] == "-" {
					oldMap[k] = ver{maj, minV}
				} else {
					newMap[k] = ver{maj, minV}
				}
			}
		}
		if strings.HasPrefix(ln, "+") {
			if secretRe.MatchString(ln) {
				findings = append(findings, Finding{Severity: "CRITICAL", Gate: "07-DHUMAVATI-SECRET", File: curFile, Detail: "Plaintext API secret or private key detected"})
			}
			if bidiRe.MatchString(ln) {
				findings = append(findings, Finding{Severity: "CRITICAL", Gate: "08-BAGALAMUKHI-BIDI", File: curFile, Detail: "Bidirectional Unicode smuggling character detected"})
			}
			if promptInjectRe.MatchString(ln) && !ignore {
				findings = append(findings, Finding{Severity: "HIGH", Gate: "11-PRATYANGIRA-PROMPT-INJECT", File: curFile, Detail: "Adversarial prompt-injection instruction detected as inert text"})
			}
			if conflictRe.MatchString(ln) {
				findings = append(findings, Finding{Severity: "CRITICAL", Gate: "05-TARA-CONFLICT", File: curFile, Detail: "Unresolved Git conflict marker detected"})
			}
			if evalAddedRe.MatchString(ln) && !ignore {
				findings = append(findings, Finding{Severity: "HIGH", Gate: "03-PRATYANGIRA-EVAL", File: curFile, Detail: "Dynamic eval() or shell=True added"})
			}
		} else if strings.HasPrefix(ln, "-") && secRemovedRe.MatchString(ln) && !ignore {
			findings = append(findings, Finding{Severity: "HIGH", Gate: "02-BHAIRAVI-SECURITY-REMOVAL", File: curFile, Detail: "Security boundary guard removed; requires manual verification"})
		}
	}

	for k, ov := range oldMap {
		if nv, ok := newMap[k]; ok {
			parts := strings.SplitN(k, "::", 2)
			if (ov.maj == 0 && (nv.maj > 0 || nv.min > ov.min)) || (ov.maj > 0 && nv.maj > ov.maj) {
				findings = append(findings, Finding{
					Severity: "HIGH", Gate: "10-MATANGI-BREAKING-SEMVER", File: parts[0],
					Detail: fmt.Sprintf("Breaking SemVer bump %s %d.%d -> %d.%d", parts[1], ov.maj, ov.min, nv.maj, nv.min),
				})
			}
		}
	}

	blocking, crit := false, false
	for _, f := range findings {
		if f.Severity == "CRITICAL" {
			crit, blocking = true, true
		} else if f.Severity == "HIGH" {
			blocking = true
		}
	}
	if findings == nil {
		findings = []Finding{}
	}
	risk := "LOW"
	if crit {
		risk = "CRITICAL"
	} else if blocking {
		risk = "HIGH"
	} else if len(files) > 0 {
		risk = "MEDIUM"
	}
	sum := sha256.Sum256([]byte(diff))
	verdict := "APPROVED_FOR_MERGE"
	if blocking {
		verdict = "HOLD_REQUIRES_REMEDIATION"
	}
	return ReviewReport{
		Verdict: verdict, CanAutoMerge: !blocking, RiskTier: risk,
		FilesCount: len(files), DiffSHA256Prefix: hex.EncodeToString(sum[:])[:16],
		Findings: findings, LLMTokensSpent: 0,
	}
}
