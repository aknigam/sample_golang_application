package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type StoryRepository struct {
}

func (repo *StoryRepository) GetStory(storyId int) (story *models.Story, err error) {
	dbstory := models.MakeDbStory()
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow(" SELECT id, story_title, sprint_id, assignee, sprint_stories_sprint_id"+
		" FROM story"+
		" WHERE (id = ?)", storyId).
		Scan(&dbstory.Id, &dbstory.StoryTitle, &dbstory.SprintId, &dbstory.Assignee, &dbstory.SprintStoriesSprintId)
	if err != nil {
		txn.Rollback()
	}
	return models.FromDbStory(dbstory), err
}

func (repo *StoryRepository) DeleteStory(storyId int) (err error) {
	_, err = Db.Exec(" DELETE FROM story"+
		" WHERE (id =  ?)", storyId)
	return
}

func (repo *StoryRepository) UpdateStory(story *models.Story) (err error) {
	dbstory := models.ToDbStory(story)
	_, err = Db.Exec(" UPDATE story"+
		" SET story_title = ?, sprint_id = ?, assignee = ?, sprint_stories_sprint_id = ?"+
		" WHERE (id = ?)", dbstory.StoryTitle, dbstory.SprintId, dbstory.Assignee, dbstory.SprintStoriesSprintId, dbstory.Id)
	if err != nil {
		common.Error.Println("Story could not be updated ")
		return
	}
	return
}

func (repo *StoryRepository) CreateStory(story *models.Story) (id int, err error) {
	dbstory := models.ToDbStory(story)
	statement := " INSERT INTO story" +
		"  (story_title, sprint_id, assignee, sprint_stories_sprint_id)" +
		" VALUES (?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dbstory.StoryTitle, dbstory.SprintId, dbstory.Assignee, dbstory.SprintStoriesSprintId)

	if err != nil {
		common.Error.Println("Could not create Story ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create Story ", err)
		return
	}
	return int(generatedId), err

}

func (repo *StoryRepository) LinkStoryTasks(storyId int, taskId int) (err error) {
	_, err = Db.Exec(" UPDATE task"+
		" SET story_id= ?"+
		" WHERE (id= ?)", storyId, taskId)
	if err != nil {
		common.Error.Println("Failed to link task with story ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) UnlinkStoryTasks(storyId int, taskId int) (err error) {
	_, err = Db.Exec(" UPDATE task"+
		" SET story_id= null"+
		" WHERE (id= ?)", taskId)
	if err != nil {
		common.Error.Println("Failed to unlink tasks from story ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) UnlinkAllStoryTasks(storyId int) (err error) {
	_, err = Db.Exec(" UPDATE task"+
		" SET story_id= null"+
		" WHERE (story_id= ?)", storyId)
	if err != nil {
		common.Error.Println("Could not delete tasks ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) DeleteAllStoryTasks(storyId int) (err error) {
	_, err = Db.Exec("delete from task where story_id = ?", storyId)
	if err != nil {
		common.Error.Println("tasks could not be deleted ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) LinkStoryComments(storyId int, commentId int) (err error) {
	_, err = Db.Exec(" UPDATE comment"+
		" SET story_id= ?"+
		" WHERE (id= ?)", storyId, commentId)
	if err != nil {
		common.Error.Println("Failed to link comment with story ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) UnlinkStoryComments(storyId int, commentId int) (err error) {
	_, err = Db.Exec(" UPDATE comment"+
		" SET story_id= null"+
		" WHERE (id= ?)", commentId)
	if err != nil {
		common.Error.Println("Failed to unlink comments from story ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) UnlinkAllStoryComments(storyId int) (err error) {
	_, err = Db.Exec(" UPDATE comment"+
		" SET story_id= null"+
		" WHERE (story_id= ?)", storyId)
	if err != nil {
		common.Error.Println("Could not delete comments ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) DeleteAllStoryComments(storyId int) (err error) {
	_, err = Db.Exec("delete from comment where story_id = ?", storyId)
	if err != nil {
		common.Error.Println("comments could not be deleted ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) LinkStoryPoComments(storyId int, poCommentId int) (err error) {
	_, err = Db.Exec(" UPDATE comment"+
		" SET po_comments_story_id= ?"+
		" WHERE (id= ?)", storyId, poCommentId)
	if err != nil {
		common.Error.Println("Failed to link poComment with story ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) UnlinkStoryPoComments(storyId int, poCommentId int) (err error) {
	_, err = Db.Exec(" UPDATE comment"+
		" SET po_comments_story_id= null"+
		" WHERE (id= ?)", poCommentId)
	if err != nil {
		common.Error.Println("Failed to unlink poComments from story ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) UnlinkAllStoryPoComments(storyId int) (err error) {
	_, err = Db.Exec(" UPDATE comment"+
		" SET po_comments_story_id= null"+
		" WHERE (po_comments_story_id= ?)", storyId)
	if err != nil {
		common.Error.Println("Could not delete poComments ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) DeleteAllStoryPoComments(storyId int) (err error) {
	_, err = Db.Exec("delete from comment where po_comments_story_id = ?", storyId)
	if err != nil {
		common.Error.Println("poComments could not be deleted ")
		panic(err)
	}
	return
}
func (repo *StoryRepository) GetAllSprintStories(sprintId int) (storys []models.Story) {
	rows, err := Db.Query(" SELECT id, story_title, sprint_id, assignee, sprint_stories_sprint_id"+
		" FROM story"+
		" WHERE (sprint_id = ?)", sprintId)
	if err != nil {
		common.Error.Println("Could not find storys ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		story := models.Story{}
		err := rows.Scan(&story.Id, &story.StoryTitle, &story.SprintId, &story.Assignee, &story.SprintStoriesSprintId)
		if err != nil {
			common.Error.Println("Could not find storys ", err)
			break
		}
		storys = append(storys, story)
	}
	err = rows.Err()
	if err != nil {
		common.Error.Fatal(err)
	}
	return
}
