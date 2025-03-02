package swagger

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/dennesshen/photon-core-starter/bean"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func init() {
	bean.Autowire(&component)
}

var component struct {
	Server *fiber.App `autowired:"true"`
}

func Start(ctx context.Context) error {
	if !config.Swagger.Enabled {
		return nil
	}
	contextPath := config.Swagger.ContextPath

	info := fmt.Sprintf("swagger inint, context-path: %s", contextPath)
	slog.InfoContext(ctx, info)
	component.Server.Static(fmt.Sprintf("%s/swagger/swagger.json", contextPath), config.Swagger.DirectoryPath+"/swagger.json")
	component.Server.Static(fmt.Sprintf("%s/swagger/swagger.yaml", contextPath), config.Swagger.DirectoryPath+"/swagger.json")
	component.Server.Get(fmt.Sprintf("%s/swagger/*", contextPath), swagger.New(swagger.Config{
		URL:         "swagger.yaml",
		DeepLinking: true,
	}))
	return nil
}
