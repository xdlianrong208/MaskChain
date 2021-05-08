<template>
  <div class="home" >
    <div class="container">
      <div class="home-head">
        <!--blockChain Statistics-->
        <div class="home-head-data margin-right-10">
          <ul>
            <li v-for="item in totalStatisticsList" :class="item.class" :key="item.label" class="lg-width">
              <span class="home-head-data-label">{{item.label}}</span>
              <span class="home-head-data-content"> {{item.value}}</span>
            </li>
          </ul>
        </div>
        <div id="myChart2" style="width:50%;height:278px;float:right;" v-loading="loading"
             element-loading-text="数据加载中..."
             element-loading-background="rgba(0, 0, 0, 0.8)"></div>
      </div>
      <div class="home-foot">
        <!--Block list-->
        <div class="home-foot-box margin-right-10">
          <div class="home-foot-box-nav">
            <div class="left">
              <span class="line"></span>
              <span class="text">区块</span>
            </div>
            <div class="right">
              <span @click="goPage('block')" class="table-link" style="padding-right: 30px"> 更多 ></span>
            </div>
          </div>
          <div class="home-foot-box-content">
            <ul>
              <li class="item" v-for="item in blocks" :key='item.number'>
                <div class="left">
                  <div  @click="goPage('blockDetail','hash',item.hash)" class="table-link">区块 {{parseInt(item.number)}}</div>
                  <div>{{item.timestamp}}</div>
                </div>
                <div class="right">
                  <div>矿工
                    <span class="block-number" >{{item.miner}}</span>
                  </div>
                </div>
              </li>
            </ul>
          </div>
        </div>
        <!--transaction list-->
        <div class="home-foot-box margin-left-10">
          <div class="home-foot-box-nav">
            <div class="left">
              <span class="line"></span>
              <span class="text">交易</span>
            </div>
            <div class="right">
              <span @click="goPage('transaction')" class="table-link" style="padding-right: 30px">更多 ></span>
            </div>
          </div>
          <div class="home-foot-box-content" v-loading="loading"
               element-loading-text="数据加载中..."
               element-loading-background="rgba(0, 0, 0, 0.8)">
            <ul>
              <li class="item" v-for="item in transactionsList[0]" :key='item.hash'>
                <div class="left">
                  <div>交易</div>
                  <div class="transaction" @click="goPage( 'transactionDetail','pkHash',item.hash)">
                    <span class="table-link" :title="item.hash">{{item.hash }}</span>
                  </div>
                </div>
                <!--后续添加交易时间戳，暂时无显示-->
                <div class="right">
                  <div>{{item.timestamp}}</div>
                </div>
              </li>
            </ul>
          </div>
        </div>
      </div>
    <router-view/>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import common from '@/common'
import router from '@/router'
import constant from '@/util/constant'
import { getBlockByNumber, getTransactionsCount, getTransactionsForFourteenDays2, getCMCount, getTxpoolPending } from '@/api/api'
import '@/assets/css/layout.css'
import '@/assets/css/public.css'

