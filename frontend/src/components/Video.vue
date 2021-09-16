<template>
    <section v-if="errored">
          <p>We're sorry, we're not able to retrieve this information at the moment, please try back later</p>
      </section>  
      <section v-else>
          <div v-if="loading">Loading...</div>
          <div v-else class="video">          
              <v-row class="lighten-4" style="height: 450px;" justify="center">
                <v-col cols=2 v-for="(video, key) in info" :key="key">
                    <v-card outlined tile height=315 md=8 >
                        <iframe width="560" height="315" :src="video.VideoId | videoUrl" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>'
                    </v-card>
                </v-col>
            </v-row>
              <span class="lighten">
                
              </span>
              <v-btn elevation="2">お気に入り解除</v-btn>
          </div>
      </section>   
           <!-- 次は動画ごとにお気に入り解除ボタンをつける！ -->
</template>

<script>
  import axios from 'axios'
  export default {
    name: 'Video',

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
          this.info = response.data
          console.log(this.info)
      })
      .catch(error => {                
          console.log(error)
          this.errored = true
      })
      .finally(() => this.loading = false)
  },

  filters: {
      videoUrl (value) {          
          return 'https://www.youtube.com/embed/' + value
      }
  }
  }
</script>
