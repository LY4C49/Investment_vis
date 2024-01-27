<template>
  <router-view/>
</template>

<script>
import {provide} from "vue"
import * as echarts from "echarts"
import axios from "axios"

// set base url
axios.defaults.baseURL = "http://124.220.7.23:8888"

export default{
  setup(){
    provide("echarts", echarts)
    provide("axios", axios)
  }
}
</script>

<style lang="less">
body{
  background: url("~@/assets/bg1.jpg") top center no-repeat;
}

*{
  margin: 0px;
  padding: 0px;
  box-sizing: border-box;
  font-family: Microsoft YaHei;
}

</style>
