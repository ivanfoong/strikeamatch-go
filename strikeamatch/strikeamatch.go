package strikeamatch

import (
  "fmt"
  "unicode"
  "github.com/ivanfoong/stringutils-go/stringutils"
)

func RunesLowerCase(words []string) (runes [][]string) {
  for _, word := range words {
    currentWordRunes := []string{} 
    for _, rune := range word {
      runeString := fmt.Sprintf("%c", unicode.ToLower(rune))
      currentWordRunes = append(currentWordRunes, runeString)
    }
    runes = append(runes, currentWordRunes)
  }
  return runes
}

func scoreStrings(s1 []string, s2 []string) (score float64) {
  size1 := len(s1)
  size2 := len(s2)
  
  occurrence1 := stringutils.StringOccurrence(s1)
  occurrence2 := stringutils.StringOccurrence(s2)
  
  smallerOccurrence := occurrence2
  largerOccurrence := occurrence1
  if size1 < size2 {
    smallerOccurrence = occurrence1
    largerOccurrence = occurrence2
  }
  
  intersectionOccurrence := 0
  
  for s, smallerOccurrence := range smallerOccurrence {
    largerOccurrence, ok := largerOccurrence[s]
    if ok && largerOccurrence > 0 {
      if smallerOccurrence < largerOccurrence {
        intersectionOccurrence += smallerOccurrence
      } else {
        intersectionOccurrence += largerOccurrence
      }
    } 
  }
  
  return (2.0 * float64(intersectionOccurrence)) / float64(size1 + size2)
}

func CompareString(text1 string, text2 string) (score float64) {
  words1 := stringutils.Words(text1)
  words2 := stringutils.Words(text2)
  
  runes1 := RunesLowerCase(words1)
  runes2 := RunesLowerCase(words2)
  
  pairs1 := []string{}
  for _, runes := range runes1 {
    pairs1 = append(pairs1, stringutils.Pairs(runes)...)
  }
  
  pairs2 := []string{}
  for _, runes := range runes2 {
    pairs2 = append(pairs2, stringutils.Pairs(runes)...)
  }
  
  score = scoreStrings(pairs1, pairs2)
  return score
}