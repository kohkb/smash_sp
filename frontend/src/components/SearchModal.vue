<template>
    <v-dialog v-model="dialog" scrollable max-width="80%">
      <v-card>
        <v-card-title>タイトル</v-card-title>
        <v-divider></v-divider>
        <v-card-text height="200px"></v-card-text>
      </v-card>
    </v-dialog>    
</template>

<script>
export default {
  name: 'SearchModal',
  data(){
    return {
      dialog: false
    }
  },
  computed: {
  }
}
</script>