export default {
  name: 'Home',
  data () {
    return {
      myChart2: '',
      blocks: [],
      transactionsList: [],
      maxBlocks: 15, // 该页面最大显示区块数
      blockNumber: 0, // 最新区块区块高度
      web3: common.web3,
      totalStatisticsList: constant.TOTAL_STATISTICS_LIST,
      chartStatistics: {
        date: [],
        dataArr: []
      },
      setInterval: {
        timer1: null,
        timer2: null,
        timer3: null,
        timer4: null,
        timer5: null
      },
      setSetIntervaling: false,
      loading: false
    }
  },
  mounted: function () {
    this.$nextTick(function () {
      this.searchTbBlockInfo()
      this.searchTbTransactionsCountInfo()
      this.searchTbCMCountInfo()
      this.searchTbPendingCountInfo()
    })
  },
  beforeDestroy: function () {
    this.clear()
  },
  watch: {
    setSetIntervaling: function () {
      this.setSetInterval()
    }
  },
  methods: {
    drawLine () {
      // 基于准备好的dom，初始化echarts实例
      this.myChart2 = this.$echarts.init(document.getElementById('myChart2'), 'shine')
      // 绘制图表
      const fontColor = '#4c9bfd'
      this.myChart2.setOption({
        title: {
          text: '近14天交易数量', // 主标题
          subtext: '', // 副标题
          x: 'center', // x轴方向对齐方式
          textStyle: {
            color: fontColor
          }
        },
        tooltip: {
          // 鼠标悬停提示内容
          trigger: 'axis', // 触发类型，默认数据触发，可选为：'axis' item
          axisPointer: {
            // 坐标轴指示器，坐标轴触发有效
            type: 'line', // 默认为直线，可选为：'line' | 'shadow'
            label: 'cross',
            show: true
          }
        },
        legend: {
          orient: 'vertical',
          x: 'right',
          y: 'top',
          textStyle: {
            color: fontColor
          },
          data: ['交易数量']
        },
        xAxis: {
          type: 'category',
          data: this.chartStatistics.date, // x轴数据
          name: '日期', // x轴名称
          axisLabel: { // x轴间隔和倾斜度
            interval: 2,
            color: '#4c9bfd',
            rotate: 30
          },
          axisLine: {
            lineStyle: {
              color: '#4c9bfd',
              width: 1
            }
          }
        },
        yAxis: {
          type: 'value',
          name: '交易数', // y轴名称
          minInterval: 1,
          axisLabel: {
            color: '#4c9bfd'
          },
          axisLine: {
            lineStyle: {
              color: '#4c9bfd',
              width: 1
            }
          },
          splitLine: {
            lineStyle: {
              color: 'rgba(0, 240, 255, 0.3)'
            }
          }
        },
        series: [
          {
            name: '交易数量',
            data: this.chartStatistics.dataArr,
            symbol: 'circle',
            symbolSize: 5,
            type: 'line',
            itemStyle: {
              normal: {
                lineStyle: {
                  color: '#00FF00'
                },
                areaStyle: {
                  color: new this.$echarts.graphic.LinearGradient(0, 1, 0, 0, [{
                    offset: 0,
                    color: 'rgba(7,44,90,0.3)'
                  }, {
                    offset: 1,
                    color: 'rgba(0,212,199,0.9)'
                  }])
                }
              }
            }
          }
        ]
      })
    },
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
    searchTbBlockInfo: function () {
      this.web3.eth.getBlockNumber()
        .then((result) => {
          this.setSetIntervaling = true
          this.blocks = []
          this.blockNumber = result
          this.totalStatisticsList[0].value = result
          if (this.maxBlocks > result) {
            this.maxBlocks = result + 1
          }
          this.searchBlocksInfo()
          this.searchTransactionsForFourteenDays()
        })
    },
    searchBlocksInfo: function () {
      const promiseArray = []
      for (let i = 0; i < this.maxBlocks; i++) {
        promiseArray.push(getBlockByNumber((this.blockNumber - i).toString(16)))
      }
      Promise.all(promiseArray).then((res) => {
        for (let i = 0; i < this.maxBlocks; i++) {
          this.blocks.push(res[i].data.result)
          this.timeTransport(this.blocks[i])
          // TODO:后续添加对transactionList中的交易进行排序，按照时间顺序;或由后端完成
        }
      })
    },
    searchTbTransactionsCountInfo: function () {
      getTransactionsCount()
        .then((res) => {
          this.totalStatisticsList[1].value = parseInt(res.data.result)
        })
    },
    searchTbPendingCountInfo: function () {
      getTxpoolPending()
        .then((res) => {
          this.totalStatisticsList[2].value = parseInt(res.data.result.pending)
        })
    },
    searchTbCMCountInfo: function () {
      getCMCount()
        .then((res) => {
          this.totalStatisticsList[3].value = parseInt(res.data.result.invalid) + parseInt(res.data.result.valid)
        })
    },
    searchTransactionsForFourteenDays: function () {
      this.loading = true
      getTransactionsForFourteenDays2()
        .then((res) => {
          // res返回前14天每天的交易量
          this.transactionsList = res.data.result
          this.loading = false
          for (let i = 0; i < this.transactionsList.length; i++) {
            if (this.transactionsList[i]) {
              this.chartStatistics.dataArr[this.transactionsList.length - i - 1] = this.transactionsList[i].length
            } else {
              this.chartStatistics.dataArr[this.transactionsList.length - i - 1] = 0
            }
            const NowTime = new Date()
            const NowDay = Math.floor(NowTime.getTime() / 1000)
            const time = NowDay - i * 24 * 60 * 60
            this.chartStatistics.date[this.transactionsList.length - i - 1] = this.format(time * 1000, 'MM-dd')
          }
          this.drawLine()
        }).catch((err) => {
          console.log(err)
          this.loading = false
          this.transactionsList = new Array(14)
          for (let i = 0; i < this.transactionsList.length; i++) {
            const NowTime = new Date()
            const NowDay = Math.floor(NowTime.getTime() / 1000)
            const time = NowDay - i * 24 * 60 * 60
            this.chartStatistics.date[this.transactionsList.length - i - 1] = this.format(time * 1000, 'MM-dd')
          }
          this.drawLine()
        })
    },
    setSetInterval: function () {
      this.setInterval.timer1 = window.setInterval(() => { this.searchTbBlockInfo() }, constant.INTERVALTIME)
      this.setInterval.timer2 = window.setInterval(() => { this.searchTbTransactionsCountInfo() }, constant.INTERVALTIME)
      this.setInterval.timer3 = window.setInterval(() => { this.searchTbCMCountInfo() }, constant.INTERVALTIME)
      this.setInterval.timer4 = window.setInterval(() => { this.searchTbPendingCountInfo() }, constant.INTERVALTIME)
    },
    // clear timer
    clear: function () {
      window.clearInterval(this.setInterval.timer1)
      window.clearInterval(this.setInterval.timer2)
      window.clearInterval(this.setInterval.timer3)
      window.clearInterval(this.setInterval.timer4)
    },
    goPage: function (name, label, data) {
      const path = {}
      path.name = name || ''
      path.query = {}
      if (label) {
        path.query[label] = data
      }
      router.push(path)
    }
  }
}
</script>

