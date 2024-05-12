package api

import (
	"net/http"
)

type AccessLevel = int32

const (
	publicAccessLevel AccessLevel = iota + 1
	internalAccessLevel
	userAccessLevel
)

type ProtectFunc = func(*http.Request, *JWT) (*JWT, *Error)

type Authorization struct {
	AccessLevel  AccessLevel
	protectFuncs []ProtectFunc
}

func newAuthorization(accessLevel AccessLevel, protectFunc ProtectFunc) Authorization {
	return Authorization{
		AccessLevel:  accessLevel,
		protectFuncs: []ProtectFunc{protectFunc},
	}
}

func (a *Authorization) Protect(r *http.Request) (jwt *JWT, aErr *Error) {
	for _, f := range a.protectFuncs {
		tempJWT, aErr := f(r, jwt)
		if aErr != nil {
			return nil, aErr
		}

		// Take first valid JWT
		if jwt == nil {
			jwt = tempJWT
		}
	}

	return
}

func (a *Authorization) AddMiddleware(f ProtectFunc) {
	a.protectFuncs = append(a.protectFuncs, f)
}

func getJWT(secret string, r *http.Request) (*JWT, *Error) {
	authHeader := r.Header.Get("Authorization")

	token, err := UnmarshalJWTFromHeader(secret, authHeader)
	if err != nil {
		return nil, NewUnauthorizedError(err, "Unauthorized")
	}

	return token, nil
}

var Public = newAuthorization(publicAccessLevel, func(r *http.Request, _ *JWT) (*JWT, *Error) {
	return nil, nil
})

//var User = newAuthorization(userAccessLevel, func(r *http.Request, _ *JWT) (*JWT, *Error) {
//	jwt, discoErr := getJWT(r)
//	if discoErr != nil {
//		return nil, discoErr
//	}
//
//	if jwt.AccessLevel != userAccessLevel && jwt.AccessLevel != internalAccessLevel {
//		return nil, NewUnauthorizedError(nil, fmt.Sprintf("insufficient access level %d", jwt.AccessLevel))
//	}
//
//	return jwt, nil
//})
//
//var Internal = newAuthorization(internalAccessLevel, func(r *http.Request, _ *JWT) (*JWT, *Error) {
//	jwt, discoErr := getJWT(r)
//	if discoErr != nil {
//		return nil, discoErr
//	}
//
//	if jwt.AccessLevel != internalAccessLevel {
//		return nil, NewUnauthorizedError(nil, fmt.Sprintf("insufficient access level %d", jwt.AccessLevel))
//	}
//
//	return jwt, nil
//})
