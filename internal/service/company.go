package service

import (
	"context"
	"fmt"
	"server/internal/entities"
	"server/internal/repository"
	"time"
)

type CompanyService struct {
	companyRepository       repository.CompanyRepository
	companyMemberRepository repository.CompanyMemberRepository
	timeout                 time.Duration
}

func NewCompanyService(companyRepository repository.CompanyRepository, companyMemberRepository repository.CompanyMemberRepository) *CompanyService {
	return &CompanyService{
		companyRepository:       companyRepository,
		companyMemberRepository: companyMemberRepository,
		timeout:                 time.Duration(10) * time.Second,
	}
}

func (s *CompanyService) CreateCompany(c context.Context, request *entities.CreateCompanyRequest) (*entities.CreateCompanyResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	comp := &entities.Company{
		Name: request.Name,
	}

	company, err := s.companyRepository.Create(ctx, comp)
	if err != nil {
		return nil, fmt.Errorf("failed to create company: %w", err)
	}

	createdCompany := &entities.CreateCompanyResponse{
		ID: company.ID,
	}
	return createdCompany, nil
}

func (s *CompanyService) GetAll(ctx context.Context, request *entities.GetCompanyMembersInfoRequest) ([]entities.CompanyMemberInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	members, err := s.companyMemberRepository.GetAll(ctx, request.CompanyID)
	if err != nil {
		return nil, fmt.Errorf("failed to get members: %w", err)
	}
	return members, nil

}
