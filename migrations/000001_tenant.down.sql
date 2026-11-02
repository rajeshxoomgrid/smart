-- =========================
-- DROP INDEXES
-- =========================

DROP INDEX IF EXISTS uix_tenant_code_active;

DROP INDEX IF EXISTS uix_tenant_contact_email_active;

DROP INDEX IF EXISTS idx_tenant_active;

DROP INDEX IF EXISTS idx_tenant_timezone;

-- =========================
-- DROP TRIGGER
-- =========================

DROP TRIGGER IF EXISTS trg_update_tenant_updated_at ON tenant;

-- =========================
-- DROP TABLE
-- =========================

DROP TABLE IF EXISTS tenant;

-- =========================
-- DROP FUNCTION (OPTIONAL)
-- =========================

-- ⚠️ Only drop if NOT used by other tables.
-- If other tables use update_updated_at_column(),
-- leave it commented.

-- DROP FUNCTION IF EXISTS update_updated_at_column();