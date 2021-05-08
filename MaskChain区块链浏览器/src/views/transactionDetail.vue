<template>
  <div>
    <div class="search-main padding-bottom-0">
      <div class="container padding-bottom-0">
        <v-nav :page='page' :hrTitle="'交易'" :navContent="PkHash" :navSubtitle="'交易'" :hrcontent="'交易信息'" :route="'transaction'"></v-nav>
      </div>
    </div>
    <div class="hash-content-info-tran">
      <div class="container">
        <div class="content">
          <el-tabs v-model="activeName" >
            <el-tab-pane :label="transactionName" name="first">
              <div v-if="ID">
                <div class="tranbox">
                  <div v-for="item in transactionData" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
              </div>
              <div v-else>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(0, 2).concat(transactionData.slice(-2, end))" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>接收方和发送方地址公钥的承诺与加密字段</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(2, 8)" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>接收方地址公钥相等证明字段</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(8, 16)" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>发送方地址公钥相等证明字段</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(16, 24)" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>发送，找零，被花费金额的承诺与加密字段</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(24, 30).concat(transactionData.slice(48, 51))" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>发送和找零金额的承诺格式相等证明字段</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(30, 46)" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>接收方公钥加密的发送金额字段</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(46, 48)" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>被花费承诺相等证明字段</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(51, 59)" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>会计平衡证明字段</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(59, 65)" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>发送出的承诺CMS的随机数，接收方公钥加密密文</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(65, 67)" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
                <span style='margin:20px;'>找零承诺CMR的随机数，发送方公钥加密密文</span>
                <div class="tranbox">
                  <div v-for="item in transactionData.slice(67, 69)" class="hash-content-label" :key='item.label'>
                    <span class="label-title">{{item.label}}:</span>
                    <span>{{item.value}}</span>
                  </div>
                </div>
              </div>
            </el-tab-pane>
            <el-tab-pane label="交易回执信息" name="second">
              <div class="tranbox">
                <div v-for="item in receiptData" class="hash-content-label" :key='item.label'>
                  <span class="label-title">{{item.label}}:</span>
                  <span class="receit-content">{{item.value}}</span>
                </div>
                <div class="hash-content-label">
                  <span class="label-title">logs:</span>
                  <div class="label-content">
                    <span class="input-data" v-if="!eventSHow">{{transactionReceiptByPkHash.logs}}</span>
                    <div class="input-data" v-for="item in eventLog" :key='item.address'>
                      <div class="item">
                        <span class="label">Address :</span>
                        <span>{{item.address}}</span>
                      </div>
                      <div class="item">
                        <span class="label">Topics :</span>
                        <div style="display: inline-block;width:800px;">
                          <div v-for="(val,index) in item.topics " :key='val'>[{{index}}] {{val}}</div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.hash-content-info-tran{
  background-color: #2a2c3b;
  color: #fff;
  font-size: 14px;

}
.hash-content-info-tran .box{
  width: 100%;
  padding: 20px 40px;
}
.hash-content-info-tran .content{
  width: 100%;
  border-top: 1px solid #999;
  box-sizing: border-box;
  background-color: #3b3e54;
}
.hash-content-info-tran .content .title{
  border-bottom: 1px solid #999;
  line-height: 40px;
  padding-top: 20px;
  padding-left: 20px;
}
.hash-content-info-tran .pre{
  white-space: pre-wrap!important;
  word-wrap: break-word!important;
  width: calc(100% - 50px)!important;
  font-size: 14px;
  line-height: 24px;
}
.hash-content-label{
  padding: 5px 30px;
}
.hash-content-label>.label-title{
  display: inline-block;
  width: 120px;
  text-align: left;
  padding-right: 10px;
  box-sizing: border-box;
  vertical-align: top;
  color: #111;
}
.hash-content-label>.receit-content{
  display: inline-block;
  width: calc(100% - 300px);
  max-width: 950px;
  white-space: pre-wrap!important;
  word-wrap: break-word!important;
  border: none;
}
.label-content{
  display: inline-block;
  width: calc(100% - 300px);
  max-width: 950px;
  height: auto;
  padding: 10px;
  border: 1px solid #eaedf3;
  border-radius: 4px;
  font-size: 12px;
}
.input-data {
  display: inline-block;
  width: 100%;
  /* padding: 10px; */
  /* max-height: 200px; */
  overflow: auto;
  word-break: break-all;
  word-wrap: break-word;
  box-sizing: border-box;
}
.input-data> .item{
  padding: 5px 0;
  font-size: 12px;
}
.input-data> .item> .label{
  display: inline-block;
  color:#111;
  font-weight: bold;
  width: 80px;
  vertical-align: top;
  font-size: 12px;
  vertical-align: top;
}
.input-label {
  padding-bottom: 15px;
}
.input-label>.label {
  display: inline-block;
  color:#111;
  font-weight: bold;
  width: 80px;
  font-size: 12px;
  vertical-align: top;
}
.detail-input-content {
  display: inline-block;
  height: auto;
  border-radius: 4px;
  font-size: 12px;
}
.tranbox{
  margin: 20px 30px;
  background-color: #fff;
  color: #42b983;
}
.content>>>.is-top{
  color: #fff;
}
.content>>>.is-active{
  color: #409EFF;
}
.content>>>#tab-first{
  padding-left: 15px;
}
.content>>>#tab-second{
  padding-left: 30px;
}
.content>>>.el-tabs__active-bar{
  padding-left: 20px;
}
</style>
<script>
import nav from '@/components/content-nav'
import { getTbTransactionByPkHash, getTbTransactionReceiptByPkHash } from '@/api/api'
import constant from '@/util/constant'

