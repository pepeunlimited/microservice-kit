package headers

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/siimooo/pepeutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	XJwtUsername 			= "X-JWT-Username"
	XJwtUserId 	 			= "X-JWT-UserId"
	XJwtEmail    			= "X-JWT-Email"
	XJwtRole 				= "X-JWT-Role"
	Authorization 			= "Authorization"
	XForwardedForH			= "X-Forwarded-For"
	BearerPrefix 			= "Bearer "
)

var (
	ErrMissingXJwtEmail 		= errors.New("nignx_jwt: missing X-JWT-Email header")
	ErrMissingXJwtUserId 		= errors.New("nignx_jwt: missing X-JWT-UserId header")
	ErrInvalidXJwtUserId 		= errors.New("nignx_jwt: invalid X-JWT-UserId header")
	ErrMissingXJwtRole 			= errors.New("nignx_jwt: missing X-JWT-Role header")
	ErrInvalidXJwtRole 			= errors.New("nignx_jwt: invalid X-JWT-Role header")
	ErrMissingXJwtUsername 		= errors.New("nignx_jwt: missing X-JWT-Username header")
	ErrMissingXForwardedFor 	= errors.New("nignx_jwt: missing X-Forwarded-For header")
	ErrMissingBearer 			= errors.New("jwt: missing bearer token")
)

func RoleDecode(ctx context.Context) (uint8, error) {
	role, valid := decode(XJwtRole, ctx)
	if !valid {
		return 0, ErrMissingXJwtRole
	}
	parsed, err := strconv.Atoi(role)
	if err != nil {
		return 0, ErrInvalidXJwtRole
	}
	return uint8(parsed), nil
}

func Role() middleware.Middleware {
	return nginxJwt(XJwtRole)
}

func UserIdDecode(ctx context.Context) (int64, error) {
	userId, valid := decode(XJwtUserId, ctx)
	if !valid {
		return -1, ErrMissingXJwtUserId
	}
	parsed, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return -1, ErrInvalidXJwtUserId
	}
	return parsed, nil
}

func UserIdHttpDecode(header http.Header) (int64, error) {
	userId, valid := decode(XJwtUserId, CreateNginxJwtHeader(header.Get(XJwtUserId), XJwtUserId))
	if !valid {
		return -1, ErrMissingXJwtUserId
	}
	parsed, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return -1, ErrInvalidXJwtUserId
	}
	return parsed, nil
}

func UserId() middleware.Middleware {
	return nginxJwt(XJwtUserId)
}

func XForwardedFor() middleware.Middleware {
	return nginxJwt(XForwardedForH)
}

func XForwardedForDecode(ctx context.Context) (string, error) {
	ip, valid := decode(XForwardedForH, ctx)
	if !valid {
		return "", ErrMissingXForwardedFor
	}
	spliced := strings.Split(ip, ",")
	if len(spliced) == 0 {
		return "", nil
	}
	return strings.TrimSpace(spliced[0]), nil
}

func EmailDecode(ctx context.Context) (string, error) {
	email, valid := decode(XJwtUsername, ctx)
	if !valid {
		return "", ErrMissingXJwtEmail
	}
	return email, nil
}

func Email() middleware.Middleware {
	return nginxJwt(XJwtEmail)
}

func UsernameDecode(ctx context.Context) (string, error) {
	username, valid := decode(XJwtUsername, ctx)
	if !valid {
		return "", ErrMissingXJwtUsername
	}
	return username, nil
}

func Username() middleware.Middleware {
	return nginxJwt(XJwtUsername)
}

func nginxJwt(header string) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return create(h, header)
	}
}

func create(h http.Handler, header string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, header, r.Header.Get(header))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func decode(header string, ctx context.Context) (string, bool)  {
	value := ctx.Value(header)
	if value == nil {
		return "", false
	}
	asStr := value.(string)
	if pepeutil.IsEmpty(asStr) {
		return "", false
	}
	return asStr, true
}

func CreateNginxJwtHeader(value string, header string) context.Context {
	return context.WithValue(context.TODO(), header, value)
}

func CreateNginxJwtHeaderXForwaredFor(ip string) context.Context {
	return CreateNginxJwtHeader(ip, XForwardedForH)
}

func CreateNginxJwtHeaderUsername(username string) context.Context {
	return CreateNginxJwtHeader(username, XJwtUsername)
}

func CreateNginxJwtHeaderEmail(email string) context.Context {
	return CreateNginxJwtHeader(email, XJwtEmail)
}

func CreateNginxJwtHeaderRole(role uint8) context.Context {
	return CreateNginxJwtHeader(strconv.Itoa(int(role)), XJwtRole)
}

func CreateNginxJwtHeaderUserId(userId int64) context.Context {
	return CreateNginxJwtHeader(strconv.FormatInt(userId, 10), XJwtUserId)
}

func SetBearer(token string, r *http.Request) {
	r.Header.Set(Authorization, BearerPrefix+token)
}

func GetBearer(bearer string) (string, error) {
	if !strings.HasPrefix(bearer, BearerPrefix) {
		return "", ErrMissingBearer
	}
	return strings.TrimPrefix(bearer, BearerPrefix), nil
}
