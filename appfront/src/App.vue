<template>
  <div id="app">
    <!--<img src="./assets/logo.png">-->
    <ul>
      <Dominio v-for="dominio in dominios" v-bind:key="dominio.ipAddress"/>
    </ul>
    <HelloWorld/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld'
import Dominio from './components/Dominio'
import getDominios from './api'

export default {
  name: 'App',
    data () {
    return {
      dominios: [],
      selectedDominio: 'www.google.com'
    }
  },
  components: {
    HelloWorld,
    Dominio
  },
  methods: {
    refreshDominio() {
      const self = this
      getDominios(this.selectedDominio)
        .then( function(dominios) {
          self.dominios = dominios
        })
    }
  },
  mounted: function () {
    this.refreshDominio()
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
