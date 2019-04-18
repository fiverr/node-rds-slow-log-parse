# @fiverr/rds-log-parse [![](https://circleci.com/gh/fiverr/node-rds-log-parse.svg?style=svg)](https://circleci.com/gh/fiverr/node-rds-log-parse) <a href="https://www.npmjs.com/package/@fiverr/rds-log-parse"><img src="https://img.shields.io/npm/v/@fiverr/rds-log-parse.svg"></a>

## 📃 RDS Log line parser

```js
const parse = require('@fiverr/rds-log-parse');

const payload = new Buffer(event.awslogs.data, 'base64');
const logs = JSON.parse(zlib.gunzipSync(payload).toString(‘ascii’));
const events = logs.map(parse);
```
