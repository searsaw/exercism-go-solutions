package house

import "bytes"

func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

func Verse(subject string, relPhrases []string, nounPhrase string) string {
	if len(relPhrases) == 0 {
		return Embed(subject, nounPhrase)
	}

	return subject + " " + Verse(relPhrases[0], relPhrases[1:], nounPhrase)
}

func Song() string {
	var buffer bytes.Buffer

	length := len(subjects)
	for i := length - 1; i >= 1; i-- {
		buffer.WriteString(Verse("This is", subjects[i:length-1], subjects[length-1]))
		buffer.WriteString("\n\n")
	}
	buffer.WriteString(Verse("This is", subjects[:length-1], subjects[length-1]))

	return buffer.String()
}

var subjects = []string{
	"the horse and the hound and the horn\nthat belonged to",
	"the farmer sowing his corn\nthat kept",
	"the rooster that crowed in the morn\nthat woke",
	"the priest all shaven and shorn\nthat married",
	"the man all tattered and torn\nthat kissed",
	"the maiden all forlorn\nthat milked",
	"the cow with the crumpled horn\nthat tossed",
	"the dog\nthat worried",
	"the cat\nthat killed",
	"the rat\nthat ate",
	"the malt\nthat lay in",
	"the house that Jack built.",
}
