package service

import (
	"context"
	"errors"

	"github.com/jmelchio/vetlab/model"
)

// VetOrg implements the api.VetOrgService interface
type VetOrg struct {
	VetOrgRepo VetOrgRepo
}

func (vetOrgService VetOrg) CreateVetOrg(ctx context.Context, vetOrg model.VetOrg) (*model.VetOrg, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	err := vetOrgService.VetOrgRepo.Create(&vetOrg)
	if err != nil {
		return nil, err
	}
	return &vetOrg, nil
}

func (vetOrgService VetOrg) UpdateVetOrg(ctx context.Context, vetOrg model.VetOrg) (*model.VetOrg, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	err := vetOrgService.VetOrgRepo.Update(&vetOrg)
	if err != nil {
		return nil, err
	}
	return &vetOrg, nil
}

func (vetOrgService VetOrg) DeleteVetOrg(ctx context.Context, vetOrg model.VetOrg) error {
	if ctx == nil {
		return errors.New(MissingContext)
	}
	return vetOrgService.VetOrgRepo.Delete(&vetOrg)
}

func (vetOrgService VetOrg) AddUserToVetOrg(ctx context.Context, user model.User, vetOrg model.VetOrg) (*model.User, error) {
	//TODO implement me
	return nil, errors.New("not yet implemented")
}

func (vetOrgService VetOrg) FindVetOrgByName(ctx context.Context, orgName string) ([]model.VetOrg, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	return vetOrgService.VetOrgRepo.GetByName(orgName)
}

func (vetOrgService VetOrg) FindVetOrgByID(ctx context.Context, orgID uint) (*model.VetOrg, error) {
	//TODO implement me
	return nil, errors.New("not yet implemented")
}
