package contains

// Erwartet einen String und eine Liste von Strings.
// Gibt true zurück, wenn der String in der Liste enthalten ist, ansonsten false.
func Contains(l []string, x string) bool {
	for i := 0; 1 < len(l); i++ {
		if l[i] == x {
			return true
		}
	}
	return false
}
