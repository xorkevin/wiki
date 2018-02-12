package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	paragraphInput  = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam suscipit diam fringilla tortor tempor bibendum. Fusce tellus libero, congue non elementum et, tristique a ante. Donec vel lectus pulvinar, vulputate ligula quis, venenatis leo. Nunc blandit a lacus id consectetur. Duis vitae elementum urna. Fusce hendrerit diam vitae feugiat aliquam. Nulla convallis id felis a pretium. Vestibulum at convallis magna. Donec ut odio mi. Curabitur malesuada purus mauris, vel ornare quam vulputate eget. Quisque vel nisi tristique, iaculis ante nec, volutpat neque. Fusce ligula diam, luctus id odio nec, ultricies mollis nunc. Ut bibendum lorem nulla, at mollis tortor semper a. Sed tempus nisl eu fringilla molestie. Praesent suscipit risus enim, ac ornare augue mollis semper.`
	paragraphInput2 = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam suscipit diam fringilla tortor tempor bibendum. Fusce tellus libero, congue non elementum et, tristique a ante. Donec vel lectus pulvinar, vulputate ligula quis, venenatis leo. 0123456789112345678921234567893123456789 Nunc blandit a lacus id consectetur. Duis vitae elementum urna. Fusce hendrerit diam vitae feugiat aliquam. Nulla convallis id felis a pretium. Vestibulum at convallis magna. Donec ut odio mi. Curabitur malesuada purus mauris, vel ornare quam vulputate eget. Quisque vel nisi tristique, iaculis ante nec, volutpat neque. Fusce ligula diam, luctus id odio nec, ultricies mollis nunc. Ut bibendum lorem nulla, at mollis tortor semper a. Sed tempus nisl eu fringilla molestie. Praesent suscipit risus enim, ac ornare augue mollis semper.`

	lineWrap = 32
)

func main() {
	{
		words := strings.Split(paragraphInput, " ")
		cost, formated := dplinebreak(words, lineWrap)
		fmt.Println("cost:", cost)
		fmt.Println(charRuler(lineWrap))
		fmt.Println(formated)
	}
	{
		words := strings.Split(paragraphInput2, " ")
		cost, formated := dplinebreak(words, lineWrap)
		fmt.Println("cost:", cost)
		fmt.Println(charRuler(lineWrap))
		fmt.Println(formated)
	}
}

func dplinebreak(words []string, width int) (int, string) {
	wordLengths := make([]int, 0, len(words))
	for _, i := range words {
		wordLengths = append(wordLengths, len(i))
	}
	cost, parts := dpCalcPartition(wordLengths, width)
	return cost, formatWords(words, parts)
}

type (
	dpPart struct {
		cost int
		prev int
	}
)

func dpCalcPartition(wordLengths []int, width int) (int, []int) {
	l := len(wordLengths)
	cache := make([]dpPart, 0, l+1)
	cache = append(cache, dpPart{
		cost: 0,
		prev: 0,
	})
	for i := 1; i < l+1; i++ {
		minCost := -1
		minPrev := 0
		for j := 0; j < i; j++ {
			c := width - sumWordSlice(wordLengths[j:i])
			if c < 0 {
				continue
			}
			c = c*c + cache[j].cost
			if minCost < 0 || c <= minCost {
				minCost = c
				minPrev = j
			}
		}
		if minCost < 0 {
			minCost = wordLengths[i-1]
			minCost = minCost*minCost + cache[i-1].cost
			minPrev = i - 1
		}
		cache = append(cache, dpPart{
			cost: minCost,
			prev: minPrev,
		})
	}
	parts := []int{}
	for i := l; i > 0; i = cache[i].prev {
		parts = append(parts, i)
	}
	for i := 0; i < len(parts)/2; i++ {
		k := parts[i]
		parts[i] = parts[len(parts)-i-1]
		parts[len(parts)-i-1] = k
	}
	return cache[l].cost, parts
}

func sumWordSlice(s []int) int {
	k := 0
	for _, i := range s {
		k += i
	}
	return k + len(s) - 1
}

func formatWords(words []string, breaks []int) string {
	lines := []string{}
	prev := 0
	for _, k := range breaks {
		if k == prev {
			continue
		}
		lines = append(lines, strings.Join(words[prev:k], " "))
		prev = k
	}
	if prev < len(words) {
		lines = append(lines, strings.Join(words[prev:], " "))
	}
	return strings.Join(lines, "\n")
}

func charRuler(width int) string {
	b := bytes.Buffer{}
	for i := 0; i < lineWrap; i++ {
		b.WriteString(fmt.Sprintf("%d", i%10))
	}
	return b.String()
}
