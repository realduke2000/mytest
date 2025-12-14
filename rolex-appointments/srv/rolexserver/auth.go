package rolexserver

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
const (
    issuer = "rolex-server"
    AUTH_AGENT_ID = "chronodigmwatch"
)

func GenerateToken() (string, error) {
    return generateToken(AUTH_AGENT_ID, time.Hour * 24 * 365, []byte(AUTH_AGENT_ID))
}

func generateToken(agentID string, ttl time.Duration, secret []byte) (string, error) {
    now := time.Now()

    claims := jwt.RegisteredClaims{
            Subject:   agentID,
            Audience: []string {agentID},
            Issuer:    issuer,
            IssuedAt:  jwt.NewNumericDate(now),
            ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
            // NotBefore: jwt.NewNumericDate(now), 
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // 使用 HS256 + secret 签名
    return token.SignedString(secret)
}

func VerifyToken(agentID, tokenString string, secret []byte) (*jwt.RegisteredClaims, error) {
    claim := jwt.RegisteredClaims{}
    _, err := jwt.ParseWithClaims(tokenString, &claim, func(t *jwt.Token) (any, error) {
        if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }
        return secret, nil
    }, 
    jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), 
    jwt.WithAudience(agentID),
    jwt.WithSubject(agentID),
    jwt.WithIssuer(issuer))
    if err != nil {
        slog.Debug("token validate err", "err", err)
    }
    return &claim, err
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Auth-Token")
        _, err := VerifyToken(AUTH_AGENT_ID, token, []byte(AUTH_AGENT_ID))
        if err != nil {
            slog.Warn("Request with invalid token.")
            w.WriteHeader(http.StatusUnauthorized)
        }  else {
            next.ServeHTTP(w, r)
        }
    })
}