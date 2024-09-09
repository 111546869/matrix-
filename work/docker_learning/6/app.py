from flask import Flask
import redis

app = Flask(__name__)

# 连接 Redis 容器，使用默认端口6379
r = redis.Redis(host='redis', port=6379)

@app.route('/')
def index():
    # 每次访问页面时，计数器自增1
    r.incr('counter')
    counter = r.get('counter').decode('utf-8')
    return f'Counter: {counter}'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
