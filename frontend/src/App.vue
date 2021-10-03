<template>
  <v-app>
    <div>
      <v-app-bar
        color="gray accent-4"
        dark
      >
        <v-app-bar-nav-icon></v-app-bar-nav-icon>
        <v-toolbar-title>Smash Sp</v-toolbar-title>
        <div class="flex-grow-1"></div>
        <v-btn icon>
          <v-icon>mdi-heart</v-icon>
        </v-btn>
        <SearchModal />
      </v-app-bar>
    </div>
    <v-main>    
      <v-container>
        <router-view></router-view> 
      </v-container>      
    </v-main>
  </v-app>
</template>

<script>
import SearchModal from './components/SearchModal';
export default {
  name: 'App',
  components: {
    SearchModal
  },

  data: () => ({    
    //
  }),
};
</script>
