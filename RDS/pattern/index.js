const matchers = [
    'User@Host: (\\w*)\\[\\w*\\] @ ([\\w\\d\\.-]*)\\s?\\[([\\d\\.]*)\\]',
    'Id: (\\d*) #',
    'Query_time: ([\\d\\.]*)',
    'Lock_time: ([\\d\\.]*)',
    'Rows_sent: (\\d*)',
    'Rows_examined: (\\d*)',
    'SET timestamp=(\\d*);',
    '(.*);',
];

module.exports = {
    pattern: new RegExp(matchers.join(' ')),
    matchers,
};
