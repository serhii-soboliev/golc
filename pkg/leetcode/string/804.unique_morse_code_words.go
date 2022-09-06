package string

/*
804. Unique Morse Code Words
https://leetcode.com/problems/unique-morse-code-words/
*/

func uniqueMorseRepresentations(words []string) int {
    morses := [] string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    transformations := make(map[string]bool)
    for _, s := range words {
        transformation := ""
        for _, r := range s {
            index := r - 'a'
            transformation += morses[index]
        }
        transformations[transformation] = true
    }
    return len(transformations)
}