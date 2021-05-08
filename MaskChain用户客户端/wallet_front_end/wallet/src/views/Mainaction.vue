<template>
    <div style="height: 100%;">
        <navmenu @changecmp="changecmps"  ref="n"></navmenu>  
        <el-row type="flex" justify="center" id="om">
            <el-col :xs="20" :sm="15" :md="8" :lg="8" :xl="8" v-show="cmp != 4">  
            <div v-show="cmp == 1">
                <p>购买数额:</p>
                <el-input maxlength="10" v-model="money" ></el-input>
                <mybutton :buttonMsg="buy" @click.native="buym"></mybutton>
            </div>

            <div v-show="cmp == 2">
                <p><span class = "t"></span>接收方公钥:</p>
                <el-input v-model="G1" placeholder="G1"></el-input>
                <el-input v-model="G2" placeholder="G2" style="margin-top:10px;"></el-input>
                <el-input v-model="P" placeholder="P" style="margin-top:10px;"></el-input>
                <el-input v-model="pub" placeholder="pub" style="margin-top:10px;"></el-input>
                <div style="margin-top:10px;">
                    <a>或选择本地账户&emsp;</a>
                    <el-select v-model="baccount" placeholder="请选择" clearable>
                        <el-option v-for="it in accountList" :value="it" :key="it.key" :label="it"></el-option>
                    </el-select>
                </div>
                <p><span class = "t"></span>转出数额:</p>
                <el-input maxlength="10" v-model="spend" ></el-input>
                <p><span class = "t"></span>使用承诺的数额</p>
                <el-input maxlength="10" v-model="transmoney" ></el-input>
                <p><span class = "t"></span>承诺cmv</p>
                <el-input v-model="moneyProm" ></el-input>
                <p><span class = "t"></span>随机数vor</p>
                <el-input v-model="r" ></el-input>
                <!-- 上面这些够了，可以返回东西了 -->
                <mybutton :buttonMsg="transfer" @click.native="transferm" style="margin-bottom: 20px"></mybutton>
            </div>

            <div v-show="cmp == 3">
                 <p>交易hash</p>
                <el-input v-model="hash" ></el-input>
                <mybutton :buttonMsg="recv" @click.native="recm"></mybutton>
            </div>
            
            <div v-show="cmp == 5">
                <mybutton :buttonMsg="showImfo" @click.native="showImfof" class="b1"></mybutton>
                <mybutton :buttonMsg="signout" @click.native="signoutf"></mybutton>
            </div>
            </el-col>
            
        </el-row>
        <el-row v-show="cmp == 4" type="flex" justify="center">
            <el-col :xs="24" :sm="20" :md="17" :lg="15" :xl="15" >
                <el-table  :data="hisList" >
                    <el-table-column
                        prop="amount"
                        label="数额">
                    </el-table-column>
                    <el-table-column
                        prop="hash"
                        label="哈希hash">
                    </el-table-column>
                    <el-table-column
                        prop="cmv"
                        label="承诺cmv">
                    </el-table-column>
                    <el-table-column
                        prop="vor"
                        label="随机数vor">
                    </el-table-column>
                </el-table>
            </el-col>    
        </el-row>
    </div>
