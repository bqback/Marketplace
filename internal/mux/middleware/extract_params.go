package middleware

import (
	"context"
	"marketplace/internal/apperrors"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ExtractFeedParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		funcName := "ExtractFeedParams"

		logger, requestID, err := utils.GetLoggerAndID(r.Context())
		if err != nil {
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
			return
		}

		// r.Get("/?sortby={sortby}&order={order}&page={page}&per={perpage}&minprice={minprice}&maxprice={maxprice}", handlers.FeedHandler.Feed)

		sortTypeParam := chi.URLParam(r, "sortby")
		sortType, ok := dto.SortTypeMap[sortTypeParam]
		if !ok {
			logger.DebugFmt("Invalid sort type", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}

		orderParam := chi.URLParam(r, "order")
		order, ok := dto.SortOrderMap[orderParam]
		if !ok {
			logger.DebugFmt("Invalid sort order", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}

		sortOpts := dto.SortOptions{Type: sortType, Order: order}

		pageParam := chi.URLParam(r, "page")
		pageNum, err := strconv.ParseUint(pageParam, 10, 64)
		if err != nil {
			logger.DebugFmt("Invalid page number", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}

		perPageParam := chi.URLParam(r, "perpage")
		perPage, err := strconv.ParseUint(perPageParam, 10, 64)
		if err != nil {
			logger.DebugFmt("Invalid number of listings per page", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}

		paginationOpts := dto.PaginationOptions{Page: &pageNum, PerPage: &perPage}

		minPriceParam := chi.URLParam(r, "minprice")
		minPrice, err := strconv.ParseUint(minPriceParam, 10, 64)
		if err != nil {
			logger.DebugFmt("Invalid minimum price", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}

		maxPriceParam := chi.URLParam(r, "maxprice")
		maxPrice, err := strconv.ParseUint(maxPriceParam, 10, 64)
		if err != nil {
			logger.DebugFmt("Invalid maximum price", requestID, funcName, nodeName)
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}

		priceOpts := dto.PriceRange{Min: &minPrice, Max: &maxPrice}

		options := dto.FeedOptions{
			Sort:  &sortOpts,
			Page:  &paginationOpts,
			Price: &priceOpts,
		}

		rCtx := context.WithValue(r.Context(), dto.FeedOptionsKey, options)
		logger.DebugFmt("Stored in context", requestID, funcName, nodeName)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
