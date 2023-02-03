package logic

import (
	"fmt"
	"strings"

	"github.com/Funmi4194/myMod/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckDocument(documentName string, w repository.WordCount) (*repository.WordCount, error) {

	text := repository.WordCount{
		DocumentName: strings.ToLower(documentName),
	}

	if err := text.FindDocument(); err != nil {
		return nil, err
	}

	if text.ID.String() == primitive.NilObjectID.String() {
		return nil, fmt.Errorf("document not found")
	}

}
