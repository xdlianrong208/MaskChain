<template>
    <div>
        <el-row type="flex" justify="center">
            <el-col :xs="20" :sm="15" :md="14.5" :lg="14.5" :xl="14.5">
                <ul class="infinite-list"  style="overflow:auto;padding:0;">
                    <li v-for="name in list" :key="name.id" class="infinite-list-item" @click="signin(name)">
                        <span>{{ name }}</span>
                        <i id="gan" class="el-icon-delete"></i>
                    </li>
                </ul>
                </el-col>
        </el-row>
        <backbutton></backbutton>
    </div>
</template>
<script>
import backbutton from '../components/Backbutton'
var list = new Array(); 
export default {    
    components: {
        backbutton
    },
    data () {
        return {
            list
        }
    },
    mounted: function () {
        // 填充 list 数组        
        var storage = window.localStorage;
        // 防止再次填充 list，清空数组
        list.length = 0;
        for(var i = 0; i < storage.length; i++) {
            if (storage.key(i) != 'loglevel:webpack-dev-server') {
                list.push(storage.key(i));
            }
        }        
    },
    methods: {
        signin (name) {
            this.$router.push({
                name: 'Mainaction', // 没有这句会 undefined
                path: '/Mainaction',
                query: {
                    account: name
                }
            })
        }
    }
}
</script>
<style>
    .infinite-list .infinite-list-item {
        /* 设置 flex float 会失效，子元素的相对定位是相对窗口？ */
        /* 设置行高会垂直居中， height 不行 */
        line-height: 56px;
        background: #e8f3fe;
        margin: 10px;
        color: #7dbcfc;
        justify-content: center;
        text-align: center;
        position: relative;
        border-radius: 6px;
        font-size: 18px;
    }
    .el-icon-delete {
        position: absolute;
        right: 5px;
        top: 50%;
        transform: translateY(-50%);
    }
</style>