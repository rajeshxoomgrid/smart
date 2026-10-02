-- ============================================================
-- DROP SHIFT HOUR SLOT
-- ============================================================

DROP INDEX IF EXISTS idx_shift_hour_slot_timing_index;

DROP INDEX IF EXISTS idx_shift_hour_slot_timing;

DROP INDEX IF EXISTS idx_shift_hour_slot_tenant;

DROP TABLE IF EXISTS shift_hour_slot;