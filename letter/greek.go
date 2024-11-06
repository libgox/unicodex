package letter

const GreekUpperMin = 'Α'
const GreekUpperMax = 'Ω'
const GreekLowerMin = 'α'
const GreekLowerMax = 'ω'

const GreekUpperCount = 24
const GreekLowerCount = 24
const GreekCount = 48

const GreekUpperLetters = "ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ"
const GreekLowerLetters = "αβγδεζηθικλμνξοπρστυφχψω"
const GreekLetters = "αβγδεζηθικλμνξοπρστυφχψωΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ"

var GreekUpperLetterRunes = []rune{'Α', 'Β', 'Γ', 'Δ', 'Ε', 'Ζ', 'Η', 'Θ', 'Ι', 'Κ', 'Λ', 'Μ', 'Ν', 'Ξ', 'Ο', 'Π', 'Ρ', 'Σ', 'Τ', 'Υ', 'Φ', 'Χ', 'Ψ', 'Ω'}
var GreekLowerLetterRunes = []rune{'α', 'β', 'γ', 'δ', 'ε', 'ζ', 'η', 'θ', 'ι', 'κ', 'λ', 'μ', 'ν', 'ξ', 'ο', 'π', 'ρ', 'σ', 'τ', 'υ', 'φ', 'χ', 'ψ', 'ω'}
var GreekLetterRunes = []rune{'α', 'β', 'γ', 'δ', 'ε', 'ζ', 'η', 'θ', 'ι', 'κ', 'λ', 'μ', 'ν', 'ξ', 'ο', 'π', 'ρ', 'σ', 'τ', 'υ', 'φ', 'χ', 'ψ', 'ω', 'Α', 'Β', 'Γ', 'Δ', 'Ε', 'Ζ', 'Η', 'Θ', 'Ι', 'Κ', 'Λ', 'Μ', 'Ν', 'Ξ', 'Ο', 'Π', 'Ρ', 'Σ', 'Τ', 'Υ', 'Φ', 'Χ', 'Ψ', 'Ω'}

func RandGreekUpper(length int) string {
	return RandLetters(GreekUpperLetterRunes, length)
}

func RandGreekLower(length int) string {
	return RandLetters(GreekLowerLetterRunes, length)
}

func RandGreek(length int) string {
	return RandLetters(GreekLetterRunes, length)
}
