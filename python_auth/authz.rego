package authz

default allow = false

token = {"valid": true, "payload": payload} {
	[_, payload, _] := io.jwt.decode(input.token)
}

allow {
	token.valid
	token.payload.roles[_] == "admin"
}
