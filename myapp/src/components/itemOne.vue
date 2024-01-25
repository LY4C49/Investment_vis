<!--  -->
<template>
  <div>
    <h2>Main Indexs</h2>
    <div class="chart" id="oneChart">
      <!-- 图表的容器 -->
      <div id="ShangHai"> </div>
      <div id="Nasdaq"> </div>
      <!-- <div id="HongKong"> </div>
      <div id="London"> </div> -->
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
      sh_chart: null,
      nq_chart: null,
      sh_data: 2700,
      nq_data: 15000,
      resp: null,
      timer: null
    };
  },

  mounted() {
    this.sh_chart = echarts.init(document.getElementById("ShangHai"))
    this.nq_chart = echarts.init(document.getElementById("Nasdaq"))

    // this.sh_chart.setOption({
    //   series: [
    //     {
    //       type: 'gauge',
    //       center: ['50%', '60%'],
    //       startAngle: 200,
    //       endAngle: -20,
    //       min: 1000,
    //       max: 6000,
    //       splitNumber: 2,
    //       itemStyle: {
    //         color: '#FFAB91'
    //       },
    //       progress: {
    //         show: true,
    //         width: 30
    //       },
    //       pointer: {
    //         show: false
    //       },
    //       axisLine: {
    //         lineStyle: {
    //           width: 30
    //         }
    //       },
    //       axisTick: {
    //         distance: -45,
    //         splitNumber: 5,
    //         lineStyle: {
    //           width: 2,
    //           color: '#999'
    //         }
    //       },
    //       splitLine: {
    //         distance: -52,
    //         length: 14,
    //         lineStyle: {
    //           width: 3,
    //           color: '#999'
    //         }
    //       },
    //       axisLabel: {
    //         distance: -10,
    //         color: '#999',
    //         fontSize: 30
    //       },
    //       anchor: {
    //         show: false
    //       },
    //       title: {
    //         show: true,
    //         fontSize: 16
    //       },
    //       detail: {
    //         valueAnimation: true,
    //         width: '60%',
    //         lineHeight: 40,
    //         borderRadius: 8,
    //         offsetCenter: [0, '-15%'],
    //         fontSize: 30,
    //         fontWeight: 'bolder',
    //         formatter: '{value}',
    //         color: 'inherit'
    //       },
    //       data: [
    //         {
    //           value: this.sh_data,
    //           name: "SHANGHAI"
    //         }
    //       ]
    //     },
    //     {
    //       type: 'gauge',
    //       center: ['50%', '60%'],
    //       startAngle: 200,
    //       endAngle: -20,
    //       min: 0,
    //       max: 60,
    //       itemStyle: {
    //         color: '#FD7347'
    //       },
    //       progress: {
    //         show: true,
    //         width: 8
    //       },
    //       pointer: {
    //         show: false
    //       },
    //       axisLine: {
    //         show: false
    //       },
    //       axisTick: {
    //         show: false
    //       },
    //       splitLine: {
    //         show: false
    //       },
    //       axisLabel: {
    //         show: false
    //       },
    //       detail: {
    //         show: false
    //       },
    //       data: [
    //         {
    //           value: 2600
    //         }
    //       ]
    //     }
    //   ]
    // })

    // this.nq_chart.setOption({
    //   series: [
    //     {
    //       type: 'gauge',
    //       center: ['50%', '60%'],
    //       startAngle: 200,
    //       endAngle: -20,
    //       min: 12000,
    //       max: 16000,
    //       splitNumber: 2,
    //       itemStyle: {
    //         color: '#FFAB91'
    //       },
    //       progress: {
    //         show: true,
    //         width: 30
    //       },
    //       pointer: {
    //         show: false
    //       },
    //       axisLine: {
    //         lineStyle: {
    //           width: 30
    //         }
    //       },
    //       axisTick: {
    //         distance: -45,
    //         splitNumber: 5,
    //         lineStyle: {
    //           width: 2,
    //           color: '#999'
    //         }
    //       },
    //       splitLine: {
    //         distance: -52,
    //         length: 14,
    //         lineStyle: {
    //           width: 3,
    //           color: '#999'
    //         }
    //       },
    //       axisLabel: {
    //         distance: -10,
    //         color: '#999',
    //         fontSize: 15
    //       },
    //       anchor: {
    //         show: false
    //       },
    //       title: {
    //         show: true,
    //         fontSize: 20
    //       },
    //       detail: {
    //         valueAnimation: true,
    //         width: '40%',
    //         lineHeight: 40,
    //         borderRadius: 8,
    //         offsetCenter: [0, '-15%'],
    //         fontSize: 20,
    //         fontWeight: 'bolder',
    //         formatter: '{value}',
    //         color: 'inherit'
    //       },
    //       data: [
    //         {
    //           value: this.nq_data,
    //           name: "NASDAQ"
    //         }
    //       ]
    //     },
    //     {
    //       type: 'gauge',
    //       center: ['50%', '60%'],
    //       startAngle: 200,
    //       endAngle: -20,
    //       min: 0,
    //       max: 60,
    //       itemStyle: {
    //         color: '#FD7347'
    //       },
    //       progress: {
    //         show: true,
    //         width: 8
    //       },
    //       pointer: {
    //         show: false
    //       },
    //       axisLine: {
    //         show: false
    //       },
    //       axisTick: {
    //         show: false
    //       },
    //       splitLine: {
    //         show: false
    //       },
    //       axisLabel: {
    //         show: false
    //       },
    //       detail: {
    //         show: false
    //       },
    //       data: [
    //         {
    //           value: 2600
    //         }
    //       ]
    //     }
    //   ]
    // })

    // this.updateall()
    this.getdata()

    this.timer = setInterval(()=>{
      setTimeout(this.updateall, 0)
    }, 2000)

  },

  methods: {

    getdata() {
      let url = '/data/indexs?sh=' + this.sh_data + '&' + 'nq=' + this.nq_data
      axios.get(url).then(response => {
        this.resp = response.data;
        this.sh_data = this.resp.data.ShangHai
        this.nq_data = this.resp.data.Nasdaq
        //console.log(this.sh_data, this.nq_data)
        this.updateNQChart()
        this.updateSHChart()
      })
    },

    updateSHChart() {
      this.sh_chart.setOption({
        series: [
          {
            type: 'gauge',
            center: ['50%', '60%'],
            startAngle: 200,
            endAngle: -20,
            min: 1000,
            max: 6000,
            splitNumber: 2,
            itemStyle: {
              color: '#FFAB91'
            },
            progress: {
              show: true,
              width: 20
            },
            pointer: {
              show: false
            },
            axisLine: {
              lineStyle: {
                width: 20
              }
            },
            axisTick: {
              distance: -45,
              splitNumber: 5,
              lineStyle: {
                width: 2,
                color: '#999'
              }
            },
            splitLine: {
              distance: -45,
              length: 45,
              lineStyle: {
                width: 3,
                color: '#999'
              }
            },
            axisLabel: {
              distance: -50,
              color: '#999',
              fontSize: 15
            },
            anchor: {
              show: false
            },
            title: {
              show: true,
              fontSize: 20,
              color:'#999'
            },
            detail: {
              valueAnimation: true,
              width: '60%',
              lineHeight: 40,
              borderRadius: 8,
              offsetCenter: [0, '-15%'],
              fontSize: 30,
              fontWeight: 'bolder',
              formatter: '{value}',
              color: 'inherit'
            },
            data: [
              {
                value: this.sh_data,
                name: "SHANGHAI"
              }
            ]
          },
          {
            type: 'gauge',
            center: ['50%', '60%'],
            startAngle: 200,
            endAngle: -20,
            min: 0,
            max: 60,
            itemStyle: {
              color: '#FD7347'
            },
            progress: {
              show: true,
              width: 8
            },
            pointer: {
              show: false
            },
            axisLine: {
              show: false
            },
            axisTick: {
              show: false
            },
            splitLine: {
              show: false
            },
            axisLabel: {
              show: false
            },
            detail: {
              show: false
            },
            data: [
              {
                value: 2600
              }
            ]
          }
        ]
      })
    },

    updateNQChart() {
      this.nq_chart.setOption({
        series: [
          {
            type: 'gauge',
            center: ['50%', '60%'],
            startAngle: 200,
            endAngle: -20,
            min: 12000,
            max: 16000,
            splitNumber: 2,
            itemStyle: {
              color: '#FFAB91'
            },
            progress: {
              show: true,
              width: 20
            },
            pointer: {
              show: false
            },
            axisLine: {
              lineStyle: {
                width: 20
              }
            },
            axisTick: {
              distance: -45,
              splitNumber: 5,
              lineStyle: {
                width: 2,
                color: '#999'
              }
            },
            splitLine: {
              distance: -45,
              length: 45,
              lineStyle: {
                width: 3,
                color: '#999'
              }
            },
            axisLabel: {
              distance: -50,
              color: '#999',
              fontSize: 15
            },
            anchor: {
              show: false
            },
            title: {
              show: true,
              fontSize: 20,
              color:'#999'
            },
            detail: {
              valueAnimation: true,
              width: '40%',
              lineHeight: 40,
              borderRadius: 8,
              offsetCenter: [0, '-15%'],
              fontSize: 30,
              fontWeight: 'bolder',
              formatter: '{value}',
              color: 'inherit'
            },
            data: [
              {
                value: this.nq_data,
                name: "NASDAQ"
              }
            ]
          },
          {
            type: 'gauge',
            center: ['50%', '60%'],
            startAngle: 200,
            endAngle: -20,
            min: 0,
            max: 60,
            itemStyle: {
              color: '#FD7347'
            },
            progress: {
              show: true,
              width: 8
            },
            pointer: {
              show: false
            },
            axisLine: {
              show: false
            },
            axisTick: {
              show: false
            },
            splitLine: {
              show: false
            },
            axisLabel: {
              show: false
            },
            detail: {
              show: false
            },
            data: [
              {
                value: 2600
              }
            ]
          }
        ]
      })
    },

    updateall(){
      this.getdata()
      ///this.updateNQChart()
      //this.updateSHChart()
    }

  },

  beforeUnmount(){
    clearInterval(this.timer)
    this.timer = null
  }
}

</script>
<style scoped>
.chart {
  height: 4.5rem;

  #ShangHai {
    height: 4rem;
    width: 4rem;
    display: inline-block;
    margin-right: 10px
      /* background-color: blue */
  }

  ;

  #Nasdaq {
    height: 4rem;
    width: 4rem;
    display: inline-block;
    /* background-color: red */
  }

  ;
}

h2 {
  height: 48px;
  color: #fff;
  line-height: 0.6rem;
  font-size: 0.3rem;
  text-align: center;
}
</style>
  