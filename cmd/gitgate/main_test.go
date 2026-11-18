// ffp: FFP-GITGATE-TEST capability-proxy go
package main

import "testing"

func TestReviewDiffAndCleanroom(t *testing.T) {
	okRep := ReviewDiff("--- a/README.md\n+++ b/README.md\n+Clean line\n", []string{"README.md"})
	if !okRep.CanAutoMerge || okRep.Verdict != "APPROVED_FOR_MERGE" {
		t.Fatalf("expected clean approval, got %+v", okRep)
	}

	tok := "sk_" + "live_" + "000000000000deadbeef"
	badRep := ReviewDiff("--- a/Cargo.toml\n+++ b/Cargo.toml\n-rand = \"0.8.5\"\n+rand = \"0.9.0\"\n+k = \""+tok+"\"\n", []string{"Cargo.toml", "leads.xlsx"})
	if badRep.CanAutoMerge || len(badRep.Findings) < 3 {
		t.Fatalf("expected blocking findings, got %+v", badRep)
	}
}
