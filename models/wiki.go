package models

import (
	// "errors"
	// "github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
)

var Wiki = mongo.NewCollection("wiki", "Title")

///////////////////////////////////////////////////////////////////////////////
// Wiki

func NewWikiEntry() *WikiEntry {
	var doc WikiEntry
	Wiki.InitDocument(&doc)
	return &doc
}

func NewWikiComment() *WikiComment {
	w := new(WikiComment)
	mongo.InitRefs(w)
	return w
}

type WikiEntry struct {
	mongo.DocumentBase `bson:",inline"`
	CreatedBy          mongo.Ref `model:"to=people"`
	CreatedAt          model.Date
	Title              model.String
	Content            model.RichText `view:"label=Wiki Content|rows=40"`
	Comments           []WikiComment  `view:"hidden"`
	Files              []media.BlobRef
	Public             model.Bool
}

type WikiComment struct {
	CommentedBy mongo.Ref `model:"to=people"`
	CommentedAt model.Date
	Content     model.Text
	Votes       model.Int
}

type WikiCommentForm struct {
	Content model.Text
}

func (self *WikiEntry) RemoveComment(i int) error {
	self.Comments = append(self.Comments[:i], self.Comments[i+1:]...)
	return self.Save()
}
