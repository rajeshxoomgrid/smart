-- ============================================================
-- UPDATE SHIFT HOUR SLOT
-- ============================================================

ALTER TABLE shift_hour_slot

ADD CONSTRAINT chk_shift_hour_slot_time
CHECK (slot_start <> slot_end);


ALTER TABLE shift_hour_slot

ADD CONSTRAINT chk_shift_hour_slot_index
CHECK (slot_index > 0);


ALTER TABLE shift_hour_slot

ADD CONSTRAINT uq_shift_hour_slot_index
UNIQUE (
    shift_timing_id,
    slot_index
);


-- ============================================================
-- INDEX
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