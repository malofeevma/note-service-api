package note

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
)

func (s *Service) GetNote(ctx context.Context, noteID int64) (*model.Note, error) {
	return s.noteStorage.GetNote(ctx, noteID)
}