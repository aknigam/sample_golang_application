package models

import "database/sql"

func MakeDbStory() *DbStory {
	return &DbStory{}
}

type (
	Story struct {
		Id                    int       `json:"id"`
		StoryTitle            string    `json:"storyTitle"`
		SprintId              int       `json:"sprintId"`
		Tasks                 []Task    `json:"tasks"`
		Assignee              int       `json:"assignee"`
		Comments              []Comment `json:"comments"`
		PoComments            []Comment `json:"poComments"`
		SprintStoriesSprintId int       `json:"sprintStoriesSprintId"`
	}

	DbStory struct {
		Id                    *sql.NullInt32  `json:"id"`
		StoryTitle            *sql.NullString `json:"storyTitle"`
		SprintId              *sql.NullInt32  `json:"sprintId"`
		Tasks                 []DbTask        `json:"tasks"`
		Assignee              *sql.NullInt32  `json:"assignee"`
		Comments              []DbComment     `json:"comments"`
		PoComments            []DbComment     `json:"poComments"`
		SprintStoriesSprintId *sql.NullInt32  `json:"sprintStoriesSprintId"`
	}
)

func GetStorysTobeDeletedAndAdded(existingStory, newStory []Story) (storysToBeUpdated, storysToBeAdded, storysToBeDeleted []Story) {

	m := make(map[int]bool)
	idStoryMap := make(map[int]Story)

	for _, item := range existingStory {
		m[item.Id] = true
		idStoryMap[item.Id] = item
	}
	for _, item := range newStory {
		if item.Id == 0 {
			storysToBeAdded = append(storysToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			storysToBeUpdated = append(storysToBeUpdated, item)
			delete(m, item.Id)
		} else {
			storysToBeDeleted = append(storysToBeDeleted, item)
		}

	}

	for k, _ := range m {
		storysToBeDeleted = append(storysToBeDeleted, idStoryMap[k])
	}
	return

}

func FromDbStory(dbStory *DbStory) (story *Story) {
	story = &Story{}
	if dbStory.Id != nil && dbStory.Id.Valid {
		story.Id = int(dbStory.Id.Int32)
	}
	if dbStory.StoryTitle != nil && dbStory.StoryTitle.Valid {
		story.StoryTitle = dbStory.StoryTitle.String
	}
	return
}

func ToDbStory(story *Story) (storyDb *DbStory) {
	storyDb = &DbStory{}
	storyDb.Id = getNullInt(story.Id)
	storyDb.StoryTitle = getNullString(story.StoryTitle)
	return
}
