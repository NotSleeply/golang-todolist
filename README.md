# Golang TodoList

一个基于 Go 后端和静态前端的待办事项管理应用。适合学习全栈开发、Go Web 服务和前后端分离架构。

## 项目结构

```
golang-todolist/
├── backend/      # Go 后端服务
│   ├── main.go   # 后端入口
│   ├── db/       # 数据模型
│   └── static/   # 后端静态资源（可选）
├── frontend/     # 前端静态页面
│   ├── index.html
│   ├── script.js
│   └── style.css
└── README.md     # 项目说明
```

## 快速开始

### 后端启动

1. 进入 `backend` 目录：
   ```bash
   cd backend
   ```
2. 安装依赖并运行：
   ```bash
   go mod tidy
   go run main.go
   ```

### 前端预览

直接用浏览器打开 `frontend/index.html` 即可访问前端页面。

## 功能简介

- 添加、删除、修改待办事项
- 前后端分离，接口通信
- 简洁易用的界面

## 贡献方式

欢迎 PR 和 Issue！请遵循以下流程：

1. Fork 本仓库
2. 新建分支并提交更改
3. 发起 Pull Request

## License

AGPL-3.0
