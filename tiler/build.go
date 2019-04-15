package tiler

import (
	"context"

	"github.com/starkandwayne/om-tiler/pattern"
	"github.com/starkandwayne/om-tiler/steps"
)

func (t *Tiler) Build(ctx context.Context, p pattern.Pattern, skipApplyChanges bool) error {
	if err := p.Validate(true); err != nil {
		return err
	}

	s := []steps.Step{
		t.stepPollTillOnline(),
		t.stepConfigureAuthentication(),
		t.stepConfigureDirector(p.Director),
		t.stepDeployDirector(skipApplyChanges),
		t.stepApplyChanges(skipApplyChanges),
	}

	s = append(s, t.stepUploadFiles(p.Tiles)...)
	s = append(s, t.stepConfigureTiles(p.Tiles)...)

	return steps.Run(ctx, s)
}
