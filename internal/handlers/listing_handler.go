package handlers

import (
	"encoding/json"
	"marketplace/internal/apperrors"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/service"
	"marketplace/internal/utils"
	"net/http"
)

type ListingHandler struct {
	ls service.IListingService
}

// @Summary Создать объявление
// @Description
// @Tags Объявления
//
// @Accept  json
// @Produce  json
//
// @Param listingData body dto.NewListingInfo true "Данные о новом объявлении"
//
// @Success 200  {object}  entities.Listing "Данные о новом объявлении"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /listing/ [post]
func (lh ListingHandler) Create(w http.ResponseWriter, r *http.Request) {
	funcName := "CreateListing"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		logger.DebugFmt("Logger or request id missing from context: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	var info dto.NewListingInfo
	err = json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		logger.DebugFmt("Failed to decode request: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}
	logger.DebugFmt("Request parsed", requestID, funcName, nodeName)

	// err = validate.Struct(newActor)
	// if err != nil {
	// 	logger.Error("Validation failed")
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
	// 	}

	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
	// 	}
	// 	apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
	// 	return
	// }

	listing, err := lh.ls.Create(rCtx, info)
	if err != nil {
		logger.DebugFmt("Failed to create listing: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Listing created", requestID, funcName, nodeName)

	if closed := respondOnErr(err, listing, "Listing not created", logger, requestID, funcName, w, r); !closed {
		r.Body.Close()
	}
}

// @Summary Получить ленту объявлений
// @Description Все параметры являются опциональными.
// @Description Тип сортировки по умолчанию - по дате объявления, в убывающем порядке (сначала новые).
// @Description В случае, если для полученных значений page и perpage объявлений не останется (например 3, 25 при <50 объявлений),
// @Description будет возвращён пустой список.
// @Tags Лента
//
// @Accept  json
// @Produce  json
//
// @Param sortby query string false "Тип сортировки (date - дата объявления, price - цена)"
// @Param order query string false "Порядок сортировки (asc - возрастающий, desc - убывающий)"
// @Param page query uint false "Номер страницы"
// @Param perpage query uint false "Число объявлений на страница (default - 25)"
// @Param minprice query uint false "Минимальная цена"
// @Param maxprice query uint false "Максимальная цена"
//
// @Success 200  {object}  []dto.FeedListingInfo "Список объявлений"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /feed/?sortby={sortby}&order={order}&page={page}&per={perpage}&minprice={minprice}&maxprice={maxprice} [get]
func (lh ListingHandler) Feed(w http.ResponseWriter, r *http.Request) {
	funcName := "GetFeed"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		logger.DebugFmt("Logger or request id missing from context: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	options, err := utils.GetFeedOpts(rCtx)
	if err != nil {
		logger.DebugFmt("Failed to get options in handler "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	// err = validate.Struct(newActor)
	// if err != nil {
	// 	logger.Error("Validation failed")
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
	// 	}

	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
	// 	}
	// 	apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
	// 	return
	// }

	listings, err := lh.ls.GetFeed(rCtx, options)
	if err != nil {
		logger.DebugFmt("Failed to get feed listings: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Got feed listings", requestID, funcName, nodeName)

	if closed := respondOnErr(err, listings, "No listings available", logger, requestID, funcName, w, r); !closed {
		r.Body.Close()
	}
}
