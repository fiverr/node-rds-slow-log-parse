# rds-slow-log-parse [![](https://circleci.com/gh/fiverr/node-rds-slow-log-parse.svg?style=svg)](https://circleci.com/gh/fiverr/node-rds-log-parse) <a href="https://www.npmjs.com/package/rds-slow-log-parse"><img src="https://img.shields.io/npm/v/rds-slow-log-parse.svg"></a>

## 📃 RDS Slow Log line parser

```js
const parse = require('rds-slow-log-parse');

const payload = new Buffer(event.awslogs.data, 'base64');
const logs = JSON.parse(zlib.gunzipSync(payload).toString('ascii'));

const events = logs.map(parse); // <= ✨ Money time
```

### Example

Record
```
# Time: 190227 9:46:00
# User@Host: rdsadmin[rdsadmin] @ localhost [127.0.0.1] Id: 9059049
# Query_time: 0.000202 Lock_time: 0.000077 Rows_sent: 1 Rows_examined: 1
use operations_production;
SET timestamp=1551260760;
SELECT count(*) from mysql.rds_table_status WHERE action = "reset slave" and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1;
```

Result
```js
{
    user: 'rdsadmin',
    host: 'localhost',
    ip: '127.0.0.1',
    id: '9059049',
    duration: 0.000202,
    lock_wait: 0.000077,
    rows_sent: 1,
    rows_examined: 1,
    context: 'operations_production', // optional
    timestamp: 1551260760000,
    query: `SELECT count(*) from mysql.rds_table_status WHERE action = "reset slave" and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1`,
}
```

### Related projects:
- [Elasticsearch Slow Log line parser](https://github.com/fiverr/node-es-slow-log-parse)
