// Package rotationalcipher implements a simple Caesar cypher.
package rotationalcipher

// RotationalCipher returns the input string rotated by an arbitrary
// number of positions. Only letters are affected. Case is maintained.
func RotationalCipher(input string, rot int) string {
	var output []rune

	for _, r := range input {
		base := 0
		switch {
		case r >= 'A' && r <= 'Z':
			base = int('A')
		case r >= 'a' && r <= 'z':
			base = int('a')
		}
		if base != 0 {
			r = rune(((int(r) - base + rot) % 26) + base)
		}
		output = append(output, r)
	}
	return string(output)
}