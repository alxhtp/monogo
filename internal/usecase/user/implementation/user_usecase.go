package userusecaseimplementation

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	entitybase "github.com/alxhtp/monogo/internal/entity/base"
	userrepository "github.com/alxhtp/monogo/internal/repository/user"
	userserializer "github.com/alxhtp/monogo/internal/serializer/user"
	userusecase "github.com/alxhtp/monogo/internal/usecase/user"
	"github.com/alxhtp/monogo/pkg/dto"
	dtobase "github.com/alxhtp/monogo/pkg/dto/base"
	errorhelper "github.com/alxhtp/monogo/pkg/helper/error"
	"github.com/alxhtp/monogo/pkg/message"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	userEntityName = "user"
)

type userUsecase struct {
	userRepository userrepository.UserRepository
	userSerializer userserializer.UserSerializer
	logger         *slog.Logger
	validator      *validator.Validate
}

func NewUserUsecase(userRepository userrepository.UserRepository, userSerializer userserializer.UserSerializer) userusecase.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		userSerializer: userSerializer,
		logger:         slog.Default().With("usecase", userEntityName),
		validator:      validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, req *dto.ReqCreateUser) dto.ResUserSingle {
	u.logger.InfoContext(ctx, "creating user")
	select {
	case <-ctx.Done():
		u.logger.ErrorContext(ctx, "CreateUser: context done", "req", req, "error", ctx.Err().Error())
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusInternalServerError, message.GetResponseMessage(message.FailedCreated, userEntityName), errorhelper.ComposeStacktrace(ctx.Err()))
	default:
	}

	if req == nil {
		u.logger.ErrorContext(ctx, "CreateUser: request is nil")
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusBadRequest, message.GetResponseMessage(message.FailedCreated, userEntityName), errorhelper.ComposeStacktrace(errors.New("request is nil")))
	}

	if err := u.validator.Struct(req); err != nil {
		u.logger.ErrorContext(ctx, "CreateUser: request validation failed", "req", req, "error", err.Error())
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusBadRequest, err.Error(), errorhelper.ComposeStacktrace(err))
	}

	user, err := u.userSerializer.CreateDTOToEntity(*req)
	if err != nil {
		u.logger.ErrorContext(ctx, "CreateUser: error creating user", "req", req, "error", err.Error())
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusInternalServerError, err.Error(), errorhelper.ComposeStacktrace(err))
	}

	output, err := u.userRepository.Create(ctx, &user)
	if err != nil {
		u.logger.ErrorContext(ctx, "CreateUser: error creating user", "req", req, "error", err.Error())
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusInternalServerError, err.Error(), errorhelper.ComposeStacktrace(err))
	}

	u.logger.InfoContext(ctx, "user created", "user", output)
	return u.userSerializer.EntityToResponseSingle(output, http.StatusCreated, message.GetResponseMessage(message.SuccessCreated, userEntityName), nil)
}

func (u *userUsecase) GetUserByID(ctx context.Context, id uuid.UUID) dto.ResUserSingle {
	u.logger.InfoContext(ctx, "getting user by id", "id", id)
	select {
	case <-ctx.Done():
		u.logger.ErrorContext(ctx, "GetUserByID: context done", "id", id, "error", ctx.Err().Error())
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusInternalServerError, message.GetResponseMessage(message.FailedGetByID, userEntityName), errorhelper.ComposeStacktrace(ctx.Err()))
	default:
	}

	if id == uuid.Nil {
		u.logger.ErrorContext(ctx, "GetUserByID: id is nil")
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusBadRequest, message.GetResponseMessage(message.FailedGetByID, userEntityName), errorhelper.ComposeStacktrace(errors.New("id is nil")))
	}

	output, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		u.logger.ErrorContext(ctx, "GetUserByID: error getting user by id", "id", id, "error", err.Error())
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusInternalServerError, err.Error(), errorhelper.ComposeStacktrace(err))
	}

	u.logger.InfoContext(ctx, "user got by id", "user", output)
	return u.userSerializer.EntityToResponseSingle(output, http.StatusOK, message.GetResponseMessage(message.SuccessGetByID, userEntityName), nil)
}

