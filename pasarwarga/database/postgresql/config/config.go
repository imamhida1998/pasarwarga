package config

const (
	DB_DRIVER   string = "postgres"
	DB_USER     string = "postgres"
	DB_PASSWORD string = "123465"
	DB_NAME     string = "template1"
	DB_SSL_MODE string = "disable"
)

//kita membagi fase-fase pembuatan software menjadi 3 yaitu dev, staging, prod
//kita mendefine variabel CURRENT_PHASE, untuk menyatakan kita sedang berada di fase mana.
const (
	DEV           = "DEV"
	STAGING       = "STAGING"
	PROD          = "PROD"
	ARIGO         = "ARIGO"
	IMAM          = "IMAM"
	CURRENT_PHASE = DEV
	ARIGO_PHASE   = ARIGO
	IMAM_POSTSQL  = IMAM
)

//buatkan struct dbConfig untuk dipakai selanjutnya di map dbConfigs
type DbConfig struct {
	Driver   string
	User     string
	Password string
	DbName   string
	SslMode  string
}


var DB_CONFIGS map[string]DbConfig = map[string]DbConfig{
	DEV: {
		Driver:   "postgres",
		User:     "postgres",
		Password: "postgres",
		DbName:   "gae02",
		SslMode:  "disable",
	},

	STAGING: {
		Driver:   "postgres",
		User:     "postgres",
		Password: "postgres",
		DbName:   "gae02",
		SslMode:  "disable",
	},

	PROD: {
		Driver:   "postgres",
		User:     "postgres",
		Password: "postgres",
		DbName:   "gae02",
		SslMode:  "disable",
	},

	IMAM: {
		Driver:   "postgres",
		User:     "postgres",
		Password: "imamhidayat123",
		DbName:   "pasarwarga",
		SslMode:  "disable",
	},
}
