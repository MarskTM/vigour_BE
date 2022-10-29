package service

import (
	"vigour/model"
	"vigour/repository"
)

type ProfileService interface {
	GetById(id int) (*model.Profile, error)
	GetAll() ([]model.Profile, error)
	GetByUserId(user_id int) (*model.Profile, error)
	Create(new *model.Profile) (*model.Profile, error)
	Update(id int, profile model.Profile) (*model.Profile, error)
	Delete(id int) error
	Upsert(profile *model.Profile) (*model.Profile, error)
}

type profileService struct {
	profileRepo model.ProfileRepository
}

// --------------------Profile module------------------------

func (s *profileService) GetById(id int) (*model.Profile, error) {
	profile, err := s.profileRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *profileService) GetAll() ([]model.Profile, error) {
	profiles, err := s.profileRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (s *profileService) GetByUserId(user_id int) (*model.Profile, error) {
	profile, err := s.profileRepo.GetByUserId(user_id)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *profileService) Create(new *model.Profile) (*model.Profile, error) {
	profile, err := s.profileRepo.Create(new)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *profileService) Update(id int, profile model.Profile) (*model.Profile, error) {
	record, err := s.profileRepo.Update(id, profile)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *profileService) Delete(id int) error {
	if err := s.profileRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *profileService) Upsert(profile *model.Profile) (*model.Profile, error) {
	return s.profileRepo.Upsert(profile)
}

func NewProfileService() ProfileService {
	profileRepo := repository.NewProfileRepository()
	return &profileService{
		profileRepo: profileRepo,
	}
}
