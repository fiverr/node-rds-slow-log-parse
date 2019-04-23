const prepqre = require('.');

describe('prepare', () => {
    it('Should replace line-breaks with white space', () => {
        const input = `Hello
I have
line
breaks`;
        expect(prepqre(input)).toEqual('Hello I have line breaks');
    });
    it('Should dedup white spaces', () => {
        const input = 'Hello I   have      white             spaces';
        expect(prepqre(input)).toEqual('Hello I have white spaces');
    });
    it('Should handle mixed white spaces and linebreaks', () => {
        const input = `Hello                    
                         I   have
                         white             
                         spaces and line breaks                `;
        expect(prepqre(input)).toEqual('Hello I have white spaces and line breaks ');
    });
});
