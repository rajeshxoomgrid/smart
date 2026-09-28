CREATE TABLE permission (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    all_perm    BOOLEAN NOT NULL DEFAULT FALSE,
    create_perm BOOLEAN NOT NULL DEFAULT FALSE,
    read_perm   BOOLEAN NOT NULL DEFAULT FALSE,
    update_perm BOOLEAN NOT NULL DEFAULT FALSE,
    delete_perm BOOLEAN NOT NULL DEFAULT FALSE,
    temp_update_permission BOOLEAN NOT NULL DEFAULT FALSE,

    created_by  BIGINT,
    updated_by  BIGINT,
    deleted_by  BIGINT,

    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    is_deleted  BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at  TIMESTAMPTZ
);