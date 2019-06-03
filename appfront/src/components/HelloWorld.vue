<template>
<div id="id">
  <form id="main" v-cloak>
    <div class="bar">
        <input type="text" v-model="busc" placeholder="Enter your Search" />
         <br/><br/>
        <input @click="buscar"  type="button" value="Search" class="btn btn-success">
        <br/><br/>
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
    -->

 <b-card-group deck v-if="posts && postsnew.length">
  <b-card header="Data">
    <b-list-group v-for="post of postsnew" v-bind:key="post.id">
      <b-list-group-item href="#"> Host: {{post.host}}</b-list-group-item>
      <b-list-group-item href="#">Port: {{post.port}}</b-list-group-item>
      <b-list-group-item href="#">Is Public: {{post.isPublic}}</b-list-group-item>
      <b-list-group-item href="#">Status: {{post.status}}</b-list-group-item>
      <b-list-group-item href="#">Start Time: {{post.startTime}}</b-list-group-item>
      <b-list-group-item href="#">Test Time: {{post.testTime}}</b-list-group-item>
      <b-list-group-item href="#">Engine Version: {{post.engineVersion}}</b-list-group-item>
      <b-list-group-item href="#">Criteria Version: {{post.criteriaVersion}}</b-list-group-item>
       <p class="card-text mt-2">
        Protocol:  {{post.protocol}}
      </p>
    </b-list-group>
  </b-card>
    <b-card header="Endpoints">
    <ul >
        <li v-for="po of posts.endpoints" v-bind:key="po.id">
           <h4>Server Name: {{po.serverName}}</h4>
           <p>Status Message: {{po.statusMessage}}</p>
           <p>Ip Address: {{po.ipAddress}}</p>
          <p><strong>Address: {{po.address}}</strong></p>
          <p>Grade: {{po.grade}}</p>
          <p><strong>Status Message: {{po.statusMessage}}</strong></p>
          <p><strong>Grade Trust Ignored: {{po.gradeTrustIgnored}}</strong></p>
          <p><strong>Has Warnings: {{po.hasWarnings}}</strong></p>
          <p><strong>Is Exceptional: {{po.isExceptional}}</strong></p>
          <p><strong>Progress: {{po.progress}}</strong></p>
          <p><strong>Duration: {{po.duration}}</strong></p>
          <p><strong>Delegation: {{po.delegation}}</strong></p>
        </li>
    </ul>
  </b-card>

</b-card-group>

</form>
 <div class="bar">
       <h1>Domain with list</h1>
         <br/><br/>
        <input @click="buscarDomains"  type="button" value="todos" class="btn btn-success">
        <br/><br/>
         <b-card-group deck v-if="posts && postdomain.length">
        <b-card header="domain">
          <b-list-group v-for="post of postdomain" v-bind:key="post.id">
            <b-list-group-item href="#"> Host: {{post.host}} Port: {{post.port}}
            <p class="card-text mt-2">
              Protocol:  {{post.protocol}}
            </p>
            </b-list-group-item>
          </b-list-group>

      </b-card-group>
    </div>

   <div class="bar">
       <h1>Domaincomparar with data</h1>
         <br/><br/>
        <input @click="buscarDomainComparars"  type="button" value="comparar" class="btn btn-success">
        <br/><br/>
 <b-card-group deck v-if="posts && postdomaincomparar.length">
  <b-card header="Data">
    <b-list-group v-for="post of postdomaincomparar[0]" v-bind:key="post.id">
      <b-list-group-item href="#"> Host: {{post.host}}</b-list-group-item>
      <b-list-group-item href="#">Port: {{post.port}}</b-list-group-item>
      <b-list-group-item href="#">Is Public: {{post.isPublic}}</b-list-group-item>
      <b-list-group-item href="#">Status: {{post.status}}</b-list-group-item>
      <b-list-group-item href="#">Start Time: {{post.startTime}}</b-list-group-item>
      <b-list-group-item href="#">Test Time: {{post.testTime}}</b-list-group-item>
      <b-list-group-item href="#">Engine Version: {{post.engineVersion}}</b-list-group-item>
      <b-list-group-item href="#">Criteria Version: {{post.criteriaVersion}}</b-list-group-item>
       <p class="card-text mt-2">
        Protocol:  {{post.protocol}}
      </p>
    </b-list-group>
  </b-card>
    <b-card header="Endpoints">
        <b-list-group v-for="post of postdomaincomparar[1]" v-bind:key="post.id">
           <div v-if="post.host.includes('Valor Editado:') === true">
              <b-list-group-item style="background-color:blue;" href="#"> Host: {{post.host}}</b-list-group-item>
          </div>
          <div v-else-if="post.host.includes('Valor Agregado:') === true">
             <b-list-group-item style="background-color:#405d27;" href="#"> Host: {{post.host}}</b-list-group-item>
          </div>
          <div v-else-if="post.host.includes('Valor Eliminado:') === true">
             <b-list-group-item style="background-color:#034f84;" href="#"> Host: {{post.host}}</b-list-group-item>
          </div>
          <div v-else>
            <b-list-group-item href="#"> Host: {{post.host}}</b-list-group-item>
          </div>

           <div v-if="post.port.includes('Valor Editado:') === true">
              <b-list-group-item style="background-color:blue;" href="#"> port: {{post.port}}</b-list-group-item>
          </div>
          <div v-else-if="post.port.includes('Valor Agregado:') === true">
             <b-list-group-item style="background-color:#405d27;" href="#"> port: {{post.port}}</b-list-group-item>
          </div>
          <div v-else-if="post.port.includes('Valor Eliminado:') === true">
             <b-list-group-item style="background-color:#034f84;" href="#"> port: {{post.port}}</b-list-group-item>
          </div>
          <div v-else>
            <b-list-group-item href="#"> port: {{post.port}}</b-list-group-item>
          </div>

           <div v-if="post.status.includes('Valor Editado:') === true">
              <b-list-group-item style="background-color:blue;" href="#"> status: {{post.status}}</b-list-group-item>
          </div>
          <div v-else-if="post.status.includes('Valor Agregado:') === true">
             <b-list-group-item style="background-color:#405d27;" href="#"> status: {{post.status}}</b-list-group-item>
          </div>
          <div v-else-if="post.status.includes('Valor Eliminado:') === true">
             <b-list-group-item style="background-color:#034f84;" href="#"> status: {{post.status}}</b-list-group-item>
          </div>
          <div v-else>
            <b-list-group-item href="#"> status: {{post.status}}</b-list-group-item>
          </div>

      <b-list-group-item href="#">Status: {{post.status}}</b-list-group-item>
      <b-list-group-item href="#">Start Time: {{post.startTime}}</b-list-group-item>
      <b-list-group-item href="#">Test Time: {{post.testTime}}</b-list-group-item>
      <b-list-group-item href="#">Engine Version: {{post.engineVersion}}</b-list-group-item>
      <b-list-group-item href="#">Criteria Version: {{post.criteriaVersion}}</b-list-group-item>
       <p class="card-text mt-2">
        Protocol:  {{post.protocol}}
      </p>
    </b-list-group>
  </b-card>

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
  name: 'app',
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
