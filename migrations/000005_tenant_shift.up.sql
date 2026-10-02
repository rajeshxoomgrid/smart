-- ============================================================
-- SHIFT TIMING
-- ============================================================

CREATE TABLE IF NOT EXISTS shift_timing (

    id SERIAL PRIMARY KEY,

    tenant_shift_id INTEGER NOT NULL,

    shift_start TIME NOT NULL,

    shift_end TIME NOT NULL,

    created_by INTEGER,

    updated_by INTEGER,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),


    -- Tenant shift
    CONSTRAINT fk_shift_timing_tenant_shift

        FOREIGN KEY (tenant_shift_id)

        REFERENCES tenant_shift(id)

        ON DELETE CASCADE,


    -- One timing configuration per shift
    CONSTRAINT uq_shift_timing_tenant_shift

        UNIQUE (tenant_shift_id),


    -- Start and end cannot be identical
    --
    -- 08:00 -> 08:00 is not allowed
    --
    -- 20:00 -> 08:00 IS allowed
    -- because this is an overnight shift.

    CONSTRAINT chk_shift_timing_time

        CHECK (shift_start <> shift_end)
);


-- ============================================================
-- INDEX
-- ============================================================

CREATE INDEX IF NOT EXISTS idx_shift_timing_tenant_shift

ON shift_timing (tenant_shift_id);


-- ============================================================
-- UPDATED_AT TRIGGER
-- ============================================================

DROP TRIGGER IF EXISTS trg_update_shift_timing_updated_at

ON shift_timing;


CREATE TRIGGER trg_update_shift_timing_updated_at

BEFORE UPDATE

ON shift_timing

FOR EACH ROW

EXECUTE FUNCTION update_updated_at_column();