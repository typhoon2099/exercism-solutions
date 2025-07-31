package parsinglogfiles

import "regexp"
import "fmt"

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)

    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)

    for _, line := range lines {
        if re.MatchString(line) {
            count += 1
        }
    }
    
	return count
}

func RemoveEndOfLineText(text string) string {
    re := regexp.MustCompile(`end-of-line\d+`)

    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    output := []string{}
    re := regexp.MustCompile(`User\s+([a-zA-Z0-9]+)\s+`)

    for _, line := range lines {
        matches := re.FindStringSubmatch(line)

        if matches != nil {
            output = append(output, fmt.Sprintf("[USR] %s %s", matches[1], line))
        } else {
        	output = append(output, line)
        }
    }

    return output
}
