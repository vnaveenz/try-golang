package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	r, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return r.MatchString(text)
}

func SplitLogLine(text string) []string {
	r, _ := regexp.Compile(`<[~*=\-]*>`)
	return r.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re, _ := regexp.Compile(`".*[pP][aA][sS][sS][wW][oO][rR][dD].*"`)
	for _, l := range lines {
		if re.MatchString(l) {
			count ++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	r, _ := regexp.Compile(`end-of-line[0-9]+`)
	return r.ReplaceAllString(text,"")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`\bUser\s+([A-Za-z0-9]+)\b`)
	for i,ln := range lines {
        if m := re.FindStringSubmatch(ln); len(m) == 2 {
            lines[i] = "[USR] " + m[1] + " " + ln
        } else {
            lines[i] = ln
        }
	}
	return lines
}