import '@/assets/css/layout.css'
import '@/assets/css/public.css'
export default {
  name: 'transactionDetail',
  components: {
    'v-nav': nav
  },
  data: function () {
    return {
      ID: null,
      PkHash: this.$route.query.pkHash,
      transactionByPkHash: '',
      transactionReceiptByPkHash: '',
      noData: false,
      noDataTransaction: false,
      contractList: [],
      groupId: localStorage.getItem('groupId'),
      transactionData: [],
      transactionBuyData: constant.TRANSACTION_BUY_ITEM_DATA,
      transactionTranferData: constant.TRANSACTION_TRANSFER_ITEM_DATA,
      activeName: 'first',
      transactionName: '交易信息',
      receiptData: constant.RECEIPT_ITEM_DATA,
      byteCode: null,
      transactionTo: null,
      methodId: null,
      buttonSHow: false,
      abiType: [],
      decodeData: null,
      inputData: [],
      funcData: null,
      buttonTitle: '解码',
      showDecode: false,
      eventLog: [],
      eventTitle: '解码',
      eventContent: true,
      eventSHow: false,
      page: {
        pageSize: this.$route.query.pageSize || 10,
        pageNumber: this.$route.query.pageNumber || 1
      }
    }
  },
  mounted: function () {
    this.searchTbTransactionByPkHash()
  },
  methods: {
    searchTbTransactionByPkHash: function () {
      this.noDataTransaction = false
      const data = this.PkHash
      getTbTransactionByPkHash(data).then(res => {
        if (res.data.result) {
          this.noDataTransaction = false
          this.transactionByPkHash = res.data.result
          this.ID = parseInt(res.data.result.ID)
          if (res.data.result.ID === '0x1') {
            this.transactionBuyData.forEach(val => {
              val.value = this.transactionByPkHash[val.label]
            })
            this.transactionName = '购币交易信息'
            this.transactionData = this.transactionBuyData
          } else if (res.data.result.ID === '0x0') {
            this.transactionTranferData.forEach(val => {
              val.value = this.transactionByPkHash[val.label]
            })
            this.transactionName = '转账交易信息'
            this.transactionData = this.transactionTranferData
          }
          this.searchTbTransactionReceiptByPkHash()
        } else {
          this.noData = true
        }
      })
    },
    searchTbTransactionReceiptByPkHash: function () {
      this.noData = false
      const data = this.PkHash
      getTbTransactionReceiptByPkHash(data).then(res => {
        if (res.data.result) {
          this.noData = false
          this.transactionReceiptByPkHash = res.data.result
          if (res.data.result) {
            this.receiptData.forEach(val => {
              val.value = this.transactionReceiptByPkHash[val.label]
            })
          }
          this.eventLog = res.data.result.logs
        } else {
          this.noData = true
        }
      })
    }
  }
}
</script>
