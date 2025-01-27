package middleware

import ports "guguzaza-users/ports/tokens"

type Middleware struct {
	jwtUtil ports.JwtUtilPort
}

func NewMiddleware(jwtUtil ports.JwtUtilPort) Middleware {
	return Middleware{
		jwtUtil: jwtUtil,
	}
}
