package repository

import (
	"context"
	"strings"

	"github.com/Funmi4194/myMod/controllers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//check for numebers of words in the text

func (w *WordCount) TotalWord() (int, error) {
	content := strings.Join(strings.Fields(w.Content), " ")
	totalCount := len(strings.Split(content, " "))

}

func (w *WordCount) FindDocument() error {
	err := controllers.Usercollection.FindOne(context.Background(), bson.D{{Key: "DocumentName", Value: w.DocumentName}}).Decode(&w)
	if err != mongo.ErrNoDocuments {
		return err
	}
	return nil
}

// func main(){

// content := strings.Join(strings.Fields(document.Content), " ")
// 		totalCount := len(strings.Split(content, " "))

// 		totalCharacters := len(strings.Split(document.Content, ""))

// 		//(total character without space)
// 		content2 := strings.Replace(document.Content, " ", "", -1)
// 		charWithoutSpace := len(strings.Split(content2, ""))

// 		//numbers of sentences
// 		sentences := len(strings.Split(document.Content, "."))

// 		// numbers of paragraphs
// 		paragraphs := len(strings.Split(document.Content, "\n\n"))

// 		//numbers of line
// 		// var line int
// 		// for i:= 0; i < len(document.Content); i++{
// 		lines := strings.SplitAfterN(document.Content, "\n", -1)
// 		line := len(lines)
// 		// }

// 		}
