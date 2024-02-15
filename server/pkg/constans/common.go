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
	Session         string = "session.id"
	ID              string = "id"
)

var (
	PgHost           = os.Getenv("PG_HOST")
	PgUSER           = os.Getenv("PG_USER")
	PgPASSWORD       = os.Getenv("PG_PASS")
	PgName           = os.Getenv("PG_DB")
	PgPORT           = os.Getenv("PG_PORT")
	RdsHOST          = os.Getenv("REDIS_HOST")
	RdsPORT          = os.Getenv("REDIS_PORT")
	RabHOST          = os.Getenv("RABBIT_HOST")
	RbtUSER          = os.Getenv("RABBIT_USER")
	RbtPASS          = os.Getenv("RABBIT_PASS")
	RbtPORT          = os.Getenv("RABBIT_PORT")
	SrvPort          = os.Getenv(PORT)
	Mode, _          = strconv.ParseBool(os.Getenv("DEBUG"))
	RdsPassword      = os.Getenv("REDIS_PASSWORD")
	ORIGINS          = os.Getenv("ORIGINS")
	SessionExpire, _ = strconv.Atoi(os.Getenv("SESSION_EXPIRE_HOUR"))
	Key              = os.Getenv("SECRET_KEY")
)
