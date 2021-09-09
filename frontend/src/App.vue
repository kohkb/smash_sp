<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <v-img
          alt="Smash Logo"
          class="shrink mr-2"
          contain
          src="/smash_logo.jpeg"
          transition="scale-transition"
          width="40"
        />
      </div>

      <v-spacer></v-spacer>
    </v-app-bar>

    <v-main>      
      <Video/>
      <v-btn elevation="2">お気に入り解除</v-btn>
    </v-main>
  </v-app>
</template>

<script>
import Video from './components/Video';

export default {
  name: 'App',

  components: {    
    Video
  },

  data: () => ({
    //
  }),
};
</script>
