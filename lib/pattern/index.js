const USERNAME = '(\\w*)\\[\\w*\\]';
const HOSTNAME = '([\\w\\d\\.-]*)';
const IP = '\\[([\\d\\.]*)\\]';
const DIGITS_SEQUENCE = '(\\d*)';
const DECIMAL_NUMBER = '([\\d\\.]*)';
const TABLE_NAME = '[\\w]*';
const delimiter = '\\s*';

const matchers = [
    `User@Host: ${USERNAME} @ ${HOSTNAME}\\s?${IP}`,
    `Id: ${DIGITS_SEQUENCE} #`,
    `Query_time: ${DECIMAL_NUMBER}`,
    `Lock_time: ${DECIMAL_NUMBER}`,
    `Rows_sent: ${DIGITS_SEQUENCE}`,
    `Rows_examined: ${DIGITS_SEQUENCE}`,
    `(?:use )?(${TABLE_NAME})?;?`, // This entire line is optional
    `SET timestamp=${DIGITS_SEQUENCE};`,
    '(.*);'
];

module.exports = {
    pattern: new RegExp(matchers.join(delimiter)),
    matchers,
    delimiter
};
