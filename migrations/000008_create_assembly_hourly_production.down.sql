-- ============================================================
-- DROP TRIGGER
-- ============================================================

DROP TRIGGER IF EXISTS update_assembly_hourly_production_updated_at ON public.assembly_hourly_production;

-- ============================================================
-- DROP INDEXES
-- ============================================================

DROP INDEX IF EXISTS public.idx_hourly_production_device_station_day;

DROP INDEX IF EXISTS public.idx_hourly_production_production_day;

DROP INDEX IF EXISTS public.idx_hourly_production_shift;

DROP INDEX IF EXISTS public.idx_hourly_production_slot;

-- ============================================================
-- DROP TABLE
--
-- This also removes:
--     Primary key
--     Foreign keys
--     Unique constraint
-- ============================================================

DROP TABLE IF EXISTS public.assembly_hourly_production;