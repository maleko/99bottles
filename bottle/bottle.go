package bottle

import "strconv"

func Verses(starting int, ending int) (verses string) {
	verses = ""
	for currentVerse := starting; currentVerse > ending; currentVerse-- {
		verses += verse(currentVerse)
	}
	return
}

func verse(number int) (verse string) {
	switch {
	case number == 0:
		verse = "No more bottles of beer on the wall, " +
			"no more bottles of beer.\n" +
			"Go to the store and buy some more, " +
			"99 bottles of beer on the wall.\n"
	case number == 1:
		verse = "1 bottle of beer on the wall, " +
			"1 bottle of beer.\n" +
			"Take it down and pass it around, " +
			"no more bottles of beer on the wall.\n"
	case number == 2:
		verse = "2 bottles of beer on the wall, " +
			"2 bottles of beer.\n" +
			"Take one down and pass it around, " +
			"1 bottle of beer on the wall.\n"
	default:
		verse = strconv.Itoa(number) + " bottles of beer on the wall, " +
			strconv.Itoa(number) + " bottles of beer.\n" +
			"Take one down and pass it around, " +
			strconv.Itoa(number-1) + " bottles of beer on the wall.\n"
	}
	return
}
