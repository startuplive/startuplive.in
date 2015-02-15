package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	// "github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/user"
	"labix.org/v2/mgo/bson"
	"strconv"
	// "github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_WikiEntry0 = &Page{
		Title: Render(
			func(ctx *Context) error {
				entry, err := getWikiEntry(ctx)
				if err != nil {
					return err
				}
				ctx.Response.WriteString(entry.Title.String())
				ctx.Response.WriteString(" | Admin")
				return nil
			},
		),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: PageScripts,
		Content: Views{
			adminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					entry, err := getWikiEntry(ctx)
					if err != nil {
						return nil, err
					}

					views := Views{
						H2(entry.Title.String()),
						HR(),
						&Form{
							SubmitButtonText:  "Update",
							SubmitButtonClass: "button",
							FormID:            "wikientry",
							ExcludedFields: []string{
								"CreatedBy",
								"CreatedAt",
								"Comments",
							},
							GetModel: FormModel(entry),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								entry := formModel.(*models.WikiEntry)
								entry.CreatedAt.SetTodayUTC()

								var author models.Person
								found, err := user.OfSession(ctx.Session, &author)
								if err != nil {
									return "", nil, err
								}
								if found {
									entry.CreatedBy.Set(&author)
								}

								return "", Admin_Wiki, entry.Save()

							},
						},
						HR(),
						DynamicView(
							func(ctx *Context) (view View, err error) {
								return ShowAnswers(entry)
							},
						),
						DynamicView(
							func(ctx *Context) (view View, err error) {
								return GetAnswerForm(entry)
							},
						),
					}
					return views, nil
				},
			),
		},
	}
}

func getWikiEntry(ctx *Context) (entry *models.WikiEntry, err error) {

	if string(ctx.URLArgs[0]) == "new" {
		entry = models.NewWikiEntry()
		return entry, nil
	} else {
		id := bson.ObjectIdHex(ctx.URLArgs[0])
		found, err := models.Wiki.TryDocumentWithID(id, &entry)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, NotFound("404: Wiki Entry not found")
		}
		return entry, nil
	}
	return nil, nil
}

func ShowAnswers(entry *models.WikiEntry) (view View, err error) {
	// entry, err := getWikiEntry(ctx)
	// entry := ctx.Data.(*PageData).WikiEntry
	// if err != nil {
	// 	return nil, err
	// }
	var views Views
	// var comment models.WikiComment
	for j := 0; j < len(entry.Comments); j++ {
		i := j
		debug.Print(len(entry.Comments))
		comment := entry.Comments[i]
		var author models.Person
		found, err := comment.CommentedBy.TryGet(&author)
		if err != nil {
			return nil, err
		}
		commentedBy := "anonymous"
		if found {
			commentedBy = author.Name.String()
		}

		views = append(views,
			DIV("wiki-answers",
				DIV("wiki-answer",
					DIV("wiki-answer-profile", ViewOrError(author.Image_50x50())),
					SPAN("wiki-answer-author", commentedBy),
					BR(),
					SPAN("wiki-answer-date", comment.CommentedAt.Get()),
					&Form{
						SubmitButtonText:    "Delete",
						SubmitButtonClass:   "delete",
						FormID:              "deleteComment" + strconv.Itoa(i),
						SubmitButtonConfirm: "Are you sure you want to delete this comment: " + comment.Content.Get() + "?",
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {

							return "", StringURL("."), entry.RemoveComment(i)
						},
					},
					DIV("wiki-answer-content", comment.Content.Get()),
					DivClearBoth(),
				),
			),
		)
	}
	return views, nil

}

func GetAnswerForm(entry *models.WikiEntry) (view View, err error) {
	// entry, err := getWikiEntry(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	return Views{
		&Form{
			SubmitButtonText:  "Add Comment",
			SubmitButtonClass: "button",
			FormID:            "wikicomment",
			GetModel: func(form *Form, ctx *Context) (interface{}, error) {
				return &models.WikiCommentForm{}, nil
			},
			OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
				comment := formModel.(*models.WikiCommentForm)

				wikiComment := models.NewWikiComment()

				wikiComment.CommentedAt.SetTodayUTC()

				var author models.Person
				found, err := user.OfSession(ctx.Session, &author)
				if err != nil {
					return "", nil, err
				}
				if found {
					wikiComment.CommentedBy.Set(&author)
				}

				wikiComment.Content = comment.Content
				wikiComment.Votes.Set(0)
				entry.Comments = append(entry.Comments, *wikiComment)

				return "", StringURL("."), entry.Save()

			},
		},
	}, nil

}
