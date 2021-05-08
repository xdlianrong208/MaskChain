<template>
    <p class="backbutton" @click="goback">返回</p>
</template>
<script>
export default {
    methods: {
        goback() {
            this.$router.go(-1)
        }
    }
}
</script>
<style>
    .backbutton {
        text-align: center;
        color: grey;
    }
</style>