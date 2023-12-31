package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.GetUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{Id}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUserById,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{Id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{Id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: false,
	},
}
