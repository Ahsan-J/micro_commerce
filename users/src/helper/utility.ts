import _ from 'lodash';

export const validate = (v : any) => !!v && !_.isUndefined(v) && !_.isNull(v) && !_.isEmpty(v) && v != "undefined" && v != "null";