package task

import (
	"github.com/pangum/pangu"
	"github.com/pangum/task/internal/core"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		core.NewAgent,
	).Build().Apply()
}
