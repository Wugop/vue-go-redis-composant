<template>
  <input v-model="message" placeholder="Entrer un mot">
  <p><button @click="addWord(message)">Ajouter un mot</button></p>

  <p>Les messages entrées sont :</p>
  <ul v-for="item,       i in items" :key="item.key">
    <li>{{ item.key + ' : ' + item.value }}</li>
    <li>
      <p><button v-on:click="removeWord(i,item.key)">Supprimer un mot</button></p>
    </li>

  </ul>
</template>

<script>
import axios from 'axios'

// Must be variable-ized and configured with a environment variable (cf. twelve factors)
const API_BASE_URL = 'http://localhost:8085/sha256';
export default {
  name: 'VueHash',
  data() {
    return {
      id: 0,
      message: "",
      items: []
    }
  },
  props: {

  },
  mounted() {
    axios.get(API_BASE_URL).then(response => {
      if(response.data != "") {
        if(response.data.toString().includes("-")) {
          var temp = response.data.split("-") 
          temp.forEach(e => {
            axios.get(API_BASE_URL + '/' + e).then(response2 => {
            var a = { 'key': e, 'value': response2.data }
            this.items.push(a)
          })
        })
        }
      else {
        axios.get(API_BASE_URL + '/' + response.data).then(response2 => {
        var a = { 'key': response.data, 'value': response2.data }
        this.items.push(a)})
      }
    }
    })
  },

  methods: {
    removeWord(i,key) {
      this.items.splice(i, 1)
      axios.delete(API_BASE_URL+"/"+key).then(response => (console.log(response)))
    },
    // for some reason some strings can't be added like : \n \d % and... there is no log :-(
    addWord(w) {
      axios.post(API_BASE_URL, { 'Key': w }).then(response => {
        console.log(response)
        axios.get(API_BASE_URL + "/" + w).then(response => {
        var a = { 'key': w, 'value': response.data }
        this.items.push(a)
      })})
    }
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
