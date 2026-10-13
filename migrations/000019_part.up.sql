CREATE TABLE part (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tenant_id BIGINT NOT NULL,
    part_no VARCHAR(100) NOT NULL,
    product_name VARCHAR(150) NOT NULL,
    color VARCHAR(50),
    rm_grade VARCHAR(100),
    part_wt NUMERIC(12,3),
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_by BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_part_part_wt
        CHECK (part_wt IS NULL OR part_wt > 0)
);

CREATE UNIQUE INDEX uq_part_tenant_part_no_active
    ON part (tenant_id, part_no)
    WHERE is_deleted = FALSE;

CREATE INDEX idx_part_tenant_id
    ON part (tenant_id);

CREATE INDEX idx_part_active
    ON part (tenant_id, is_active)
    WHERE is_deleted = FALSE;

CREATE INDEX idx_part_deleted_by
    ON part (deleted_by)
    WHERE deleted_by IS NOT NULL;