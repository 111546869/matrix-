# 使用 Python 官方镜像
FROM python:3.9-slim

# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY . /app

# 安装所需的 Python 依赖
RUN pip install --no-cache-dir flask redis

# 暴露应用的端口
EXPOSE 5000

# 运行 Flask 应用
CMD ["python", "app.py"]
