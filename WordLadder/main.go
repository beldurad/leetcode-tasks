type queueElem struct {
	word    string
	pathLen int
	next    *queueElem
}

type wordPermutations struct {
	word  string
	bytes []byte
	i     int
	r     int
}

func newWordPermutations(word string) *wordPermutations {
	return &wordPermutations{
		word:  word,
		bytes: []byte(word),
	}
}

func (w *wordPermutations) next() (word []byte, ok bool) {
	if w.r+'a' > 'z' {
		w.i += 1
		w.r = 0
		w.bytes = []byte(w.word)
	}
	if w.i >= len(w.bytes) {
		word, ok = nil, false
		return
	}
	w.bytes[w.i] = byte(w.r + 'a')
	w.r += 1
	word, ok = w.bytes, true
	return
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordsMap := make(map[string]int)
	for i, w := range wordList {
		wordsMap[w] = i
	}
	cur := &queueElem{
		word:    beginWord,
		pathLen: 1,
	}
	tail := cur
	for cur != nil {
		perms := newWordPermutations(cur.word)
		for bytes, ok := perms.next(); ok; bytes, ok = perms.next() {
			wordIndex, exists := wordsMap[string(bytes)]
			if !exists {
				continue
			}
			word := wordList[wordIndex]
			if word == endWord {
				return cur.pathLen + 1
			}
			tail.next = &queueElem{
				word:    word,
				pathLen: cur.pathLen + 1,
			}
			tail = tail.next
			delete(wordsMap, word)

		}
		delete(wordsMap, cur.word)
		cur = cur.next
	}
	return 0
}
