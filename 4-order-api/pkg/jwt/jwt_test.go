package jwt_test

import (
	"go/order-api/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const phone = "+79991234567"
	jwtService := jwt.NewJWT("ad771ba220d2bebaf8008f5f32223f6fe4a68dc651a315f52b1d790445b65a41")
	token, err := jwtService.Create(&jwt.TokenData{
		Phone: phone,
	})
	if err != nil {
		t.Fatal(err)
	}

	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is not valid")
	}

	if data.Phone != phone {
		t.Fatalf("Expected email %s, got %s", phone, data.Phone)
	}
}