<style type="text/css">
#building{
  background:url("../assets/images/header-ng.png");
  position:fixed;
  background-size:100% 100%;}
.home {
  width: 100%;
  background-color: #2a2c3b;
}

.home-head {
  font-size: 0;
  width: 100%;
}

.home-head-data {
  display: inline-block;
  width: calc(50% - 20px);
  box-sizing: border-box;
  padding: 30px;
  background-color: #3b3e54;
  vertical-align: middle;
}

.lg-width {
  width: 47%;
}
.home-head-data ul {
  list-style: none;
  margin: 0;
  padding: 0;
}
.home-head-data ul li {
  display: inline-block;
  height: 110px;
  padding: 25px 15px;
  font-size: 14px;
  cursor: pointer;
  color: #fff;
  box-sizing: border-box;
}

.home-head-chart {
  width: calc(59% - 10px) !important;
  height: 310px;
  vertical-align: middle;
}

.home-head-chart .el-input__inner {
  background-color: #3b3e54;
  color: #fff;
}

.home-head-chart .el-range-input {
  background-color: #3b3e54;
  color: #fff;
}
.el-range-separator {
  color: #fff !important;
}

.home-head-data-label {
  color: white;
}
.home-head-data-content {
  display: block;
  color: white;
  font-size: 38px;
  text-align: right;
}

.home-center {
  margin-top: 20px;
}
.home-center-table {
  background-color: #3b3e54;
}
.home-center .el-table .table-style {
  background-color: #3b3e54;
  color: #fff;
  text-align: center;
  overflow: hidden;
  text-overflow: ellipsis;
  -o-text-overflow: ellipsis;
  -webkit-text-overflow: ellipsis;
  -moz-text-overflow: ellipsis;
  white-space: nowrap;
  border-color: #666;
}
.home-center .el-table .table-style:hover {
  color: #333;
}
.el-table {
  border-color: #666 !important;
}
.el-table__body-wrapper {
  border-color: #666 !important;
}
.el-table table {
  border-color: #666 !important;
}
.el-table tbody {
  border-color: #666 !important;
}
.el-table tr {
  border-color: #666 !important;
}
.el-table th {
  border-color: #666 !important;
}
.el-table td {
  border-color: #666 !important;
}
.el-table--border:after,
.el-table--group:after,
.el-table:before {
  background-color: #666 !important;
}
.home-center .el-table .name-wrapper {
  width: 100%;
  height: 100%;
  display: inline-block;
  /*width: 450px;*/
  overflow: hidden;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}
