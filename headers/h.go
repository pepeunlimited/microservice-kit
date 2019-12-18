package headers

import (
	"github.com/pepeunlimited/microservice-kit/middleware"
)

const (
	XForwardedFor			= "X-Forwarded-For"
	XJwtUsername 			= "X-JWT-Username"
	XJwtUserId 	 			= "X-JWT-UserId"
	XJwtRoles 				= "X-JWT-Roles"
	XJwtEmail    			= "X-JWT-Email"
)

func Roles() middleware.Middleware {
	return middlewarez(XJwtRoles)
}

func Email() middleware.Middleware {
	return middlewarez(XJwtEmail)
}

func UserId() middleware.Middleware {
	return middlewarez(XJwtUserId)
}

func Username() middleware.Middleware {
	return middlewarez(XJwtUsername)
}

func XForwardedForz() middleware.Middleware {
	return middlewarez(XForwardedFor)
}