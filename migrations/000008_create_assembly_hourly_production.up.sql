-- ============================================================
-- CREATE TABLE: assembly_hourly_production
--
-- Stores completed hourly production summaries.
--
-- Source:
--     assembly_production_log
--
-- Configuration:
--     tenant_shift
--     shift_timing
--     shift_hour_slot
--
-- One row represents:
--     One device
--     One machine
--     One station
--     One production day
--     One shift
--     One hourly slot
-- ============================================================

CREATE TABLE IF NOT EXISTS public.assembly_hourly_production
(
    id BIGSERIAL PRIMARY KEY,

-- --------------------------------------------------------
-- Tenant / Device Information
-- --------------------------------------------------------

tenant_id VARCHAR(50) NOT NULL,
customer_id VARCHAR(50),
device_id VARCHAR(50) NOT NULL,
machine_id VARCHAR(50) NOT NULL,
station VARCHAR(20) NOT NULL,
variant VARCHAR(100),

-- --------------------------------------------------------
-- Production Day
--
-- Important for overnight Shift B.
--
-- Example:
-- Production Day = 2026-08-12
--
-- Shift B:
-- 20:00 -> 21:00
-- ...
-- 23:00 -> 00:00
-- 00:00 -> 01:00
-- ...
-- 07:00 -> 08:00
--
-- All belong to production_day = 2026-08-12
-- --------------------------------------------------------

production_day DATE NOT NULL,

-- --------------------------------------------------------
-- Shift / Slot References
-- --------------------------------------------------------

tenant_shift_id BIGINT NOT NULL,
shift_timing_id BIGINT NOT NULL,
shift_hour_slot_id BIGINT NOT NULL,
slot_index INTEGER NOT NULL,
slot_start TIME NOT NULL,
slot_end TIME NOT NULL,

-- --------------------------------------------------------
-- Actual Timestamp Boundary
--
-- These are the actual date/time boundaries of the slot.
-- TIMESTAMPTZ is used because production_time is TIMESTAMPTZ.
-- --------------------------------------------------------

actual_start TIMESTAMPTZ NOT NULL, actual_end TIMESTAMPTZ NOT NULL,

-- --------------------------------------------------------
-- Production Counter
--
-- assembly_production_log.production_count is cumulative.
--
-- Example:
--
-- start = 150
-- end   = 208
--
-- hourly production = 58
-- --------------------------------------------------------

start_production_count BIGINT,
end_production_count BIGINT,
hourly_production_count BIGINT NOT NULL DEFAULT 0,

-- --------------------------------------------------------
-- Cycle Time Statistics
-- --------------------------------------------------------

cycle_count INTEGER NOT NULL DEFAULT 0,
min_cycle_time_sec NUMERIC(10, 3),
avg_cycle_time_sec NUMERIC(10, 3),
max_cycle_time_sec NUMERIC(10, 3),

-- --------------------------------------------------------
-- Audit
-- --------------------------------------------------------


created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- FOREIGN KEYS
-- ============================================================

ALTER TABLE public.assembly_hourly_production
ADD CONSTRAINT fk_hourly_production_tenant_shift FOREIGN KEY (tenant_shift_id) REFERENCES public.tenant_shift (id);

ALTER TABLE public.assembly_hourly_production
ADD CONSTRAINT fk_hourly_production_shift_timing FOREIGN KEY (shift_timing_id) REFERENCES public.shift_timing (id);

ALTER TABLE public.assembly_hourly_production
ADD CONSTRAINT fk_hourly_production_shift_hour_slot FOREIGN KEY (shift_hour_slot_id) REFERENCES public.shift_hour_slot (id);

-- ============================================================
-- UNIQUE CONSTRAINT
--
-- Prevents duplicate hourly records.
--
-- IMPORTANT:
-- slot_index alone cannot be used because:
--
-- 2026-08-10 / Shift A / Slot 1
-- 2026-08-11 / Shift A / Slot 1
--
-- are both valid records.
--
-- Therefore production_day + slot ID are required.
-- ============================================================

ALTER TABLE public.assembly_hourly_production
ADD CONSTRAINT uq_assembly_hourly_production_slot UNIQUE (
    tenant_id,
    device_id,
    machine_id,
    station,
    production_day,
    shift_timing_id,
    shift_hour_slot_id
);

-- ============================================================
-- INDEXES
-- ============================================================

-- Fast lookup for current device/station/shift history
CREATE INDEX IF NOT EXISTS idx_hourly_production_device_station_day ON public.assembly_hourly_production (
    tenant_id,
    device_id,
    station,
    production_day
);

-- Fast lookup by production day
CREATE INDEX IF NOT EXISTS idx_hourly_production_production_day ON public.assembly_hourly_production (tenant_id, production_day);

-- Fast lookup by shift
CREATE INDEX IF NOT EXISTS idx_hourly_production_shift ON public.assembly_hourly_production (
    tenant_shift_id,
    production_day
);

-- Fast lookup by slot
CREATE INDEX IF NOT EXISTS idx_hourly_production_slot ON public.assembly_hourly_production (
    shift_hour_slot_id,
    production_day
);

-- ============================================================
-- UPDATED_AT TRIGGER
--
-- Assumes your project already has:
--
-- update_updated_at_column()
--
-- If this function already exists in your database,
-- this trigger will automatically update updated_at.
-- ============================================================

DROP TRIGGER IF EXISTS update_assembly_hourly_production_updated_at
ON public.assembly_hourly_production;

CREATE TRIGGER update_assembly_hourly_production_updated_at
BEFORE UPDATE
ON public.assembly_hourly_production
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();