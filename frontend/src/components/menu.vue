<template lang="html">
<div v-if:"menuMode">
×<br>
<span>ログアウト</span><br>
<span>ヘルプ</span><br>
<span>このソフトについて</span><br>
</div>
<div v-else>
≡
</div>
</template>

<script>
export default {
data:{
menuMode:false
}
}
</script>

<style lang="css" scoped>
</style>
