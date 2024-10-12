package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	transpositionMap := map[string]string{
		"C":   "F",
		"C#":  "F#",
		"D♭":  "G♭",
		"D":   "G",
		"D#":  "G#",
		"E♭":  "A♭",
		"E":   "A",
		"F":   "B♭",
		"F#":  "B",
		"G♭":  "B",
		"G":   "C",
		"G#":  "C#",
		"A♭":  "D♭",
		"A":   "D",
		"A#":  "D#",
		"B♭":  "E♭",
		"B":   "E",
		"C'":  "F'",
		"C#'": "F#'",
		"D♭'": "G♭'",
		"D'":  "G'",
		"D#'": "G#'",
		"E♭'": "A♭'",
		"E'":  "A'",
		"F'":  "B♭'",
		"F#'": "B'",
		"G♭'": "B'",
		"G'":  "C'",
	}

	fmt.Print("🎼 ")
	reader := bufio.NewReader(os.Stdin)
	notes, _ := reader.ReadString('\n')
	notes = strings.TrimSpace(notes)
	notes = strings.ToUpper(notes)
	notes = strings.ReplaceAll(notes, "’", "'")

	notesArray := strings.Fields(notes)

	for i, note := range notesArray {
		if len(note) > 1 && note[1] == 'B' {
			notesArray[i] = note[:1] + "♭" + note[2:]
		}
	}

	var transposedNotesArray []string
	for i, note := range notesArray {
		if transposedNote, exists := transpositionMap[note]; exists {
			if len(transposedNote) > len(note) {
				notesArray[i] += " "
			} else if len(transposedNote) < len(note) {
				transposedNote += " "
			}
			transposedNotesArray = append(transposedNotesArray, transposedNote)
		} else {
			transposedNotesArray = append(transposedNotesArray, note)
		}
	}

	notesNormalized := strings.Join(notesArray, " ")
	newNotes := strings.Join(transposedNotesArray, " ")

	fmt.Println()
	fmt.Println("1: ", notesNormalized)
	fmt.Println("2: ", newNotes)
}
