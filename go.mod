module planets-api

go 1.15

// replace (
// 	github.com/Ocelani/planets-api/api => ./api
// 	github.com/Ocelani/planets-api/pkg => ./pkg
// )

require (
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/google/gofuzz v1.2.0
	github.com/stretchr/testify v1.6.1
	github.com/tsenart/vegeta/v12 v12.8.4 // indirect
	go.mongodb.org/mongo-driver v1.4.6
)
