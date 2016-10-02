package common

const (
	EMAIL_REGX = ""
	// Response message
	BAD_REQUEST_MESSAGE        = "Something went wrong"
	INTERNAL_EXCEPTION_MESSAGE = "INTERNAL_EXCEPTION_MESSAGE"
	LOGIN_WITH_FB_SUCCESS      = "LOGIN_WITH_FB_SUCCESS"
)

// Initialize application config.
var appConfig = NewConfiguration()

// Constants used in whole application
var FB_CLIENT_ID = appConfig.ValueOf("FBClientId")
var FB_CLIENT_SECRECT = appConfig.ValueOf("FBClientSecrect")
var FB_FETCH_INFO_URL = appConfig.ValueOf("FBFetchInfoUrl")
var FB_FIELDS = appConfig.ValueOf("FBFields")
var FB_AUTH_URL = appConfig.ValueOf("FBAuthURL")
var FB_TOKEN_URL = appConfig.ValueOf("FBTokenURL")
var FB_REDIRECT_URL = appConfig.ValueOf("RedirectURL")
