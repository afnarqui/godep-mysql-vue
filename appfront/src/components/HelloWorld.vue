<template>

<div id="id">
  <form id="main" v-cloak>

    <div class="bar">
        <!-- Create a binding between the searchString model and the text field -->

        <input type="text" v-model="searchString" placeholder="Enter your search" />
    </div>

    <ul v-if="posts && posts.length" >
        <!-- Render a li element for every entry in the computed filteredArticles array. -->
        <!--https://api.ssllabs.com/api/v3/analyze?host=google.com-->
        <li v-for="post of posts" v-bind:key="post.id">
           <p><strong>{{post.address}}</strong></p>
    <p>{{post.grade}}</p>
        </li>
    </ul>
<input @click="buscar" type="button" value="Añadir" class="btn btn-success">
</form>



</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'app',
  data () {
    return {
      posts: []
    }
  },
  mounted() {
    axios.get('http://localhost:8081/public').then((response) => {
      console.log(response)
      this.posts = response.data.endpoints
    })
    .catch((e) => {
      console.error(e)
    })
  },
  methods: {
    buscar: function () {
      console.log('entro a buscar')
          axios.get('http://localhost:8081/public').then((response) => {
      console.log(response.data.endpoints)
      this.posts = response.data.endpoints
    })
    .catch((e) => {
      console.error(e)
    })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
