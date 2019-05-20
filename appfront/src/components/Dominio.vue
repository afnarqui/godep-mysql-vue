<template>
<div id="app">
  <h1>Dominio vue</h1>
  <p>{{dominio.ipAddress}}</p>
</div>
</template>

<script>
export default {
  name: 'dominio',
  props: [
    'dominio'
  ]
}
</script>
