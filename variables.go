package error

// All custom error code should greater than 999 and start three numbers must use http status code like 400, 403, 500
// Custom error codes of business start from 400200 and messages are inserted to ErrorMessages by AddErrorMessages (in main function)
const (
	ERROR_HTTP_TOKEN_NOT_VALID        = 401000
	ERROR_HTTP_FORBIDDEN              = 403000
	ERROR_HTTP_NOT_FOUND              = 404000
	ERROR_HTTP_TOKEN_EXPIRED          = 401001
	ERROR_HTTP_TOKEN_NOT_ISSUED       = 401002
	ERROR_HTTP_TOKEN_NOT_VALID_YET    = 401003
	ERROR_HTTP_TOKEN_INVALID_SEGMENTS = 401004
	ERROR_HTTP_INTERNAL_SERVER_ERROR  = 500000
	ERROR_GLOBAL_MYSQL_ERROR          = 500001
	ERROR_PARAMETERS_INVALID          = 400000

	ERROR_INVALID_LOGIN    = 400001
	ERROR_INVALID_USERNAME = 400002
	ERROR_INVALID_MOBILE   = 400003
	ERROR_INVALID_EMAIL    = 400004
	ERROR_INVALID_STATUS   = 400006

	// password errors
	ERROR_INVALID_PASSWORD                       = 400005
	ERROR_INVALID_RESET_PASSWORD_CODE            = 400007
	ERROR_RESET_PASSWORD_CODE_EXPIRED            = 400008
	ERROR_RESET_PASSWORD_RESEND_TIME_NOT_ELAPSED = 400009

	ERROR_DUPLICATE_USERNAME                          = 400010
	ERROR_DUPLICATE_MOBILE                            = 400011
	ERROR_DUPLICATE_EMAIL                             = 400012
	ERROR_DUPLICATE_KEY                               = 400013
	ERROR_FOREIGN_KEY_CONSTRAINT_FAILED_DELETE_UPDATE = 400014
	ERROR_FOREIGN_KEY_CONSTRAINT_FAILED_ADD_UPDATE    = 400015
	ERROR_SYSTEMIC_RECORD_UPDATE_DELETE               = 400016
	ERROR_USER_STATUS_IS_NOT_ACTIVE                   = 400017
	ERROR_USER_STATUS_IS_NOT_PENDING                  = 400018
	ERROR_UPLOADED_FILE_INCORRECT                     = 400021
	ERROR_UPLOADED_FILE_NOT_SUPPORTED                 = 400022
	ERROR_DOWNLOAD_FILE_INCORRECT                     = 400023
	ERROR_INVALID_TIME_RANGE                          = 400024
	ERROR_DELETE_ADMIN_USER                           = 400025
	ERROR_ADMIN_UPDATE_HIMSELF_STATUS                 = 400026
	ERROR_ADMIN_DELETE_HIMSELF_ADMIN_ROLE             = 400027
	ERROR_INVALID_TYPE                                = 400028
	ERROR_INVALID_USER_ROLE                           = 400029
	ERROR_INVALID_AMOUNT                              = 400030
	ERROR_UPDATE_ADMIN_PASSWORD                       = 400031
	ERROR_INVALID_CURRENT_PASSWORD                    = 400032
	ERROR_INVALID_NEW_PASSWORD                        = 400033
	ERROR_SYSTEMIC_SLIDER_DELETE                      = 400034
	ERROR_INVALID_TIME                                = 400035
	ERROR_USER_NOT_EXISTS                             = 400037
	ERROR_FIREBASE_TOKEN_DUPLICATED                   = 400038
	ERROR_LOGIN_INVALID                               = 400039
	ERROR_INVALID_MESSAGE_ACTION                      = 400041
	ERROR_INVALID_MESSAGE_OBJECT                      = 400042
	ERROR_DUPLICATE_STATUS                            = 400043
	ERROR_PRICE_CURRENCY_IS_INVALID                   = 400044
	ERROR_MONGO_DATABASE_VALIDATION_FAIELD            = 400045
)

