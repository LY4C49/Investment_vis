<!--  -->
<template>
  <div>
    <h2>Your Portfolio</h2>
    <div class="chart" id="fourChart">
    </div>
  </div>
</template>
<script>
import { } from 'vue';
import * as echarts from 'echarts';
import axios from 'axios';
export default {

  data() {
    return {
      chart: null,
      stock: 10,
      option: 10,
      cash: 10,
      resp: null,
      timer: null
    };
  },

  mounted() {
    this.chart = echarts.init(document.getElementById("fourChart"))
    this.updateChart()

    this.timer = setInterval(() => {
      setTimeout(this.getdata, 0)
    }, 2000)
  },

  methods: {
    getdata() {
      let url = '/data/rank?stock=' + this.stock + '&' + 'option=' + this.option + '&' + 'cash=' + this.cash
      axios.get(url).then(response => {
        this.resp = response.data
        this.stock = this.resp.data.stock
        this.option = this.resp.data.option
        this.cash = this.resp.data.cash
        this.updateChart()
      })

    },

    updateChart() {
      this.chart.setOption(
        {
          tooltip: {
            trigger: 'item'
          },
          legend: {
            top: '1%',
            left: 'center',
            itemGap:25,
            textStyle: {
              color: "#999"
            }
          },
          textStyle: {
            fontSize: 20,
            color: "#999"
          },
          series: [
            {
              name: 'Access From',
              type: 'pie',
              radius: ['50%', '70%'],
              avoidLabelOverlap: false,
              itemStyle: {
                borderRadius: 10,
                borderColor: '#fff',
                borderWidth: 2
              },
              label: {
                show: false,
                position: 'center'
              },
              emphasis: {
                label: {
                  show: true,
                  fontSize: 40,
                  fontWeight: 'bold'
                }
              },
              labelLine: {
                show: false
              },
              data: [
                { value: this.stock, name: 'Stock' },
                { value: this.option, name: 'Option' },
                { value: this.cash, name: 'Cash' },
              ]
            }
          ]
        }
      )

    }

  },

  beforeUnmount() {
    clearInterval(this.timer)
    this.timer = null
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
  