package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func checkRsMessage(msg string) bool {
	rsWords := map[string]bool{
		"надо": false,
		"было": false,
		"рс":   false,
	}
	words := strings.Fields(strings.ToLower(msg))
	for _, word := range words {
		word = strings.Map(
			func(r rune) rune {
				if !unicode.IsLetter(r) {
					return -1
				}

				return r
			},
			word,
		)

		if _, ok := rsWords[word]; ok {
			rsWords[word] = true

			isRsMsg := true
			for _, v := range rsWords {
				if !v {
					isRsMsg = false
					break
				}
			}

			if isRsMsg {
				return true
			}
		}
	}

	return false
}

func formatReply(rsCounts map[string]int) string {
	keys := make([]string, 0, len(rsCounts))
	for k := range rsCounts {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		return "Статистика пустая!"
	}

	sort.Slice(
		keys,
		func(i, j int) bool {
			return rsCounts[keys[i]] > rsCounts[keys[j]]
		},
	)

	msg := "\"Надо было на РС\" Global Count:\n"
	for _, k := range keys {
		msg += fmt.Sprintf("%s: %d\n", k, rsCounts[k])
	}

	return msg
}
