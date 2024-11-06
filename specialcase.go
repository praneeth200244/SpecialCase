package specialcase

import ("unicode/utf8"
"re")

func Split(source string) (splitEntries []string) {
	// Don't split invalid utf8
	if !utf8.ValidString(source) {
		return []string{source}
	}
	splitEntries = re.sub(r'(?<=[a-z])(?=[A-Z])', ' ', text).lower()
	return
}
