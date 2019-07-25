package label

import (
	"context"
	"fmt"
	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/pkg/jsonschema"
	"github.com/pkg/errors"
)

//go:generate mockery -name=LabelRepository -output=automock -outpkg=automock -case=underscore
type LabelRepository interface {
	Upsert(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string, label *model.Label) error
}

//go:generate mockery -name=LabelDefinitionRepository -output=automock -outpkg=automock -case=underscore
type LabelDefinitionRepository interface {
	Create(ctx context.Context, tenant string, labelDefinition *model.LabelDefinition) error
	Exists(ctx context.Context, tenant string, key string) (bool, error)
	GetByKey(ctx context.Context, tenant string, key string) (*model.LabelDefinition, error)
}

type labelUpsertService struct{
	labelRepo LabelRepository
	labelDefinitionRepo LabelDefinitionRepository
}

func NewLabelUpsertService(labelRepo LabelRepository, labelDefinitionRepo LabelDefinitionRepository) *labelUpsertService {
	return &labelUpsertService{labelRepo: labelRepo, labelDefinitionRepo: labelDefinitionRepo}
}

func (s *labelUpsertService) UpsertMultipleLabels(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string, labels map[string]interface{}) error {
	for key, val := range labels {
		err := s.UpsertLabel(ctx, tenant, objectType, objectID, &model.Label{Key: key, Value: val})
		if err != nil {
			return errors.Wrap(err, "while upserting multiple labels")
		}
	}

	return nil
}

func (s *labelUpsertService) UpsertLabel(ctx context.Context, tenant string, objectType model.LabelableObject, objectID string, label *model.Label) error {
	exists, err := s.labelDefinitionRepo.Exists(ctx, tenant, label.Key)
	if err != nil {
		return errors.Wrapf(err, "while checking if LabelDefinition for key %s exists", label.Key)
	}

	if !exists {
		err := s.labelDefinitionRepo.Create(ctx, tenant, &model.LabelDefinition{
			Key: label.Key,
			Schema: nil,
		})
		if err != nil {
			return errors.Wrapf(err, "while creating empty LabelDefinition for %s", label.Key)
		}
	}

	err = s.validateLabelValue(ctx, tenant, label)
	if err != nil {
		return errors.Wrap(err, "while validating label value")
	}

	err = s.labelRepo.Upsert(ctx, tenant, objectType, objectID, label)
	if err != nil {
		return errors.Wrapf(err, "while creating label %s for %s %s", label.Key, objectType, objectID)
	}

	return nil
}

func (s *labelUpsertService) validateLabelValue(ctx context.Context, tenant string, label *model.Label) error {
	labelDefSchema, err := s.labelDefinitionRepo.GetByKey(ctx, tenant, label.Key)
	if err != nil {
		return errors.Wrapf(err, "while reading JSON schema for LabelDefinition for key %s", label.Key)
	}

	if labelDefSchema == nil {
		// nothing to validate
		return nil
	}

	validator, err := jsonschema.NewValidatorFromRawSchema(labelDefSchema)
	if err != nil {
		return errors.Wrapf(err, "while creating JSON Schema validator for schema %+v", labelDefSchema)
	}

	valid, err := validator.ValidateRaw(label.Value)
	if err != nil {
		return errors.Wrapf(err, "while validating value %+v against JSON Schema: %+v", label.Value, labelDefSchema)
	}

	if !valid {
		return fmt.Errorf("Label value for key %s is not valid", label.Key)
	}

	return nil
}
