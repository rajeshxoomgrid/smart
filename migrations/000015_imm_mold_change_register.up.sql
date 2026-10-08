CREATE TABLE IF NOT EXISTS imm_mold_change_register (
    id BIGSERIAL PRIMARY KEY,

    tenant_id BIGINT NOT NULL,

    device_id VARCHAR(100) NOT NULL,

    machine_id VARCHAR(100) NOT NULL,

    old_mold_no VARCHAR(100),

    new_mold_no VARCHAR(100) NOT NULL,

    effective_from TIMESTAMPTZ NOT NULL,

    reason VARCHAR(100),

    created_by BIGINT NOT NULL,

    updated_by BIGINT NOT NULL,

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    deleted_by BIGINT,

    deleted_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT chk_imm_mold_change_new_mold_not_empty
        CHECK (LENGTH(TRIM(new_mold_no)) > 0),

    CONSTRAINT chk_imm_mold_change_device_not_empty
        CHECK (LENGTH(TRIM(device_id)) > 0),

    CONSTRAINT chk_imm_mold_change_machine_not_empty
        CHECK (LENGTH(TRIM(machine_id)) > 0)
);

CREATE INDEX IF NOT EXISTS idx_imm_mold_change_tenant_device
ON imm_mold_change_register (
    tenant_id,
    device_id,
    effective_from DESC
);

CREATE INDEX IF NOT EXISTS idx_imm_mold_change_lookup
ON imm_mold_change_register (
    tenant_id,
    device_id,
    effective_from DESC
)
WHERE is_deleted = FALSE;

CREATE INDEX IF NOT EXISTS idx_imm_mold_change_tenant_machine
ON imm_mold_change_register (
    tenant_id,
    machine_id,
    effective_from DESC
)
WHERE is_deleted = FALSE;

CREATE INDEX IF NOT EXISTS idx_imm_mold_change_created_at
ON imm_mold_change_register (
    created_at DESC
)
WHERE is_deleted = FALSE;