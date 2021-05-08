<template>
    <div id="o" class="main">
        <!-- 输入信息，需要用 el-col 实现响应式布局 -->
        <div class="imfo">
            <p style="margin-top: 0;">我们将通过您的个人信息为您生成账户公钥并在本地储存相关信息</p>
            <p style="font-size: 17px; color: #666666;">为保证您的账户安全，并让您在多个设备上打开您的账户，请及时备份并安全保管相关文件</p>
<!--                    <p>文件储存地址为.....</p>-->
<!--                    <p>文件名为.....</p>-->
        </div>
        <el-row type="flex" justify="center">
            <el-col :xs="20" :sm="15" :md="8" :lg="8" :xl="8"  @click="register">                
                <p><span class = "t"></span>姓名：</p>
                <el-input maxlength="12" v-model="name" minlength="1"></el-input>
                <p><span class = "t"></span>身份证号：</p>
                <el-input maxlength="18" minlength="18" v-model="id"></el-input>
                <p><span class = "t"></span>自定义字符串：</p>
                <el-input maxlength="255" v-model="string" minlength="1"></el-input>
                <mybutton :buttonMsg="bm" @click.native="register">创建账户</mybutton>
            </el-col>
        </el-row>
        <backbutton></backbutton>
    </div>
</template>

<script>
// @ is an alias to /src
import mybutton from '../components/Mybutton.vue'
import backbutton from '../components/Backbutton.vue'

export default {
    components: {
        mybutton,
        backbutton
    },
    data() {
        return {
            bm: '创建账户',
            id: '',
            name: '',
            string: ''
        }
    },
    mounted: function() {
        // 小屏适配
        if(document.documentElement.clientHeight < 870) {
            console.log(document.documentElement.clientHeight);
            document.getElementById("o").style.top = "10px";
            document.getElementById("o").style.transform = "none";
        }
    },
    methods: {
        Account(imfo) {
            this.imfo = imfo;
            this.history = new Array();
        },
        register() {
            if (this.id == '' || this.name == '' || this.string == '') {
                this.$message.error ('提交的信息不能为空');
            } else {
                this.axios.post('http://39.105.58.136:4396/wallet/register', {
                    name: this.name,
                    id: this.id,
                    str: this.string
                }).then((response)=>{
                    console.log(response);
                    this.$message.success({
                        message: '创建成功',
                        duration: 1500
                    }); 
                    // 创建成功后加入 localstorage
                    var storage = window.localStorage;
                    response.data;
                    storage.setItem(this.name, JSON.stringify(new this.Account(response.data)));
                    console.log(storage);
                    // 跳转
                    setTimeout(() => {
                        this.$router.push({
                            name: 'Mainaction',
                            path: '/Mainaction',
                            query: {
                                account: this.name,
                            }
                        })
                    }, 1500);
                }).catch((response)=>{
                    this.$message.error('创建失败，请重试');
                    console.log(response);
                })
            }
        }
    }
}
</script>
<style>
    #o {
        position: relative;
        top: 45%;
        transform: translateY(-50%);
    }   
    .imfo {
        text-align: center; 
        margin-bottom: 34px;
    }
    .imfo p {
        margin-top: 19px;   
    }
    .b1 {
        background: white;
        color: #007FD8;
    }
    
    .t {
        width: 5px;
        height: 36px;
        background: #007FD8;
        border-radius: 5px;
        margin-right: 13px;
        display: inline-block;
        vertical-align: middle; /* 设置 inline-block 元素对齐基准线 */
        /* transform: translateY(50%); */
    }
    .el-input {
        margin-bottom: 20px;        
        font-size: 17px !important;
    }
    .el-input__inner {
        height: 55px !important;
    }
</style>
<style scoped>
    .el-col p {
        font-size: 20px;
        margin-top: 35px;
    }
</style>