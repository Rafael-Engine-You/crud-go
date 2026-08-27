package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// Estrutura para receber a requisição de login com e-mail e senha
type LoginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type AuthService struct {
	db *pgxpool.Pool
}

func NewAuthService(db *pgxpool.Pool) *AuthService {
	return &AuthService{db: db}
}

// LoginHandler godoc
// @Summary Realiza a autenticação do usuário
// @Description Autentica um usuário com e-mail e senha e retorna um token JWT de acesso
// @Tags auth
// @Accept json
// @Produce json
// @Param credenciais body LoginRequest true "Credenciais de Acesso (E-mail e Senha)"
// @Success 200 {object} TokenResponse
// @Failure 400 {string} string "Dados de login inválidos"
// @Failure 401 {string} string "E-mail ou senha incorretos"
// @Router /login [post]
func (s *AuthService) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Email == "" || req.Senha == "" {
		http.Error(w, "Dados de login inválidos", http.StatusBadRequest)
		return
	}

	// Busca a senha criptografada no banco de dados pelo e-mail
	var senhaHash string
	sql := `SELECT senha FROM usuarios WHERE email = $1`
	err = s.db.QueryRow(context.Background(), sql, req.Email).Scan(&senhaHash)
	if err != nil {
		http.Error(w, "E-mail ou senha incorretos", http.StatusUnauthorized)
		return
	}

	// Compara a senha enviada com o hash salvo no banco usando bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(req.Senha))
	if err != nil {
		http.Error(w, "E-mail ou senha incorretos", http.StatusUnauthorized)
		return
	}

	// Se a senha estiver correta, gera o JWT
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"email": req.Email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(), // Token expira em 2 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TokenResponse{Token: tokenString})
}

// Middleware para verificar se o token JWT é válido nas rotas protegidas
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token de autorização não fornecido", http.StatusUnauthorized)
			return
		}

		// O formato esperado é "Bearer <token>"
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			http.Error(w, "Formato do token inválido", http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[7:]
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido ou expirado", http.StatusUnauthorized)
			return
		}

		// Se o token for válido, deixa a requisição seguir adiante
		next.ServeHTTP(w, r)
	})
}
