package hooks

import (
	"github.com/bluenviron/mediamtx/internal/conf"
	"github.com/bluenviron/mediamtx/internal/externalcmd"
	"github.com/bluenviron/mediamtx/internal/logger"
)

type OnGetParams struct {
	Logger          logger.Writer
	ExternalCmdPool *externalcmd.Pool
	Conf            *conf.Path
	ExternalCmdEnv  externalcmd.Environment
	Query           string
}

func OnGet(params OnGetParams) func() {
	var env externalcmd.Environment
	var onGetCmd *externalcmd.Cmd

	if params.Conf.RunOnGet != "" {
		env = params.ExternalCmdEnv
		env["MTX_QUERY"] = params.Query
	}

	if params.Conf.RunOnGet != "" {
		params.Logger.Log(logger.Info, "runOnGet command started")
		onGetCmd = externalcmd.NewCmd(
			params.ExternalCmdPool,
			params.Conf.RunOnGet,
			false,
			env,
			func(err error) {
				params.Logger.Log(logger.Info, "runOnGet command exited: %v", err)
			})
	}

	return func() {
		if onGetCmd != nil {
			onGetCmd.Close()
			params.Logger.Log(logger.Info, "runOnGet command stopped")
		}

		/*if params.Conf.RunOnUnread != "" {
			params.Logger.Log(logger.Info, "runOnUnread command launched")
			externalcmd.NewCmd(
				params.ExternalCmdPool,
				params.Conf.RunOnUnread,
				false,
				env,
				nil)
		}*/
	}
}
