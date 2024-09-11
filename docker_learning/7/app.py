from flask import Flask
import redis
import os
from datetime import datetime

app = Flask(__name__)

# 连接 Redis 容器
r = redis.Redis(host='redis', port=6379)

# 设置文件保存路径
file_save_path = '/app/timestamp_files'

# 确保目录存在
if not os.path.exists(file_save_path):
    os.makedirs(file_save_path)

@app.route('/')
def index():
    # 每次访问时计数器自增
    r.incr('counter')
    counter = r.get('counter').decode('utf-8')

    # 获取当前时间戳并创建文件
    timestamp = datetime.now().strftime('%Y-%m-%d_%H-%M-%S')
    file_name = f"{file_save_path}/{timestamp}.txt"
    
    # 创建文件并写入内容
    with open(file_name, 'w') as f:
        f.write(f"Counter: {counter}, Timestamp: {timestamp}\n")
    
    return f'Counter: {counter}, Timestamp: {timestamp}'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
