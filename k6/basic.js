import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 10,          // 并发虚拟用户数
  duration: '30s',  // 持续时长
};

const BASE = __ENV.BASE_URL || 'http://localhost:8080';

export default function () {
  const payload = JSON.stringify([
    { Index: 1, ArrivalTime: 0, BurstTime: 5 },
    { Index: 2, ArrivalTime: 1, BurstTime: 3 },
  ]);
  const headers = { 'Content-Type': 'application/json', 'X-User-ID': 'user1' };
  const res = http.post(`${BASE}/submitasks`, payload, { headers });

  check(res, { 'status is 200': r => r.status === 200 });
  sleep(0.2); // 每次请求间隔，避免纯压测打爆
}