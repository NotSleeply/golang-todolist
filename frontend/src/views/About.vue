<template>
  <v-container class="py-8">
    <v-row justify="center">
      <v-col cols="12" md="10" lg="8">
        <!-- 页面标题 -->
        <v-card class="mb-6 text-center" elevation="8" color="primary" dark>
          <v-card-title class="text-h3 pa-6">
            <v-icon left size="large">mdi-information</v-icon>
            关于 TodoList
          </v-card-title>
        </v-card>

        <!-- 应用描述 -->
        <v-card class="mb-6" elevation="4">
          <v-card-text class="pa-6">
            <v-icon left size="large" color="primary" class="mb-4">mdi-application</v-icon>
            <p class="text-h6 mb-0">这是一个基于 Vue 3 + Go 的现代化 TodoList 应用。</p>
          </v-card-text>
        </v-card>

        <!-- 技术栈 -->
        <v-card class="mb-6" elevation="4">
          <v-card-title class="text-h4 pa-4">
            <v-icon left color="primary">mdi-layers</v-icon>
            技术栈
          </v-card-title>
          <v-card-text>
            <v-row>
              <v-col cols="12" md="6">
                <v-card variant="outlined" class="h-100">
                  <v-card-title class="text-h5">
                    <v-icon left color="primary">mdi-web</v-icon>
                    前端
                  </v-card-title>
                  <v-card-text>
                    <v-list>
                      <v-list-item>
                        <v-list-item-title>Vue 3 (Composition API)</v-list-item-title>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Vuetify 3</v-list-item-title>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Vite</v-list-item-title>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Pinia (状态管理)</v-list-item-title>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Vue Router</v-list-item-title>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Axios</v-list-item-title>
                      </v-list-item>
                    </v-list>
                  </v-card-text>
                </v-card>
              </v-col>

              <v-col cols="12" md="6">
                <v-card variant="outlined" class="h-100">
                  <v-card-title class="text-h5">
                    <v-icon left color="success">mdi-server</v-icon>
                    后端
                  </v-card-title>
                  <v-card-text>
                    <v-list>
                      <v-list-item>
                        <v-list-item-title>Go (Golang)</v-list-item-title>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Gin 框架</v-list-item-title>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>GORM</v-list-item-title>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>MySQL</v-list-item-title>
                      </v-list-item>
                    </v-list>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>

        <!-- 功能特性 -->
        <v-card class="mb-6" elevation="4">
          <v-card-title class="text-h4 pa-4">
            <v-icon left color="primary">mdi-star</v-icon>
            功能特性
          </v-card-title>
          <v-card-text>
            <v-list>
              <v-list-item>
                <v-list-item-prepend>
                  <v-icon color="primary">mdi-pencil</v-icon>
                </v-list-item-prepend>
                <v-list-item-title>创建、编辑、删除任务</v-list-item-title>
              </v-list-item>
              <v-list-item>
                <v-list-item-prepend>
                  <v-icon color="success">mdi-check-circle</v-icon>
                </v-list-item-prepend>
                <v-list-item-title>标记任务完成状态</v-list-item-title>
              </v-list-item>
              <v-list-item>
                <v-list-item-prepend>
                  <v-icon color="info">mdi-chart-line</v-icon>
                </v-list-item-prepend>
                <v-list-item-title>任务统计</v-list-item-title>
              </v-list-item>
              <v-list-item>
                <v-list-item-prepend>
                  <v-icon color="purple">mdi-palette</v-icon>
                </v-list-item-prepend>
                <v-list-item-title>现代化的用户界面</v-list-item-title>
              </v-list-item>
              <v-list-item>
                <v-list-item-prepend>
                  <v-icon color="orange">mdi-responsive</v-icon>
                </v-list-item-prepend>
                <v-list-item-title>响应式设计</v-list-item-title>
              </v-list-item>
              <v-list-item>
                <v-list-item-prepend>
                  <v-icon color="teal">mdi-sync</v-icon>
                </v-list-item-prepend>
                <v-list-item-title>实时数据同步</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-card-text>
        </v-card>

        <!-- 版本信息 -->
        <v-card elevation="4" color="grey-lighten-4">
          <v-card-title class="text-h5 pa-4">
            <v-icon left color="primary">mdi-information-outline</v-icon>
            版本信息
          </v-card-title>
          <v-card-text>
            <v-row>
              <v-col cols="12" sm="6">
                <v-chip color="primary" prepend-icon="mdi-tag">
                  版本: 1.0.0
                </v-chip>
              </v-col>
              <v-col cols="12" sm="6">
                <v-chip color="success" prepend-icon="mdi-calendar">
                  构建时间: {{ new Date().toLocaleDateString() }}
                </v-chip>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
// 这里可以添加组件逻辑
</script>

