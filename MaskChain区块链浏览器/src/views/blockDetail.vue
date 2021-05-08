<template>
  <div>
    <div class="search-main padding-bottom-0">
      <div class="container padding-bottom-0">
        <v-nav :page='page' :hrTitle="'区块'" :navContent="blockName" :navSubtitle="'区块'" :hrcontent="'区块信息'" :route="'block'"></v-nav>
          <div class="detail-content">
          <div class="c-title">
          <div>区块信息</div>
          </div>
        </div>
      </div>
    </div>
    <div class="hash-content-info">
      <div class="container">
        <div class="content">
          <div >
            <json-viewer
              :value="blockContent"
              :expand-depth='5'
              copyable
              sort></json-viewer>
            <div style="text-align: center" v-if="noData">暂无数据</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getBlockByHash } from '@/api/api'
import nav from '@/components/content-nav'
import '@/assets/css/layout.css'
import '@/assets/css/public.css'

export default {
  name: 'blockDetail',
  components: {
    'v-nav': nav
  },
  data: function () {
    return {
      blockName: '',
      blockContent: '',
      noData: false,
      page: {
        pageSize: this.$route.query.pageSize || 20,
        pageNumber: this.$route.query.pageNumber || 1
      }
    }
  },
  created: function () {
    this.blockName = '#' + this.$route.query.hash
  },
  mounted: function () {
    this.searchBlockDetail()
  },
  methods: {
    searchBlockDetail: function () {
      this.noData = false
      const data = this.$route.query.hash || ''
      getBlockByHash(data).then(res => {
        if (res.data.result) {
          this.noData = false
          this.blockContent = res.data.result
        } else {
          this.noData = true
        }
      })
    }
  }
}
</script>

<style scoped>
.content>>>.jv-code{
  max-height: none;
}
</style>
