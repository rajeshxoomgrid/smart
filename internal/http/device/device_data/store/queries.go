package store

const GetLatestProductionByDevice = `
SELECT
    tenant_id,
    event_id,
    customer_id,
    device_id,
    machine_id,
    station,
    variant,
    production_count,
    cycle_time_sec,
    production_time,
    created_at
FROM assembly_production_log
WHERE tenant_id = $1
  AND device_id = $2
  AND ($3 = '' OR station = $3)
  AND ($4 = '' or variant=$4)
ORDER BY created_at DESC
LIMIT $5;
`

const GetFirstShiftProductionCount = `
SELECT production_count
FROM assembly_production_log
WHERE tenant_id = $1
  AND device_id = $2
  AND ($3 = '' OR station = $3)
  AND ($4 = '' OR variant = $4)
  AND production_time >= $5
  AND production_time < $6
ORDER BY production_time ASC
LIMIT 1;
`

const GetFirstDayProductionCount = `
SELECT production_count
FROM assembly_production_log
WHERE tenant_id = $1
  AND device_id = $2
  AND ($3 = '' OR station = $3)
  AND ($4 = '' OR variant = $4)
  AND production_time >= $5
  AND production_time < $6
ORDER BY production_time ASC
LIMIT 1;
`

const GetTenantCountByDevice = `
SELECT
    event_id,
    tenant_id,
    shift_name,
    customer_id,
    device_id,
    machine_id,
    station,
    variant,
    production_count,
    cycle_time_sec,
    production_time,
    created_at
FROM assembly_production_log
WHERE tenant_id = $1
  AND device_id = $2
  AND ($3 = '' OR station = $3)
  AND ($4 = '' OR variant = $4)
ORDER BY production_time DESC
LIMIT 1;
`

// const GetFirstShiftProductionCount = `
// SELECT production_count
// FROM assembly_production_log
// WHERE tenant_id=$1
// AND device_id=$2
// AND ($3='' OR station=$3)
// AND production_time BETWEEN $4 AND $5
// ORDER BY production_time ASC
// LIMIT 1;
// `

// const GetFirstDayProductionCount = `
// SELECT
//     production_count
// FROM assembly_production_log
// WHERE tenant_id = $1
//   AND device_id = $2
//   AND ($3 = '' OR station = $3)
//   AND production_time >= $4
//   AND production_time < $5
// ORDER BY production_time ASC
// LIMIT 1;
// `

// const GetShiftStartProductionCount = `
// SELECT production_count
// FROM assembly_production_log
// WHERE tenant_id=$1
// AND device_id=$2
// AND ($3='' OR station=$3)
// AND production_time BETWEEN $4 AND $5
// ORDER BY production_time ASC
// LIMIT 1;
// `

// const GetLatestProductionByDevice = `
// SELECT
//     tenant_id,
//     event_id,
//     customer_id,
//     device_id,
//     machine_id,
//     station,
//     production_count,
//     cycle_time_sec,
//     production_time,
//     created_at
// FROM assembly_production_log
// WHERE tenant_id = $1
//   AND device_id = $2
//   AND ($3 = '' OR station = $3)
// ORDER BY created_at DESC
// LIMIT $4;
// `

// const GetLatestProductionByDevice = `
// SELECT
//
//	tenant_id,
//	event_id,
//	customer_id,
//	device_id,
//	machine_id,
//	station,
//	production_count,
//	cycle_time_sec,
//	production_time,
//	created_at
//
// FROM assembly_production_log
// WHERE tenant_id = $1
//
//	AND device_id = $2
//	AND ($3 = '' OR station = $3)
//
// ORDER BY production_time DESC
// LIMIT 1;
// `
const GetFirstProductionInShift = `
SELECT
    tenant_id,
    event_id,
    customer_id,
    device_id,
    machine_id,
    station,
    production_count,
    cycle_time_sec,
    production_time,
    created_at
FROM assembly_production_log
WHERE tenant_id = $1
  AND device_id = $2
  AND ($3 = '' OR station = $3)
  AND production_time >= $4
  AND production_time <  $5
ORDER BY production_time ASC
LIMIT 1;
`
const GetLastProductionInShift = `
SELECT
    tenant_id,
    event_id,
    customer_id,
    device_id,
    machine_id,
    station,
    production_count,
    cycle_time_sec,
    production_time,
    created_at
FROM assembly_production_log
WHERE tenant_id = $1
  AND device_id = $2
  AND ($3 = '' OR station = $3)
  AND production_time >= $4
  AND production_time <  $5
ORDER BY production_time DESC
LIMIT 1;
`
