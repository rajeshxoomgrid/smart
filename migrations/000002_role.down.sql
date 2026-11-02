
-- ============================================
-- Drop trigger
-- ============================================

DROP TRIGGER IF EXISTS trg_update_user_role_updated_at
ON user_role;

-- ============================================
-- Drop unique index
-- ============================================

DROP INDEX IF EXISTS uq_user_role_tenant_active;

-- ============================================
-- Drop user_role table
-- ============================================

DROP TABLE IF EXISTS user_role;

-- ============================================
-- Do not drop update_updated_at_column()
-- ============================================
-- This function may be used by other tables.