</template>
<script>
import navmenu from '../components/Navmenu'
import mybutton from '../components/Mybutton'
var account;
var accountList = new Array();
export default {
    components: {
        navmenu,
        mybutton,
    },
    data() {
        return {
            transfer: '发起转出',
            buy: '兑换',
            recv: '接收',
            signout: '登出',
            showImfo: '显示账户信息',
            money: '',
            cmp: '1', // 用来改变显示的组件
            transmoney: '',
            moneyProm: '',
            r: '',
            G1: '',
            G2: '',
            P: '',
            pub: '',
            hash: '',
            hisList: '',
            spend: '',
            nowm: '',
            accountList,
            baccount: ''
        }
    },
    created: function () {
        account = this.$route.query.account;
        if (account == undefined) {
            this.$message.error({
                message: '请登录账户',
                duration: 1400
            }); 
            setTimeout(() => {
                this.$router.push({
                    path: '/',
                    name: 'Main',
                })
            }, 1500);   
        }
        this.hisList = JSON.parse(window.localStorage.getItem(account)).history;
    },
    mounted: function () {
        this.showCoin();
        // 获取本地账户列表
        for(var i = 0; i < window.localStorage.length; i++) {
            var name = window.localStorage.key(i);
            if (name != 'loglevel:webpack-dev-server' && name != account) {
                accountList.push(name);
            }
        }
        this.$refs.n.changename(account);     
    },
    watch: {
        baccount(val) {
            if(val == '') {
                this.G1 = '';
                this.G2 = '';
                this.P = '';
                this.pub = '';
            } else {
                var b = JSON.parse(window.localStorage.getItem(val)).imfo;
                console.log(b);
                this.G1 = b.G1;
                this.G2 = b.G2;
                this.P = b.P;
                this.pub = b.publickey;
            }
        }
    },
    methods: {
        getPri() {
            var pri = JSON.parse(window.localStorage.getItem(account)).imfo;
            return pri;
        },
        storeImfo(response, amount) {
            // 更新信息
            // 取出 history 并修改
            console.log("?");
            var old = JSON.parse(window.localStorage.getItem(account));
            var neww = response.data;
            neww.vm = amount;
            old.history.push(neww); // 喜加一
            console.log(neww);
            window.localStorage.setItem(account, JSON.stringify(old));
            console.log(window.localStorage.getItem(account));
            this.showCoin();
        },
        Pub(G1, G2, P, H) {
            this.G1 = G1;
            this.G2 = G2;
            this.P = P;
            this.H = H;
        },
        transferm() {
            console.log("我要转账");
            var pri = this.getPri();
            this.$message('正在生成：会计平衡证明、监管相等证明、范围证明、密文格式正确证明');
            this.axios.post('http://192.168.0.104:4396/wallet/exchange',{
                    sg1: pri.G1,
                    sg2: pri.G2,
                    sp: pri.P,
                    sh: pri.publickey,
                    sx:pri.privatekey,
                    amount: this.transmoney,
                    rg1: this.G1,
                    rg2: this.G2,
                    rp: this.P,
                    rh: this.pub,
                    cmv: this.moneyProm,
                    vor: this.r,
                    spend: this.spend
            }).then((response)=>{
                this.storeImfo(response, -this.spend);
            }).catch((response)=>{
                    this.$message.error(response);
                    console.log(response);
            });
            
        },
        buym() {
            console.log("我要购币");
            var pri = this.getPri();
            console.log(pri.privatekey);
            this.axios({
                url: 'http://192.168.0.104:4396/wallet/buycoin',
                method: 'post',
                data: {
                    g1: pri.G1,
                    g2: pri.G2,
                    p: pri.P,
                    h: pri.publickey,
                    x: pri.privatekey,
                    amount: this.money
                    },
                timeout: '600000'
            }).then((response)=>{
                this.storeImfo(response, this.money);
            }).catch((response)=>{
                    this.$message.error(response);
                    console.log(response);
            });
            this.$message.success({
                        message: '金额加密正确',
                        duration: 1000
                    });
            setTimeout(() => {
                this.$message.success({
                        message: '公钥加密正确',
                        duration: 1000
                    }); 
            }, 2000);
        },
        recm() {
            console.log("我要收款");
            var pri = this.getPri();
            this.axios({
                url: 'http://192.168.0.104:4396/wallet/receive',
                method: 'post',
                data: {
                    g1: pri.G1,
                    g2: pri.G2,
                    p: pri.P,
                    h: pri.publickey,
                    x:pri.privatekey,
                    hash: this.hash
                } ,
                timeout: '600000'
            }).then((response)=>{
                response.data.amount = parseInt(response.data.amount);
                this.storeImfo(response, parseInt(response.data.amount));
            }).catch((response)=>{
                    this.$message.error(response);
                    console.log(response);
            });
        },
        changecmps(index) {
            this.cmp = index;
            // 防止小手机转账界面崩坏
            if (this.cmp == 2) {
                document.getElementById("om").style.top = "10px";
                document.getElementById("om").style.transform = "none";
            } else {
                document.getElementById("om").style.top = "30%";
                document.getElementById("om").style.transform = "translateY(-50%)";
            }
            // 加载历史
            if (this.cmp == 4) {
                // 更新
                this.hisList = JSON.parse(window.localStorage.getItem(account)).history;               
            }
        },
        signoutf() {
            account = undefined;
            accountList.length = 0;
            this.$router.push({
                path: '/'
            })
        },
        showImfof() {
            this.showCoin();
            var G1 = JSON.stringify((JSON.parse(window.localStorage.getItem(account))).imfo.G1);
            var G2 = JSON.stringify((JSON.parse(window.localStorage.getItem(account))).imfo.G2);
            var P = JSON.stringify((JSON.parse(window.localStorage.getItem(account))).imfo.P);
            var pub = JSON.stringify((JSON.parse(window.localStorage.getItem(account))).imfo.publickey);
            var pri = JSON.stringify((JSON.parse(window.localStorage.getItem(account))).imfo.privatekey);
            this.$alert("<p>G1:" + G1 + "</p>" +
                "<p>G2:" + G2 + "</p>" +
                "<p>P:" + P + "</p>" +
                "<p>pub:" + pub + "</p>" +
                "<p>pri:" + pri + "</p>", {
                confirmButtonText: '确定',
                dangerouslyUseHTMLString: true,
                customClass:'message_box_alert'
            });
                        window.localStorage.setItem(2,window.localStorage.getItem("1"));

            // var twqee = {
            //     hash: "ASBWJAKFA",
            //     cmv: "SDUIFUISAASK",
            //     r: "DSFSAFSA",
            //     amount: 100,
            //     vm: 100
            // };
            // var old = JSON.parse(window.localStorage.getItem(account));            old.history.push(twqee); // 喜加一
            // window.localStorage.setItem(account, JSON.stringify(old));
            // console.log(old);
        },
        showCoin(){
            // 刷新余额
            console.log("改余额");
            var sum = 0;
            var his = JSON.parse(window.localStorage.getItem(account)).history;
            for (var i = 0; i < his.length; i++){
                if(his[i].vm != undefined){
                    sum = sum + parseInt(his[i].vm);
                    console.log(sum);
                }
            }
            // 上当了,因为少了个s找了好久问题
            this.$refs.n.changenm(sum);    
        }
    }
}
</script>
<style>
.message_box_alert {
    word-break: break-all !important;
}
.el-message-box{
    width: 80%;
}
#om {
    position: relative;
    top: 35%;
    transform: translateY(-50%);
}  
.el-table td, .el-table th {
    text-align: center !important;
}
</style>
<style scoped>
.el-col p {
    margin-top: 25px !important;
}
</style>