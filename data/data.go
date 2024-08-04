package data

type Word struct {
	Word string `json:"word" validate:"required"`
}

type Synonym struct {
	Synonym string `json:"synonym" validate:"required"`
}

type Synonyms struct {
	Synonyms []string `json:"synonyms" validate:"required"`
}

type CreateSynonymRequest struct {
	Word
	Synonym
}

type CompleteWordInfo struct {
	Word
	Synonyms
}

func NewCompleteWordInfo(word string, synonyms []string) CompleteWordInfo {
	return CompleteWordInfo{
		Word: Word{
			Word: word,
		},
		Synonyms: Synonyms{
			Synonyms: synonyms,
		},
	}
}

type WordsInfo struct {
	Words []string `json:"words"`
}
