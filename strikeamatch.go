package strikeamatch

import (
  "fmt"
  "strings"
  "unicode"
  "regexp"
)

func runeInSlice(r rune, runes []rune) bool {
  for _, cr := range runes {
    if cr == r {
      return true
    }
  }
  return false
}

func splitWord(text string) ([]string) {
  var words []string 
  f := func(c rune) bool {
    delimiters := []rune{
      ' ',
      ',',
      ':',
      ';',
      '(',
      ')',
      '-',
    }
    return runeInSlice(c, delimiters)
  }
  for _, word := range strings.FieldsFunc(text, f) {
    words = append(words, word)
  }
  return words
}

func filterVolumeWords(words []string) ([]string, []string) {
  var cleanWords []string
  var volumeWords []string
  
  for _, word := range words {
    r, err := regexp.Compile(`\d+([Mm][Ll]|[Ll]|[Mm][Gg]|[Gg])$`)
    
    if err != nil {
      fmt.Printf("There is a problem with your regexp.\n")
      break
    }
    
    if r.MatchString(word) {
      volumeWords = append(volumeWords, word)
    } else {
      cleanWords = append(cleanWords, word)
    }
  }
  
  return cleanWords, volumeWords
}

func generateCharacterPairsForWord(word string) ([]string) {
  var pairs []string 
  runes := []rune(word)
  for index, rune := range runes {
    if index == 0 {
      continue
    }
    
    pair := fmt.Sprintf("%c%c", unicode.ToLower(runes[index-1]), unicode.ToLower(rune))
    pairs = append(pairs, pair)
  }
  return pairs
}

func generateCharacterPairs(text string) ([]string) {
  var allPairs []string
  
  words := splitWord(text)
  
  for _, word := range words {
    pairs := generateCharacterPairsForWord(word)
    allPairs = append(allPairs, pairs...)
  }
  
  return allPairs
}

func countCharacterPairs(pairs []string) (map[string]int) {
  pairCounts := make(map[string]int)
  
  for _, pair := range pairs {
    count, exist := pairCounts[pair]
    if exist {
      count++
    } else {
      count = 1
    }
    pairCounts[pair] = count
  }
  
  return pairCounts
}

func CompareString(text1 string, text2 string) (float64) {
  score := 0.0
  
  pairs1 := generateCharacterPairs(text1)
  pairs2 := generateCharacterPairs(text2)
  
  pairsSize1 := len(pairs1)
  pairsSize2 := len(pairs2)
  
  pairsCount1 := countCharacterPairs(pairs1)
  pairsCount2 := countCharacterPairs(pairs2)
  
  // fmt.Printf("%v\n", pairsCount1)
  // fmt.Printf("%v\n", pairsCount2)
  
  smallerPairCount := pairsCount2
  largerPairCount := pairsCount1
  if pairsSize1 < pairsSize2 {
    smallerPairCount = pairsCount1
    largerPairCount = pairsCount2
  }
  
  intersectionCount := 0
  
  for pair, smallerCount := range smallerPairCount {
    largerCount, ok := largerPairCount[pair]
    if ok && largerCount > 0 {
      if smallerCount < largerCount {
        intersectionCount += smallerCount
      } else {
        intersectionCount += largerCount
      }
    } 
  }
  
  score = (2.0 * float64(intersectionCount)) / float64(pairsSize1 + pairsSize2)
  
  return score
}