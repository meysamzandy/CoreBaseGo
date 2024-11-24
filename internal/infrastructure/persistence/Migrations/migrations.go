package migrations

import "CoreBaseGo/internal/domain/importer/entity"

func Tables() []interface{} {
	modelsToMigrate := []interface{}{
		&entity.Plan{},
	}
	return modelsToMigrate
}
