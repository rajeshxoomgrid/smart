CREATE TABLE IF NOT EXISTS imm_production_log (
    id BIGSERIAL PRIMARY KEY,

    event_id VARCHAR(150) NOT NULL,
    tenant_id INTEGER NOT NULL,
    device_id VARCHAR(100) NOT NULL,
    machine_id VARCHAR(100) NOT NULL,

    part_name VARCHAR(100) NOT NULL DEFAULT '',

    production_count BIGINT NOT NULL DEFAULT 0,
    cycle_time_sec DOUBLE PRECISION NOT NULL DEFAULT 0,
    production_time TIMESTAMPTZ NOT NULL,

    reset BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT uk_imm_production_log_event_id
        UNIQUE (event_id),

    CONSTRAINT chk_imm_production_log_count
        CHECK (production_count >= 0),

    CONSTRAINT chk_imm_production_log_cycle_time
        CHECK (cycle_time_sec >= 0)
);

CREATE INDEX IF NOT EXISTS idx_imm_production_log_tenant
    ON imm_production_log (tenant_id);

CREATE INDEX IF NOT EXISTS idx_imm_production_log_device
    ON imm_production_log (device_id);

CREATE INDEX IF NOT EXISTS idx_imm_production_log_machine
    ON imm_production_log (machine_id);

CREATE INDEX IF NOT EXISTS idx_imm_production_log_production_time
    ON imm_production_log (production_time DESC);

CREATE INDEX IF NOT EXISTS idx_imm_production_log_tenant_time
    ON imm_production_log (tenant_id, production_time DESC);

CREATE INDEX IF NOT EXISTS idx_imm_production_log_tenant_device_time
    ON imm_production_log (
        tenant_id,
        device_id,
        production_time DESC
    );