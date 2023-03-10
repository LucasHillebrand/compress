package main

import "os"

func NextChars(org string, start, length int) string {
	out := ""
	for i := start; i < start+length && i < len(org); i++ {
		out += string(org[i])
	}
	return out
}

func Count(org, keyword string) int {
	out := 0

	for i := 0; i < len(org); i++ {
		if NextChars(org, i, len(keyword)) == keyword {
			out++
		}
	}

	return out
}

func Split(org, keyword string) []string {
	out := make([]string, Count(org, keyword)+1)

	for i, col := 0, 0; i < len(org); i++ {
		if NextChars(org, i, len(keyword)) == keyword {
			col++
			i += len(keyword) - 1
		} else {
			out[col] += string(org[i])
		}
	}
	return out
}

func ReadFileList(filename string) (out []string, err string) {

	bytes, errstr := os.ReadFile(filename)
	if errstr != nil {
		err = errstr.Error()
	} else {
		str := ""
		for i := 0; i < len(bytes); i++ {
			str += string(bytes[i])
		}
		return Split(str, "\n"), ""
	}

	return
}

func toString(org []string) string {
	out := ""
	for i := 0; i < len(org); i++ {
		out += org[i]
		if i+1 < len(org) {
			out += "\n"
		}
	}
	//fmt.Printf("%s,--\n", out)
	return out
}
