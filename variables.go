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
			ERROR_HTTP_FORBIDDEN:                              "دسترسی غیر مجاز",
			ERROR_HTTP_TOKEN_EXPIRED:                          "توکن منقضی شده است",
			ERROR_PARAMETERS_INVALID:                          "پارامترهای ارسالی نامعتبر است",
			ERROR_GLOBAL_MYSQL_ERROR:                          "خطای داخلی دیتابیس اتفاق افتاده است",
			ERROR_INVALID_LOGIN:                               "نام کاربری یا رمز عبور نامعتبر است",
			ERROR_INVALID_MOBILE:                              "موبایل نامعتبر است",
			ERROR_INVALID_EMAIL:                               "ایمیل نامعتبر است",
			ERROR_INVALID_USERNAME:                            "نام کاربری نامعتبر است",
			ERROR_INVALID_PASSWORD:                            "رمز عبور نامعتبر است",
			ERROR_INVALID_STATUS:                              "وضعیت نامعتبر است",
			ERROR_INVALID_RESET_PASSWORD_CODE:                 "کد اعتبارسنجی نامعتبر است",
			ERROR_RESET_PASSWORD_CODE_EXPIRED:                 "کد اعتبارسنجی منقضی شده است",
			ERROR_RESET_PASSWORD_RESEND_TIME_NOT_ELAPSED:      "زمان امکان ارسال مجدد کد اعتبارسنجی سپری نشده است",
			ERROR_DUPLICATE_USERNAME:                          "نام کاربری تکراری است",
			ERROR_DUPLICATE_MOBILE:                            "موبایل تکراری است",
			ERROR_DUPLICATE_EMAIL:                             "ایمیل تکراری است",
			ERROR_DUPLICATE_KEY:                               "خطای کلید تکراری",
			ERROR_FOREIGN_KEY_CONSTRAINT_FAILED_DELETE_UPDATE: "خطا در حذف یا ویرایش رکورد",
			ERROR_FOREIGN_KEY_CONSTRAINT_FAILED_ADD_UPDATE:    "خطا در افزودن یا ویرایش رکورد",
			ERROR_SYSTEMIC_RECORD_UPDATE_DELETE:               "امکان ویرایش یا حذف رکورد سیستمی وجود ندارد",
			ERROR_USER_STATUS_IS_NOT_ACTIVE:                   "حساب کاربر فعال نیست",
			ERROR_USER_STATUS_IS_NOT_PENDING:                  "حساب کاربر در انتظار تایید نیست",
			ERROR_UPLOADED_FILE_INCORRECT:                     "فایل آپلود صحیح نیست",
			ERROR_UPLOADED_FILE_NOT_SUPPORTED:                 "فایل آپلود پشتیبانی نمی شود",
			ERROR_DOWNLOAD_FILE_INCORRECT:                     "فایل دانلود صحیح نیست",
			ERROR_INVALID_TIME_RANGE:                          "بازه زمانی نامعتبر است",
			ERROR_DELETE_ADMIN_USER:                           "امکان حذف مدیر وجود ندارد؛ ابتدا نقش آن را حذف کنید",
			ERROR_ADMIN_UPDATE_HIMSELF_STATUS:                 "شما نمی توانید وضعیت خود را تغییر دهید",
			ERROR_ADMIN_DELETE_HIMSELF_ADMIN_ROLE:             "شما نمی توانید نقش مدیر خود را حذف کنید",
			ERROR_INVALID_TYPE:                                "نوع نامعتبر است",
			ERROR_INVALID_USER_ROLE:                           "نقش کاربر نامعتبر است",
			ERROR_INVALID_AMOUNT:                              "مبلغ نامعتبر است",
			ERROR_UPDATE_ADMIN_PASSWORD:                       "امکان تغییر رمز عبور مدیر وجود ندارد",
			ERROR_INVALID_CURRENT_PASSWORD:                    "رمز عبور فعلی اشتباه است",
			ERROR_INVALID_NEW_PASSWORD:                        "رمز عبور جدید نامعتبر است",
			ERROR_SYSTEMIC_SLIDER_DELETE:                      "امکان حذف اسلایدر سیستمی وجود ندارد",
			ERROR_INVALID_TIME:                                "زمان نامعتبر است",
			ERROR_INVALID_MESSAGE_ACTION:                      "اکشن نامعتبر است.",
			ERROR_INVALID_MESSAGE_OBJECT:                      "آبجکت نامعتبر است.",
			ERROR_DUPLICATE_STATUS:                            "وضعیت تکراری است.",
			ERROR_PRICE_CURRENCY_IS_INVALID:                   "واحد پولی (%s) ارسالی صحیح نیست",
			ERROR_MONGO_DATABASE_VALIDATION_FAIELD:            "خطای پایگاه داده",
		},
	}
)
