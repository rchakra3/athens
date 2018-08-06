package download

import (
	"io"

	"github.com/bketelsen/buffet"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/log"
	p "github.com/gomods/athens/pkg/protocol"
)

// PathVersionZip URL.
const PathVersionZip = "/{module:.+}/@v/{version}.zip"

// VersionZipHandler implements GET baseURL/module/@v/version.zip
func VersionZipHandler(dp p.Protocol, lggr log.Entry, eng *render.Engine) buffalo.Handler {
	const op errors.Op = "download.VersionZipHandler"

	return func(c buffalo.Context) error {
		sp := buffet.SpanFromContext(c).SetOperationName("versionZipHandler")
		defer sp.Finish()
		mod, ver, verInfo, err := getModuleVersion(c, lggr, dp)
		if err != nil {
			err := errors.E(op, errors.M(mod), errors.V(ver), err)
			lggr.SystemErr(err)
			return c.Render(errors.Kind(err), nil)
		}
		defer verInfo.Zip.Close()

		// Calling c.Response().Write will write the header directly
		// and we would get a 0 status in the buffalo logs.
		c.Render(200, nil)
		_, err = io.Copy(c.Response(), verInfo.Zip)
		if err != nil {
			lggr.SystemErr(errors.E(op, errors.M(mod), errors.V(ver), err))
		}

		return nil
	}
}
