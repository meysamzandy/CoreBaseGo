package migrations

import sampleFeatureEntity "CoreBaseGo/internal/domain/sampleFeature/entity"

func Tables() []interface{} {
	modelsToMigrate := []interface{}{
		&sampleFeatureEntity.SampleFeature{},
	}
	return modelsToMigrate
}
