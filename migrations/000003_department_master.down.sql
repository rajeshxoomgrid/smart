
-- ============================================================
-- DROP DEPARTMENT MASTER
-- ============================================================

DROP TRIGGER IF EXISTS trg_update_department_master_updated_at
ON department_master;

DROP INDEX IF EXISTS uq_department_master_tenant_department_active;

DROP INDEX IF EXISTS idx_department_master_tenant_id;

DROP INDEX IF EXISTS idx_department_master_is_deleted;

DROP TABLE IF EXISTS department_master;
