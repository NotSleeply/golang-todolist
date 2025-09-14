<template>
  <div class="about">
    <div class="container">
      <h1>关于 TodoList</h1>
      <div class="content">
        <p>这是一个基于 Vue 3 + Go 的现代化 TodoList 应用。</p>
        
        <h2>技术栈</h2>
        <div class="tech-stack">
          <div class="tech-section">
            <h3>前端</h3>
            <ul>
              <li>Vue 3 (Composition API)</li>
              <li>Vite</li>
              <li>Pinia (状态管理)</li>
              <li>Vue Router</li>
              <li>Axios</li>
            </ul>
          </div>
          
          <div class="tech-section">
            <h3>后端</h3>
            <ul>
              <li>Go (Golang)</li>
              <li>Gin 框架</li>
              <li>GORM</li>
              <li>MySQL</li>
            </ul>
          </div>
        </div>

        <h2>功能特性</h2>
        <ul class="features">
          <li>📝 创建、编辑、删除任务</li>
          <li>✅ 标记任务完成状态</li>
          <li>📊 任务统计</li>
          <li>🎨 现代化的用户界面</li>
          <li>📱 响应式设计</li>
          <li>🔄 实时数据同步</li>
        </ul>

        <div class="version-info">
          <p><strong>版本:</strong> 1.0.0</p>
          <p><strong>构建时间:</strong> {{ new Date().toLocaleDateString() }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// 这里可以添加组件逻辑
</script>

<style scoped>
.about {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 2rem 0;
}

.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 0 1rem;
}

h1 {
  text-align: center;
  color: white;
  font-size: 2.5rem;
  margin-bottom: 2rem;
  text-shadow: 0 2px 4px rgba(0,0,0,0.3);
}

.content {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 2rem;
  color: white;
}

h2 {
  color: #fff;
  margin-top: 2rem;
  margin-bottom: 1rem;
  font-size: 1.5rem;
}

h3 {
  color: #fff;
  margin-bottom: 0.5rem;
}

.tech-stack {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
  margin: 1rem 0;
}

.tech-section {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  padding: 1.5rem;
}

ul {
  list-style: none;
  padding: 0;
}

ul li {
  padding: 0.5rem 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

ul li:last-child {
  border-bottom: none;
}

.features li {
  font-size: 1.1rem;
  padding: 0.75rem 0;
}

.version-info {
  margin-top: 2rem;
  padding-top: 1rem;
  border-top: 1px solid rgba(255, 255, 255, 0.2);
}

.version-info p {
  margin: 0.5rem 0;
}

@media (max-width: 768px) {
  h1 {
    font-size: 2rem;
  }
  
  .content {
    padding: 1.5rem;
  }
  
  .tech-stack {
    grid-template-columns: 1fr;
  }
}
</style>
