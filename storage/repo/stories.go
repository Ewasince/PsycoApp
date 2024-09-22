package repo

import (
	. "StorageModule/models"
	"time"
)

func GetStories(patientId uint, dateStart, dateFinish time.Time) (*[]Story, error) {
	var stories = make([]Story, 0)
	err := DB.
		Where("patient_id = ?", patientId).
		Where("date >= ?", dateStart).
		Where("date < ?", dateFinish).
		Find(&stories).
		Error

	return &stories, err
}

func GetStoryMinDate(patientId uint) (time.Time, error) {
	var story Story

	err := DB.
		Where("patient_id = ?", patientId).
		Order("date desc").
		First(&story).
		Error

	return story.Date, err
}

func CreateStory(story *Story) error {
	return DB.Create(story).Error
}
