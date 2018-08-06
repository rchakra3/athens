package download

import (
	"context"
	"io"

	"github.com/gomods/athens/pkg/errors"
	p "github.com/gomods/athens/pkg/protocol"
	"github.com/gomods/athens/pkg/storage"
)

type protocol struct {
	s  storage.Backend
	dp p.Protocol
}

// New takes an upstream Protocol and storage
// it always prefers storage, otherwise it goes to upstream
// and fills the storage with the results.
func New(dp p.Protocol, s storage.Backend) p.Protocol {
	return &protocol{dp: dp, s: s}
}

func (p *protocol) List(ctx context.Context, mod string) ([]string, error) {
	return p.dp.List(ctx, mod)
}

func (p *protocol) Info(ctx context.Context, mod, ver string) ([]byte, error) {
	const op errors.Op = "protocol.Info"
	v, err := p.s.Get(ctx, mod, ver)
	if errors.IsNotFoundErr(err) {
		v, err = p.fillCache(ctx, mod, ver)
	}
	if err != nil {
		return nil, errors.E(op, err)
	}

	return v.Info, nil
}

func (p *protocol) fillCache(ctx context.Context, mod, ver string) (*storage.Version, error) {
	const op errors.Op = "protocol.fillCache"
	v, err := p.dp.Version(ctx, mod, ver)
	if err != nil {
		return nil, errors.E(op, err)
	}
	defer v.Zip.Close()
	err = p.s.Save(ctx, mod, ver, v.Mod, v.Zip, v.Info)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return p.s.Get(ctx, mod, ver)
}

func (p *protocol) Latest(ctx context.Context, mod string) (*storage.RevInfo, error) {
	const op errors.Op = "protocol.Latest"
	info, err := p.dp.Latest(ctx, mod)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return info, nil
}

func (p *protocol) GoMod(ctx context.Context, mod, ver string) ([]byte, error) {
	const op errors.Op = "protocol.GoMod"
	v, err := p.s.Get(ctx, mod, ver)
	if errors.IsNotFoundErr(err) {
		v, err = p.fillCache(ctx, mod, ver)
	}
	if err != nil {
		return nil, errors.E(op, err)
	}

	return v.Mod, nil
}

func (p *protocol) Zip(ctx context.Context, mod, ver string) (io.ReadCloser, error) {
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

func (p *protocol) Version(ctx context.Context, mod, ver string) (*storage.Version, error) {
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
