<template>
  <div>
    <p v-if="averageScore !== null">Average Score (based on 9 exams) : {{ averageScore.toFixed(2) }} points </p> 
    <p> Total Score : 45 points </p>
    <div v-if="data && data.length > 0">
      <canvas ref="chartCanvas"></canvas>
    </div>
    <p v-else>No exam data available.</p>
  </div>
</template>

<script>
import { ref, onMounted, watch, computed } from 'vue';
import Chart from 'chart.js/auto';
import ChartDataLabels from 'chartjs-plugin-datalabels'; // 导入插件

export default {
  name: 'GradesChart',
  props: {
    data: {
      type: Array,
      required: true,
      default: () => []  // 默认值为空数组
    },
    studentName: {
      type: String,
      default: ''
    }
  },
  setup(props) {
    const chartCanvas = ref(null);
    let chartInstance = null;

    // 计算平时分
    const averageScore = computed(() => {
      if (props.data.length === 0) return null;
      const totalScore = props.data.reduce((sum, item) => sum + item.score, 0);
      return totalScore / 20; // 除以 20
    });

    const renderChart = () => {
      if (chartInstance) {
        chartInstance.destroy();  // 如果图表已存在，先销毁
      }
      
      chartInstance = new Chart(chartCanvas.value, {
        type: 'bar',
        data: {
          labels: props.data.map(item => item.examName),  // 使用 props.data 映射标签
          datasets: [{
            label: 'Score',
            data: props.data.map(item => item.score),  // 使用 props.data 映射分数
            backgroundColor: 'rgba(54, 162, 235, 0.2)', // 胶水蓝色的背景
            borderColor: 'rgba(54, 162, 235, 1)', // 胶水蓝色的边框
            borderWidth: 1
          }]
        },
        options: {
          responsive: true,
          plugins: {
            legend: {
              display: true,
              position: 'top',
              labels: {
                font: {
                  size: 16 // 增加图例字体大小
                }
              }
            },
            tooltip: {
              callbacks: {
                label: function(tooltipItem) {
                  return `Score: ${tooltipItem.raw}`;
                }
              }
            },
            datalabels: {
              color: 'black',
              display: true,
              anchor: 'end',
              align: 'top',
              formatter: (value) => value, // 显示分数
              font: {
                size: 14, // 设置标签字体大小
                weight: 'bold'
              }
            }
          },
          scales: {
            x: {
              title: {
                display: true,
                text: 'Exams',
                font: {
                  size: 16 // 增加x轴标题字体大小
                }
              },
              ticks: {
                font: {
                  size: 14 // 增加x轴刻度字体大小
                }
              }
            },
            y: {
              title: {
                display: true,
                text: 'Score',
                font: {
                  size: 16 // 增加y轴标题字体大小
                }
              },
              ticks: {
                font: {
                  size: 14 // 增加y轴刻度字体大小
                }
              }
            }
          }
        },
        plugins: [ChartDataLabels] // 添加插件
      });
    };

    onMounted(() => {
      if (props.data && props.data.length > 0) {
        renderChart();  // 当组件挂载时渲染图表
      }
    });

    watch(() => props.data, (newData) => {
      if (chartInstance) {
        chartInstance.destroy(); // 销毁旧图表
      }
      if (newData && newData.length > 0) {
        renderChart();  // 监控数据变化并重新渲染图表
      }
    }, { deep: true });

    return {
      chartCanvas,
      averageScore
    };
  }
};
</script>

<style scoped>
/* 图表容器样式 */
.chart-container {
  position: relative;
  height: 400px; /* 调整为适合的高度 */
  width: 100%;
}

/* 平时分的样式 */
p {
  text-align: center; 
  font-size: 24px; /* 设置字体大小 */
  font-weight: bold; /* 设置字体加粗 */
  margin-bottom: 1rem; /* 设置底部外边距 */
}
</style>
