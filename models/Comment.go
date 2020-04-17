package models

import "database/sql"
import "time"
import "github.com/go-sql-driver/mysql"

func MakeDbComment() *DbComment {
	return &DbComment{}
}

type (
	Comment struct {
		Id                int        `json:"id"`
		Blurb             string     `json:"blurb"`
		AddedDateTime     *time.Time `json:"addedDateTime"`
		StoryId           int        `json:"storyId"`
		PoCommentsStoryId int        `json:"poCommentsStoryId"`
	}

	DbComment struct {
		Id                *sql.NullInt32  `json:"id"`
		Blurb             *sql.NullString `json:"blurb"`
		AddedDateTime     *mysql.NullTime `json:"addedDateTime"`
		StoryId           *sql.NullInt32  `json:"storyId"`
		PoCommentsStoryId *sql.NullInt32  `json:"poCommentsStoryId"`
	}
)

func GetCommentsTobeDeletedAndAdded(existingComment, newComment []Comment) (commentsToBeUpdated, commentsToBeAdded, commentsToBeDeleted []Comment) {

	m := make(map[int]bool)
	idCommentMap := make(map[int]Comment)

	for _, item := range existingComment {
		m[item.Id] = true
		idCommentMap[item.Id] = item
	}
	for _, item := range newComment {
		if item.Id == 0 {
			commentsToBeAdded = append(commentsToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			commentsToBeUpdated = append(commentsToBeUpdated, item)
			delete(m, item.Id)
		} else {
			commentsToBeDeleted = append(commentsToBeDeleted, item)
		}

	}

	for k, _ := range m {
		commentsToBeDeleted = append(commentsToBeDeleted, idCommentMap[k])
	}
	return

}

func FromDbComment(dbComment *DbComment) (comment *Comment) {
	comment = &Comment{}
	if dbComment.Id != nil && dbComment.Id.Valid {
		comment.Id = int(dbComment.Id.Int32)
	}
	if dbComment.Blurb != nil && dbComment.Blurb.Valid {
		comment.Blurb = dbComment.Blurb.String
	}
	if dbComment.AddedDateTime != nil && dbComment.AddedDateTime.Valid {
		comment.AddedDateTime = &dbComment.AddedDateTime.Time
	}
	return
}

func ToDbComment(comment *Comment) (commentDb *DbComment) {
	commentDb = &DbComment{}
	commentDb.Id = getNullInt(comment.Id)
	commentDb.Blurb = getNullString(comment.Blurb)
	commentDb.AddedDateTime = getNullTime(comment.AddedDateTime)
	return
}
