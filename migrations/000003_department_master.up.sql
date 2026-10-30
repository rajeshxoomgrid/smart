
-- ============================================================
-- CREATE DEPARTMENT MASTER
-- ============================================================

CREATE TABLE IF NOT EXISTS department_master (
    id BIGSERIAL PRIMARY KEY,

    tenant_id BIGINT NOT NULL,
    department VARCHAR(100) NOT NULL,

    created_by BIGINT,
    updated_by BIGINT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_department_master_tenant
        FOREIGN KEY (tenant_id)
        REFERENCES tenant(id)
        ON DELETE CASCADE,

    CONSTRAINT chk_department_master_name
        CHECK (LENGTH(TRIM(department)) > 0)
);

-- ============================================================
-- UNIQUE ACTIVE DEPARTMENT PER TENANT
-- Case-insensitive.
-- Deleted department names can be reused.
-- ============================================================

CREATE UNIQUE INDEX IF NOT EXISTS
    uq_department_master_tenant_department_active
ON department_master (
    tenant_id,
    LOWER(TRIM(department))
)
WHERE is_deleted = FALSE;

-- ============================================================
-- INDEXES
-- ============================================================

CREATE INDEX IF NOT EXISTS idx_department_master_tenant_id
    ON department_master (tenant_id);

CREATE INDEX IF NOT EXISTS idx_department_master_is_deleted
    ON department_master (is_deleted);

-- ============================================================
-- UPDATED_AT TRIGGER
-- Requires update_updated_at_column() to exist.
-- ============================================================

DROP TRIGGER IF EXISTS trg_update_department_master_updated_at
ON department_master;

CREATE TRIGGER trg_update_department_master_updated_at
BEFORE UPDATE ON department_master
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
