/*
 * http://en.wikipedia.org/wiki/Pig_Latin
 *
 *   1. In words that begin with consonant  sounds, the initial consonant or
 *      consonant cluster is moved to the end of the word, and "ay" is added, as
 *      in the following examples:
 *          * beast → east-bay
 *          * dough → ough-day
 *          * happy → appy-hay
 *          * question → estion-quay
 *   2. In words that begin with vowel sounds or silent consonants, the syllable
 *      "way" is simply added to the end of the word. In some variants, the
 *      syllable "ay" is added, without the "w" in front.
 *          * another→ another-way or another-ay
 *          * if→ if-way or if-ay
 *   3. In compound words or words with two distinct syllables, each component
 *      word or syllable is sometimes transcribed separately. For example:
 *      birdhouse would be ird-bay-ouse-hay.
 *
 * Transcription varies. A hyphen or apostrophe is sometimes used to facilitate
 * translation back into English. Ayspray, for instance, is ambiguous, but
 * ay-spray means "spray" whereas ays-pray means "prays."
 */

package main

import (
	// "log"
	"strings"
)

func IsVowel(s string) bool {
	// [0] is the backing int, [0:1] is the string we need
	return strings.Index("aeiou", strings.ToLower(s[0:1])) != -1
}

func piglatin(s string) string {
	if IsVowel(s) {
		return s + "-way"
	}
	return s[1:] + "-" + s[0:1] + "ay"
}
