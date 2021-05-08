<template>
<!--  <div class="search-main" style="height: auto;">-->
    <div class="transaction">
<!--      <v-nav :hrTitle="title" :hrcontent="title" :route="'transaction'"></v-nav>-->
      <div class="search-nav">
        <div class="hashInput">
          <el-input placeholder="请输入交易哈希或块高" v-model="searchKeyValue" class="input-with-select">
            <el-button slot="append" icon="el-icon-search" @click="search" :disabled="submitDisabled"></el-button>
          </el-input>
        </div>
      </div>
      <div class="search-table">
        <el-table :data="transactionList"  v-loading="loading"
                  element-loading-text="数据加载中..."
                  element-loading-background="rgba(0, 0, 0, 0.8)">
          <el-table-column prop="hash" label="哈希"  :show-overflow-tooltip="true" align="center" min-width="350px">
            <template slot-scope="scope">
              <span @click="goPage('transactionDetail','pkHash',scope.row.hash)" class="table-link" >{{scope.row.hash}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="blockNumber" label="所属块" align="center" :show-overflow-tooltip="true" min-width="100px">
            <template slot-scope="scope">
              <span @click="goPage('blockDetail','hash',scope.row.blockHash)" class="table-link" >{{parseInt(scope.row.blockNumber)}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="nonce" label="交易块内ID" align="center" :show-overflow-tooltip="true">
            <template slot-scope="scope">{{ parseInt(scope.row.nonce) }}</template>
          </el-table-column>
          <el-table-column prop="from" label="发送者"  :show-overflow-tooltip="true" align="center"></el-table-column>
          <el-table-column   align="center" width="30px">
            <template>
              <img src="@/assets/images/s-right.png">
            </template>
          </el-table-column>
          <el-table-column label="接受者"  :show-overflow-tooltip="true" align="center">
            <template slot-scope="scope">
              <span>{{scope.row.to}}</span>
            </template>
          </el-table-column>
        </el-table>
        <div class="search-pagation">
          <div style="line-height: 40px">
            <span>查询结果 : </span>
            <span>共计{{pagination.total}}条数据</span>
          </div>
        </div>
      </div>
    </div>
<!--  </div>-->
</template>
<script type="es6">
// import nav from '@/components/content-nav'
import { getBlockByNumber, getTbTransactionInfo } from '@/api/api'
import router from '@/router'

import '@/assets/css/layout.css'
import '@/assets/css/public.css'
import common from '@/common'

export default {
  name: 'transaction',
  // components: {
  //   'v-nav': nav
  // },
  data: function () {
    return {
      web3: common.web3,
      searchKeyValue: '',
      blockNum: this.$route.query.blockNum || '',
      title: '交易',
      transactionData: this.$route.query.transactionData || '',
      blockHeight: this.$route.query.blockHeight || '',
      transactionList: [],
      pagination: {
        currentPage: this.$route.query.pageNumber || 1,
        pageSize: this.$route.query.pageSize || 10,
        total: 0
      },
      submitDisabled: false,
      loading: false
    }
  },
  mounted: function () {
    this.searchTbTransactionInfo()
    // this.pagination.currentPage = this.$route.query.pageNumber || 1
    // this.pagination.pageSize = this.$route.query.pageSize || 10
  },
  methods: {
    goPage: function (name, label, data) {
      const path = {}
      path.name = name || ''
      path.query = {}
      if (label) {
        path.query[label] = data
      }
      router.push(path)
    },
    search: function () {
      this.submitDisabled = true
      const reg = /^[0-9]+.?[0-9]*$/
      if (this.searchKeyValue.length > 60) {
        this.blockHeight = ''
        this.transactionData = this.searchKeyValue
        this.searchTransaction()
      } else if (reg.test(this.searchKeyValue) && this.searchKeyValue.substring(0, 2) !== '0x') {
        this.blockHeight = this.searchKeyValue
        this.transactionData = ''
        this.searchTransaction()
      } else if (this.searchKeyValue === '') {
        alert('请输入块高或完整的哈希')
        this.$router.go(0)
      }
      this.searchKeyValue = ''
    },
    searchTbTransactionInfo: function () {
      this.loading = true
      if (this.blockHeight) {
        this.searchTransaction()
      } else {
        this.web3.eth.getBlockNumber()
          .then((result) => {
            this.blockNum = result
            this.searchTransactionsInfo()
          })
      }
    },
    searchTransactionsInfo: function () {
      const num = parseInt(this.blockNum)
      const promiseArray = []
      for (let i = 0; i < num; i++) {
        promiseArray.push(getBlockByNumber((num - i).toString(16)))
      }
      Promise.all(promiseArray).then((res) => {
        for (let i = 0; i < this.blockNum; i++) {
          this.transactionList = this.transactionList.concat(res[i].data.result.transactions)
          this.loading = false
        }
        this.pagination.total = this.transactionList.length
      })
    },
    searchTransaction: function () {
      this.transactionList = []
      getTbTransactionInfo(this.blockHeight, this.transactionData)
        .then((res) => {
          this.submitDisabled = false
          this.loading = false
          if (this.blockHeight) {
            this.transactionList = res.data.result.transactions
            this.pagination.total = this.transactionList.length
          } else if (this.transactionData) {
            this.transactionList.push(res.data.result)
            this.pagination.total = 1
          }
          this.blockHeight = ''
          this.transactionData = ''
        })
    }
  }
}
</script>
