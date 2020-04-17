package service

import (
	"context"
	"database/sql"
	"sample_golang_application/errors"
	"sample_golang_application/models"
	"sample_golang_application/repository"
	"time"
)

type ISprintService interface {
	GetSprint(int) (sprint *models.Sprint, appErr *errors.AppError)
	DeleteSprint(int) (appErr *errors.AppError)
	UpdateSprint(*models.Sprint) (appErr *errors.AppError)
	CreateSprint(sprint *models.Sprint) (sprintId int, appErr *errors.AppError)
	CreateSprintStories(sprintId int, story *models.Story) (storyId int, appErr *errors.AppError)
	UpdateSprintStories(sprintId int, story *models.Story) (appErr *errors.AppError)
	DeleteSprintStories(sprintId, storyId int) (appErr *errors.AppError)
	GetSprintStories(sprintId, storyId int) (stories *models.Story, appErr *errors.AppError)
	UnlinkSprintStories(sprintId int, storiesId int) (appError *errors.AppError)
	LinkSprintStories(sprintId int, storiesId int) (appError *errors.AppError)
}

type SprintServiceTxn struct {
	SprintService ISprintService
}

func (service *SprintServiceTxn) CreateSprintStories(sprintId int, story *models.Story) (storyId int, appErr *errors.AppError) {
	panic("implement me")
}

func (service *SprintServiceTxn) UpdateSprintStories(sprintId int, story *models.Story) (appErr *errors.AppError) {
	panic("implement me")
}

func (service *SprintServiceTxn) DeleteSprintStories(sprintId, storyId int) (appErr *errors.AppError) {
	panic("implement me")
}

func (service *SprintServiceTxn) GetSprintStories(sprintId, storyId int) (stories *models.Story, appErr *errors.AppError) {
	panic("implement me")
}

func (service *SprintServiceTxn) UnlinkSprintStories(sprintId int, storiesId int) (appError *errors.AppError) {
	panic("implement me")
}

func (service *SprintServiceTxn) LinkSprintStories(sprintId int, storiesId int) (appError *errors.AppError) {
	panic("implement me")
}
var (
	ctx context.Context
	db  *sql.DB
)

// refer: https://golang.org/src/database/sql/example_test.go
func (service *SprintServiceTxn) UpdateSprint(sprint *models.Sprint) (appErr *errors.AppError) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	opts := &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	}
	txn, err := repository.Db.BeginTx(ctx, opts)
	if err != nil {
		return
	}
	defer txn.Commit()
	appErr = service.SprintService.UpdateSprint(sprint)
	if appErr != nil {
		txn.Rollback()
		return
	}


	return
}

func (service *SprintServiceTxn) CreateSprint(sprint *models.Sprint) (sprintId int, appErr *errors.AppError) {
	txn, err := repository.Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	appErr = service.SprintService.UpdateSprint(sprint)
	if appErr != nil {
		txn.Rollback()
		return
	}
	return
}

func (service *SprintServiceTxn) DeleteSprint(sprintId int) ( appErr *errors.AppError) {
	txn, err := repository.Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	appErr = service.SprintService.DeleteSprint(sprintId)
	if appErr != nil {
		txn.Rollback()
		return
	}
	return
}

func (service *SprintServiceTxn) GetSprint(sprintId int) ( sprint *models.Sprint, appErr *errors.AppError) {
	txn, err := repository.Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	_, appErr = service.SprintService.GetSprint(sprintId)
	if appErr != nil {
		txn.Rollback()
		return
	}
	return
}



type SprintService struct {
	SprintRepo  repository.SprintRepository
	TaskRepo    repository.TaskRepository
	PersonRepo  repository.PersonRepository
	CommentRepo repository.CommentRepository
	StoryRepo   repository.StoryRepository
}

func (service *SprintService) GetSprint(sprintId int) (sprint *models.Sprint, appErr *errors.AppError) {
	sprint, err := service.SprintRepo.GetSprint(sprintId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "sprint not found", Code: -1}
		return
	}
	if sprint.Owner != nil {
		owner, _ := service.GetSprintOwner(sprintId, sprint.Owner.Id)
		sprint.Owner = owner
	}
	return
}

// ------------------ start get -----------------
func (service *SprintService) GetSprintOwner(sprintId, personId int) (owner *models.Person, appErr *errors.AppError) {
	owner, err := service.PersonRepo.GetPerson(personId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "sprint owner not found", Code: -1}
		return
	}
	return
}

// ------------------ end get-----------------

func (service *SprintService) DeleteSprint(sprintId int) (appErr *errors.AppError) {
	sprint, appErr := service.GetSprint(sprintId)
	if appErr != nil {
		return appErr
	}
	service.SprintRepo.UnlinkAllSprintStories(sprintId)

	service.SprintRepo.DeleteSprint(sprintId)

	service.DeleteSprintOwner(sprintId, sprint.Owner.Id)

	return
}
func (service *SprintService) DeleteSprintOwner(sprintId, personId int) (appErr *errors.AppError) {
	_, appErr = service.GetSprintOwner(sprintId, personId)
	if appErr != nil {
		return appErr
	}
	service.PersonRepo.DeletePerson(personId)

	return
}



