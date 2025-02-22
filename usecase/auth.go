package usecase

import (
	"github.com/kakiyuta/golang-clean-architecture/app/domain/dto/input"
	"github.com/kakiyuta/golang-clean-architecture/app/domain/dto/output"
	"github.com/kakiyuta/golang-clean-architecture/app/infra/db"
	"github.com/kakiyuta/golang-clean-architecture/app/library/weberrors"
)

type AuthUsecase struct {
	connectionController db.Connector
}

func NewAuthUsecase(cc db.Connector) AuthUsecase {
	return AuthUsecase{
		connectionController: cc,
	}
}

func (a *AuthUsecase) Login(input *input.Login) (*output.AuthLogin, error) {
	// 簡易的なログイン処理
	if input.Email != "hoge@example.com" || input.Password != "password" {
		return nil, weberrors.New(401, "Invalid email or password")
	}

	// 本来はJWTなどを生成して返す
	return output.NewAuthLogin("dummy_token"), nil
}
