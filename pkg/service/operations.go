package service

import (
	"coinkeeper/models"
	"coinkeeper/pkg/repository"
)

func GetAllOperations(userID uint) (operations []models.Operation, err error) {
	operations, err = repository.GetAllOperations(userID)
	if err != nil {
		return nil, err
	}

	return operations, nil
}

func GetOperationByID(userID, operationID uint) (o models.Operation, err error) {
	o, err = repository.GetOperationByID(userID, operationID)
	if err != nil {
		return models.Operation{}, err
	}

	return o, nil
}
