<template>
  <div id="app">
    <t-header class="header">
      <t-head-menu theme="light" value="home">
        <template #logo>
          <h2 class="logo">TodoList</h2>
        </template>
        <t-menu-item value="home" @click="$router.push('/')">
          <template #icon>
            <t-icon name="home" />
          </template>
          首页
        </t-menu-item>
        <t-menu-item value="about" @click="$router.push('/about')">
          <template #icon>
            <t-icon name="info-circle" />
          </template>
          关于
        </t-menu-item>
      </t-head-menu>
    </t-header>

    <main class="main-content">
      <router-view />
    </main>
  </div>
</template>

<style scoped>
  .header {
    border-bottom: 1px solid var(--td-component-border);
  }

  .logo {
    margin: 0;
    padding: 0 16px;
    color: var(--td-brand-color);
    font-size: 20px;
    font-weight: 600;
  }

  .main-content {
    min-height: calc(100vh - 64px);
    background: var(--td-bg-color-page);
  }
</style>
