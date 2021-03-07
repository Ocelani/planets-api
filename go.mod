module planets-api

go 1.15

// replace (
// 	github.com/Ocelani/planets-api/api => ./api
// 	github.com/Ocelani/planets-api/pkg => ./pkg
// )

require (
	github.com/gofiber/fiber/v2 v2.5.0
	github.com/stretchr/testify v1.6.1
	go.mongodb.org/mongo-driver v1.4.6
)
