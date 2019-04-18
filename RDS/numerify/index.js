/**
 * Mutate a giver object - convert specified fields to floating point numbers
 * @param  {Object}    obj
 * @param  {[Strings]} [fields]
 * @return {Object}
 */
module.exports = function numerify(obj, fields = []) {
    fields.forEach(function convert(field) {
        obj[field] = parseFloat(obj[field]);
    });

    return obj;
}
