<template>

        <div class="button">{{ buttonMsg }}</div>
</template>
<style>
    .button {
        text-align: center;
        background: #007FD8;
        font-size: 22px;
        border-radius: 6px;
        border: #007FD8 1px solid;
        color: white;
        line-height: 66px;
        margin-top: 2.8rem;
    }

</style>
<script>
export default {
    props: {
        buttonMsg: {
            type: String,
            require: true
        }
    }
}
</script>