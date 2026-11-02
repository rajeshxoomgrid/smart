
-- ============================================
-- Create user_role table
-- ============================================

CREATE TABLE IF NOT EXISTS user_role (
    id SERIAL PRIMARY KEY,
    user_role VARCHAR(100) NOT NULL,
    tenant_id INTEGER NOT NULL,
    created_by INTEGER,
    updated_by INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_user_role_tenant
        FOREIGN KEY (tenant_id)
        REFERENCES tenant(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

-- ============================================
-- Unique active role per tenant
-- Case-insensitive uniqueness
-- ============================================

CREATE UNIQUE INDEX uq_user_role_tenant_active
ON user_role (tenant_id, LOWER(user_role))
WHERE is_deleted = FALSE;

-- ============================================
-- Function to update updated_at
-- ============================================

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- ============================================
-- Trigger to update updated_at
-- ============================================

DROP TRIGGER IF EXISTS trg_update_user_role_updated_at
ON user_role;

CREATE TRIGGER trg_update_user_role_updated_at
BEFORE UPDATE ON user_role
FOR EACH ROW
WHEN (OLD IS DISTINCT FROM NEW)
EXECUTE FUNCTION update_updated_at_column();
