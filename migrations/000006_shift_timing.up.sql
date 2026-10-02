-- ============================================================
-- SHIFT HOUR SLOT
-- ============================================================

CREATE TABLE IF NOT EXISTS shift_hour_slot (

    id BIGSERIAL PRIMARY KEY,

    tenant_id BIGINT NOT NULL,

    shift_timing_id BIGINT NOT NULL,

    slot_start TIME NOT NULL,

    slot_end TIME NOT NULL,

    slot_index INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),


    -- Tenant
    CONSTRAINT fk_shift_hour_slot_tenant

        FOREIGN KEY (tenant_id)

        REFERENCES tenant(id)

        ON DELETE CASCADE,


    -- Shift timing
    CONSTRAINT fk_shift_hour_slot_timing

        FOREIGN KEY (shift_timing_id)

        REFERENCES shift_timing(id)

        ON DELETE CASCADE,


    -- Slot start/end cannot be identical
    CONSTRAINT chk_shift_hour_slot_time

        CHECK (slot_start <> slot_end),


    -- Slot index must start from 1
    CONSTRAINT chk_shift_hour_slot_index

        CHECK (slot_index > 0),


    -- One slot index per shift
    CONSTRAINT uq_shift_hour_slot

        UNIQUE (
            shift_timing_id,
            slot_index
        ),


    -- Same tenant + shift + start time
    -- cannot be duplicated
    CONSTRAINT uq_shift_hour_slot_start

        UNIQUE (
            tenant_id,
            shift_timing_id,
            slot_start
        )
);


-- ============================================================
-- INDEXES
-- ============================================================

CREATE INDEX IF NOT EXISTS idx_shift_hour_slot_tenant

ON shift_hour_slot (tenant_id);


CREATE INDEX IF NOT EXISTS idx_shift_hour_slot_timing

ON shift_hour_slot (shift_timing_id);


CREATE INDEX IF NOT EXISTS idx_shift_hour_slot_timing_index

ON shift_hour_slot (
    shift_timing_id,
    slot_index
);