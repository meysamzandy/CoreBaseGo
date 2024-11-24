package entity

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Plan struct {
	ID          uint      `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Name        string    `gorm:"size:255;not null"`
	Slug        string    `gorm:"size:255;unique;not null"`
	Description string    `gorm:"type:text"`
	Price       uint      `gorm:"not null"`
	IsActive    bool      `gorm:"default:true"`
}

// StorePlan creates a new plan in the database
func StorePlan(db *gorm.DB, plan *Plan) (*Plan, error) {
	if err := db.Create(plan).Error; err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	return plan, nil
}

// UpdatePlan updates an existing plan
func UpdatePlan(db *gorm.DB, id uint, updatedPlan *Plan) (*Plan, error) {
	var plan Plan
	fmt.Println(updatedPlan.IsActive)
	fmt.Println(updatedPlan.Price)
	if err := db.First(&plan, id).Error; err != nil {
		return nil, fmt.Errorf("plan not found: %w", err)
	}

	if err := db.Model(&plan).Updates(updatedPlan).Error; err != nil {
		return nil, fmt.Errorf("failed to update plan: %w", err)
	}

	if result := db.Preload("PlanFeature").First(&plan, plan.ID); result.Error != nil {
		return nil, fmt.Errorf("error fetching Plan or Feature: %w", result.Error)
	}

	return &plan, nil
}

// DeletePlan removes a plan from the database
func DeletePlan(db *gorm.DB, id uint) error {
	var plan Plan
	if err := db.First(&plan, id).Error; err != nil {
		return fmt.Errorf("plan not found: %w", err)
	}

	if err := db.Delete(&plan).Error; err != nil {
		return fmt.Errorf("failed to delete plan: %w", err)
	}

	return nil
}

// GetPlans retrieves all plans
func GetPlans(db *gorm.DB) ([]Plan, error) {
	var plans []Plan
	if err := db.Find(&plans).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch plans: %w", err)
	}

	return plans, nil
}

// GetPlanByID fetches a single plan by its ID
func GetPlanByID(db *gorm.DB, id uint) (*Plan, error) {
	var plan Plan
	if err := db.First(&plan, id).Error; err != nil {
		return nil, fmt.Errorf("plan not found: %w", err)
	}
	if result := db.Preload("PlanFeature").First(&plan, plan.ID); result.Error != nil {
		return nil, fmt.Errorf("error fetching Plan or Feature: %w", result.Error)
	}
	return &plan, nil
}
