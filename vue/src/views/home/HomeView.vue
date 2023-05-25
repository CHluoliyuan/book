<template>
  <div>
    <el-card>
      <div id="line" style="width: 100%; height: 700px"></div>
    </el-card>
  </div>
</template>

<script>
import Cookies from 'js-cookie'
import request from "@/utils/request";
import * as echarts from 'echarts'

const option = {
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'cross',
      crossStyle: {
        color: '#999'
      }
    }
  },
  legend: {
    data: ['借书量', '还书量', '比例']
  },
  xAxis: [
    {
      type: 'category',
      data:[],
      axisPointer: {
        type: 'shadow'
      }
    }
  ],

  yAxis: [
    {
      type: 'value',
      name: '借书量',
      min: 0,
      max: 10,
      interval: 50,
      axisLabel: {
        formatter: '{value} 册'
      }
    },
    {
      type: 'value',
      name: '比例',
      min: 0,
      max: 25,
      interval: 10,
      axisLabel: {
        formatter: '{value} %'
      }
    }
  ],
  series: [
    {
      name: '借书量',
      type: 'bar',
      tooltip: {
        valueFormatter: function (value) {
          return value + '册';
        }
      },
      data: []
    },
    {
      name: '还书量',
      type: 'bar',
      tooltip: {
        valueFormatter: function (value) {
          return value + '册';
        }
      },
      data: []
    },
    {
      name: '比例',
      type: 'line',
      yAxisIndex: 1,
      tooltip: {
        valueFormatter: function (value) {
          return value + '%';
        }
      },
      data: []
    }
  ]
};


export default {
    data() {
      return {
        admin: Cookies.get('admin') ? JSON.parse(Cookies.get('admin')) : {},
        lineBox: null,
        timeRange: 'week',
        options: [
          {label: '最近一周', value: 'week'},
          {label: '最近一个月', value: 'month'},
          {label: '最近两个月', value: 'month2'},
          {label: '最近三个月', value: 'month3'},
        ]
      }
    },
    mounted() {  // 等页面元素全部初始化好
      this.load()
    },
    methods: {
      load() {
        if (!this.lineBox) {
          this.lineBox = echarts.init(document.getElementById('line'))  // 初始化echarts容器
        }
        // 从后台获取数据
        request.get('/echart_data').then(res => {
          const title_data=[]
          const borrow_data=[]
          const retur_data=[]
          const linelist=[]
          for (let key in res.data) {
            title_data.push(res.data[key][0])
            borrow_data.push(res.data[key][1])
            retur_data.push(res.data[key][2])
          }
          const sum = borrow_data.reduce((acc, curr) => acc + parseInt(curr), 0);
          for (const i of borrow_data) {
            linelist.push((i/sum*100).toFixed(2))
          }

          option.xAxis[0].data=title_data
          option.series[0].data=borrow_data
          option.series[1].data=retur_data
          option.series[2].data=linelist
          option.yAxis[1].max= Math.max(...linelist)+10
          option.yAxis[0].max= Math.max(...borrow_data)+10
          this.lineBox.setOption(option)  // 设置容器的属性值，当你的数据发生变化的时候，一定要重新setOption
        })
      }
    }
  }
</script>

<style>
.input {
  width: 300px;
}
</style>
