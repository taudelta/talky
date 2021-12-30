package storage

import (
	"github.com/taudelta/talky/internal/dto"
	"github.com/taudelta/talky/internal/models"
	"github.com/taudelta/talky/internal/storage/ifaces"
)

type PhraseRepository interface {
	Find(ifaces.StorageQuery, dto.FindParams) ([]*models.Phrase, error)
}
