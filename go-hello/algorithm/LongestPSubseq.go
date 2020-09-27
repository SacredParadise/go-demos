package main

import "bytes"

func findLPS(input string) int {
	if input == "" {
		return 0
	}

	runes := ([]rune)(input)
	var buffer bytes.Buffer
	maxLen := findLPSRecursive(&buffer, runes, 0, len(runes) - 1)
	println(maxLen)
	var lps string
	if maxLen % 2 == 0 {
		lps = buffer.String()
		lps += reverseString(buffer.String())
	} else {
		l := buffer.Len()
		var lpsRunes []rune
		for i := 0; i < l - 1; i++ {
			v, _, _ := buffer.ReadRune()
			lpsRunes = append(lpsRunes, v)
		}
		lps = string(lpsRunes)
		center, _, _ := buffer.ReadRune()
		lps = lps + string(center)+ reverseString(lps)
	}

	println(lps)

	return maxLen
}

func findLPSRecursive(buffer *bytes.Buffer, runes []rune, start int, end int) int {
    len := 0
    for (start <= end) {

    	if (start == end) {
    		buffer.WriteRune(runes[start])
    		len += 1
    		return len
		}

		if (runes[start] == runes[end]) {
			buffer.WriteRune(runes[start])
			len += 2
			start++
			end--
			continue
		}

		var buffer1 bytes.Buffer
    	var buffer2 bytes.Buffer
    	len1 := findLPSRecursive(&buffer1, runes, start + 1, end)
    	len2 := findLPSRecursive(&buffer2, runes, start, end -1)
		if len1 > len2 {
			len += len1
			buffer.Write(buffer1.Bytes())
		} else {
			len += len2
			buffer.Write(buffer2.Bytes())
		}
		return len
	}

	return len
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func reverseString(str string) string {
	runes := ([]rune)(str)
	for from, to := 0, len(runes) -1; from < to; from, to = from +1, to -1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func main() {
	findLPS("abdbca")
	findLPS("cddpd")
}