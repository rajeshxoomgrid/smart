CREATE TABLE permission (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    user_id       BIGINT NOT NULL,
    role_id       BIGINT NOT NULL,
    department_id BIGINT NOT NULL,

    all_perm BOOLEAN NOT NULL DEFAULT FALSE,
    create_perm BOOLEAN NOT NULL DEFAULT FALSE,
    read_perm BOOLEAN NOT NULL DEFAULT FALSE,
    update_perm BOOLEAN NOT NULL DEFAULT FALSE,
    delete_perm BOOLEAN NOT NULL DEFAULT FALSE,
    temp_update_permission BOOLEAN NOT NULL DEFAULT FALSE,

    created_by BIGINT,
    updated_by BIGINT,
    deleted_by BIGINT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_permission_user
        FOREIGN KEY (user_id)
        REFERENCES "user"(id),

    CONSTRAINT fk_permission_role
        FOREIGN KEY (role_id)
        REFERENCES role(id),

    CONSTRAINT fk_permission_department
        FOREIGN KEY (department_id)
        REFERENCES department_master(id)
);

CREATE UNIQUE INDEX uq_permission_user_role_department
ON permission(user_id, role_id, department_id)
WHERE is_deleted = FALSE;

CREATE INDEX idx_permission_user_id
ON permission(user_id);

CREATE INDEX idx_permission_role_id
ON permission(role_id);

CREATE INDEX idx_permission_department_id
ON permission(department_id);