.home-center .el-table .el-table__empty-block {
  background-color: #3b3e54;
}
.home-foot {
  margin-top: 20px;
  font-size: 0;
}
.home-foot-box {
  display: inline-block;
  width: calc(50% - 10px);
  padding-bottom: 20px;
  background-color: #3b3e54;
  font-size: 14px;
  color: #fff;
  vertical-align: top;
  box-sizing: border-box;
}
.home-foot-box-nav {
  width: 100%;
  height: 50px;
  line-height: 50px;
  margin-left: 15px;
  padding-right: 15px;
  overflow: hidden;
  box-sizing: border-box;
}
.home-foot-box-nav .left {
  float: left;
  width: 50%;
}
.home-foot-box-nav .left .line {
  display: inline-block;
  width: 3px;
  height: 14px;
  margin-right: 12px;
  background-color: #2196f3;
  vertical-align: middle;
}
.home-foot-box-nav .right {
  float: right;
  width: 50%;
  text-align: right;
  cursor: pointer;
}
.home-foot-box-content ul {
  padding: 0;
  margin: 0;
  list-style: none;
}
.home-foot-box-content .item {
  border-bottom: 1px solid #999;
  overflow: hidden;
  line-height: 28px;
  padding: 12px 30px;
}
.home-foot-box-content .item .left {
  float: left;
}
.home-foot-box-content .item .right {
  float: right;
  text-align: right;
  overflow: hidden;
}
.home-foot-box-content .item .block-number {
  display: inline-block;
  width: 116px;
  overflow: hidden;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
  vertical-align: middle;
}
.home-foot-box-content .item .txn {
  float: right;
  width: 70px;
  height: 28px;
  line-height: 28px;
  background-color: #2196f3;
  color: #fff;
  text-align: center;
  cursor: pointer;
}
.home-foot-box-content .item .transaction {
  width: 320px;
  overflow: hidden;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
  vertical-align: middle;
}
.home-foot-box-content .number {
  display: inline-block;
  width: 148px;
  overflow: hidden;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
  vertical-align: middle;
}
.home-foot-box-content .image {
  vertical-align: middle;
}

table {
  border-color: #999 !important;
}
.node-false {
  color: #f00 !important;
}
@media screen and (max-width: 1200px) {
  .home-head-data {
    display: inline-block;
    width: calc(41% - 10px);
    box-sizing: border-box;
    padding: 20px 20px;
    background-color: #3b3e54;
    vertical-align: middle;
  }
  .home-head-data-content {
    display: block;
    color: white;
    padding-top: 10px;
    font-size: 28px;
    text-align: right;
  }
  .home-head-chart {
    width: calc(59% - 10px) !important;
    height: 290px;
    vertical-align: middle;
  }
  .home-foot-box-content .item {
    border-bottom: 1px solid #999;
    overflow: hidden;
    line-height: 28px;
    padding: 12px 20px;
  }
  .home-foot-box-content .item .transaction {
    width: 280px;
    overflow: hidden;
    text-overflow: ellipsis !important;
    white-space: nowrap !important;
    vertical-align: middle;
  }
  .home-foot-box-content .number {
    display: inline-block;
    width: 130px;
    overflow: hidden;
    text-overflow: ellipsis !important;
    white-space: nowrap !important;
    vertical-align: middle;
  }
}
@media screen and (max-width: 1000px) {
  .home-head-data {
    display: inline-block;
    width: calc(41% - 10px);
    box-sizing: border-box;
    padding: 10px 10px;
    background-color: #3b3e54;
    vertical-align: middle;
  }
  .home-head-chart {
    width: calc(59% - 10px) !important;
    height: 270px;
    vertical-align: middle;
  }
  .home-head-data-content {
    display: block;
    color: white;
    padding-top: 10px;
    font-size: 24px;
    text-align: right;
  }
  .home-head-data ul li {
    display: inline-block;
    height: 110px;
    padding: 25px 10px;
    font-size: 14px;
    cursor: pointer;
    color: #fff;
    box-sizing: border-box;
  }
  .home-foot-box-content .item {
    border-bottom: 1px solid #999;
    overflow: hidden;
    line-height: 28px;
    padding: 12px 20px;
  }
  .home-foot-box-content .item .transaction {
    width: 200px;
    overflow: hidden;
    text-overflow: ellipsis !important;
    white-space: nowrap !important;
    vertical-align: middle;
  }
  .home-foot-box-content .number {
    display: inline-block;
    width: 85px;
    overflow: hidden;
    text-overflow: ellipsis !important;
    white-space: nowrap !important;
    vertical-align: middle;
  }
  .home-foot-box-content .item .txn {
    float: right;
    width: 70px;
    height: 28px;
    line-height: 28px;
    background-color: #2196f3;
    color: #fff;
    text-align: center;
    cursor: pointer;
  }
}
</style>
