package label

import (
	"context"
	"fmt"
	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/internal/persistence"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

const tableName string = `"public"."labels"`
const fields string = `"id", "tenant_id", "key", "value", "app_id", "runtime_id"`

type dbRepository struct {
}

func NewRepository() *dbRepository {
	return &dbRepository{}
}

func (r *dbRepository) Upsert(ctx context.Context, label *model.Label) error {
	if label == nil {
		return errors.New("Item cannot be empty")
	}

	persist, err := persistence.FromCtx(ctx)
	if err != nil {
		return errors.Wrap(err, "while fetching persistence from context")
	}

	entity, err := EntityFromModel(label)
	if err != nil {
		return errors.Wrap(err, "while creating Label entity from model")
	}

	stmt := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (:id, :tenant_id, :key, :value, :app_id, :runtime_id) ON CONFLICT target action`,
		fields, tableName)

	_, err = persist.NamedExec(stmt, entity)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == persistence.UniqueViolation {
			return errors.New("Unique Violation error")
		}

		return errors.Wrap(err, "while inserting the runtime entity to database")
	}

	return nil
}

func (r *dbRepository) GetByKey(ctx context.Context, tenant string, objectType model.LabelableObject, objectID, key string) (*model.Label, error) {
	persist, err := persistence.FromCtx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "while fetching DB from context")
	}

	stmt := fmt.Sprintf(`SELECT %s FROM %s WHERE "key" = $1 AND "tenant_id" = $2`,
		fields, tableName)

	var entity Entity
	err = persist.Get(&entity, stmt, key, tenant)
	if err != nil {
		return nil, errors.Wrap(err, "while getting Entity from DB")
	}

	labelModel := entity.ToModel()

	return labelModel, nil
}

func (r *dbRepository) List(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string) (map[string]*model.Label, error) {
	persist, err := persistence.FromCtx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "while fetching DB from context")
	}

	stmt := fmt.Sprintf(`SELECT %s FROM %s WHERE "tenant_id" = $1`,
		fields, tableName)

	var entities []Entity
	err = persist.Select(&entities, stmt, tenant)
	if err != nil {
		return nil, errors.Wrap(err, "while fetching Labels from DB")
	}

	labelsMap := make(map[string]*model.Label)

	for _, entity := range entities {
		model := entity.ToModel()

		labelsMap[model.Key] = model
	}

	return labelsMap, nil
}

func (r *dbRepository) Delete(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string, key string) error {
	persist, err := persistence.FromCtx(ctx)
	if err != nil {
		return errors.Wrap(err, "while fetching persistence from context")
	}

	stmt := fmt.Sprintf(`DELETE FROM %s WHERE "key" = $1 AND "%s" = $2 AND "tenant_id" = $3`, tableName, r.objectField(objectType))
	_, err = persist.Exec(stmt, key, objectID, tenant)

	return errors.Wrap(err, "while deleting the Label entity from database")
}

func (r *dbRepository) DeleteAll(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string) error {
	persist, err := persistence.FromCtx(ctx)
	if err != nil {
		return errors.Wrap(err, "while fetching persistence from context")
	}

	stmt := fmt.Sprintf(`DELETE FROM %s WHERE "%s" = $1 AND "tenant_id" = $2`, tableName, r.objectField(objectType))
	_, err = persist.Exec(stmt, objectID, tenant)

	return errors.Wrapf(err, "while deleting all Label entities from database for %s %s", objectType, objectID)
}

func (r *dbRepository) objectField(objectType model.LabelableObject) string {
	switch objectType {
	case model.ApplicationLabelableObject:
		return "app_id"
	case model.RuntimeLabelableObject:
		return "runtime_id"
	}

	return ""
}
