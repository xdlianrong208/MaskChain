<template>
<!-- 其实可以用 Tab 实现 -->
    <div>
        <img src="../assets/m.png" class="s">
        <el-card>
            <strong style="position: absolute; right: 10px;font-size: 1.3rem">{{name}}</strong>            
            <span style="font-size: 0.8rem;">积分余额</span>
            <span style="font-size: 1.5rem; ">&emsp;{{nowm}}</span>
        </el-card>
        <el-row type="flex" justify="center">
            <el-col :xs="24" :sm="20" :md="17" :lg="15" :xl="15">
                <el-menu mode="horizontal" class="emenu" background-color="white" active-text-color="#007FD8" text-color="#007FD8">
                    <el-menu-item index="1" @click="con(1)" ref="no">
                            <i class="el-icon-coin"></i>兑换
                        </el-menu-item>
                    <el-menu-item index="2" @click="con(2)"><i class="el-icon-s-promotion"></i>转出</el-menu-item>
                    <el-menu-item index="3" @click="con(3)"><i class="el-icon-receiving"></i>接收</el-menu-item>
                    <el-menu-item index="4" @click="con(4)"><i class="el-icon-time"></i>历史</el-menu-item>
                    <el-menu-item index="5" @click="con(5)"><i class="el-icon-user"></i>我的</el-menu-item>
                </el-menu>
            </el-col>
        </el-row>
    </div>
</template>
<script>
export default {

    data() {
        return {
            nowm: 0,
            name: '',
        }
    },
    methods: {
        con(index) {
            console.log(index);
            // 修改父组件的值
            this.$emit('changecmp', index);    
        },
        changenm(vm) {
            this.nowm = vm;
        },
        changename(name) {
            this.name = name;
        }
    },
    mounted: function () {
        // 初始化点击，不能放在 created 里。因为还没创建出来是拿不到的
        this.$refs.no.$el.click();
    }
}
</script>
<style>
/* 开发者模式拿到的 css 类名 */
    
    .el-menu--horizontal>.el-menu-item {
        text-align: center;
        width: 20%;
    }
    .el-menu-item.is-active {
        background-color: white !important;
    }
    .el-card {
        border: #007FD8 1px solid !important;
        background-color:#007FD8 !important;
        color: white !important;
        text-align: center;
        /* display: flex;
        justify-content: center; */
        border-bottom-left-radius: 16px !important;
        border-bottom-right-radius: 16px !important;
        border-top-left-radius: 0 !important;
        border-top-right-radius: 0 !important;
        box-shadow:  0 2px 4px #007FD8;
    }
    /* .el-card, .el-message {
        border-radius: 0 !important;
    } */
    .s {
        position: absolute;
        height: 4.5rem;
        top: 10px;
        left: 20px;
    }
    .el-card__body {
        line-height: 98px;
        padding: 0 !important;
    }
    .el-menu.el-menu--horizontal {
        border-bottom: none !important;
        margin-bottom: 15px;
        margin-top: 10px;
    }
</style>
