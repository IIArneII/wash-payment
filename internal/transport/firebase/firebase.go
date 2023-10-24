package firebase

import (
	"context"
	"path/filepath"
	"strings"
	"time"
	"wash-payment/internal/app"

	opErrors "github.com/go-openapi/errors"
	"go.uber.org/zap"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

const authTimeout = time.Second * 10

var ErrUnauthorized = opErrors.New(401, "unauthorized")

type FirebaseService interface {
	Auth(token string) (*app.Auth, error)
}

type firebaseService struct {
	userSvc app.UserService
	auth    *auth.Client
	l       *zap.SugaredLogger
}

func NewFirebaseService(l *zap.SugaredLogger, keyFilePath string, userSvc app.UserService) (FirebaseService, error) {
	keyFilePath, err := filepath.Abs(keyFilePath)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsFile(keyFilePath)

	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &firebaseService{
		auth:    auth,
		userSvc: userSvc,
		l:       l,
	}, nil
}

func (svc *firebaseService) Auth(bearer string) (*app.Auth, error) {
	svc.l.Infof("token: %s", bearer)

	ctx, cancel := context.WithTimeout(context.Background(), authTimeout)
	defer cancel()

	idToken := strings.TrimSpace(strings.Replace(bearer, "Bearer", "", 1))

	if idToken == "" {
		return nil, ErrUnauthorized
	}

	token, err := svc.auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, ErrUnauthorized
	}
	svc.l.Infof("uid: %s", token.UID)

	fbUser, err := svc.auth.GetUser(ctx, token.UID)
	if err != nil {
		return nil, ErrUnauthorized
	}
	svc.l.Infof("uid: %s", fbUser.UID)

	user, err := svc.userSvc.Get(ctx, fbUser.UID)
	if err != nil {
		svc.l.Infof("err: %w", err)
		return nil, ErrUnauthorized
	}
	svc.l.Infof("user: %s", user.ID)

	return &app.Auth{
		User:         user,
		Disabled:     fbUser.Disabled,
		UserMetadata: (app.AuthUserMeta)(*fbUser.UserMetadata),
	}, nil
}
