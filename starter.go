package swaggerStarter

import (
	"github.com/Phofuture/photon-core-starter/core"
	"github.com/Phofuture/photon-swagger-starter/swagger"
)

func init() {
	core.RegisterCoreDependency(swagger.Start)
}
