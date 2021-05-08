<template>
  <div class="block">
    <!--搜索框-->
    <div class="search-nav">
      <div class="hashInput">
        <el-input placeholder="请输入区块哈希或块高" v-model="input" class="input-with-select">
          <el-button slot="append" icon="el-icon-search" @click="search" :disabled="submitDisabled"></el-button>
        </el-input>
      </div>
    </div>
    <div class="search-pagation">
      <el-pagination style="display: inline-block"
                     layout="sizes,prev, pager, next"
                     :total="pagination.total"
                     @size-change="handleSizeChange"
                     @current-change="handleCurrentChange"
                     :current-page.sync="pagination.currentPage"
                     :page-sizes="[10, 20, 30, 50]"
                     :page-size="pagination.pageSize">
      </el-pagination>
    </div>
    <!--搜索结果，默认显示全部区块，按高度从高到到低排序-->
    <div class="search-table">
      <el-table :data="blockList"
                v-loading="loading"
                element-loading-text="数据加载中..."
                element-loading-background="rgba(0, 0, 0, 0.8)">
        <el-table-column  prop="number" label="块高" align="center" :class-name="'table-link'" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <span @click="goPage('blockDetail','hash',scope.row.hash)">{{parseInt(scope.row.number)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="timestamp" label="生成时间" min-width="120px" align="center" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column prop="transactions" label="交易数量" align="center" :class-name="'table-link'" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <span @click="goPage('transaction','blockHeight',scope.row.number)">{{scope.row.transactions.length}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="miner" label="出块者" min-width="100px" :show-overflow-tooltip="true" align="center"></el-table-column>
        <el-table-column prop="hash" label="哈希" min-width="350px" :show-overflow-tooltip="true" align="center" :class-name="'table-link'">
          <template slot-scope="scope">
            <span @click="goPage('blockDetail','hash',scope.row.hash)">{{scope.row.hash}}</span>
          </template>
        </el-table-column>
      </el-table>
      <!--查询结果总结，后续添加换页功能-->
      <div class="search-pagation">
        <div style="line-height: 40px;">
          <span>查询结果 : </span>
          <span>共计{{pagination.total}}条数据</span>
        </div>
<!--        <el-pagination style="display: inline-block"-->
<!--                       layout="sizes,prev, pager, next"-->
<!--                       :total="pagination.total"-->
<!--                       @size-change="handleSizeChange"-->
<!--                       @current-change="handleCurrentChange"-->
<!--                       :current-page.sync="pagination.currentPage"-->
<!--                       :page-sizes="[10, 20, 30, 50]"-->
<!--                       :page-size="pagination.pageSize">-->
<!--        </el-pagination>-->
      </div>
    </div>
  </div>
</template>

<script>
import '@/assets/css/layout.css'
import '@/assets/css/public.css'
import common from '@/common'
import { getBlockByNumber, getBlockByHash } from '@/api/api'
import router from '@/router'

export default {
  name: 'block',
  data () {
    return {
      input: '',
      hashData: null,
      blockNumber: null,
      pagination: {
        currentPage: this.$route.query.pageNumber || 1,
        pageSize: this.$route.query.pageSize || 20,
        total: 0
      },
      web3: common.web3,
      maxBlocks: 0,
      totalBlockNumber: 0,
      blockList: [],
      submitDisabled: false,
      loading: false
    }
  },
  mounted: function () {
    this.searchTbBlockInfo()
  },
  methods: {
    format: function (time, format) {
      const t = new Date(time)
      const tf = function (i) {
        return (i < 10 ? '0' : '') + i
      }
      return format.replace(/yyyy|MM|dd|HH|mm|ss/g, function (a) {
        switch (a) {
          case 'yyyy':
            return tf(t.getFullYear())
          case 'MM':
            return tf(t.getMonth() + 1)
          case 'mm':
            return tf(t.getMinutes())
          case 'dd':
            return tf(t.getDate())
          case 'HH':
            return tf(t.getHours())
          case 'ss':
            return tf(t.getSeconds())
        }
      })
    },
    timeTransport: function (block) {
      const time = parseInt(block.timestamp, 16) * 1000
      block.timestamp = this.format(time, 'yyyy-MM-dd HH:mm:ss')
    },
    search: function () {
      const reg = /^[0-9]+.?[0-9]*$/
      if (this.input.length > 60) {
        this.hashData = this.input
        this.blockNumber = ''
        this.searchBlock(this.hashData, 2)
      } else if (reg.test(this.input) && this.input.substring(0, 2) !== '0x') {
        this.hashData = ''
        this.blockNumber = this.input
        this.searchBlock(this.blockNumber, 1)
      } else if (this.input === '') {
        alert('请输入块高或完整的哈希')
        this.$router.go(0)
      }
      this.input = ''
    },
    searchBlock: function (data, type) {
      if (type === 2) {
        getBlockByHash(this.input)
          .then((res) => {
            this.pagination.total = 1
            console.log(res.data.result)
            this.blockList = []
            this.blockList.push(res.data.result)
            // 根据hash查找区块默认只能找到一个区块
            this.timeTransport(this.blockList[0])
          })
      } else if (type === 1) {
        getBlockByNumber(parseInt(this.input).toString(16))
          .then((res) => {
            this.pagination.total = 1
            this.blockList = []
            this.blockList.push(res.data.result)
            for (let i = 0; i < this.blockList.length; i++) {
              this.timeTransport(this.blockList[i])
            }
          })
      }
    },
    searchTbBlockInfo: function () {
      this.loading = true
      this.web3.eth.getBlockNumber()
        .then((result) => {
          this.pagination.total = result
          this.totalBlockNumber = result
          this.searchBlocksInfo()
        })
    },
    searchBlocksInfo: function () {
      this.maxBlocks = this.pagination.pageSize
      const pageSize = this.pagination.pageSize
      const currentPage = this.pagination.currentPage
      if (this.maxBlocks > this.totalBlockNumber - pageSize * (currentPage - 1)) {
        this.maxBlocks = this.totalBlockNumber - pageSize * (currentPage - 1) + 1
      }
      const promiseArray = []
      this.blockList = []
      for (let i = 0; i < this.maxBlocks; i++) {
        promiseArray.push(getBlockByNumber((this.totalBlockNumber - i - pageSize * (currentPage - 1)).toString(16)))
      }
      Promise.all(promiseArray).then((res) => {
        for (let i = 0; i < this.maxBlocks; i++) {
          this.blockList.push(res[i].data.result)
          this.timeTransport(this.blockList[i])
        }
        this.loading = false
      })
    },
    goPage: function (name, label, data) {
      const path = {}
      path.name = name || ''
      path.query = {}
      if (label) {
        path.query[label] = data
      }
      router.push(path)
    },
    handleSizeChange (val) {
      this.pagination.pageSize = val
      this.maxBlocks = this.pagination.pageSize
      this.pagination.currentPage = 1
      router.push({
        // name: "block",
        query: {
          pageSize: this.pagination.pageSize,
          pageNumber: this.pagination.currentPage
        }
      })
      this.searchBlocksInfo()
    },
    handleCurrentChange (val) {
      this.pagination.currentPage = val
      router.push({
        // name: "block",
        query: {
          pageSize: this.pagination.pageSize,
          pageNumber: this.pagination.currentPage
        }
      })
      this.searchBlocksInfo()
    }
  }
}
</script>
