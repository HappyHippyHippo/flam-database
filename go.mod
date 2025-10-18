module github.com/happyhippyhippo/flam-database

go 1.24.0

replace github.com/happyhippyhippo/flam => ../flam

replace github.com/happyhippyhippo/flam-config => ../flam-config

replace github.com/happyhippyhippo/flam-filesystem => ../flam-filesystem

replace github.com/happyhippyhippo/flam-time => ../flam-time

require (
	github.com/golang/mock v1.6.0
	github.com/happyhippyhippo/flam v0.3.0
	github.com/happyhippyhippo/flam-config v0.3.0
	github.com/happyhippyhippo/flam-filesystem v0.3.0
	github.com/happyhippyhippo/flam-time v0.3.0
	github.com/stretchr/testify v1.8.4
	go.uber.org/dig v1.19.0
	gorm.io/driver/mysql v1.6.0
	gorm.io/driver/postgres v1.6.0
	gorm.io/driver/sqlite v1.6.0
	gorm.io/gorm v1.30.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
