package repository

import (
	"context"
	"time"

	"go_grpc/lib"
	"go_grpc/model"
)

func (repo *Repository) CreateSession(ctx context.Context, session model.Session) (model.Session, error) {
	err := repo.db.Create(&session).Error
	if err != nil {
		LogError(ctx, "error create session", err)
		return model.Session{}, err
	}

	return session, nil
}

func (repo *Repository) GetSessionByAccessToken(ctx context.Context, accessToken string) (model.Session, error) {
	var session model.Session

	err := repo.db.Where("token = ?", accessToken).Take(&session).Error
	if IsNotFound(err) {
		LogWarn(ctx, "can't find session by session_id")
		return model.Session{}, lib.ErrorNotFound
	}

	if err != nil {
		LogError(ctx, "error find session by access_token", err)
	}

	return session, err
}

func (repo *Repository) GetSessionByID(ctx context.Context, id uint) (model.Session, error) {
	var session model.Session

	err := repo.db.Take(&session, id).Error
	if IsNotFound(err) {
		LogWarn(ctx, "can't find session by session_id")
		return model.Session{}, lib.ErrorNotFound
	}

	if err != nil {
		LogError(ctx, "error get session by session_id", err)
		return model.Session{}, err
	}

	return session, nil
}

func (repo *Repository) UpdateSession(ctx context.Context, session *model.Session) error {
	err := repo.db.Save(session).Error
	if err != nil {
		LogError(ctx, "error save session", err)
	}

	return err
}

func (repo *Repository) GetActiveSessionsByUserID(ctx context.Context, userID uint) ([]model.Session, error) {
	var sessions []model.Session

	err := repo.db.Where("sessions.resource_type = 'users' AND sessions.resource_id = ? AND expires_at > ?", userID, time.Now()).Find(&sessions).Error
	if err != nil {
		LogError(ctx, "error get active sessions by user_id", err)
		return []model.Session{}, err
	}

	return sessions, nil
}
