package redirect

import (
	"context"
	"io"

	"github.com/gomods/athens/pkg/errors"
	p "github.com/gomods/athens/pkg/protocol"
	"github.com/gomods/athens/pkg/storage"
)

type redirect struct {
	s storage.Backend
}

// New accepts storage and a worker
// it always prefers storage, otherwise it goes to upstream
// and fills the storage with the results.
func New(s storage.Backend) p.Protocol {
	return &redirect{s: s}
}

func (r *redirect) List(ctx context.Context, mod string) ([]string, error) {
	const op errors.Op = "redirect.List"
	// redirect. There is no work to be done here
	return nil, errors.E(op, mod, errors.KindTempRedirect)
}

func (r *redirect) Info(ctx context.Context, mod, ver string) ([]byte, error) {
	const op errors.Op = "redirect.Info"
	v, err := r.s.Get(ctx, mod, ver)
	if err != nil {
		if errors.IsNotFoundErr(err) {
			// fill the cache asynchronously
			r.fillCacheAsync(ctx, mod, ver)
			return nil, errors.E(op, mod, errors.KindTempRedirect)
		}
		return nil, errors.E(op, err)
	}

	return v.Info, nil
}

func (r *redirect) Latest(ctx context.Context, mod string) (*storage.RevInfo, error) {
	const op errors.Op = "protocol.Latest"
	// redirect. There is no work to be done here
	return nil, errors.E(op, mod, errors.KindTempRedirect)
}

func (r *redirect) GoMod(ctx context.Context, mod, ver string) ([]byte, error) {
	const op errors.Op = "protocol.GoMod"
	v, err := r.s.Get(ctx, mod, ver)
	if errors.IsNotFoundErr(err) {
		r.fillCacheSync(ctx, mod, ver)
	}
	if err != nil {
		return nil, errors.E(op, err)
	}

	return v.Mod, nil
}

func (r *redirect) Zip(ctx context.Context, mod, ver string) (io.ReadCloser, error) {
	const op errors.Op = "protocol.Zip"
	v, err := p.s.Get(ctx, mod, ver)
	if errors.IsNotFoundErr(err) {
		v, err = p.fillCache(ctx, mod, ver)
	}
	if err != nil {
		return nil, errors.E(op, err)
	}

	return v.Zip, nil
}

func (r *redirect) Version(ctx context.Context, mod, ver string) (*storage.Version, error) {
	const op errors.Op = "protocol.Version"
	v, err := p.s.Get(ctx, mod, ver)
	if errors.IsNotFoundErr(err) {
		v, err = p.fillCache(ctx, mod, ver)
	}
	if err != nil {
		return nil, errors.E(op, err)
	}

	return v, nil
}

func (r *redirect) fillCacheAsync(ctx context.Context, mod, ver string) error {
	const op errors.Op = "protocol.fillCache"
	// schedule work
	return nil
}
