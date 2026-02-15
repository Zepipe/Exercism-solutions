package proverb

import "fmt"

func Proverb(rhyme []string) []string {
    if len(rhyme) == 0 {
        return []string{}
    }
    
    lines := []string{}
    
    for i := 0; i < len(rhyme)-1; i++ {
        lines = append(lines, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
    }
    
    lines = append(lines, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
    
    return lines
}
