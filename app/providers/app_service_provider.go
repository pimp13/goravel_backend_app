package providers

import (
	"goravel_by_gin/app/services"
	"log"

	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type AppServiceProvider struct {
}

func (receiver *AppServiceProvider) Register(app foundation.Application) {
	app.Singleton("services.UploadService", func(app foundation.Application) (any, error) {
		return services.NewUploadService(), nil
	})

	app.Singleton("services.PostService", func(app foundation.Application) (any, error) {
		instance, err := app.Make("services.UploadService")
		if err != nil {
			log.Fatalf("failed to make services.UploadService: %v", err)
			return nil, err
		}
		return services.NewPostService(facades.App().MakeOrm(), instance.(services.UploadService)), nil
	})

	app.Singleton("services.AuthService", func(app foundation.Application) (any, error) {
		return services.NewAuthService(facades.App().MakeOrm()), nil
	})

	app.Singleton("services.TagService", func(app foundation.Application) (any, error) {
		return services.NewTagService(facades.App().MakeOrm()), nil
	})
}

func (receiver *AppServiceProvider) Boot(app foundation.Application) {

}