func (u *userUsecase) GetUsersByFilter(ctx context.Context, filter *dto.ReqGetUser) dto.ResUserList {
	u.logger.InfoContext(ctx, "getting users by filter", "filter", filter)
	select {
	case <-ctx.Done():
		u.logger.ErrorContext(ctx, "GetUsersByFilter: context done", "filter", filter, "error", ctx.Err().Error())
		return u.userSerializer.EntityToResponseList(nil, entitybase.BasePaginationResult{}, http.StatusInternalServerError, message.GetResponseMessage(message.FailedList, userEntityName), errorhelper.ComposeStacktrace(ctx.Err()))
	default:
	}

	if filter == nil {
		u.logger.ErrorContext(ctx, "GetUsersByFilter: filter is nil")
		return u.userSerializer.EntityToResponseList(nil, entitybase.BasePaginationResult{}, http.StatusBadRequest, message.GetResponseMessage(message.FailedList, userEntityName), errorhelper.ComposeStacktrace(errors.New("filter is nil")))
	}

	userFilter, err := u.userSerializer.FilterDTOToEntity(*filter)
	if err != nil {
		u.logger.ErrorContext(ctx, "GetUsersByFilter: error converting filter to entity", "filter", filter, "error", err.Error())
		return u.userSerializer.EntityToResponseList(nil, entitybase.BasePaginationResult{}, http.StatusInternalServerError, err.Error(), errorhelper.ComposeStacktrace(err))
	}

	output, paginationResult, err := u.userRepository.GetByFilter(ctx, &userFilter)
	if err != nil {
		u.logger.ErrorContext(ctx, "GetUsersByFilter: error getting users by filter", "filter", filter, "error", err.Error())
		return u.userSerializer.EntityToResponseList(nil, entitybase.BasePaginationResult{}, http.StatusInternalServerError, err.Error(), errorhelper.ComposeStacktrace(err))
	}

	u.logger.InfoContext(ctx, "users got by filter", "users", output)
	return u.userSerializer.EntityToResponseList(output, paginationResult, http.StatusOK, message.GetResponseMessage(message.SuccessList, userEntityName), nil)
}

func (u *userUsecase) UpdateUser(ctx context.Context, id uuid.UUID, req *dto.ReqUpdateUser) dto.ResUserSingle {
	u.logger.InfoContext(ctx, "updating user", "id", id, "req", req)
	select {
	case <-ctx.Done():
		u.logger.ErrorContext(ctx, "UpdateUser: context done", "id", id, "req", req, "error", ctx.Err().Error())
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusInternalServerError, message.GetResponseMessage(message.FailedUpdated, userEntityName), errorhelper.ComposeStacktrace(ctx.Err()))
	default:
	}

	if id == uuid.Nil {
		u.logger.ErrorContext(ctx, "UpdateUser: id is nil")
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusBadRequest, message.GetResponseMessage(message.FailedUpdated, userEntityName), errorhelper.ComposeStacktrace(errors.New("id is nil")))
	}

	if req == nil {
		u.logger.ErrorContext(ctx, "UpdateUser: request is nil")
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusBadRequest, message.GetResponseMessage(message.FailedUpdated, userEntityName), errorhelper.ComposeStacktrace(errors.New("request is nil")))
	}

	updateMap, err := u.userSerializer.UpdateDTOToMap(*req)
	if err != nil {
		u.logger.ErrorContext(ctx, "UpdateUser: error converting update to map", "req", req, "error", err.Error())
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusInternalServerError, err.Error(), errorhelper.ComposeStacktrace(err))
	}

	output, err := u.userRepository.Update(ctx, id, updateMap)
	if err != nil {
		u.logger.ErrorContext(ctx, "UpdateUser: error updating user", "id", id, "req", req, "error", err.Error())
		return u.userSerializer.EntityToResponseSingle(nil, http.StatusInternalServerError, err.Error(), errorhelper.ComposeStacktrace(err))
	}

	u.logger.InfoContext(ctx, "user updated", "user", output)
	return u.userSerializer.EntityToResponseSingle(output, http.StatusOK, message.GetResponseMessage(message.SuccessUpdated, userEntityName), nil)
}

func (u *userUsecase) DeleteUser(ctx context.Context, id uuid.UUID) dtobase.BaseRes {
	u.logger.InfoContext(ctx, "deleting user", "id", id)
	select {
	case <-ctx.Done():
		u.logger.ErrorContext(ctx, "DeleteUser: context done", "id", id, "error", ctx.Err().Error())
		return dtobase.BaseRes{Success: false, Code: http.StatusInternalServerError, Message: message.GetResponseMessage(message.FailedDeleted, userEntityName)}
	default:
	}

	if id == uuid.Nil {
		u.logger.ErrorContext(ctx, "DeleteUser: id is nil")
		return dtobase.BaseRes{Success: false, Code: http.StatusBadRequest, Message: message.GetResponseMessage(message.FailedDeleted, userEntityName)}
	}

	err := u.userRepository.Delete(ctx, id)
	if err != nil {
		u.logger.ErrorContext(ctx, "DeleteUser: error deleting user", "id", id, "error", err.Error())
		return dtobase.BaseRes{Success: false, Code: http.StatusInternalServerError, Message: message.GetResponseMessage(message.FailedDeleted, userEntityName)}
	}

	u.logger.InfoContext(ctx, "user deleted", "id", id)
	return dtobase.BaseRes{Success: true, Code: http.StatusOK, Message: message.GetResponseMessage(message.SuccessDeleted, userEntityName)}
}
