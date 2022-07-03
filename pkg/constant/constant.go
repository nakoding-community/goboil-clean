package constant

// General
const (
	APP                  = "APP"
	APP_NAME             = "github.com/nakoding-community/goboil-clean"
	PORT                 = "PORT"
	ENV                  = "ENV"
	VERSION              = "VERSION"
	HOST                 = "HOST"
	SCHEME               = "SCHEME"
	JWT_KEY              = "JWT_KEY"
	FIRESTORE_PROJECT_ID = "FIRESTORE_PROJECT_ID"
	IS_RUN_MIGRATION     = "IS_RUN_MIGRATION"
	IS_RUN_SEEDER        = "IS_RUN_SEEDER"
	IS_RUN_CRON          = "IS_RUN_CRON"
)

// Db
const (
	DB_DEFAULT_CREATED_BY = "system"
	DB_HOST               = "DB_HOST"
	DB_USER               = "DB_USER"
	DB_PASS               = "DB_PASS"
	DB_PORT               = "DB_PORT"
	DB_NAME               = "DB_NAME"
	DB_SSLMODE            = "DB_SSLMODE"
	DB_TZ                 = "DB_TZ"
	DB_GOBOIL_CLEAN       = "goboil_clean_db"
)

const (
	LENGTH_CODE        = 20
	MAX_DATA_FIRESTORE = 50
)

type (
	contextKey string
	reqIDKey   string
)

const (
	CONTEXT_KEY contextKey = "context_key"
	REQ_ID_KEY  reqIDKey   = "req_id_key"
)
