package helper

import "github.com/golang-jwt/jwt/v5"

const jwtSecret = "test"

type JWTData struct {
	Id   int64
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
		"id":   data.Id,
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
	idFloat64 := t.Claims.(jwt.MapClaims)["id"].(float64)
	name := t.Claims.(jwt.MapClaims)["name"]
	return t.Valid, &JWTData{
		Id:   int64(idFloat64),
		Name: name.(string),
	}

}
