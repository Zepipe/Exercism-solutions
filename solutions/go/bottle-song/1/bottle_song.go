package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	if startBottles <= 0 || takeDown <= 0 || takeDown > startBottles || startBottles > 10 {
		return []string{}
	}

	songLyrics := make([]string, 0)
	bottlesNumToName := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for i := 0; i < takeDown; i++ {
		currentBottles := startBottles - i
		nextBottles := currentBottles - 1
		
		// First two lines
		if currentBottles == 1 {
			songLyrics = append(songLyrics, fmt.Sprintf("%s green bottle hanging on the wall,", bottlesNumToName[currentBottles]))
			songLyrics = append(songLyrics, fmt.Sprintf("%s green bottle hanging on the wall,", bottlesNumToName[currentBottles]))
		} else {
			songLyrics = append(songLyrics, fmt.Sprintf("%s green bottles hanging on the wall,", bottlesNumToName[currentBottles]))
			songLyrics = append(songLyrics, fmt.Sprintf("%s green bottles hanging on the wall,", bottlesNumToName[currentBottles]))
		}
		
		// Third line
		songLyrics = append(songLyrics, "And if one green bottle should accidentally fall,")
        
		// Fourth line
		if nextBottles == 1 {
			songLyrics = append(songLyrics, fmt.Sprintf("There'll be one green bottle hanging on the wall."))
		} else if nextBottles == 0 {
			songLyrics = append(songLyrics, "There'll be no green bottles hanging on the wall.")
		} else {
			songLyrics = append(songLyrics, fmt.Sprintf("There'll be %s green bottles hanging on the wall.", 
				strings.ToLower(bottlesNumToName[nextBottles])))
		}
		
		// Add blank line between verses except after the last verse
		if i < takeDown-1 {
			songLyrics = append(songLyrics, "")
		}
	}
	
	return songLyrics
}