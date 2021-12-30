package storage

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/taudelta/talky/internal/dto"
	"github.com/taudelta/talky/internal/models"
	"github.com/taudelta/talky/internal/storage/ifaces"
)

type phraseRepository struct {
}

func NewPhraseRepository() *phraseRepository {
	return &phraseRepository{}
}

func (r *phraseRepository) Find(q ifaces.StorageQuery, params dto.FindParams) ([]*models.Phrase, error) {
	query := sq.Select("id", "text", "category_id", "created_at").From("phrases").Where(
		sq.Eq{"category_id": params.CategoryIDList},
	).Limit(uint64(params.Limit)).Offset(uint64(params.Offset)).OrderBy("created_at").PlaceholderFormat(sq.Dollar)

	result := make([]*models.Phrase, 0)
	err := q.GetAll(query, func() []interface{} {
		model := &models.Phrase{}
		result = append(result, model)
		return []interface{}{&model.ID, &model.Text, &model.CategoryID, &model.CreatedAt}
	})
	if err != nil {
		return result, errors.WithStack(fmt.Errorf("find phrases error: %w", err))
	}

	return result, nil
}
