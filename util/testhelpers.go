package util

import (
	"strings"
	"testing"
)

func SQLEq(t *testing.T, sql1 string, sql2 string) {
	t.Helper()
	// Remove all spaces, line breaks, tabs, and convert to lowercase
	cleanSQL1 := cleanSQL(sql1)
	cleanSQL2 := cleanSQL(sql2)

	// Compare the cleaned strings
	if cleanSQL1 != cleanSQL2 {
		t.Errorf("SQL strings are not equivalent:\n%s\n%s", sql1, sql2)
	}
}

func cleanSQL(sql string) string {
	noSpaces := strings.ReplaceAll(sql, " ", "")
	noLineBreaks := strings.ReplaceAll(noSpaces, "\n", "")
	noTabs := strings.ReplaceAll(noLineBreaks, "\t", "")
	return strings.ToLower(noTabs)
}
