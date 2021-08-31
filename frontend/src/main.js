import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'

Vue.config.productionTip = false

new Vue({
  el: '#app',
  data () {
      return {
          info: null,
          loading: true,
          errored: false
      }
  },
  mounted() {
      axios.get('http://localhost:8000/videos/favorites')
      .then(response => {
          this.info = response.data.bpi
      })
      .catch(error => {                
          console.log(error)
          this.errored = true
      })
      .finally(() => this.loading = false)
  },

  filters: {
      currencydecimal (value) {
          return value.toFixed(2)
      }
  }
})
