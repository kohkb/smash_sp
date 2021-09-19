import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import Video from './components/Video';
import vuetify from './plugins/vuetify'

Vue.config.productionTip = false

// TODO: Routerを別ファイルに移す
Vue.use(VueRouter)

const routes = [
  {path: '/', component: Video}
]

const router = new VueRouter({
  routes
})

new Vue({  
  vuetify,
  render: h => h(App),
  router: router
}).$mount('#app')
