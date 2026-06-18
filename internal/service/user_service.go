package service

import (
	"context"
	"time"

	"github.com/ksploitx/user-age-api/internal/models"
	"github.com/ksploitx/user-age-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func (s *UserService) Create(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.Create(ctx, req.Name, dob)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *UserService) GetByID(ctx context.Context, id int32) (models.UserDetailResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.UserDetailResponse{}, err
	}

	return models.UserDetailResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}, nil
}

func (s *UserService) Update(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.Update(ctx, id, req.Name, dob)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *UserService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) List(ctx context.Context) ([]models.UserDetailResponse, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var result []models.UserDetailResponse
	for _, u := range users {
		result = append(result, models.UserDetailResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.Dob.Format("2006-01-02"),
			Age:  calculateAge(u.Dob),
		})
	}
	return result, nil
}
