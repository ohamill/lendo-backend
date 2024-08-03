package data

type WordInfo struct {
	Word string `json:"word"`
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

type SynonymsInfo struct {
	Synonyms []string `json:"synonyms"`
}

func NewSynonymsInfo(synonyms []string) SynonymsInfo {
	return SynonymsInfo{
		Synonyms: synonyms,
	}
}
