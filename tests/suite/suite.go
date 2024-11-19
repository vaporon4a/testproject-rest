package suite

import (
	"context"
	"testing"
	"testproject-rest/internal/config"

	"github.com/go-resty/resty/v2"
)

type Suite struct {
	*testing.T
	Cfg    *config.Config
	Client resty.Client
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoad()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ApiTimeout)

	t.Cleanup(func() {
		t.Helper()
		cancel()
	})

	cc := resty.New()
	cc.SetBaseURL("http://" + cfg.ApiAddres)

	return ctx, &Suite{
		T:      t,
		Cfg:    cfg,
		Client: *cc,
	}

}
