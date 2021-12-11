package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* Original Segments:
 *
 *  aaaa
 * b    c
 * b    c
 *  dddd
 * e    f
 * e    f
 *  gggg
 */

func parseLine(line string) (signals, outputs []string) {
	halfs := strings.Split(line, "|")
	s := strings.Fields(halfs[0])
	o := strings.Fields(halfs[1])

	signals = append(signals, s...)
	outputs = append(outputs, o...)

	return
}

func removeChars(from, remove string) string {
	for _, r := range remove {
		c := string(r)
		from = strings.Replace(from, c, "", -1)
	}
	return from
}

func removeDuplicateChars(str string) string {
	fixedStr := ""
	for _, c := range str {
		res := strings.Contains(fixedStr, string(c))
		if !res {
			fixedStr += string(c)
		}
	}
	return fixedStr
}

func getSegmentsByLength(segments []string, length int) []string {
	result := []string{}
	for _, s := range segments {
		if len(s) == length {
			result = append(result, s)
		}
	}
	return result
}

func union(segments ...string) string {
	res := ""
	for _, v := range segments {
		res += v
	}
	return removeDuplicateChars(res)
}

func compare(one, two string) bool {
	if len(one) != len(two) {
		return false
	}

	for _, c := range one {
		if !strings.Contains(two, string(c)) {
			return false
		}
	}
	return true
}

func exponent(exp, base int) int {
	output := 1
	for exp != 0 {
		output *= base
		exp--
	}
	return output
}

func fixMapping(inputs []string) map[int]string {
	fixedToBroken := map[rune]rune{}
	digits := map[int]string{}

	// Find unique: l2: 1, l4: 4, l3: 7, l7: 8
	digits[1] = getSegmentsByLength(inputs, 2)[0]
	digits[4] = getSegmentsByLength(inputs, 4)[0]
	digits[7] = getSegmentsByLength(inputs, 3)[0]
	digits[8] = getSegmentsByLength(inputs, 7)[0]

	// a = 7 - 1
	newA := removeChars(digits[7], digits[1])
	fixedToBroken['a'] = rune(newA[0])

	// g = 9 - (4+7)
	// 9 is the only digit that has exactly one extra segment to union(4,7)
	fourAndSeven := union(digits[4], digits[7])
	nines := getSegmentsByLength(inputs, 6)
	for _, nine := range nines {
		if n := removeChars(nine, fourAndSeven); len(n) == 1 {
			fixedToBroken['g'] = rune(n[0])
			digits[9] = nine
			break
		}
	}

	// e = 8 - 9
	newE := removeChars(digits[8], digits[9])
	fixedToBroken['e'] = rune(newE[0])

	// 6 => 8 - 6 must be contained in 1
	sixes := getSegmentsByLength(inputs, 6)
	for _, six := range sixes {
		intersection := removeChars(digits[8], six)
		intersection = removeChars(digits[1], intersection)
		if len(intersection) == 1 {
			digits[6] = six
			break
		}
	}

	// c = 8 - 6
	newC := removeChars(digits[8], digits[6])
	fixedToBroken['c'] = rune(newC[0])

	// f = 1 - c
	newF := removeChars(digits[1], string(fixedToBroken['c']))
	fixedToBroken['f'] = rune(newF[0])

	// 5 = 6 - e
	five := removeChars(digits[6], string(fixedToBroken['e']))
	digits[5] = five

	// 2 => 5 - c - e must be contained in segment with length 5 (with 2
	// addinitonal segments)
	twos := getSegmentsByLength(inputs, 5)
	ce := union(string(fixedToBroken['c']), string(fixedToBroken['e']))
	twoMustContain := removeChars(digits[5], ce)
	for _, two := range twos {
		if t := removeChars(two, twoMustContain); len(t) == 2 {
			digits[2] = two
			break
		}
	}

	// 3 = 2 + f - e
	twof := union(digits[2], string(fixedToBroken['f']))
	three := removeChars(twof, string(fixedToBroken['e']))
	digits[3] = three

	// d = 2 - a - c - e - g
	aceg := union(string(fixedToBroken['a']), string(fixedToBroken['c']),
		string(fixedToBroken['e']), string(fixedToBroken['g']))
	newD := removeChars(digits[2], aceg)
	fixedToBroken['d'] = rune(newD[0])

	// 0 = 8 - d
	zero := removeChars(digits[8], string(fixedToBroken['d']))
	digits[0] = zero

	return digits
}

func decode(outputs []string, digits map[int]string) int {
	res := 0
	max := len(outputs) - 1

	for i, out := range outputs {
		for j := 0; j < 10; j++ {
			digit := digits[j]
			isEqual := compare(out, digit)
			if isEqual {
				res += j * exponent(max-i, 10)
				break
			}
		}
	}

	return res
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	totalNr := 0

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		inputs, outputs := parseLine(line)

		digits := fixMapping(inputs)
		value := decode(outputs, digits)

		totalNr += value
	}

	fmt.Println("total:", totalNr)
}
