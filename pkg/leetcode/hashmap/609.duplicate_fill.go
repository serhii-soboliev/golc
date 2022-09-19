package hashmap

import (
	"strings"
)

/*
609. Find Duplicate File in System
https://leetcode.com/problems/find-duplicate-file-in-system/
*/

func findDuplicate(paths []string) [][]string {
	result := [][]string{}
	contentToFile := make(map[string][]string)
	for _, path := range paths {
		sp := strings.Split(path, " ")
		folder := sp[0]
		for i:=1; i<len(sp);i++ {
			file := sp[i]
			contentStartIdx := strings.Index(file, "(")
			filePath := file[:contentStartIdx]
			content := file[contentStartIdx:]
			fullFilePath := folder + "/" + filePath
			contentToFile[content] = append(contentToFile[content], fullFilePath)
		}
	}
	for _,v := range contentToFile {
		if len(v) > 1 {
			result = append(result, v)
		}
	}
	return result
}

func FindDuplicate(paths []string) [][]string {
	return findDuplicate(paths)
}