<template>
    <section v-if="errored">
          <p>We're sorry, we're not able to retrieve this information at the moment, please try back later</p>
      </section>  
      <section v-else>
          <div v-if="loading">Loading...</div>
          <div v-else class="video">
              <v-row justify="center">
                <v-col cols=12 v-for="(video, key) in info" :key="key" sm=10 md=6 lg=6 xl=3 >
                    <v-card height=300 class="darken-4">
                        <iframe width="100%" height="100%" :src="video.VideoId | videoUrl" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>                                                                          
                        <v-btn elevation="2" @click='removeFavorite(video.ID)'>お気に入り解除</v-btn>                                                                                             
                    </v-card>
                </v-col>
            </v-row>              
          </div>
      </section>              
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
      axios.get('/api/videos/favorites')
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
  methods: {
      removeFavorite: function(id) {        
        axios.delete('/api/videos/favorites/' + id)
        .then(response => {
        console.log(response.data)
      })
      .catch(error => {                
        console.log(error)
        this.errored = true
      })
      .finally(() => this.loading = false)              
      }
  },
  filters: {
      videoUrl (value) {          
          return 'https://www.youtube.com/embed/' + value
      }
    }
  }
</script>
