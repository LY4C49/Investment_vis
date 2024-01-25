<!--  -->
<template>
  <div>
    <h2>History</h2>
    <div class="chart" id="twoChart">
    </div>
  </div>
</template>
<script>
import { } from 'vue';
import * as echarts from 'echarts';
import axios from 'axios'
export default {
  data() {
    return {
      x_Axis: null,
      stock_asset: null,
      option_asset: null,
      total_asset: null,
      yield: null,
      chart: null,
      resp: null
    }
  },

  mounted() {
    this.chart = echarts.init(document.getElementById("twoChart"))

    this.getdata()
    //this.updateChart()
  },

  methods: {
    getdata() {
      let url = '/data/history'
      axios.get(url)
        .then(response => {
          this.resp = response.data
          this.x_Axis = this.resp.data.xaxis
          this.stock_asset = this.resp.data.stock
          this.option_asset = this.resp.data.option
          this.total_asset = this.resp.data.total_asset
          this.yield = this.resp.data.yield
          this.updateChart()
        })
    },

    updateChart() {
      this.chart.setOption(

        {
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'cross',
              crossStyle: {
                color: '#999'
              }
            }
          },
          toolbox: {
            feature: {
              dataView: { show: false, readOnly: false },
              magicType: { show: false, type: ['line', 'bar'] },
              restore: { show: true },
              saveAsImage: { show: true }
            }
          },
          legend: {
            //data: ['Evaporation', 'Precipitation', 'Temperature']
            data: ['Stock', 'Option', 'Total Asset', 'Yield'],
            itemGap: 25,
            textStyle:{
              color: "#999",
              fontSize: 18
            }
          },
          xAxis: [
            {
              type: 'category',
              // data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
              data : this.x_Axis,
              axisPointer: {
                type: 'shadow'
              }
            }
          ],
          yAxis: [
            {
              type: 'value',
              name: 'Stock',
              min: 0,
              max: 200,
              interval: 50,
              axisLabel: {
                formatter: '${value}K'
              }
            },
            {
              type: 'value',
              name: 'Yield',
              min: -100,
              max: 100,
              interval: 50,
              axisLabel: {
                formatter: '{value}%'
              }
            }
          ],
          series: [
            {
              name: 'Stock',
              type: 'bar',
              tooltip: {
                valueFormatter: function (value) {
                  return '$' +value + 'K';
                }
              },
              // data: [
              //   2.0, 4.9, 7.0, 23.2, 25.6, 76.7, 135.6, 162.2, 32.6, 20.0, 6.4, 3.3
              // ]
              data : this.stock_asset,
              yAxisIndex: 0,
            },
            {
              name: 'Option',
              type: 'bar',
              tooltip: {
                valueFormatter: function (value) {
                  return '$' +value + 'K';
                }
              },
              // data: [
              //   2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3
              // ]
              data: this.option_asset,
              yAxisIndex: 0,
            },
            {
              name: 'Total Asset',
              type: 'bar',
              tooltip: {
                valueFormatter: function (value) {
                  return '$' +value + 'K';
                }
              },
              // data: [
              //   2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3
              // ]
              data: this.total_asset,
              yAxisIndex: 0,
            },
            {
              name: 'Yield',
              type: 'line',
              yAxisIndex: 1,
              tooltip: {
                valueFormatter: function (value) {
                  return value + ' %';
                }
              },
              //data: [2.0, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3, 23.4, 23.0, 16.5, 12.0, 6.2]
              data: this.yield,
            }
          ]
        }

      )
    },

    // renderData(){
    //   this.getdata()
    //   this.updateChart()
    // }
  }

}
</script>
<style scoped>
.chart {
  height: 4.5rem;
}

h2 {
  height: 48px;
  color: #fff;
  line-height: 0.6rem;
  font-size: 0.25rem;
  text-align: center;
}
</style>
  