module bot

go 1.17

replace github.com/diable201/GolangDevelopment/tree/master/HW_04/weather => ./weather

require (
	github.com/diable201/GolangDevelopment/tree/master/HW_04/weather v0.0.0-00010101000000-000000000000
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
)

require (
	github.com/briandowns/openweathermap v0.16.0 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
)
