package constans

import (
	"os"
	"strconv"
)

const (
	Response        string = "response"
	PORT            string = "PORT"
	InternalError   string = "Internal Server Error"
	ContentType     string = "Content-Type"
	JsonContentType string = "application/json"
	TextContentType string = "text/plain"
	Session         string = "session.id"
	ID              string = "id"
	StartSrvLog     string = "Server started on port %s..."
	SrvStr          string = ":%s"
	Origin          string = "Origin"
	ENV             string = "../.env"
	False           string = "false"
	EmptyStr        string = ""
)

var (
	PgHost           = os.Getenv("PG_HOST")
	PgUSER           = os.Getenv("PG_USER")
	PgPASSWORD       = os.Getenv("PG_PASS")
	PgName           = os.Getenv("PG_DB")
	PgPORT           = os.Getenv("PG_PORT")
	RabHOST          = os.Getenv("RABBIT_HOST")
	RbtUSER          = os.Getenv("RABBIT_USER")
	RbtPASS          = os.Getenv("RABBIT_PASS")
	RbtPORT          = os.Getenv("RABBIT_PORT")
	SrvPort          = os.Getenv(PORT)
	Mode, _          = strconv.ParseBool(os.Getenv("DEBUG"))
	ModeStr          = os.Getenv("DEBUG")
	ORIGINS          = os.Getenv("ORIGINS")
	SessionExpire, _ = strconv.Atoi(os.Getenv("SESSION_EXPIRE_HOUR"))
	Key              = os.Getenv("SECRET_KEY")
)
