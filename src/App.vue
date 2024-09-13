<template>
  <div id="app">
    <header class="header">
      <h1>Principles of Computer Composition Exam Grade Query System</h1> <!-- Updated Header -->
    </header>
    <!-- <nav>
      <router-link to="/admin">Admin</router-link>   
    </nav>
    <router-view/> -->
    <!-- 登录界面 -->
    <section v-if="!isLoggedIn">
      <div v-if="showLogin" class="login-section">
        <h2>Login</h2>
        <form @submit.prevent="login">
          <label for="loginName">Student Name:</label>
          <input v-model="loginName" type="text" id="loginName" placeholder="Student Name" required>
          <label for="loginPassword">Password:</label>
          <input v-model="loginPassword" type="password" id="loginPassword" placeholder="Password" required>
          <button type="submit">Login</button>
          <p v-if="loginErrorMessage" style="color:red;">{{ loginErrorMessage }}</p>
          <p>
            <a href="#" @click.prevent="switchToChangePassword">Change Password</a>
          </p>
        </form>
      </div>

      <div v-if="!showLogin" class="change-password-section">
        <h2>Change Password</h2>
        <form @submit.prevent="changePassword">
          <label for="name">Student Name:</label>
          <input v-model="name" type="text" id="name" placeholder="Student Name" required>
          
          <label for="currentPassword">Current Password:</label>
          <input v-model="oldPassword" type="password" id="currentPassword" placeholder="Current Password" required>

          <label for="newPassword">New Password:</label>
          <input v-model="newPassword" type="password" id="newPassword" placeholder="New Password" required>

          <label for="confirmPassword">Confirm New Password:</label>
          <input v-model="confirmPassword" type="password" id="confirmPassword" placeholder="Confirm New Password" required>

          <button type="submit">Change Password</button>
          <p v-if="changePasswordErrorMessage" style="color:red;">{{ changePasswordErrorMessage }}</p>
          <p v-if="changePasswordSuccessMessage" style="color:green;">{{ changePasswordSuccessMessage }}</p>
          <p>
            <a href="#" @click.prevent="switchToLogin">Back to Login</a>
          </p>
        </form>
      </div>
    </section>

    <!-- 结果视图 -->
    <section v-else>
      <div class="chart-container">
        <!-- 显示学生姓名和成绩 -->
        <header class="results-header">{{ studentName }}'s Exam Results</header>
        <div v-if="examData.length > 0">
          <GradesChart :data="examData" />
        </div>
        <p v-else>No exam data available.</p>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref } from 'vue'; 
import config from '@/config.js';  // 导入配置文件
import GradesChart from './components/GradesChart.vue'; // 确保路径正确


const isLoggedIn = ref(false);
const showLogin = ref(true);

const loginName = ref('');
const loginPassword = ref('');
const loginErrorMessage = ref('');

const name = ref('');//修改界面的用户名
const showChangePassword = ref(false);
const oldPassword = ref('');
const newPassword = ref('');
const confirmPassword = ref('');
const changePasswordErrorMessage = ref('');
const changePasswordSuccessMessage = ref('');

const examData = ref([]);
const studentName = ref(''); // 新增学生姓名变量

//测试数据
// const mockExamData = [         
//       { examName: 'test1', score: 85 },
//       { examName: 'test2', score: 0 },
//       { examName: 'test3', score: 1 }
// ];
    
const switchToChangePassword = () => {
  showLogin.value = false;
  showChangePassword.value = true;
};

const switchToLogin = () => {
  showLogin.value = true;
  changePasswordErrorMessage.value = ''; // Optional: Clear any existing error messages
  changePasswordSuccessMessage.value = ''; // Optional: Clear success message
};

const switchToScore = () => {
  isLoggedIn.value = true;
};

