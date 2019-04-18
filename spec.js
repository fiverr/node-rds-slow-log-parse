const fixtures = require('./fixtures');
const rds = require('.');

describe('RDS', () => {
    fixtures.forEach(
        ([entry, result], index) => {
            test(`Should parse entry ${index} correctly`, () => {
                expect(rds(entry)).toEqual(result);
            });
        }
    );
});
