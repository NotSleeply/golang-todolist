<template>
  <div class="about-page">
    <t-space direction="vertical" size="24px" style="width: 100%">
      <!-- 页面标题 -->
      <t-card class="title-card" theme="primary">
        <div class="title-content">
          <t-icon name="info-circle" size="48px" />
          <h1 class="page-title">关于 TodoList</h1>
        </div>
      </t-card>

      <!-- 应用描述 -->
      <t-card class="description-card">
        <template #header>
          <div class="card-header">
            <t-icon name="desktop" />
            <span class="header-title">应用描述</span>
          </div>
        </template>
        <p class="description-text">这是一个基于 Vue 3 + Go 的现代化 TodoList 应用。</p>
      </t-card>

      <!-- 技术栈 -->
      <t-card class="tech-stack-card">
        <template #header>
          <div class="card-header">
            <t-icon name="layers" />
            <span class="header-title">技术栈</span>
          </div>
        </template>

        <t-row :gutter="24">
          <t-col :span="6">
            <t-card variant="outline" class="tech-card">
              <template #header>
                <div class="tech-header">
                  <t-icon name="logo-chrome" />
                  <span>前端</span>
                </div>
              </template>
              <t-list split>
                <t-list-item>Vue 3 (Composition API)</t-list-item>
                <t-list-item>TDesign Vue Next</t-list-item>
                <t-list-item>Vite</t-list-item>
                <t-list-item>Pinia (状态管理)</t-list-item>
                <t-list-item>Vue Router</t-list-item>
                <t-list-item>Axios</t-list-item>
              </t-list>
            </t-card>
          </t-col>

          <t-col :span="6">
            <t-card variant="outline" class="tech-card">
              <template #header>
                <div class="tech-header">
                  <t-icon name="server" />
                  <span>后端</span>
                </div>
              </template>
              <t-list split>
                <t-list-item>Go (Golang)</t-list-item>
                <t-list-item>Gin 框架</t-list-item>
                <t-list-item>GORM</t-list-item>
                <t-list-item>MySQL</t-list-item>
              </t-list>
            </t-card>
          </t-col>
        </t-row>
      </t-card>

      <!-- 功能特性 -->
      <t-card class="features-card">
        <template #header>
          <div class="card-header">
            <t-icon name="star" />
            <span class="header-title">功能特性</span>
          </div>
        </template>

        <t-list split>
          <t-list-item>
            <template #prefix>
              <t-icon name="edit" class="feature-icon" />
            </template>
            创建、编辑、删除任务
          </t-list-item>
          <t-list-item>
            <template #prefix>
              <t-icon name="check-circle" class="feature-icon" />
            </template>
            标记任务完成状态
          </t-list-item>
          <t-list-item>
            <template #prefix>
              <t-icon name="chart-line" class="feature-icon" />
            </template>
            任务统计
          </t-list-item>
          <t-list-item>
            <template #prefix>
              <t-icon name="palette" class="feature-icon" />
            </template>
            现代化的用户界面
          </t-list-item>
          <t-list-item>
            <template #prefix>
              <t-icon name="mobile" class="feature-icon" />
            </template>
            响应式设计
          </t-list-item>
          <t-list-item>
            <template #prefix>
              <t-icon name="refresh" class="feature-icon" />
            </template>
            实时数据同步
          </t-list-item>
        </t-list>
      </t-card>

      <!-- 版本信息 -->
      <t-card class="version-card">
        <template #header>
          <div class="card-header">
            <t-icon name="info" />
            <span class="header-title">版本信息</span>
          </div>
        </template>

        <t-row :gutter="16">
          <t-col :span="6">
            <t-tag theme="primary" size="large">
              <template #icon>
                <t-icon name="tag" />
              </template>
              版本: 1.0.0
            </t-tag>
          </t-col>
          <t-col :span="6">
            <t-tag theme="success" size="large">
              <template #icon>
                <t-icon name="calendar" />
              </template>
              构建时间: {{ new Date().toLocaleDateString() }}
            </t-tag>
          </t-col>
        </t-row>
      </t-card>
    </t-space>
  </div>
</template>

<script setup>
  // 这里可以添加组件逻辑
</script>

<style scoped>
  .about-page {
    max-width: 1200px;
    margin: 0 auto;
    padding: 24px;
  }

  .title-card {
    text-align: center;
  }

  .title-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 16px;
    padding: 48px 24px;
  }

  .page-title {
    margin: 0;
    font-size: 32px;
    font-weight: 600;
    color: white;
  }

  .card-header {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .header-title {
    font-size: 18px;
    font-weight: 600;
  }

  .description-text {
    font-size: 16px;
    line-height: 1.6;
    margin: 0;
  }

  .tech-card {
    height: 100%;
  }

  .tech-header {
    display: flex;
    align-items: center;
    gap: 8px;
    font-weight: 600;
  }

  .feature-icon {
    color: var(--td-brand-color);
  }
</style>
