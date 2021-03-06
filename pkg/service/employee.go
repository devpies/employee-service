package service

import (
	"context"

	"github.com/devpies/employee-service/pkg/model"
	"github.com/devpies/employee-service/pkg/repository"
	"github.com/devpies/employee-service/pkg/trace"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// EmployeeService is responsible for managing employee.
type EmployeeService struct {
	logger       *zap.Logger
	employeeRepo employeeRepository
}

// NewEmployeeService creates a new instance of EmployeeService.
func NewEmployeeService(logger *zap.Logger, employeeRepo employeeRepository) *EmployeeService {
	return &EmployeeService{
		logger:       logger,
		employeeRepo: employeeRepo,
	}
}

// GetEmployeeByID find an employee by id.
func (es *EmployeeService) GetEmployeeByID(ctx context.Context, id string) (model.Employee, error) {
	ctx, span := trace.NewSpan(ctx, "service.employee.GetEmployeeByID", nil)
	defer span.End()

	if _, err := uuid.Parse(id); err != nil {
		return model.Employee{}, repository.ErrInvalidID
	}

	return es.employeeRepo.FindEmployeeByID(ctx, id)
}

// CreateEmployee find an employee by id.
func (es *EmployeeService) CreateEmployee(ctx context.Context, newEmployee model.NewEmployee) error {
	_, span := trace.NewSpan(ctx, "service.employee.GetEmployeeByID", nil)
	defer span.End()

	// perform business logic
	// call repository layer etc.

	return nil
}

// UpdateEmployee find an employee by id.
func (es *EmployeeService) UpdateEmployee(ctx context.Context, employee model.UpdateEmployee) (model.Employee, error) {
	_, span := trace.NewSpan(ctx, "service.employee.GetEmployeeByID", nil)
	defer span.End()

	// perform business logic
	// call repository layer etc.

	return model.Employee{}, nil
}
