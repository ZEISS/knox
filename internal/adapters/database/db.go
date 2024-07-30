package database

import (
	"context"
	"errors"

	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/knox/internal/models"
	"github.com/zeiss/knox/internal/ports"

	openfga "github.com/openfga/go-sdk/client"
	fga "github.com/zeiss/fiber-authz/openfga"
	seed "github.com/zeiss/gorm-seed"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ ports.ReadTx = (*readTxImpl)(nil)

type readTxImpl struct {
	conn *gorm.DB
}

// NewReadTx ...
func NewReadTx() seed.ReadTxFactory[ports.ReadTx] {
	return func(db *gorm.DB) (ports.ReadTx, error) {
		return &readTxImpl{conn: db}, nil
	}
}

// GetLock ...
func (r *readTxImpl) GetLock(ctx context.Context, lock *models.Lock) error {
	return r.conn.Where(lock).First(lock).Error
}

// GetProject ...
func (r *readTxImpl) GetProject(ctx context.Context, project *models.Project) error {
	return r.conn.Where(project).First(project).Error
}

// GetEnvironment ...
func (r *readTxImpl) GetEnvironment(ctx context.Context, teamName string, projectName string, environment *models.Environment) error {
	return r.conn.
		Where("project_id = (?)", r.conn.Model(&models.Project{}).Where("name = ?", projectName).Where("owner_id = (?)", r.conn.Model(&models.Team{}).Where("name = ?", teamName).Select("id")).Select("id")).
		Where(environment).
		First(environment).Error
}

// GetTeam ...
func (r *readTxImpl) GetTeam(ctx context.Context, team *models.Team) error {
	return r.conn.Where(team).First(team).Error
}

// GetState ...
func (r *readTxImpl) GetState(ctx context.Context, state *models.State) error {
	return r.conn.Where(state).Last(state).Error
}

// ListProjects ...
func (r *readTxImpl) ListProjects(ctx context.Context, teamName string, results *tables.Results[models.Project]) error {
	return r.conn.Scopes(tables.PaginatedResults(&results.Rows, results, r.conn)).
		Where("owner_id = (?)", r.conn.Model(&models.Team{}).Where("name = ?", teamName).Select("id")).
		Find(&results.Rows).Error
}

// ListEnvironments ...
func (r *readTxImpl) ListEnvironments(ctx context.Context, teamName, projetName string, results *tables.Results[models.Environment]) error {
	return r.conn.Scopes(tables.PaginatedResults(&results.Rows, results, r.conn)).
		Where("project_id = (?)", r.conn.Model(&models.Project{}).Where("name = ?", projetName).Where("owner_id = (?)", r.conn.Model(&models.Team{}).Where("name = ?", teamName).Select("id")).Select("id")).
		Find(&results.Rows).Error
}

// ListTeams ...
func (r *readTxImpl) ListTeams(ctx context.Context, results *tables.Results[models.Team]) error {
	return r.conn.Scopes(tables.PaginatedResults(&results.Rows, results, r.conn)).Find(&results.Rows).Error
}

// AuthenticateClient ...
func (r *readTxImpl) AuthenticateClient(ctx context.Context, teamId, projectId, environmentId, username, password string) error {
	environment := models.Environment{
		Name:     environmentId,
		Username: username,
	}

	err := r.conn.Debug().
		Model(&models.Environment{}).
		Where("project_id = (?)", r.conn.Model(&models.Project{}).Where("name = ?", projectId).Where("owner_id = (?)", r.conn.Model(&models.Team{}).Where("name = ?", teamId).Select("id")).Select("id")).
		Where(&environment).
		First(&environment).Error
	if err != nil {
		return err
	}

	return environment.ComparePassword(password)
}

// ListStates ...
func (r *readTxImpl) ListStates(ctx context.Context, teamName, projectName, environmentName string, results *tables.Results[models.State]) error {
	return r.conn.Scopes(tables.PaginatedResults(&results.Rows, results, r.conn)).
		Where("environment_id = (?)", r.conn.Model(&models.Environment{}).Where("name = ?", environmentName).Where("project_id = (?)", r.conn.Model(&models.Project{}).Where("name = ?", projectName).Where("owner_id = (?)", r.conn.Model(&models.Team{}).Where("name = ?", teamName).Select("id")).Select("id")).Select("id")).
		Find(&results.Rows).Error
}

type writeTxImpl struct {
	conn *gorm.DB
	fga  *openfga.OpenFgaClient
	readTxImpl
}

// NewWriteTx ...
func NewWriteTx(fga *openfga.OpenFgaClient) seed.ReadWriteTxFactory[ports.ReadWriteTx] {
	return func(db *gorm.DB) (ports.ReadWriteTx, error) {
		return &writeTxImpl{conn: db, fga: fga}, nil
	}
}

// CreateLock creates a new lock.
func (rw *writeTxImpl) CreateLock(ctx context.Context, lock *models.Lock) error {
	return rw.conn.Create(lock).Error
}

// DeleteLock deletes a lock.
func (rw *writeTxImpl) DeleteLock(ctx context.Context, lock *models.Lock) error {
	return rw.conn.Delete(lock).Error
}

// CreateProject creates a new project.
func (rw *writeTxImpl) CreateProject(ctx context.Context, project *models.Project) error {
	return rw.conn.Create(project).Error
}

// DeleteProject deletes a project.
func (rw *writeTxImpl) DeleteProject(ctx context.Context, teamName string, project *models.Project) error {
	return rw.conn.
		Where("owner_id = (?)", rw.conn.Model(&models.Team{}).Where("name = ?", teamName).Select("id")).
		Delete(project).
		Error
}

// UpdateState...
func (rw *writeTxImpl) UpdateState(ctx context.Context, state *models.State) error {
	latest := &models.State{}

	result := rw.conn.Debug().
		Where("project_id = (?)", state.ProjectID).
		Where("environment_id = (?)", state.EnvironmentID).
		Last(latest)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	if latest.Version > 0 {
		state.Version = latest.Version + 1
	}

	if latest.Version > 0 {
		err := rw.conn.Delete(&latest).Error
		if err != nil {
			return err
		}
	}

	return rw.conn.Select("*").Create(&state).Error
}

// CreateSnapshot creates a new snapshot.
func (rw *writeTxImpl) CreateSnapshot(ctx context.Context, snapshot *models.Snapshot) error {
	return rw.conn.Create(snapshot).Error
}

// GetTeam ...
func (rw *writeTxImpl) GetTeam(ctx context.Context, team *models.Team) error {
	return rw.conn.Where(team).First(team).Error
}

// CreateTeam creates a new team.
func (rw *writeTxImpl) CreateTeam(ctx context.Context, team *models.Team) error {
	return rw.conn.Create(team).Error
}

// DeleteTeam deletes a team.
func (rw *writeTxImpl) DeleteTeam(ctx context.Context, team *models.Team) error {
	return rw.conn.Delete(team).Error
}

// CreateEnvironment creates a new environment.
func (rw *writeTxImpl) CreateEnvironment(ctx context.Context, environment *models.Environment) error {
	err := rw.conn.Create(environment).Error
	if err != nil {
		return err
	}

	err = rw.conn.Preload(clause.Associations).Preload("Project.Owner").Where(environment).First(environment).Error
	if err != nil {
		return err
	}

	body := openfga.ClientWriteRequest{
		Writes: []openfga.ClientTupleKey{
			{
				User:     fga.NewUser(fga.Namespace("project"), fga.Join(fga.DefaultSeparator, environment.Project.Owner.Name, environment.Project.Name)).String(),
				Relation: fga.NewRelation(fga.String("owner")).String(),
				Object:   fga.NewObject(fga.Namespace("environment"), fga.Join(fga.DefaultSeparator, environment.Project.Owner.Name, environment.Project.Name, environment.Name)).String(),
			},
		},
	}

	_, err = rw.fga.Write(ctx).Body(body).Execute()
	if err != nil {
		return err
	}

	return nil
}

// DeleteEnvironment deletes an environment.
func (rw *writeTxImpl) DeleteEnvironment(ctx context.Context, environment *models.Environment) error {
	err := rw.conn.Preload(clause.Associations).Preload("Project.Owner").Where(environment).First(environment).Error
	if err != nil {
		return err
	}

	err = rw.conn.Delete(environment).Error
	if err != nil {
		return err
	}

	body := openfga.ClientWriteRequest{
		Deletes: []openfga.ClientTupleKeyWithoutCondition{
			{
				User:     fga.NewUser(fga.Namespace("project"), fga.Join(fga.DefaultSeparator, environment.Project.Owner.Name, environment.Project.Name)).String(),
				Relation: fga.NewRelation(fga.String("owner")).String(),
				Object:   fga.NewObject(fga.Namespace("environment"), fga.Join(fga.DefaultSeparator, environment.Project.Owner.Name, environment.Project.Name, environment.Name)).String(),
			},
		},
	}

	_, err = rw.fga.Write(ctx).Body(body).Execute()
	if err != nil {
		return err
	}

	return nil
}
