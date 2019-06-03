<template>
<div id="id">
 <div class="bar">
         <br/>

         <b-card-group deck v-if="posts && postdomain.length">
        <b-card header="Domain all">
          <b-list-group v-for="post of postdomain" v-bind:key="post.id">
            <b-list-group-item href="#"> Host: {{post.host}} Port: {{post.port}}
            <p class="card-text mt-2">
              Protocol:  {{post.protocol}}
            </p>
            </b-list-group-item>
          </b-list-group>

      </b-card-group>
    </div>
</div>
</template>
<script>
import axios from 'axios';
import getDominios from '../api/info'
import buscardomain from '../api/domain'
import buscardomaincomparar from '../api/domaincomparar'

export default {
  name: 'Domainall',
  data () {
    return {
      posts: [],
      postsnew: [],
      postdomain: [],
      postdomaincomparar: [],
      color: 'blue',
    }
  },
  mounted() {
    this.buscarDomains()
  },
  methods: {
    buscar: function () {
      console.log('entro a buscar')
        const self = this
        const url = `http://localhost:8081/public?nombre=${self.busc}`
        getDominios(url)
        .then( function(dominios) {
          let data = JSON.stringify(dominios)
          let datanew = `[${data}]`
          self.postsnew = JSON.parse(datanew)
          console.log(self.postsnew)
          console.log(JSON.parse(data))
          self.posts = JSON.parse(data)
        })
    .catch((e) => {
      console.error(e)
    })
    },
    buscarDomains: function () {
        console.log('entro a buscardomain')
        const self = this
        const url = `http://localhost:8081/buscardomain`
    buscardomain(url)
    .then( function(buscar){
          self.postdomain = buscar

    }).catch((e)=> {
      console.log(e)
    })
    },
    buscarDomainComparars: function () {
      const self = this
      const url = `http://localhost:8081/buscardomaincomparar`
    buscardomaincomparar(url)
    .then( function(buscar){
          self.postdomaincomparar = buscar
    }).catch((e)=> {
      console.log(e)
    })

    }
  }
}
</script>
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
