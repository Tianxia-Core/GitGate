// ffp: FFP-GITGATE-CLI capability-proxy go
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

const Version = "1.0.0"

func runGit(dir string, args ...string) (string, string, int) {
	c := exec.Command("git", args...)
	c.Dir = dir
	var o, e bytes.Buffer
	c.Stdout, c.Stderr = &o, &e
	rc := 0
	if err := c.Run(); err != nil {
		if ex, ok := err.(*exec.ExitError); ok {
			rc = ex.ExitCode()
		} else {
			rc = 1
		}
	}
	return o.String(), e.String(), rc
}

func ShipRepo(repoDir, customMsg, customBranch string, execute, openPR bool) map[string]any {
	out, _, rc := runGit(repoDir, "status", "--porcelain=v1", "-z")
	if rc != 0 {
		return map[string]any{"ok": false, "error": "not_a_git_repository"}
	}
	var files []string
	groups := map[string][]string{}
	quarantined := 0
	for _, entry := range strings.Split(out, "\x00") {
		if len(entry) < 4 {
			continue
		}
		rel := filepath.ToSlash(entry[3:])
		if idx := strings.Index(rel, " -> "); idx >= 0 {
			rel = rel[idx+4:]
		}
		if bad, _ := IsForbiddenPublicPath(rel); bad {
			quarantined++
			continue
		}
		files = append(files, rel)
		top := "root"
		if s := strings.IndexByte(rel, '/'); s > 0 {
			top = rel[:s]
		}
		groups[top] = append(groups[top], rel)
	}
	if len(files) == 0 {
		return map[string]any{"ok": true, "clean": true, "quarantined": quarantined}
	}
	var gnames []string
	for g := range groups {
		gnames = append(gnames, g)
	}
	sort.Strings(gnames)
	scope := strings.Trim(slugRe.ReplaceAllString(strings.ToLower(gnames[0]), "-"), "-")
	if scope == "" {
		scope = "core"
	}
	branch := customBranch
	if branch == "" {
		branch = "feat/" + scope + "-update"
	}
	headline := customMsg
	if headline == "" {
		headline = fmt.Sprintf("feat(%s): update %d verified Clean-Room file(s)", scope, len(files))
	}
	diffOut, _, _ := runGit(repoDir, append([]string{"diff", "HEAD", "--"}, files...)...)
	rev := ReviewDiff(diffOut, files)
	res := map[string]any{
		"ok": rev.CanAutoMerge, "execute": execute, "files_count": len(files),
		"quarantined": quarantined, "branch": branch, "headline": headline, "review": rev,
	}
	if !rev.CanAutoMerge || !execute {
		return res
	}
	curOut, _, _ := runGit(repoDir, "branch", "--show-current")
	cur := strings.TrimSpace(curOut)
	if (cur == "main" || cur == "master") && openPR {
		_, _, _ = runGit(repoDir, "checkout", "-B", branch)
		cur = branch
	}
	_, _, _ = runGit(repoDir, append([]string{"add", "--"}, files...)...)
	_, _, _ = runGit(repoDir, "commit", "-m", headline+"\n\n[GitGate: 14-Gate & Clean-Room Verified]")
	_, _, _ = runGit(repoDir, "push", "-u", "origin", cur)
	if openPR && cur != "main" && cur != "master" {
		if gh, err := exec.LookPath("gh"); err == nil && gh != "" {
			pr := exec.Command(gh, "pr", "create", "--title", headline, "--body", "Verified by `gitgate` (0 LLM tokens).")
			pr.Dir = repoDir
			if b, err := pr.Output(); err == nil {
				res["pr_url"] = strings.TrimSpace(string(b))
				mg := exec.Command(gh, "pr", "merge", "--squash", "--delete-branch", "--auto")
				mg.Dir = repoDir
				_ = mg.Run()
			}
		}
	}
	return res
}

func AuditCleanroom(root string) map[string]any {
	var violations []Finding
	scanned := 0
	skip := map[string]bool{".git": true, "node_modules": true, "target": true, "dist": true}
	_ = filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if skip[d.Name()] && p != root {
				return filepath.SkipDir
			}
			return nil
		}
		rel, _ := filepath.Rel(root, p)
		if bad, msg := IsForbiddenPublicPath(rel); bad {
			violations = append(violations, Finding{Severity: "CRITICAL", Gate: "07-DHUMAVATI-CLEANROOM", File: rel, Detail: msg})
			return nil
		}
		scanned++
		return nil
	})
	return map[string]any{"ok": len(violations) == 0, "edition": "PUBLIC_CLEANROOM_SEED", "scanned": scanned, "violations": violations}
}

