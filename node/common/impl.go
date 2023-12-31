package common

import (
	"context"
	"fmt"

	"github.com/LMF709268224/titan-vps/api"
	"github.com/LMF709268224/titan-vps/api/types"
	"github.com/LMF709268224/titan-vps/build"
	"github.com/LMF709268224/titan-vps/journal/alerting"
	"github.com/LMF709268224/titan-vps/node/repo"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/gbrlsnchs/jwt/v3"

	"github.com/google/uuid"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
)

var session = uuid.New()

// CommonAPI api o
type CommonAPI struct {
	Alerting     *alerting.Alerting
	APISecret    *jwt.HMACSHA
	ShutdownChan chan struct{}
}

// MethodGroup: Auth

type jwtPayload struct {
	Allow []auth.Permission
}

// NewCommonAPI initializes a new CommonAPI
func NewCommonAPI(lr repo.LockedRepo, secret *jwt.HMACSHA) (CommonAPI, error) {
	commAPI := CommonAPI{
		APISecret: secret,
	}

	return commAPI, nil
}

// AuthVerify verifies a JWT token and returns the permissions associated with it
func (a *CommonAPI) AuthVerify(ctx context.Context, token string) (*types.JWTPayload, error) {
	var payload types.JWTPayload
	if _, err := jwt.Verify([]byte(token), a.APISecret, &payload); err != nil {
		return nil, xerrors.Errorf("JWT Verification failed: %w", err)
	}

	return &payload, nil
}

// AuthNew generates a new JWT token with the provided permissions
func (a *CommonAPI) AuthNew(ctx context.Context, payload *types.JWTPayload) (string, error) {
	tk, err := jwt.Sign(&payload, a.APISecret)
	if err != nil {
		return "", err
	}

	return string(tk), nil
}

// LogList returns a list of available logging subsystems
func (a *CommonAPI) LogList(context.Context) ([]string, error) {
	return logging.GetSubsystems(), nil
}

// LogSetLevel sets the log level for a given subsystem
func (a *CommonAPI) LogSetLevel(ctx context.Context, subsystem, level string) error {
	return logging.SetLogLevel(subsystem, level)
}

// LogAlerts returns an empty list of alerts
func (a *CommonAPI) LogAlerts(ctx context.Context) ([]alerting.Alert, error) {
	return []alerting.Alert{}, nil
}

// Version provides information about API provider
func (a *CommonAPI) Version(context.Context) (api.APIVersion, error) {
	v, err := api.VersionForType(types.RunningNodeType)
	if err != nil {
		return api.APIVersion{}, err
	}

	return api.APIVersion{
		Version:    build.UserVersion(),
		APIVersion: v,
	}, nil
}

// Discover returns an OpenRPC document describing an RPC API.
func (a *CommonAPI) Discover(ctx context.Context) (types.OpenRPCDocument, error) {
	return nil, fmt.Errorf("not implement")
}

// Shutdown trigger graceful shutdown
func (a *CommonAPI) Shutdown(context.Context) error {
	a.ShutdownChan <- struct{}{}
	return nil
}

// Session returns a UUID of api provider session
func (a *CommonAPI) Session(ctx context.Context) (uuid.UUID, error) {
	return session, nil
}

// Closing jsonrpc closing
func (a *CommonAPI) Closing(context.Context) (<-chan struct{}, error) {
	return make(chan struct{}), nil // relies on jsonrpc closing
}
