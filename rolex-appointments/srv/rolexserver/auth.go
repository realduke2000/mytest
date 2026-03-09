package rolexserver

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
    issuer = "rolex-server"
    tokenSecret = "chronodigmwatch" // agent id, per node
)

func KnownAgentIDs() map[string]bool {
    return map[string]bool {
        "chronodigmwatch": true,
    }
}

// Tool function 
func GenerateToken() (string, error) {
    return generateToken("chronodigmwatch", time.Hour * 24 * 365, []byte(tokenSecret))
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

func VerifyToken(knownSubjects map[string]bool, tokenString string, secret []byte) (*jwt.RegisteredClaims, error) {
    if splits := strings.SplitN(tokenString, " ", 2); len(splits) != 2 || !strings.EqualFold(splits[0], "Bearer") || strings.TrimSpace(splits[1]) == "" {
        slog.Error("Invalid token format")
        return nil, errors.New("invalid token format")
    }

    claims := jwt.RegisteredClaims{}
    _, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (any, error) {
        if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }
        return secret, nil
    }, 
    jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), 
    jwt.WithIssuer(issuer))

    if err != nil {
        slog.Debug("token validate err", "err", err)
        return &claims, err
    }

    if sub, ok := knownSubjects[claims.Subject]; !ok || !sub {
        slog.Error("Unknown subject", "subject", claims.Subject)
        err = errors.New("unknown token subject")
    }

    return &claims, err
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        claims, err := VerifyToken(KnownAgentIDs(), token, []byte(tokenSecret))
        if err != nil {
            slog.Warn("Request with invalid token.")
            w.WriteHeader(http.StatusUnauthorized)
        }  else {
            ctx := context.WithValue(r.Context(), "agent-id", claims.Subject)
            next.ServeHTTP(w, r.WithContext(ctx))
        }
    })
}