package helper

import "github.com/golang-jwt/jwt/v5"

const jwtSecret = "test"

type JWTData struct {
	Name string
}

type jWT struct {
	Secret string
}

func NewJWT() *jWT {
	return &jWT{
		Secret: jwtSecret,
	}
}

func (j *jWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": data.Name,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *jWT) Parse(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	name := t.Claims.(jwt.MapClaims)["name"]
	return t.Valid, &JWTData{
		Name: name.(string),
	}

}
