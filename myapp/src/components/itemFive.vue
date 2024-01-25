<!--  -->
<template>
  <div>
    <h2>Investment Yield</h2>
    <div class="chart" id="fiveChart">
    </div>
  </div>
</template>

<script>
import * as echarts from 'echarts';
import axios from 'axios';
export default {

  data() {
    return {
      panel: null,
      stock: 10,
      option: -5,
      total: 5,
      resp: null,
      gaugeData: null
    }
  },


  mounted() {
    this.panel = echarts.init(document.getElementById("fiveChart"))

    this.updatePanel
    this.updateGauge()
    
    this.timer = setInterval(() => {
      setTimeout(this.getdata, 0)
    }, 2000)
  },

  methods: {
    getdata() {
      let url = '/data/panel?stock=' + this.stock + '&' + 'option=' + this.option + '&' + 'total=' + this.total
      axios.get(url).then(response => {
        this.resp = response.data
        this.stock = this.resp.data.stock
        this.option = this.resp.data.option
        this.total = this.resp.data.total
        this.updateGauge()
        this.updatePanel()
      })
    },

    updateGauge() {
      this.gaugeData = [
        {
          value: this.stock,
          name: 'Stock',
          title: {
            offsetCenter: ['-60%', '90%']
          },
          detail: {
            offsetCenter: ['-60%', '105%']
          }
        },
        {
          value: this.option,
          name: 'Option',
          title: {
            offsetCenter: ['0%', '90%']
          },
          detail: {
            offsetCenter: ['0%', '105%']
          }
        },
        {
          value: this.total,
          name: 'Total',
          title: {
            offsetCenter: ['60%', '90%'],
          },
          detail: {
            offsetCenter: ['60%', '105%']
          }
        }
      ]
    },

    updatePanel() {
      this.panel.setOption(
        {
          series: [
            {
              type: 'gauge',
              min: -100,
              max: 100,
              anchor: {
                show: true,
                showAbove: true,
                size: 18,
                itemStyle: {
                  color: '#FAC858'
                }
              },
              pointer: {
                icon: 'path://M2.9,0.7L2.9,0.7c1.4,0,2.6,1.2,2.6,2.6v115c0,1.4-1.2,2.6-2.6,2.6l0,0c-1.4,0-2.6-1.2-2.6-2.6V3.3C0.3,1.9,1.4,0.7,2.9,0.7z',
                width: 8,
                length: '80%',
                offsetCenter: [0, '8%']
              },
              progress: {
                show: true,
                overlap: true,
                roundCap: true
              },
              axisLine: {
                roundCap: true
              },
              data: this.gaugeData,
              title: {
                fontSize: 20,
                color: "#999"
              },
              axisLabel: {
                textStyle: {
                  fontSize: 18
                },
              },
              detail: {
                width: 70,
                height: 20,
                fontSize: 18,
                color: '#fff',
                backgroundColor: 'inherit',
                borderRadius: 3,
                formatter: '{value}%'
              }
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
  height: 8rem;
}

h2 {
  height: 48px;
  color: #fff;
  line-height: 0.6rem;
  font-size: 0.5rem;
  text-align: center;
}
</style>
  