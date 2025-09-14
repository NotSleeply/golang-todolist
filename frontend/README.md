# Vue 3 TodoList Frontend

## 项目结构
```
frontend/
├── index.html          # 入口 HTML 文件
├── package.json        # 项目依赖配置
├── vite.config.js      # Vite 构建配置
├── style.css          # 全局样式
└── src/
    ├── main.js        # Vue 应用入口
    ├── App.vue        # 主组件
    └── api/
        └── todoApi.js # API 请求封装
```

## 功能特性
- ✅ Vue 3 Composition API
- ✅ 现代化 UI 设计
- ✅ 响应式布局
- ✅ API 请求封装
- ✅ 错误处理
- ✅ 加载状态
- ✅ 表单验证

## 安装和运行

1. 安装依赖：
```bash
cd frontend
npm install
```

2. 启动开发服务器：
```bash
npm run dev
```

3. 构建生产版本：
```bash
npm run build
```

## API 接口
- GET `/api/v1/todos` - 获取所有任务
- POST `/api/v1/todos` - 创建任务
- PUT `/api/v1/todos` - 更新任务
- DELETE `/api/v1/todos` - 删除任务
