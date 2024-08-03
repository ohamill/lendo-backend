package data

type WordInfo struct {
	Word string `json:"word" validate:"required"`
	SynonymsInfo
}

func NewWordInfo(word string, synonyms []string) WordInfo {
	return WordInfo{
		Word: word,
		SynonymsInfo: SynonymsInfo{
			Synonyms: synonyms,
		},
	}
}

type WordsInfo struct {
	Words []string `json:"words"`
}

func NewWordsInfo(words []string) WordsInfo {
	return WordsInfo{
		Words: words,
	}
}

type SynonymsInfo struct {
	Synonyms []string `json:"synonyms" validate:"required"`
}

func NewSynonymsInfo(synonyms []string) SynonymsInfo {
	return SynonymsInfo{
		Synonyms: synonyms,
	}
}
