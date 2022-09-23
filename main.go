package main

import (
	"clinicapp/pkg/auth"
	"clinicapp/pkg/booking"
	"clinicapp/pkg/canceling"
	"clinicapp/pkg/config"
	"clinicapp/pkg/deleting"
	"clinicapp/pkg/editing"
	"clinicapp/pkg/handler"
	"clinicapp/pkg/listing"
	"clinicapp/pkg/storage/cache"
	"clinicapp/pkg/storage/postgres"

	_ "github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func main() {
	// migration.MigrateDown()
	// migration.MigrateUp()

	// define the storages
	s, _ := postgres.NewStorage()
	c := cache.NewCacheMem()

	// define the services
	lister := listing.NewService(s, c)
	booker := booking.NewService(s)
	canceler := canceling.NewService(s)
	deleter := deleting.NewService(s)
	editer := editing.NewService(s)
	authenticator := auth.NewService(s)

	// define the handlers
	handler.Handler(lister, booker, canceler, deleter, editer, authenticator)
}
