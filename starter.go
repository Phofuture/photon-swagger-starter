package swaggerStarter

import (
	"github.com/dennesshen/photon-core-starter/core"
	"github.com/dennesshen/photon-swagger-starter/swagger"
)

func init() {
	core.RegisterCoreDependency(swagger.Start)
}
