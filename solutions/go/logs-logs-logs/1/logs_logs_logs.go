package logs

//import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	applications := map[string]string{
        "❗": "recommendation",
        "🔍": "search",
        "☀": "weather",
    }

    for _, char := range log {
        value, exists := applications[string(char)]
                                      
        if exists {
            return value
        }
    }

    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    newLog := []rune{}
	for _, char := range log {
        if rune(char) == oldRune {                          
        	newLog = append(newLog, newRune)
        } else {
            newLog = append(newLog, char)
        }
    }

    return string(newLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
