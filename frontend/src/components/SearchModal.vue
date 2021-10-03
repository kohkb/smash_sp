<template>
<div>
  <v-btn icon @click.stop="dialog = true">
    <v-icon>mdi-magnify</v-icon>                
  </v-btn>              
  <v-dialog
      v-model="dialog"
      max-width="290"
      persistent>
    <v-card>
      
      <v-form v-model="valid">
        <v-container>
          <v-text-field
            v-model="query"
            :rules="nameRules"        
            label="検索キーワード"
            required
          ></v-text-field>
        </v-container>
      </v-form>
    </v-card>
    </v-dialog>  
  </div>
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
