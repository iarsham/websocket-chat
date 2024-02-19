package constans

const (
	DbName            string = "postgres"
	PostgresConnected string = "Postgres Successfully Connected..."
	RedisConnected    string = "Redis Successfully Connected..."
	RabbitMQConnected string = "RabbitMQ Successfully Connected..."
	PgStr             string = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	AmqpStr           string = "amqp://%s:%s@%s:%s/"
	RedisStr          string = "%s:%s"
)
