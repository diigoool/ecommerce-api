package config

var JWTSecret []byte

func InitJWT() {
	secret := GetEnv("JWT_SECRET")

	if secret == "" {
		panic("JWT_SECRET is not set")
	}

	JWTSecret = []byte(secret)
}
