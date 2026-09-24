package deptstore

// ============================================================
// COMMON SELECT COLUMNS
// ============================================================

const departmentSelectColumns = `
    id,
    department,
    created_by,
    updated_by,
    created_at,
    updated_at,
    is_deleted,
    deleted_at
`

// ============================================================
// CREATE
// ============================================================

const queryCreateDepartment = `
    INSERT INTO department_master (
        department,
        created_by,
        updated_by
    )
    VALUES (
        $1,
        $2,
        $3
    )
    RETURNING
` + departmentSelectColumns

// ============================================================
// GET BY ID
// ============================================================

const queryGetDepartmentByID = `
    SELECT
` + departmentSelectColumns + `
    FROM department_master
    WHERE
        id = $1
        AND is_deleted = FALSE
`

// ============================================================
// GET BY Department Name
// ============================================================
const queryGetDepartmentByName = `
    SELECT 
        ` + departmentSelectColumns + `
    FROM department_master
    WHERE
        department = $1
        AND is_deleted = FALSE
`

// ============================================================
// LIST
// ============================================================

const queryListDepartment = `
    SELECT
` + departmentSelectColumns + `
    FROM department_master
    WHERE
        is_deleted = FALSE
`

// ============================================================
// COUNT
// ============================================================

const queryCountDepartment = `
    SELECT COUNT(*)
    FROM department_master
    WHERE
        is_deleted = FALSE
`

// ============================================================
// UPDATE
// ============================================================

const queryUpdateDepartment = `
    UPDATE department_master
    SET
        department = $1,
        updated_by = $2,
        updated_at = CURRENT_TIMESTAMP
    WHERE
        id = $3
        AND is_deleted = FALSE
    RETURNING
` + departmentSelectColumns

// ============================================================
// SOFT DELETE
// ============================================================

const querySoftDeleteDepartment = `
    UPDATE department_master
    SET
        is_deleted = TRUE,
        deleted_at = CURRENT_TIMESTAMP,
        updated_by = $1,
        updated_at = CURRENT_TIMESTAMP
    WHERE
        id = $2
        AND is_deleted = FALSE
`
