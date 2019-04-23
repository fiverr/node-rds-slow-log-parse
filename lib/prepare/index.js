/**
 * Prepare string for RegExp processing: Replace consecutive line-breaks or white spaces with one white space
 * @param  {String} string
 * @return {String}
 */
module.exports = (string) => string.replace(/[\n\s]+/g, ' ');
