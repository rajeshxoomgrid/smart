package partstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	partconstant "github.com/rajeshbond/smart/internal/http/part/part_constant"
	partdto "github.com/rajeshbond/smart/internal/http/part/part_dto"
)

func (s *PartStore) GetPartByID(ctx context.Context, tenantID, id int64) (*partdto.PartResponse, error) {

	getPartQuery := `
	SELECT ` + partconstant.PartColumns + `
	FROM part 
	WHERE tenant_id = $1
	AND id = $2
	AND is_deleted = FALSE
	`

	part, err := scanPart(s.db.QueryRowContext(
		ctx,
		getPartQuery,
		tenantID,
		id,
	))

	if errors.Is(err, sql.ErrNoRows) {
		return nil, partconstant.ErrNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("get part:%w", err)
	}

	return part, nil

}
