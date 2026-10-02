-- ============================================================
-- DROP SHIFT TIMING
-- ============================================================

DROP TRIGGER IF EXISTS trg_update_shift_timing_updated_at
ON shift_timing;

DROP INDEX IF EXISTS idx_shift_timing_tenant_shift;

DROP TABLE IF EXISTS shift_timing;