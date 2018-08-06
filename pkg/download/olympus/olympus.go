package olympus

import (
	"context"
	"io"

	"github.com/gomods/athens/pkg/download"
	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/storage"
)

type olympus struct {
	//some way of submitting jobs
}

// New returns a download protocol by using
// go get. You must have a modules supported
// go binary for this to work.
func New() (download.Protocol, error) {
	const op errors.Op = "olympus.New"
	return &olympus{}, nil
}

func (o *olympus) List(ctx context.Context, mod string) ([]string, error) {
	const op errors.Op = "olympus.List"
	// Schedule job to pull from olympus
	return nil, errors.E(op, errors.M(mod), errors.KindNotFound)
}

func (gg *olympus) Info(ctx context.Context, mod string, ver string) ([]byte, error) {
	const op errors.Op = "olympus.Info"
	// Schedule job to pull from olympus
	return nil, errors.E(op, errors.M(mod), errors.V(ver), errors.KindNotFound)
}

func (gg *olympus) Latest(ctx context.Context, mod string) (*storage.RevInfo, error) {
	const op errors.Op = "olympus.Latest"
	// Schedule job to pull from olympus
	return nil, errors.E(op, errors.M(mod), errors.KindNotFound)
}

func (gg *olympus) GoMod(ctx context.Context, mod string, ver string) ([]byte, error) {
	const op errors.Op = "olympus.GoMod"
	// Schedule job to pull from olympus
	return nil, errors.E(op, errors.M(mod), errors.V(ver), errors.KindNotFound)
}

func (gg *olympus) Zip(ctx context.Context, mod, ver string) (io.ReadCloser, error) {
	const op errors.Op = "olympus.Zip"
	// Schedule job to pull from olympus
	return nil, errors.E(op, errors.M(mod), errors.V(ver), errors.KindNotFound)
}

func (gg *olympus) Version(ctx context.Context, mod, ver string) (*storage.Version, error) {
	const op errors.Op = "olympus.Version"
	// Schedule job to pull from olympus
	return nil, errors.E(op, errors.M(mod), errors.V(ver), errors.KindNotFound)
}
