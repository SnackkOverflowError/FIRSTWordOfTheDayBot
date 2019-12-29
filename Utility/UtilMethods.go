package Utility

func GetIndex(s string, list []string) int {
	for x := 0; x < len(list); x++ {
		if list[x] == s {
			return x
		}
	}
	return -1
}

func Remove(s []string, i int) []string {
	if i == -1 {
		return s
	}
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func StartsWith(full string, starts string) bool {
	if len(full)  < len(starts) || len(full) == 0 || len(starts) == 0 {
		return false
	}
	for x := 0; x < len(starts);x++ {
		if full[x] != starts[x] {
			return false;
		}
	}
	return true;

}
