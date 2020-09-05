import Vue from "vue";
import Vuetify from "vuetify";
import App from "./App.vue";
import store from "./store/store";

import "vuetify/dist/vuetify.min.css";
import "./assets/style/app.global.scss";

const vuetifyOptions = {};
Vue.use(Vuetify);

Vue.config.productionTip = false;

new Vue({
  store,
  render: h => h(App),
  vuetify: new Vuetify(vuetifyOptions)
}).$mount("#app");
