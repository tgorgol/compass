package labeldef

import (
	"context"
	"fmt"
	"github.com/kyma-incubator/compass/components/director/pkg/graphql"
	"strings"

	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/internal/persistence"
	"github.com/pkg/errors"
)

const (
	tableName = `"public"."label_definitions"`
)

type repo struct {
	conv Converter
}
//go:generate mockery -name=Converter -output=automock -outpkg=automock -case=underscore
type Converter interface {
	FromGraphQL(input graphql.LabelDefinitionInput, tenant string) model.LabelDefinition
	ToGraphQL(definition model.LabelDefinition) graphql.LabelDefinition
	ToEntity(in model.LabelDefinition) (Entity, error)
}

func NewRepository(conv Converter) *repo {
	return &repo{conv: conv}
}

func (r *repo) Create(ctx context.Context, def model.LabelDefinition) error {
	db, err := persistence.FromCtx(ctx)
	if err != nil {
		return err
	}

	entity, err := r.conv.ToEntity(def)
	if err != nil {
		return errors.Wrap(err, "while converting Label Definition to insert")
	}
	_, err = db.NamedExec(fmt.Sprintf("insert into %s (id, tenant_id, key, schema) values(:id, :tenant_id, :key, :schema)", tableName), entity)
	if err != nil {
		return errors.Wrap(err, "while inserting Label Definition")
	}
	return nil
}

func (r *repo) GetByKey(ctx context.Context, tenant string, key string) (*model.LabelDefinition, error) {
	persist, err := persistence.FromCtx(ctx)
	if err != nil {
		return nil, err
	}

	stmt := fmt.Sprintf(`SELECT "id", "tenant_id", "key", "schema" FROM %s WHERE "key" = $1 AND "tenant_id" = $2`,
		tableName)

	var runtimeEnt Entity
	err = persist.Get(&runtimeEnt, stmt, key, tenant)
	if err != nil {
		return nil, errors.Wrap(err, "while fetching runtime from DB")
	}

	return nil, nil
}

func (r *repo) Exists(ctx context.Context, tenant string, key string) (bool, error) {
	def, err := r.GetByKey(ctx, tenant, key)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return false, nil
		}
		return false, err
	}
	if def != nil {
		return true, nil
	}
	return false, nil
}

type Entity struct {
	ID         string `db:"id"`
	TenantID   string `db:"tenant_id"`
	Key        string `db:"key"`
	SchemaJSON interface{} `db:"schema"`
}