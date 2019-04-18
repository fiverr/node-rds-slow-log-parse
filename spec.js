const fixtures = require('../fixtures/RDS');
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
