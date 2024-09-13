<template>
  <div id="app">
    <!-- 页面标题 -->
    <header class="header">
      <h1>Welcome to Your principles of computer composition test check</h1>
    </header>

    <!-- 查询表单 -->
    <section class="search-section">
      <form @submit.prevent="fetchGrades">
        <label for="studentId">Please input your student number：</label>
        <input v-model="studentId" type="text" id="studentId" placeholder="student number" required>
        <button type="submit">check score</button>
      </form>
    </section>

    <!-- 查询结果显示区域 -->
    <section class="result-section" v-if="grades.length > 0">
      <h2>查询结果</h2>
      <div class="result">
        <div v-for="(grade, index) in grades" :key="index">
          <p>科目: {{ grade.subject }}, 成绩: {{ grade.score }}</p>
        </div>
      </div>
    </section>

    <section class="result-section" v-else-if="errorMessage">
      <p style="color:red;">{{ errorMessage }}</p>
    </section>
  </div>
</template>

<script>
import { ref } from 'vue';

export default {
  name: 'App',
  setup() {
    const studentId = ref('');
    const grades = ref([]);
    const errorMessage = ref('');

    const fetchGrades = () => {
      if (!studentId.value.trim()) {
        errorMessage.value = '请输入有效的学号。';
        return;
      }

      // 清空之前的错误信息和结果
      errorMessage.value = '';
      grades.value = [];

      // 模拟一个异步查询过程，通常这里是对API的调用
      fetch(`https://api.example.com/grades?studentId=${studentId.value}`)  // 替换为你的实际API端点
        .then(response => response.json())
        .then(data => {
          if (data && data.grades) {
            grades.value = data.grades;
          } else {
            errorMessage.value = '未找到相关成绩记录。';
          }
        })
        .catch(error => {
          console.error('查询时出错:', error);
          errorMessage.value = '查询时出错，请稍后重试。';
        });
    };

    return {
      studentId,
      grades,
      errorMessage,
      fetchGrades
    };
  }
};
</script>

<style scoped>
body {
  font-family: Arial, sans-serif;
  background-color: #f0f0f0;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.header {
  background-color: #007bff;
  color: white;
  width: 100%;
  padding: 1rem;
  text-align: center;
}

.search-section {
  margin: 2rem 0;
  padding: 1rem;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  width: 300px;
  text-align: center;
  /* Center align text */
  display: flex;
  /* Add flexbox for centering */
  flex-direction: column;
  /* Stack children vertically */
  align-items: center;
  /* Center items horizontally */
}

.search-section input {
  width: 100%;
  /* Make input take the full width */
  padding: 0.5rem;
  margin: 1rem 0;
  border: 1px solid #ddd;
  border-radius: 4px;
  box-sizing: border-box;
  /* Include padding in width */
}

.search-section button {
  padding: 0.5rem 1rem;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 1rem;
  /* Add margin to separate button from input */
}

.search-section button:hover {
  background-color: #0056b3;
}

.result-section {
  margin: 1rem;
  padding: 1rem;
  background-color: white;
  border-radius: 8px;
  width: 300px;
  text-align: center;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.result {
  margin-top: 1rem;
  font-size: 1.2rem;
  color: #333;
}
</style>


