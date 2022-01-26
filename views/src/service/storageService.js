// 本地缓存服务
const PREFIX = 'ginessential_';

// user 模块
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// 存储信息
const set = (key, data) => {
  localStorage.setItem(key, data);
};

// 读取信息
const get = (key) => {
  const value = localStorage.getItem(key);
  return value;
};

export default {
  set,
  get,
  USER_TOKEN,
  USER_INFO,
};