var (
	DefaultLanguage = LANG_EN

	errorServiceCode int

	// -->language:{code:message}<--
	I18nTranslates = map[string]map[int]string{
		LANG_FA: {
			ERROR_HTTP_FORBIDDEN:                              "???????????? ?????? ????????",
			ERROR_HTTP_TOKEN_EXPIRED:                          "???????? ?????????? ?????? ??????",
			ERROR_PARAMETERS_INVALID:                          "???????????????????? ???????????? ?????????????? ??????",
			ERROR_GLOBAL_MYSQL_ERROR:                          "???????? ?????????? ?????????????? ?????????? ???????????? ??????",
			ERROR_INVALID_LOGIN:                               "?????? ???????????? ???? ?????? ???????? ?????????????? ??????",
			ERROR_INVALID_MOBILE:                              "???????????? ?????????????? ??????",
			ERROR_INVALID_EMAIL:                               "?????????? ?????????????? ??????",
			ERROR_INVALID_USERNAME:                            "?????? ???????????? ?????????????? ??????",
			ERROR_INVALID_PASSWORD:                            "?????? ???????? ?????????????? ??????",
			ERROR_INVALID_STATUS:                              "?????????? ?????????????? ??????",
			ERROR_INVALID_RESET_PASSWORD_CODE:                 "???? ???????????????????? ?????????????? ??????",
			ERROR_RESET_PASSWORD_CODE_EXPIRED:                 "???? ???????????????????? ?????????? ?????? ??????",
			ERROR_RESET_PASSWORD_RESEND_TIME_NOT_ELAPSED:      "???????? ?????????? ?????????? ???????? ???? ???????????????????? ???????? ???????? ??????",
			ERROR_DUPLICATE_USERNAME:                          "?????? ???????????? ???????????? ??????",
			ERROR_DUPLICATE_MOBILE:                            "???????????? ???????????? ??????",
			ERROR_DUPLICATE_EMAIL:                             "?????????? ???????????? ??????",
			ERROR_DUPLICATE_KEY:                               "???????? ???????? ????????????",
			ERROR_FOREIGN_KEY_CONSTRAINT_FAILED_DELETE_UPDATE: "?????? ???? ?????? ???? ???????????? ??????????",
			ERROR_FOREIGN_KEY_CONSTRAINT_FAILED_ADD_UPDATE:    "?????? ???? ???????????? ???? ???????????? ??????????",
			ERROR_SYSTEMIC_RECORD_UPDATE_DELETE:               "?????????? ???????????? ???? ?????? ?????????? ???????????? ???????? ??????????",
			ERROR_USER_STATUS_IS_NOT_ACTIVE:                   "???????? ?????????? ???????? ????????",
			ERROR_USER_STATUS_IS_NOT_PENDING:                  "???????? ?????????? ???? ???????????? ?????????? ????????",
			ERROR_UPLOADED_FILE_INCORRECT:                     "???????? ?????????? ???????? ????????",
			ERROR_UPLOADED_FILE_NOT_SUPPORTED:                 "???????? ?????????? ???????????????? ?????? ??????",
			ERROR_DOWNLOAD_FILE_INCORRECT:                     "???????? ???????????? ???????? ????????",
			ERROR_INVALID_TIME_RANGE:                          "???????? ?????????? ?????????????? ??????",
			ERROR_DELETE_ADMIN_USER:                           "?????????? ?????? ???????? ???????? ???????????? ?????????? ?????? ???? ???? ?????? ????????",
			ERROR_ADMIN_UPDATE_HIMSELF_STATUS:                 "?????? ?????? ???????????? ?????????? ?????? ???? ?????????? ????????",
			ERROR_ADMIN_DELETE_HIMSELF_ADMIN_ROLE:             "?????? ?????? ???????????? ?????? ???????? ?????? ???? ?????? ????????",
			ERROR_INVALID_TYPE:                                "?????? ?????????????? ??????",
			ERROR_INVALID_USER_ROLE:                           "?????? ?????????? ?????????????? ??????",
			ERROR_INVALID_AMOUNT:                              "???????? ?????????????? ??????",
			ERROR_UPDATE_ADMIN_PASSWORD:                       "?????????? ?????????? ?????? ???????? ???????? ???????? ??????????",
			ERROR_INVALID_CURRENT_PASSWORD:                    "?????? ???????? ???????? ???????????? ??????",
			ERROR_INVALID_NEW_PASSWORD:                        "?????? ???????? ???????? ?????????????? ??????",
			ERROR_SYSTEMIC_SLIDER_DELETE:                      "?????????? ?????? ?????????????? ???????????? ???????? ??????????",
			ERROR_INVALID_TIME:                                "???????? ?????????????? ??????",
			ERROR_INVALID_MESSAGE_ACTION:                      "???????? ?????????????? ??????.",
			ERROR_INVALID_MESSAGE_OBJECT:                      "?????????? ?????????????? ??????.",
			ERROR_DUPLICATE_STATUS:                            "?????????? ???????????? ??????.",
			ERROR_PRICE_CURRENCY_IS_INVALID:                   "???????? ???????? (%s) ???????????? ???????? ????????",
			ERROR_MONGO_DATABASE_VALIDATION_FAIELD:            "???????? ???????????? ????????",
		},
	}
)
