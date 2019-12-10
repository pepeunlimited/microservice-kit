package headers

import (
	"github.com/pepeunlimited/microservice-kit/middleware"
)

const (
	XForwardedFor			= "X-Forwarded-For"
	XJwtUsername 			= "X-JWT-Username"
	XJwtUserId 	 			= "X-JWT-UserId"
	XJwtRole 				= "X-JWT-Role"
	XJwtEmail    			= "X-JWT-Email"
)

func Role() middleware.Middleware {
	return middlewarez(XJwtRole)
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