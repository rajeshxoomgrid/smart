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

const GetYesterdayProductionQuery = `
WITH params AS (
    SELECT $1::date AS production_date
),
time_ranges AS (
    SELECT
        'Complete Day' AS shift,
        production_date + TIME '08:00:00' AS start_time,
        production_date + INTERVAL '1 day'
            + TIME '08:00:00' AS end_time,
        1 AS sort_order
    FROM params

    UNION ALL

    SELECT
        'Shift A',
        production_date + TIME '08:00:00',
        production_date + TIME '20:00:00',
        2
    FROM params

    UNION ALL

    SELECT
        'Shift B',
        production_date + TIME '20:00:00',
        production_date + INTERVAL '1 day'
            + TIME '08:00:00',
        3
    FROM params
)
SELECT
    tr.shift,
    COALESCE(
        MAX(apl.production_count)
        - MIN(apl.production_count) + 1,
        0
    )::bigint AS production_count
FROM time_ranges tr
LEFT JOIN public.assembly_production_log apl
    ON apl.production_time >= tr.start_time
    AND apl.production_time < tr.end_time
    AND apl.device_id = $2
    AND ($3::text IS NULL OR apl.station = $3)
    AND ($4::text IS NULL OR apl.variant = $4)
GROUP BY
    tr.shift,
    tr.start_time,
    tr.end_time,
    tr.sort_order
ORDER BY tr.sort_order;
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
