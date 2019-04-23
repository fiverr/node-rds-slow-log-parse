const {pattern, matchers} = require('.');

describe('RDS/pattern', () => {
    test('Should create a patterns comprised of all matchers, separated by 0 or more spaces', () => {
        expect(pattern.source).toEqual(matchers.join('\\s*'));
    });

    const [
        userhost,
        id,
        duration,
        lock_wait,
        rows_sent,
        rows_examined,
        context,
        timestamp,
        query,
    ] = matchers.map((part) => new RegExp(part));

    describe('userhost', () => {
        test('Should extract username and hostname', () => {
            const [, user, host, ip] = userhost.exec('User@Host: admin[wheel] @ localhost [127.0.0.1]');
            expect(user).toEqual('admin');
            expect(host).toEqual('localhost');
            expect(ip).toEqual('127.0.0.1');
        });
        test('Should extract username and hostname', () => {
            const [, user, host, ip] = userhost.exec('User@Host: admin[wheel] @ [127.0.0.1]');
            expect(user).toEqual('admin');
            expect(host).toEqual('');
            expect(ip).toEqual('127.0.0.1');
        });
        test('Should extract username and hostname', () => {
            const [, user, host, ip] = userhost.exec('User@Host: admin2[wheel] @ some-domain.com []');
            expect(user).toEqual('admin2');
            expect(host).toEqual('some-domain.com');
            expect(ip).toEqual('');
        });
    });
    describe('id', () => {
        test('Should extract id', () => {
            const [, match] = id.exec('Id: 123456 #');
            expect(match).toEqual('123456');
        });
    });
    describe('duration', () => {
        test('Should extract duration', () => {
            const [, match] = duration.exec('Query_time: 0.000202');
            expect(match).toEqual('0.000202');
        });
    });
    describe('lock_wait', () => {
        test('Should extract lock_wait', () => {
            const [, match] = lock_wait.exec('Lock_time: 0.000077');
            expect(match).toEqual('0.000077');
        });
    });
    describe('rows_sent', () => {
        test('Should extract rows_sent', () => {
            const [, match] = rows_sent.exec('Rows_sent: 12');
            expect(match).toEqual('12');
        });
    });
    describe('rows_examined', () => {
        test('Should extract rows_examined', () => {
            const [, match] = rows_examined.exec('Rows_examined: 12');
            expect(match).toEqual('12');
        });
    });
    describe('context', () => {
        test('Should ignore SQL context content', () => {
            const [, match] = context.exec('use submiterr_production;');
            expect(match).toEqual('submiterr_production');
        });
    });
    describe('timestamp', () => {
        test('Should extract timestamp', () => {
            const [, match] = timestamp.exec('SET timestamp=1551260760;');
            expect(match).toEqual('1551260760');
        });
    });
    describe('query', () => {
        test('Should extract query with single quotes', () => {
            const [, match] = query.exec("SELECT count(*) from mysql.rds_replication_status WHERE action = 'reset slave' and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1;");
            expect(match).toEqual("SELECT count(*) from mysql.rds_replication_status WHERE action = 'reset slave' and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1");
        });
        test('Should extract query with double quotes', () => {
            const [, match] = query.exec('SELECT count(*) from mysql.rds_replication_status WHERE action = "reset slave" and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1;');
            expect(match).toEqual('SELECT count(*) from mysql.rds_replication_status WHERE action = "reset slave" and master_host is NULL and master_port is NULL ORDER BY action_timestamp LIMIT 1');
        });
    });
});
