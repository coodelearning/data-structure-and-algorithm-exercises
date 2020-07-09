package _35

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) == 1 {
		if s == t {
			return true
		}
		return false
	}
	sExists := [26]int{}
	tExists := [26]int{}
	for _, v := range s {
		sExists[v-'a'] += 1
	}
	for _, v := range t {
		tExists[v-'a'] += 1
	}
	for i := 0; i < 26; i++ {
		if sExists[i] != tExists[i] {
			return false
		}
	}
	return true
}
