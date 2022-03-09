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
	//TODO implement me
	return nil, errors.New("not yet implemented")
}

func (vetOrgService VetOrg) UpdateVetOrg(ctx context.Context, vetOrg model.VetOrg) (*model.VetOrg, error) {
	//TODO implement me
	return nil, errors.New("not yet implemented")
}

func (vetOrgService VetOrg) DeteleVetOrg(ctx context.Context, vetOrg model.VetOrg) error {
	//TODO implement me
	return errors.New("not yet implemented")
}

func (vetOrgService VetOrg) AddUserToVetOrg(ctx context.Context, user model.User, vetOrg model.VetOrg) (*model.User, error) {
	//TODO implement me
	return nil, errors.New("not yet implemented")
}

func (vetOrgService VetOrg) FindVetOrgByName(ctx context.Context, orgName string) (*model.VetOrg, error) {
	//TODO implement me
	return nil, errors.New("not yet implemented")
}

func (vetOrgService VetOrg) FindVetOrgByID(ctx context.Context, orgID uint) (*model.VetOrg, error) {
	//TODO implement me
	return nil, errors.New("not yet implemented")
}
