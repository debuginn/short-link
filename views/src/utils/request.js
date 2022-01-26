import axios from 'axios';
import storageService from '../service/storageService';

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 1000 * 5,
});

service.interceptors.request.use((config) => {
  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` });
  return config;
}, (error) => {
  return Promise.reject(error);
});

export default service;
