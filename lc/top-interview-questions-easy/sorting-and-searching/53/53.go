package _53

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	s, e := 1, n
	for e > s {
		mid := (s + e - 1) / 2
		if isBadVersion(mid) {
			e = mid
		} else {
			s = mid + 1
		}

	}
	return e
}

func isBadVersion(version int) bool {
	return true
}
