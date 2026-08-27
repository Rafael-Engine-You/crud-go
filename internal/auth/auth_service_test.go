package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcryptPasswordHashing(t *testing.T) {
	password := "minhasenhasegura"

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("falha ao gerar hash da senha: %v", err)
	}

	senhaHash := string(hashedBytes)

	err = bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(password))
	if err != nil {
		t.Error("esperava que a senha correta fosse validada com sucesso pelo bcrypt")
	}

	err = bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte("senhadiferente"))
	if err == nil {
		t.Error("esperava que a senha incorreta fosse rejeitada pelo bcrypt")
	}
}
