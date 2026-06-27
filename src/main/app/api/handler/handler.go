package handler

import "todo_app/src/main/app/service"

type Handler struct {
	ServiceUser service.UserServiceI
}
