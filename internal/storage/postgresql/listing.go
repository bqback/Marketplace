package postgresql

import (
	"context"
	"database/sql"
	"marketplace/internal/apperrors"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/pkg/entities"
	"marketplace/internal/utils"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PgListingStorage struct {
	db *sqlx.DB
}

func NewListingStorage(db *sqlx.DB) *PgListingStorage {
	return &PgListingStorage{
		db: db,
	}
}

func (s *PgListingStorage) GetListings(ctx context.Context, opts dto.FeedOptions) ([]*entities.Listing, error) {
	return nil, nil
}

func (s *PgListingStorage) Create(ctx context.Context, info dto.NewListingInfo) (*entities.Listing, error) {
	logger, requestID, err := utils.GetLoggerAndID(ctx)
	if err != nil {
		return nil, err
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	funcName := "CreateListing"

	newListing := &entities.Listing{
		NewListingInfo: info,
	}

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.DebugFmt("Failed to start transaction with error "+err.Error(), requestID, funcName, nodeName)
		return nil, apperrors.ErrCouldNotBeginTransaction
	}
	logger.DebugFmt("Transaction started", requestID, funcName, nodeName)

	query1, args, err := squirrel.
		Insert(listingTable).
		Columns(allListingInsertFields...).
		Values(info.Title, info.Description, info.ImageLink, info.Price, info.DateCreated).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(returnIDSuffix).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	var listingID int
	row := tx.QueryRow(query1, args...)
	if err := row.Scan(&listingID); err != nil {
		logger.DebugFmt("Listing insert failed with error "+err.Error(), requestID, funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return nil, apperrors.ErrListingNotCreated
	}
	logger.DebugFmt("Listing created", requestID, funcName, nodeName)

	query2, args, err := squirrel.
		Insert(userListingTable).
		Columns(userListingFields...).
		Values(userID, listingID).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(returnIDSuffix).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID, funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Query built", requestID, funcName, nodeName)

	_, err = tx.Exec(query2, args)
	if err != nil {
		logger.DebugFmt("Failed to link listing to user with error "+err.Error(), requestID, funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return nil, apperrors.ErrCouldNotLinkListing
	}
	logger.DebugFmt("Listing linked", requestID, funcName, nodeName)

	err = tx.Commit()
	if err != nil {
		logger.DebugFmt("Failed to commit changes", requestID, funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return nil, apperrors.ErrListingNotCreated
	}
	logger.DebugFmt("Changes commited", requestID, funcName, nodeName)

	newListing.ID = uint64(listingID)

	return newListing, nil
}
