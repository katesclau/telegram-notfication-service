module scada-lts.org/telegramsvc

go 1.16

require (
	github.com/profclems/go-dotenv v0.1.1
)

replace handler => ./handler
replace clinet => ./client
