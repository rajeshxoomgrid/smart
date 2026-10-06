-- ============================================================================
-- IMM MOLD CHANGE COMMAND
-- ============================================================================

CREATE TABLE imm_mold_change_command (
    id BIGSERIAL PRIMARY KEY,

    tenant_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,

    device_id VARCHAR(100) NOT NULL,
    machine_id VARCHAR(100) NOT NULL,

    old_mold_no VARCHAR(100),
    requested_mold_no VARCHAR(100) NOT NULL,

    status VARCHAR(30) NOT NULL DEFAULT 'PENDING',

    failure_reason TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    confirmed_at TIMESTAMPTZ,

    CONSTRAINT chk_imm_mold_change_status
        CHECK (
            status IN (
                'PENDING',
                'SUCCESS',
                'FAILED'
            )
        )
);

-- ============================================================================
-- INDEXES
-- ============================================================================

CREATE INDEX idx_imm_mold_change_tenant_device
ON imm_mold_change_command (
    tenant_id,
    device_id
);

CREATE INDEX idx_imm_mold_change_pending
ON imm_mold_change_command (
    tenant_id,
    device_id,
    status
);

CREATE INDEX idx_imm_mold_change_created_at
ON imm_mold_change_command (
    created_at DESC
);

-- ============================================================================
-- ONLY ONE PENDING COMMAND PER DEVICE
-- ============================================================================

CREATE UNIQUE INDEX ux_imm_mold_change_one_pending
ON imm_mold_change_command (
    tenant_id,
    device_id
)
WHERE status = 'PENDING';