func (service *SprintService) UpdateSprint(sprint *models.Sprint) (appErr *errors.AppError) {
	//existingSprint, err := service.GetSprint(sprint.Id)
	//if err != nil {
	//	return err
	//}
	//service.upsertSprintOwner(sprint.Id, existingSprint.Owner, sprint.Owner)
	sprintUpdateErr := service.SprintRepo.UpdateSprint(sprint)
	if sprintUpdateErr != nil {
		appErr = &errors.AppError{Error: sprintUpdateErr, Message: "sprint could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

func (service *SprintService) upsertSprintOwner(sprintId int, existingOwner *models.Person, newOwner *models.Person) {
	//var zeroPerson = &models.Person{}

	if newOwner != nil {
		if existingOwner == nil {
			service.CreateSprintOwner(sprintId, newOwner)
		} else {
			service.UpdateSprintOwner(sprintId, newOwner)
		}
	} else if existingOwner != nil {
		service.DeleteSprintOwner(sprintId, existingOwner.Id)
	}
}

// ------------------ start update -----------------
func (service *SprintService) UpdateSprintOwner(sprintId int, person *models.Person) (appErr *errors.AppError) {
	_, err := service.GetSprintOwner(sprintId, person.Id)
	if err != nil {
		return err
	}
	personUpdateErr := service.PersonRepo.UpdatePerson(person)
	if personUpdateErr != nil {
		appErr = &errors.AppError{Error: personUpdateErr, Message: "person could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

// ------------------ end update-----------------

func (service *SprintService) CreateSprint(sprint *models.Sprint) (sprintId int, appErr *errors.AppError) {
	sprintId, err := service.SprintRepo.CreateSprint(sprint)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "sprint could not be created", Code: -1}
	}
	var emptyOwner = &models.Person{}
	if sprint.Owner != emptyOwner {
		service.CreateSprintOwner(sprintId, sprint.Owner)
		service.SprintRepo.LinkSprintOwner(sprint.Id, sprint.Owner.Id)
	}
	return sprintId, appErr
}

func (service *SprintService) CreateSprintOwner(sprintId int, person *models.Person) (personId int, appErr *errors.AppError) {
	personId, err := service.PersonRepo.CreatePerson(person)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "sprint could not be created", Code: -1}
	}
	return sprintId, appErr
}

func (service *SprintService) GetSprintStories(sprintId, storyId int) (stories *models.Story, appErr *errors.AppError) {
	stories, err := service.StoryRepo.GetStory(storyId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "sprint not found", Code: -1}
		return
	}
	if stories.SprintId != sprintId {
		return nil, nil
	}
	return
}

func (service *SprintService) DeleteSprintStories(sprintId, storyId int) (appErr *errors.AppError) {
	_, appErr = service.GetSprintStories(sprintId, storyId)
	if appErr != nil {
		return appErr
	}
	service.StoryRepo.UnlinkAllStoryTasks(storyId)

	service.StoryRepo.UnlinkAllStoryComments(storyId)

	service.StoryRepo.UnlinkAllStoryPoComments(storyId)

	service.StoryRepo.DeleteStory(storyId)

	return
}

func (service *SprintService) UpdateSprintStories(sprintId int, story *models.Story) (appErr *errors.AppError) {
	_, err := service.GetSprintStories(sprintId, story.Id)
	if err != nil {
		return err
	}
	storyUpdateErr := service.StoryRepo.UpdateStory(story)
	if storyUpdateErr != nil {
		appErr = &errors.AppError{Error: storyUpdateErr, Message: "story could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

func (service *SprintService) CreateSprintStories(sprintId int, story *models.Story) (storyId int, appErr *errors.AppError) {
	story.SprintId = sprintId
	storyId, err := service.StoryRepo.CreateStory(story)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "sprint could not be created", Code: -1}
	}
	return sprintId, appErr
}

func (service *SprintService) LinkSprintStories(sprintId int, storiesId int) (appError *errors.AppError) {
	err := service.SprintRepo.LinkSprintStories(sprintId, storiesId)
	if err != nil {
		return &errors.AppError{Error: err, Message: "Sprint could not be linked with stories", Code: -1}
	}
	return
}

func (service *SprintService) UnlinkSprintStories(sprintId int, storiesId int) (appError *errors.AppError) {
	err := service.SprintRepo.UnlinkSprintStories(sprintId, storiesId)
	if err != nil {
		return &errors.AppError{Error: err, Message: "Sprint could not be linked with stories", Code: -1}
	}
	return
}

//    ONE TO MANY END
