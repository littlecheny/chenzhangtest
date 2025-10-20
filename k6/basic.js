import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 10,          // 并发虚拟用户数
  duration: '30s',  // 持续时长
};

const BASE = __ENV.BASE_URL || 'http://localhost:8080';

export default function () {
  const randBurst = () => Math.floor(1 + Math.random() * 100); // 1-100 随机
  const payload = JSON.stringify([
    { Index: 1, BurstTime: randBurst() },
    { Index: 2, BurstTime: randBurst() },
  ]);
  const headers = { 'Content-Type': 'application/json', 'X-User-ID': 'user1' };
  const res = http.post(`${BASE}/submitasks`, payload, { headers });

  check(res, { 'status is 200': r => r.status === 200 });
  sleep(0.2); // 每次请求间隔，避免纯压测打爆
}