<template>
  <div id="o">
    <el-row>
            <!-- 图片也可以换成 element 自带的 -->
            <el-row type="flex" justify="center"><img src="../assets/w.png" style="width: 10rem; height: 10rem"></el-row> 
            <el-row type="flex" justify="center" style="margin-top: 0.5rem;">
                <el-col :xs="10" :sm="8" :md="7.5" :lg="7.5" :xl="7.5">
                    <router-link to="/Newwallet" style="text-decoration: none;">
                        <mybutton :buttonMsg="newbutton"></mybutton>
                    </router-link>
                </el-col>
            </el-row>
            <el-row type="flex" justify="center">
                <el-col :xs="10" :sm="8" :md="7.5" :lg="7.5" :xl="7.5">
                    <router-link to="/Loadwallet" style="text-decoration: none;">
                        <mybutton :buttonMsg="addbutton" class="b1"></mybutton>
                    </router-link>
                </el-col>
            </el-row>
        </el-row>  
    </div>
</template>
<script>
import mybutton from "../components/Mybutton"
export default {
    components: {
        mybutton
    },
    data() {
        return {
            addbutton: "加载账户",
            newbutton: "新建账户"
        }
    },
    created: function () {
        // 浏览器不支持，警告
        if(!window.localStorage){
            this.$alert('浏览器版本过低，请更换浏览器或关闭隐私模式', '警告', {
          confirmButtonText: '确定'});
        }
    }
}
</script>
