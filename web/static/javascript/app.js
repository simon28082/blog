const _ = require('lodash');

/* axios */
const axios = require('axios');
axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';

const Vue = require('vue');
import App from '../components/app'

const app = new Vue({
    el:'#app',
    components: {
        App
    },
    template:'<App />'
});