<template lang="html">
<script>
  console.log(item)
  console.log(list)
  </script>
  <tr><td>{{ item.name }}</td></tr>
</template>

<script>
export default {
  name:'monoItem',
  props:{
    item : ''
  }
}
</script>

<style lang="css" scoped>
</style>
