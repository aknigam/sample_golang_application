package service

import (
	"sample_golang_application/errors"
	"sample_golang_application/models"
	"sample_golang_application/repository"
)

type StoryService struct {
	StoryRepo   repository.StoryRepository
	TaskRepo    repository.TaskRepository
	CommentRepo repository.CommentRepository
}

func (service *StoryService) GetStoryTasks(storyId, taskId int) (tasks *models.Task, appErr *errors.AppError) {
	tasks, err := service.TaskRepo.GetTask(taskId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "story not found", Code: -1}
		return
	}
	if tasks.StoryId != storyId {
		return nil, nil
	}
	return
}

func (service *StoryService) DeleteStoryTasks(storyId, taskId int) (appErr *errors.AppError) {
	_, appErr = service.GetStoryTasks(storyId, taskId)
	if appErr != nil {
		return appErr
	}
	service.TaskRepo.DeleteTask(taskId)

	return
}

func (service *StoryService) UpdateStoryTasks(storyId int, task *models.Task) (appErr *errors.AppError) {
	_, err := service.GetStoryTasks(storyId, task.Id)
	if err != nil {
		return err
	}
	taskUpdateErr := service.TaskRepo.UpdateTask(task)
	if taskUpdateErr != nil {
		appErr = &errors.AppError{Error: taskUpdateErr, Message: "task could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

func (service *StoryService) CreateStoryTasks(storyId int, task *models.Task) (taskId int, appErr *errors.AppError) {
	task.StoryId = storyId
	taskId, err := service.TaskRepo.CreateTask(task)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "story could not be created", Code: -1}
	}
	return storyId, appErr
}

func (service *StoryService) LinkStoryTasks(storyId int, tasksId int) (appError *errors.AppError) {
	err := service.StoryRepo.LinkStoryTasks(storyId, tasksId)
	if err != nil {
		return &errors.AppError{Error: err, Message: "Story could not be linked with tasks", Code: -1}
	}
	return
}

func (service *StoryService) UnlinkStoryTasks(storyId int, tasksId int) (appError *errors.AppError) {
	err := service.StoryRepo.UnlinkStoryTasks(storyId, tasksId)
	if err != nil {
		return &errors.AppError{Error: err, Message: "Story could not be linked with tasks", Code: -1}
	}
	return
}

//    ONE TO MANY END

func (service *StoryService) GetStoryComments(storyId, commentId int) (comments *models.Comment, appErr *errors.AppError) {
	comments, err := service.CommentRepo.GetComment(commentId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "story not found", Code: -1}
		return
	}
	if comments.StoryId != storyId {
		return nil, nil
	}
	return
}

func (service *StoryService) DeleteStoryComments(storyId, commentId int) (appErr *errors.AppError) {
	_, appErr = service.GetStoryComments(storyId, commentId)
	if appErr != nil {
		return appErr
	}
	service.CommentRepo.DeleteComment(commentId)

	return
}

func (service *StoryService) UpdateStoryComments(storyId int, comment *models.Comment) (appErr *errors.AppError) {
	_, err := service.GetStoryComments(storyId, comment.Id)
	if err != nil {
		return err
	}
	commentUpdateErr := service.CommentRepo.UpdateComment(comment)
	if commentUpdateErr != nil {
		appErr = &errors.AppError{Error: commentUpdateErr, Message: "comment could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

func (service *StoryService) CreateStoryComments(storyId int, comment *models.Comment) (commentId int, appErr *errors.AppError) {
	comment.StoryId = storyId
	commentId, err := service.CommentRepo.CreateComment(comment)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "story could not be created", Code: -1}
	}
	return storyId, appErr
}

func (service *StoryService) LinkStoryComments(storyId int, commentsId int) (appError *errors.AppError) {
	err := service.StoryRepo.LinkStoryComments(storyId, commentsId)
	if err != nil {
		return &errors.AppError{Error: err, Message: "Story could not be linked with comments", Code: -1}
	}
	return
}

func (service *StoryService) UnlinkStoryComments(storyId int, commentsId int) (appError *errors.AppError) {
	err := service.StoryRepo.UnlinkStoryComments(storyId, commentsId)
	if err != nil {
		return &errors.AppError{Error: err, Message: "Story could not be linked with comments", Code: -1}
	}
	return
}

//    ONE TO MANY END

func (service *StoryService) GetStoryPoComments(storyId, commentId int) (poComments *models.Comment, appErr *errors.AppError) {
	poComments, err := service.CommentRepo.GetComment(commentId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "story not found", Code: -1}
		return
	}
	if poComments.PoCommentsStoryId != storyId {
		return nil, nil
	}
	return
}

func (service *StoryService) DeleteStoryPoComments(storyId, commentId int) (appErr *errors.AppError) {
	_, appErr = service.GetStoryPoComments(storyId, commentId)
	if appErr != nil {
		return appErr
	}
	service.CommentRepo.DeleteComment(commentId)

	return
}

func (service *StoryService) UpdateStoryPoComments(storyId int, comment *models.Comment) (appErr *errors.AppError) {
	_, err := service.GetStoryPoComments(storyId, comment.Id)
	if err != nil {
		return err
	}
	commentUpdateErr := service.CommentRepo.UpdateComment(comment)
	if commentUpdateErr != nil {
		appErr = &errors.AppError{Error: commentUpdateErr, Message: "comment could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

func (service *StoryService) CreateStoryPoComments(storyId int, comment *models.Comment) (commentId int, appErr *errors.AppError) {
	comment.PoCommentsStoryId = storyId
	commentId, err := service.CommentRepo.CreateComment(comment)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "story could not be created", Code: -1}
	}
	return storyId, appErr
}

func (service *StoryService) LinkStoryPoComments(storyId int, poCommentsId int) (appError *errors.AppError) {
	err := service.StoryRepo.LinkStoryPoComments(storyId, poCommentsId)
	if err != nil {
		return &errors.AppError{Error: err, Message: "Story could not be linked with poComments", Code: -1}
	}
	return
}

func (service *StoryService) UnlinkStoryPoComments(storyId int, poCommentsId int) (appError *errors.AppError) {
	err := service.StoryRepo.UnlinkStoryPoComments(storyId, poCommentsId)
	if err != nil {
		return &errors.AppError{Error: err, Message: "Story could not be linked with poComments", Code: -1}
	}
	return
}

//    ONE TO MANY END
