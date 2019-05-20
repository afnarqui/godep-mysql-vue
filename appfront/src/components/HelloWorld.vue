<template>
<ul v-if="posts && posts.length">
  <li v-for="post of posts" v-bind:key="post.id">
    <p><strong>{{post.title}}</strong></p>
    <p>{{post.description}}</p>
  </li>
</ul>
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
  created() {
    axios.get('http://localhost/notes').then((response) => {
      this.posts = response
    })
    .catch((e) => {
      console.error(e)
    })
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