const login = async () => {
  if (!loginName.value.trim() || !loginPassword.value.trim()) {
    loginErrorMessage.value = 'Please enter Student Name and Password.';
    return;
  }

  loginErrorMessage.value = '';

  try {
    const requestBody = {
      name: loginName.value,
      password: loginPassword.value
    };

    console.log('Sending request:', requestBody); // 打印请求体
    console.log('Request URL:', `${config.API_BASE_URL}/grades`);
    const response = await fetch(`${config.API_BASE_URL}/grades`, {  // Updated URL
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestBody)// 确保请求体已正确序列化为JSON
    });

    console.log('Response status:', response.status); // 打印响应状态

    if (!response.ok) {
      throw new Error(`Network response was not OK: ${response.statusText}`);
    }

    const data = await response.json();

    console.log('Response data:', data); // 打印响应数据

    if (data.success) {
      isLoggedIn.value = true;
      examData.value = data.data; // 假设后端返回的数据结构包含 `data.data`
      
      studentName.value = loginName.value; // 登录成功后保存学生姓名
      localStorage.setItem('jwtToken', data.token); // 保存 JWT Token 到 localStorage

      console.log('data:', data.data);//打印data分数数据
      console.log('examData.value:', examData.value);
      

      // 将对象转换为数组格式
      const formattedData = Object.entries(data.data).map(([examName, score]) => ({
        examName,
        score
      }));

      examData.value = formattedData; // 更新为转换后的数组
      console.log('Formatted Exam Data:', examData.value); // 打印转换后的数据

      // examData.value = mockExamData;
      // console.log('Formatted Exam Data:', examData.value); // 打印转换后的数据

      switchToScore(); // 切换到分数视图
    } else {
      loginErrorMessage.value = 'Login failed. Please check Student Name and Password.';
    }
  } catch (error) {
    console.error('Error during login:', error);
    loginErrorMessage.value = 'Error during login. Please try again later.';
  }
};

const changePassword = async () => {
  if (
    !oldPassword.value.trim() ||
    !newPassword.value.trim() ||
    !confirmPassword.value.trim()
  ) {
    changePasswordErrorMessage.value = 'Please enter all fields: current password, new password, and confirm new password.';
    return;
  }

  if (newPassword.value !== confirmPassword.value) {
    changePasswordErrorMessage.value = 'New password and confirm password do not match.';
    return;
  }

  changePasswordErrorMessage.value = '';

  try {
    const response = await fetch(`${config.API_BASE_URL}/change_password`, {  // 更新URL
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: name.value,  // 用户名
        old_password: oldPassword.value,  // 旧密码
        new_password: newPassword.value  // 新密码
      })
    });

    if (!response.ok) {
      throw new Error('Network response was not OK');
    }

    const data = await response.json();
    console.log('Response data:', data); // 打印响应数据

    if (data.success) {
      changePasswordSuccessMessage.value = 'Password changed successfully!';
      changePasswordErrorMessage.value = '';
      // 这里可以重置输入框
      oldPassword.value = '';
      newPassword.value = '';
      confirmPassword.value = '';
    } else {
      changePasswordErrorMessage.value = data.message || 'Password change failed. Please check your current password.';
    }
  } catch (error) {
    console.error('Error during password change:', error);
    changePasswordErrorMessage.value = 'Error during password change. Please try again later.';
  }
};   

</script>

<style>
.header {
  text-align: center;
  background-color: blue;
  color: white;
  padding: 1rem;
}

.results-header {
  font-size: 2rem; /* Increase the font size */
  font-weight: bold; /* Optionally make the text bold */
  text-align: center; /* Center align text */
  margin-top: 1rem; /* Add top margin for spacing */
}

.login-section,
.register-section {
  max-width: 400px;
  margin: 0 auto;
  padding: 1rem;
  border: 1px solid #ddd;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.login-section h2,
.register-section h2 {
  text-align: center;
}

form {
  display: flex;
  flex-direction: column;
}

form label {
  margin-bottom: 0.5rem;
  font-size: 1.1rem;
}

form input {
  margin-bottom: 1.5rem;
  padding: 0.8rem;
  font-size: 1.2rem;
  width: 90%;
  border-radius: 5px;
  border: 1px solid #ddd;
}

form button {
  margin-bottom: 1rem;
  padding: 0.8rem;
  font-size: 1.2rem;
  background-color: blue;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  width: 100%;
}

form button:hover {
  background-color: darkblue;
}

section {
  max-width: 800px;
  margin: 2rem auto;
}

.chart-container {
  position: relative;
  height: 400px;
  width: 100%;
}
</style>
