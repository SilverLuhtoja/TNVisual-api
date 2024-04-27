package route

import (
	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"github.com/SilverLuhtoja/TNVisual/internal/infrastructure"
	"github.com/SilverLuhtoja/TNVisual/src/api/auth"
	"github.com/SilverLuhtoja/TNVisual/src/api/project"
	"github.com/SilverLuhtoja/TNVisual/src/api/user"
)

type Controllers struct {
	UserController    user.UserController
	AuthController    auth.AuthController
	ProjectController project.ProjectController
}

func InitializeAllControllers() *Controllers {
	db := infrastructure.NewDatabase()

	return &Controllers{
		UserController:    getUserController(db),
		AuthController:    getAuthController(db),
		ProjectController: getProjectController(db),
	}
}

func getUserController(db *database.Queries) user.UserController {
	userRepo := user.NewUserRepostitory(db)
	userInteractor := user.NewUserInteractor(userRepo)
	return *user.NewUserController(*userInteractor)
}

func getAuthController(db *database.Queries) auth.AuthController {
	loginRepo := auth.NewLoginRepository(db)
	loginInteractor := auth.NewLoginInteractor(loginRepo)
	return *auth.NewLoginController(*loginInteractor)
}

func getProjectController(db *database.Queries) project.ProjectController {
	loginRepo := project.NewLoginRepository(db)
	loginInteractor := project.NewProjectInteractor(loginRepo)
	return *project.NewProjectController(*loginInteractor)
}
