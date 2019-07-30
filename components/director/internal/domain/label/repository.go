package label

import (
	"context"
	"fmt"
	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/internal/persistence"
	"github.com/pkg/errors"
)

const labelTable string = `"public"."labels"`

type dbRepository struct {
}

func NewRepository() *dbRepository {
	return &dbRepository{}
}

func (r *dbRepository) Upsert(ctx context.Context, label *model.Label) error {
	panic("not implemented")
}

func (r *dbRepository) GetByKey(ctx context.Context, tenant string, objectType model.LabelableObject, objectID, key string) (*model.Label, error) {
	persist, err := persistence.FromCtx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "while fetching DB from context")
	}

	stmt := fmt.Sprintf(`SELECT "id", "tenant_id", "key", "label_def_id", "app_id", "runtime_id", "value" FROM %s WHERE "key" = $1 AND "tenant_id" = $2`,
		labelTable)

	var entity Entity
	err = persist.Get(&entity, stmt, key, tenant)
	if err != nil {
		return nil, errors.Wrap(err, "while getting Entity from DB")
	}

	labelModel, err := entity.ToModel()
	if err != nil {
		return nil, errors.Wrap(err, "while creating Entity model from entity")
	}

	return labelModel, nil
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
