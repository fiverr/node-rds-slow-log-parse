# rds-slow-log-parse [![](https://circleci.com/gh/fiverr/node-rds-slow-log-parse.svg?style=svg)](https://circleci.com/gh/fiverr/node-rds-log-parse) <a href="https://www.npmjs.com/package/rds-slow-log-parse"><img src="https://img.shields.io/npm/v/rds-slow-log-parse.svg"></a>

## 📃 RDS Log line parser

```js
const parse = require('rds-slow-log-parse');

const payload = new Buffer(event.awslogs.data, 'base64');
const logs = JSON.parse(zlib.gunzipSync(payload).toString(‘ascii’));
const events = logs.map(parse);
```
