// Package RansomNote
// 383. Ransom Note
// https://leetcode.com/problems/ransom-note/
package RansomNote

func canConstruct(ransomNote string, magazine string) bool {
	mask := make([]bool, len(magazine))
	for i := range ransomNote {
		hasChar := false
		for j := range magazine {
			if mask[j] {
				continue
			}
			if ransomNote[i] == magazine[j] {
				hasChar = true
				mask[j] = true
				break
			}

			mask[j] = false
		}

		if !hasChar {
			return false
		}
	}
	return true
}
