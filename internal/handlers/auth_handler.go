package handlers

import (
	"encoding/json"
	"marketplace/internal/apperrors"
	"marketplace/internal/pkg/dto"
	"marketplace/internal/service"
	"marketplace/internal/utils"
	"net/http"
)

type AuthHandler struct {
	as service.IAuthService
}

// @Summary Авторизоваться
// @Description
// @Tags Авторизация
//
// @Accept  json
// @Produce  json
//
// @Param loginData body dto.LoginInfo true "Данные для авторизации"
//
// @Success 200  {object}  dto.JWT "JWT-токен для аутентификации пользователя"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/login/ [post]
func (ah AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	funcName := "Login"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	var info dto.LoginInfo
	err = json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		logger.DebugFmt("Failed to decode request: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
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

	token, err := ah.as.Authorize(rCtx, info)
	if err != nil {
		logger.DebugFmt("Failed to authorize: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.UnauthorizedResponse, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Authorization", token.Token)

	r.Body.Close()
}

// @Summary Зарегистрироваться
// @Description
// @Tags Авторизация
//
// @Accept  json
// @Produce  json
//
// @Param signupData body dto.SignupInfo true "Данные для регистрации"
//
// @Success 200  {object}  dto.UserInfo "Информация о пользователе"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/singup/ [post]
func (ah AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	funcName := "Signup"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		logger.DebugFmt("Logger or request id missing from context: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	var info dto.SignupInfo
	err = json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		logger.DebugFmt("Failed to decode request: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
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

	user, err := ah.as.Register(rCtx, info)
	if err != nil {
		logger.DebugFmt("Failed to authorize: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.UnauthorizedResponse, w, r)
		return
	}
	logger.DebugFmt("Signed up successfully", requestID, funcName, nodeName)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Authorization", user.Token)

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		logger.Error("Failed to marshal response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Failed to return response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Sent response", requestID, funcName, nodeName)

	r.Body.Close()
}
