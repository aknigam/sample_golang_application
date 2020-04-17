package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type CommentRepository struct {
}

func (repo *CommentRepository) GetComment(commentId int) (comment *models.Comment, err error) {
	dbcomment := models.MakeDbComment()
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow(" SELECT id, blurb, added_date_time, story_id, po_comments_story_id"+
		" FROM comment"+
		" WHERE (id = ?)", commentId).
		Scan(&dbcomment.Id, &dbcomment.Blurb, &dbcomment.AddedDateTime, &dbcomment.StoryId, &dbcomment.PoCommentsStoryId)
	if err != nil {
		txn.Rollback()
	}
	return models.FromDbComment(dbcomment), err
}

func (repo *CommentRepository) DeleteComment(commentId int) (err error) {
	_, err = Db.Exec(" DELETE FROM comment"+
		" WHERE (id =  ?)", commentId)
	return
}

func (repo *CommentRepository) UpdateComment(comment *models.Comment) (err error) {
	dbcomment := models.ToDbComment(comment)
	_, err = Db.Exec(" UPDATE comment"+
		" SET blurb = ?, added_date_time = ?, story_id = ?, po_comments_story_id = ?"+
		" WHERE (id = ?)", dbcomment.Blurb, dbcomment.AddedDateTime, dbcomment.StoryId, dbcomment.PoCommentsStoryId, dbcomment.Id)
	if err != nil {
		common.Error.Println("Comment could not be updated ")
		return
	}
	return
}

func (repo *CommentRepository) CreateComment(comment *models.Comment) (id int, err error) {
	dbcomment := models.ToDbComment(comment)
	statement := " INSERT INTO comment" +
		"  (blurb, added_date_time, story_id, po_comments_story_id)" +
		" VALUES (?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dbcomment.Blurb, dbcomment.AddedDateTime, dbcomment.StoryId, dbcomment.PoCommentsStoryId)

	if err != nil {
		common.Error.Println("Could not create Comment ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create Comment ", err)
		return
	}
	return int(generatedId), err

}

func (repo *CommentRepository) GetAllStoryComments(storyId int) (comments []models.Comment) {
	rows, err := Db.Query(" SELECT id, blurb, added_date_time, story_id, po_comments_story_id"+
		" FROM comment"+
		" WHERE (story_id = ?)", storyId)
	if err != nil {
		common.Error.Println("Could not find comments ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(&comment.Id, &comment.Blurb, &comment.AddedDateTime, &comment.StoryId, &comment.PoCommentsStoryId)
		if err != nil {
			common.Error.Println("Could not find comments ", err)
			break
		}
		comments = append(comments, comment)
	}
	err = rows.Err()
	if err != nil {
		common.Error.Fatal(err)
	}
	return
}
func (repo *CommentRepository) GetAllStoryPoComments(storyId int) (comments []models.Comment) {
	rows, err := Db.Query(" SELECT id, blurb, added_date_time, story_id, po_comments_story_id"+
		" FROM comment"+
		" WHERE (po_comments_story_id = ?)", storyId)
	if err != nil {
		common.Error.Println("Could not find comments ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(&comment.Id, &comment.Blurb, &comment.AddedDateTime, &comment.StoryId, &comment.PoCommentsStoryId)
		if err != nil {
			common.Error.Println("Could not find comments ", err)
			break
		}
		comments = append(comments, comment)
	}
	err = rows.Err()
	if err != nil {
		common.Error.Fatal(err)
	}
	return
}
