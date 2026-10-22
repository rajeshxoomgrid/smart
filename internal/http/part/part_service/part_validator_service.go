package partservice

import (
	"fmt"
	"strings"

	partserviceconstant "github.com/rajeshbond/smart/internal/http/part/part_service/partServiceConstant"
)

func valodateField(partNo, productName string, color, rmGrade *string, partWt *float64) error {
	if strings.TrimSpace(partNo) == "" || len(strings.TrimSpace(partNo)) > 6 {
		return fmt.Errorf("%w: part_no is required and must be atleast more than 6 characters", partserviceconstant.ErrInvalidRequest)
	}
	if strings.TrimSpace(productName) == "" || len(strings.TrimSpace(productName)) > 6 {
		return fmt.Errorf("%w: part_name is required and must be atleast more than 6 characters", partserviceconstant.ErrInvalidRequest)
	}
	if color != nil && len(strings.TrimSpace(*color)) > 50 {
		return fmt.Errorf("%w: color must be at most 50 characters", partserviceconstant.ErrInvalidRequest)
	}
	if rmGrade != nil && len(strings.TrimSpace(*rmGrade)) > 10 {
		return fmt.Errorf("%w: rm_grade must be at most 100 characters", partserviceconstant.ErrInvalidRequest)
	}
	if partWt != nil && *partWt <= 0 {
		return fmt.Errorf("%w: part_wt must be greater than zero", partserviceconstant.ErrInvalidRequest)
	}
	return nil
}

func nornalizeNullableStrings(color, rmGrade, description *string) {
	if color != nil {
		v := strings.TrimSpace(*color)
		*color = v
	}
	if rmGrade != nil {
		v := strings.TrimSpace(*rmGrade)
		*rmGrade = v
	}
	if description != nil {
		v := strings.TrimSpace(*description)
		*description = v
	}
}
