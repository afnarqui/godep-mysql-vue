<template>

<div id="id">
  <form id="main" v-cloak>

    <div class="bar">
        <!-- Create a binding between the searchString model and the text field -->

        <input type="text" v-model="searchString" placeholder="Enter your Search" />
        <br>
        <input @click="buscar" v-model="busc" type="button" value="Search" class="btn btn-success">
    </div>
<!--
  "host": "www.google.com",
  "port": 443,
  "protocol": "http",
  "isPublic": false,
  "status": "READY",
  "startTime": 1558411159030,
  "testTime": 1558411584928,
  "engineVersion": "1.34.2",
  "criteriaVersion": "2009p",
"endpoints": [
    {
      "ipAddress": "2607:f8b0:4005:809:0:0:0:2004",
      "serverName": "sfo03s08-in-x04.1e100.net",
      "statusMessage": "Ready",
      "grade": "A+",
      "gradeTrustIgnored": "A+",
      "hasWarnings": false,
      "isExceptional": true,
      "progress": 100,
      "duration": 331020,
      "delegation": 2
    },
     <ul v-if="posts && posts.length" >
        <li v-for="post of posts" v-bind:key="post.id">
           <h2></h2>
           <p><strong>{{post.address}}</strong></p>
    <p>{{post.grade}}</p>
        </li>
    </ul>



<b-card-group deck v-for="post of posts" v-bind:key="post.id">
  <b-card header="Data">
    <b-list-group>
      <b-list-group-item href="#">{{post.host}}</b-list-group-item>
      <b-list-group-item href="#">{{post.port}}</b-list-group-item>

    </b-list-group>
      <p class="card-text mt-2">
         {{post.status}}
      </p>

      <strong>{{post.protocol}}</strong>

  </b-card>
   <ul >
        <li v-for="po of post.endpoints" v-bind:key="po.id">
           <h1>{{po.serverName}}</h1>
           <h2>{{po.statusMessage}}</h2>

           <p><strong>{{po.address}}</strong></p>
    <p>{{po.grade}}</p>

        </li>
    </ul>
</b-card-group>
-->
<b-card-group deck v-for="post of posts" v-bind:key="post.id">
  <b-card header="Data">
    <b-list-group>
      <b-list-group-item href="#">{{post.address}}</b-list-group-item>
      <b-list-group-item href="#">{{post.grade}}</b-list-group-item>
      <b-list-group-item href="#">{{post.statusMessage}}</b-list-group-item>

    </b-list-group>
      <p class="card-text mt-2">
         {{post.serverName}}
      </p>

  </b-card>
</b-card-group>

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
      //this.posts = response.data.endpoints
      this.posts = response.data.endpoints
    })
    .catch((e) => {
      console.error(e)
    })
  },
  methods: {
    buscar: function () {
      console.log('entro a buscar')
          const self = this
          url = `http://localhost:8081/buscar?${self.busc}`
          axios.get(url).then((response) => {
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
