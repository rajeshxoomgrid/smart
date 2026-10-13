

CREATE TABLE apmold (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tenant_id BIGINT NOT NULL,
    part_id BIGINT NOT NULL,
    mold_no VARCHAR(50) NOT NULL,
    mold_name VARCHAR(150),
    cavity_count INTEGER NOT NULL DEFAULT 1,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_by BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_apmold_part
        FOREIGN KEY (part_id)
        REFERENCES part(id),

    CONSTRAINT chk_apmold_cavity_count
        CHECK (cavity_count > 0)
);

CREATE UNIQUE INDEX uq_apmold_tenant_mold_part_active
    ON apmold (tenant_id, mold_no, part_id)
    WHERE is_deleted = FALSE;

CREATE INDEX idx_apmold_tenant_id
    ON apmold (tenant_id);

CREATE INDEX idx_apmold_active
    ON apmold (tenant_id, is_active)
    WHERE is_deleted = FALSE;

CREATE INDEX idx_apmold_deleted_by
    ON apmold (deleted_by)
    WHERE deleted_by IS NOT NULL;