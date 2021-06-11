module scada-lts.org/telegramsvc

go 1.16

require (
	github.com/profclems/go-dotenv v0.1.1
	gorm.io/driver/mysql v1.1.0 // indirect
	gorm.io/gorm v1.21.10
)

replace webhook => ./webhook

replace client => ./client

replace dispatcher => ./dispatcher

replace routes => ./routes
