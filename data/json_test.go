package data

import (
	"strings"
	"testing"
)

func TestDecodeJson_WordInfo(t *testing.T) {
	r := strings.NewReader(`
		{
			"word": "happy",
			"synonyms": [
				"glad",
				"joyful"
			]
		}`,
	)

	wordInfo, err := DecodeJson[WordInfo](r)
	if err != nil {
		t.Fatalf("error decoding wordInfo: %v\n", err)
	}

	if wordInfo.Word != "happy" {
		t.Fatalf("wordInfo.Word expected 'happy', got '%v'\n", wordInfo.Word)
	}
	if len(wordInfo.Synonyms) != 2 {
		t.Fatalf("wordInfo.Synonyms length expected 2, got %v\n", len(wordInfo.Synonyms))
	}
	if wordInfo.Synonyms[0] != "glad" {
		t.Fatalf("wordInfo.Synonyms[0] expected 'glad', got '%v'\n", wordInfo.Synonyms[0])
	}
	if wordInfo.Synonyms[1] != "joyful" {
		t.Fatalf("wordInfo.Synonyms[1] expected 'joyful', got '%v'\n", wordInfo.Synonyms[1])
	}
}

func TestDecodeJson_SynonymsInfo(t *testing.T) {
	f := strings.NewReader(`
		{
			"synonyms": [
				"glad",
				"joyful"
			]
		}`,
	)

	synonymsInfo, err := DecodeJson[SynonymsInfo](f)
	if err != nil {
		t.Fatalf("error decoding synonymsInfo: %v\n", err)
	}

	if len(synonymsInfo.Synonyms) != 2 {
		t.Fatalf("synonymsInfo.Synonyms length expected 2, got %v\n", len(synonymsInfo.Synonyms))
	}
	if synonymsInfo.Synonyms[0] != "glad" {
		t.Fatalf("wordInfo.Synonyms[0] expected 'glad', got '%v'\n", synonymsInfo.Synonyms[0])
	}
	if synonymsInfo.Synonyms[1] != "joyful" {
		t.Fatalf("wordInfo.Synonyms[1] expected 'joyful', got '%v'\n", synonymsInfo.Synonyms[1])
	}
}