func runMCP() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		var req map[string]any
		if json.Unmarshal(sc.Bytes(), &req) != nil {
			continue
		}
		id, ok := req["id"]
		if !ok {
			continue
		}
		method, _ := req["method"].(string)
		var result any
		switch method {
		case "initialize":
			result = map[string]any{"protocolVersion": "2024-11-05", "capabilities": map[string]any{"tools": map[string]any{}}, "serverInfo": map[string]any{"name": "gitgate", "version": Version}}
		case "tools/list":
			result = map[string]any{"tools": []map[string]any{
				{"name": "gitgate_review_diff", "description": "Audit a unified git diff across 14 security, SemVer & Clean-Room gates (0 LLM tokens).", "inputSchema": map[string]any{"type": "object", "properties": map[string]any{"diff": map[string]any{"type": "string"}}, "required": []string{"diff"}}},
			}}
		case "tools/call":
			params, _ := req["params"].(map[string]any)
			args, _ := params["arguments"].(map[string]any)
			d, _ := args["diff"].(string)
			b, _ := json.MarshalIndent(ReviewDiff(d, nil), "", "  ")
			result = map[string]any{"content": []map[string]any{{"type": "text", "text": string(b)}}}
		default:
			result = map[string]any{}
		}
		out, _ := json.Marshal(map[string]any{"jsonrpc": "2.0", "id": id, "result": result})
		fmt.Println(string(out))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("GitGate v%s (MIT) — Zero-Token GitHub Commit, 14-Gate Differential PR Review & Clean-Room Auto-Merge CLI\n", Version)
		fmt.Println("Usage: gitgate <ship|review|cleanroom|mcp-stdio|--version>")
		return
	}
	switch os.Args[1] {
	case "--version", "-v", "version":
		fmt.Printf("gitgate v%s (%s/%s)\n", Version, runtime.GOOS, runtime.GOARCH)
	case "review", "review-diff":
		fs := flag.NewFlagSet("review", flag.ExitOnError)
		repo := fs.String("repo-dir", ".", "Git repository")
		base := fs.String("base", "HEAD~1", "Base ref")
		head := fs.String("head", "HEAD", "Head ref")
		stdin := fs.Bool("stdin", false, "Read diff from stdin")
		_ = fs.Parse(os.Args[2:])
		var d string
		var files []string
		if *stdin {
			raw, _ := io.ReadAll(os.Stdin)
			d = string(raw)
		} else {
			d, _, _ = runGit(*repo, "diff", *base, *head)
			nf, _, _ := runGit(*repo, "diff", "--name-only", *base, *head)
			for _, l := range strings.Split(strings.TrimSpace(nf), "\n") {
				if l != "" {
					files = append(files, l)
				}
			}
		}
		rep := ReviewDiff(d, files)
		b, _ := json.MarshalIndent(rep, "", "  ")
		fmt.Println(string(b))
		if !rep.CanAutoMerge {
			os.Exit(1)
		}
	case "ship":
		fs := flag.NewFlagSet("ship", flag.ExitOnError)
		repo := fs.String("repo-dir", ".", "Git repository")
		msg := fs.String("message", "", "Commit message")
		br := fs.String("branch", "", "Branch name")
		noPR := fs.Bool("no-pr", false, "Skip PR")
		execFlag := fs.Bool("execute", false, "Execute commit/push/PR")
		_ = fs.Parse(os.Args[2:])
		b, _ := json.MarshalIndent(ShipRepo(*repo, *msg, *br, *execFlag, !*noPR), "", "  ")
		fmt.Println(string(b))
	case "cleanroom":
		target := "."
		if len(os.Args) > 2 {
			target = os.Args[2]
		}
		rep := AuditCleanroom(target)
		b, _ := json.MarshalIndent(rep, "", "  ")
		fmt.Println(string(b))
		if ok, _ := rep["ok"].(bool); !ok {
			os.Exit(1)
		}
	case "mcp-stdio", "mcp":
		runMCP()
	default:
		os.Exit(2)
	}
}
