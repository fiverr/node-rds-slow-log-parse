# @fiverr/rds-log-parse

## 📃 RDS Log line parser

```js
const parse = require('@fiverr/rds-log-parse');

const logs = new Buffer(event.awslogs.data, 'base64');
const events = logs.map(parse);
```
