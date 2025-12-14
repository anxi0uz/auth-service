package model

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Login    string
	Email    string
	Password string
	FullName string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

type TokenClaims struct {
	UserID      uuid.UUID
	Email       string
	Permissions []string
	ExpiresAt   time.Time
}

type AuthUser struct {
	ID          uuid.UUID  `json:"id"`
	Login       string     `json:"login"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`

	FailedLoginAttempts int        `json:"-"`
	LastLoginAt         *time.Time `json:"last_login_at,omitempty"`
	PasswordChangedAt   *time.Time `json:"password_changed_at,omitempty"`
}
