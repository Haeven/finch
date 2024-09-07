// pkg/account/account.go
package account

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type AccountService struct {
	db *bun.DB
}

type Account struct {
	bun.BaseModel `bun:"table:accounts"`

	ID             int64     `bun:"id,pk,autoincrement"`
	Username       string    `bun:"username,notnull,unique"`
	Email          string    `bun:"email,notnull,unique"`
	HashedPassword string    `bun:"hashed_password,notnull"`
	ProfilePicture string    `bun:"profile_picture"`
	DisplayName    string    `bun:"display_name"`
	Language       string    `bun:"language,notnull"`
	CreatedAt      time.Time `bun:"created_at,notnull"`
	UpdatedAt      time.Time `bun:"updated_at,notnull"`
}

func NewAccountService(db *bun.DB) *AccountService {
	return &AccountService{
		db: db,
	}
}

func (s *AccountService) CreateAccount(username, email, hashedPassword, language string) (*Account, error) {
	account := &Account{
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
		Language:       language,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	_, err := s.db.NewInsert().Model(account).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) GetAccountByID(id int64) (*Account, error) {
	account := new(Account)
	err := s.db.NewSelect().Model(account).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountService) UpdateProfilePicture(id int64, profilePicture string) error {
	_, err := s.db.NewUpdate().
		Model((*Account)(nil)).
		Set("profile_picture = ?", profilePicture).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(context.Background())
	return err
}

func (s *AccountService) UpdateDisplayName(id int64, displayName string) error {
	_, err := s.db.NewUpdate().
		Model((*Account)(nil)).
		Set("display_name = ?", displayName).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(context.Background())
	return err
}

func (s *AccountService) UpdateLanguage(id int64, language string) error {
	_, err := s.db.NewUpdate().
		Model((*Account)(nil)).
		Set("language = ?", language).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(context.Background())
	return err
}

func (s *AccountService) UpdateEmail(id int64, email string) error {
	_, err := s.db.NewUpdate().
		Model((*Account)(nil)).
		Set("email = ?", email).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(context.Background())
	return err
}

func (s *AccountService) UpdatePassword(id int64, hashedPassword string) error {
	_, err := s.db.NewUpdate().
		Model((*Account)(nil)).
		Set("hashed_password = ?", hashedPassword).
		Set("updated_at = ?", time.Now()).
		Where("id = ?", id).
		Exec(context.Background())
	return err
}

func (s *AccountService) DeleteAccount(id int64) error {
	_, err := s.db.NewDelete().
		Model((*Account)(nil)).
		Where("id = ?", id).
		Exec(context.Background())
	return err
}
