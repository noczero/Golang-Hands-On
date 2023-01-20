package database

var connection string

func init() {
	// this function invoke automatically when its imported
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
