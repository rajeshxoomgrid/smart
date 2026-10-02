-- ============================================================
-- ROLLBACK SHIFT HOUR SLOT UPDATE
-- ============================================================

ALTER TABLE shift_hour_slot
DROP CONSTRAINT IF EXISTS uq_shift_hour_slot_index;

ALTER TABLE shift_hour_slot
DROP CONSTRAINT IF EXISTS chk_shift_hour_slot_index;

ALTER TABLE shift_hour_slot
DROP CONSTRAINT IF EXISTS chk_shift_hour_slot_time;

DROP INDEX IF EXISTS idx_shift_hour_slot_timing_index;

DROP INDEX IF EXISTS idx_shift_hour_slot_timing;

DROP INDEX IF EXISTS idx_shift_hour_slot_tenant;