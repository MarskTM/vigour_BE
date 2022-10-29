package repository

import (
	"vigour/infrastructure"
	"vigour/model"
	"time"
	"gorm.io/gorm/clause"
)

type profileRepository struct{}

func (r *profileRepository) GetById(id int) (*model.Profile, error) {
	db := infrastructure.GetDB()

	var record model.Profile
	if err := db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *profileRepository) GetAll() ([]model.Profile, error) {
	db := infrastructure.GetDB()

	var records []model.Profile
	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (r *profileRepository) GetByUserId(user_id int) (*model.Profile, error) {
	db := infrastructure.GetDB()

	var record model.Profile
	if err := db.Where("user_id = ?", user_id).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *profileRepository) Create(new *model.Profile) (*model.Profile, error) {
	db := infrastructure.GetDB()

	newRecod := new
	if err := db.Create(newRecod).Error; err != nil {
		return nil, err
	}
	return newRecod, nil
}

func (r *profileRepository) Update(id int, profile model.Profile) (*model.Profile, error) {
	db := infrastructure.GetDB()

	if err := db.Where("id = ?", id).Updates(profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepository) Delete(id int) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Profile{Id: id}).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (r *profileRepository) Upsert(profile *model.Profile) (*model.Profile, error) {
	db := infrastructure.GetDB()

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func NewProfileRepository() model.ProfileRepository {
	return &profileRepository{}
}
