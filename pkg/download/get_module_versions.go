package download

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/log"
	"github.com/gomods/athens/pkg/paths"
	p "github.com/gomods/athens/pkg/protocol"
	"github.com/gomods/athens/pkg/storage"
)

func getModuleVersion(c buffalo.Context, lggr log.Entry, dp p.Protocol) (string, string, *storage.Version, error) {
	const op errors.Op = "download.getModuleVersion"
	params, err := paths.GetAllParams(c)
	if err != nil {
		return "", "", nil, errors.E(op, err, errors.KindBadRequest)
	}

	versionInfo, err := dp.Version(c, params.Module, params.Version)
	if err != nil {
		return "", "", nil, errors.E(op, err)
	}
	return params.Module, params.Version, versionInfo, nil
}
