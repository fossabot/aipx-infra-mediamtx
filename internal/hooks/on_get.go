package hooks

import (
	"fmt"
	"github.com/bluenviron/mediamtx/internal/conf"
	"github.com/bluenviron/mediamtx/internal/externalcmd"
	"github.com/bluenviron/mediamtx/internal/logger"
	"strings"
	"time"
)

type OnGetParams struct {
	Logger          logger.Writer
	ExternalCmdPool *externalcmd.Pool
	Conf            *conf.Path
	ExternalCmdEnv  externalcmd.Environment
	Query           string
	Start           time.Time
	Duration        time.Duration
	PathName        string
	SegmentPaths    []string
}

func OnGet(params OnGetParams) func() {
	var env externalcmd.Environment
	var onGetCmd *externalcmd.Cmd

	if params.Conf.RunOnGet != "" {
		const layout = "2006-01-02T15:04:05.000Z"
		env = params.ExternalCmdEnv
		env["MTX_PATH"] = params.PathName
		env["MTX_QUERY"] = params.Query
		env["MTX_START"] = params.Start.Format(layout)
		env["MTX_DURATION"] = fmt.Sprintf("%d", int(params.Duration.Seconds()))
		env["MTX_SEGMENT_PATHS"] = strings.Join(params.SegmentPaths, ",")
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
