module github.com/katesclau/telegramsvc

go 1.16

require (
	github.com/brianvoe/gofakeit/v6 v6.5.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/profclems/go-dotenv v0.1.1
	github.com/stretchr/testify v1.7.0 // indirect
	gorm.io/driver/mysql v1.1.0 // indirect
	gorm.io/gorm v1.21.10
)

replace webhook => ./webhook

replace client => ./client

replace dispatcher => ./dispatcher

replace routes => ./routes
