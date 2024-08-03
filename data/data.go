package data

type CreateWordRequest struct {
	Word     string   `json:"word"`
	Synonyms []string `json:"synonyms"`
}

type CreateWordResponse struct {
	Word     string   `json:"word"`
	Synonyms []string `json:"synonyms"`
}

type CreateSynonymsRequest struct {
	Synonyms []string `json:"synonyms"`
}

type CreateSynonymsResponse struct {
	Word     string   `json:"word"`
	Synonyms []string `json:"synonyms"`
}

type GetSynonymsResponse struct {
	Synonyms []string `json:"synonyms"`
}
