package label

import (
	"context"
	"github.com/kyma-incubator/compass/components/director/internal/model"
)

type dbRepository struct {
}

func NewRepository() *dbRepository {
	return &dbRepository{}
}

func (r *dbRepository) Upsert(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string, label *model.Label) error {
	panic("not implemented")
}

func (r *dbRepository) GetByKey(ctx context.Context, tenant string, objectType model.LabelableObject, objectID, key string) (*model.Label, error) {
	panic("not implemented")
}

func (r *dbRepository) List(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string) (map[string]interface{}, error) {
	panic("not implemented")
}

func (r *dbRepository) Delete(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string, key string) error {
	panic("not implemented")
}

func (r *dbRepository) DeleteAll(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string) error {
	panic("not implemented")
}
