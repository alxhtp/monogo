package userrepositoryimplementation

import (
	"context"
	"errors"

	"github.com/alxhtp/monogo/internal/entity"
	entitybase "github.com/alxhtp/monogo/internal/entity/base"
	userrepository "github.com/alxhtp/monogo/internal/repository/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db   *gorm.DB
	user entity.User
}

func NewUserRepository(db *gorm.DB) userrepository.UserRepository {
	return &userRepository{db: db, user: entity.User{}}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) (output *entity.User, err error) {
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	err = r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, user.ID)
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (output *entity.User, err error) {
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}

	if err := r.db.WithContext(ctx).Table(r.user.TableName()).First(&output, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return
}

func (r *userRepository) GetByFilter(ctx context.Context, filter *entity.UserFilter) (output []entity.User, paginationResult entitybase.BasePaginationResult, err error) {
	if r.db == nil {
		return nil, entitybase.BasePaginationResult{}, errors.New("database connection is not initialized")
	}

	query := r.db.WithContext(ctx).Model(&output)
	query, err = r.applyFilter(query, *filter)
	if err != nil {
		return nil, entitybase.BasePaginationResult{}, err
	}

	query = entitybase.PaginateEntityQuery(query, r.user.TableName(), r.user.OrderMap(), &filter.PaginationFilter, &paginationResult)

	if err = query.Find(&output).Error; err != nil {
		return
	}

	return output, paginationResult, nil
}

func (r *userRepository) applyFilter(db *gorm.DB, filter entity.UserFilter) (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("database connection is not initialized")
	}

	table := r.user.TableName()
	if filter.IDs != nil {
		db = db.Where(table+".id IN (?)", filter.IDs)
	}

	if filter.Name != nil {
		// ignore case
		db = db.Where(table+".name ILIKE ?", "%"+*filter.Name+"%")
	}

	if filter.Email != nil {
		db = db.Where(table+".email = ?", filter.Email)
	}

	if filter.Status != nil {
		db = db.Where(table+".status = ?", filter.Status)
	}

	if filter.Sex != nil {
		db = db.Where(table+".metadata->>'sex' = ?", filter.Sex)
	}

	if filter.Address != nil {
		db = db.Where(table+".metadata->>'address' ILIKE ?", "%"+*filter.Address+"%")
	}

	if filter.Phone != nil {
		db = db.Where(table+".metadata->>'phone' = ?", filter.Phone)
	}

	return db, nil
}

func (r *userRepository) Update(ctx context.Context, id uuid.UUID, updateMap map[string]any) (output *entity.User, err error) {
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}

	err = r.db.WithContext(ctx).Model(&r.user).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, id)
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	if r.db == nil {
		return errors.New("database connection is not initialized")
	}

	err = r.db.WithContext(ctx).Where("id = ?", id).Delete(&r.user).Error
	return err
}
