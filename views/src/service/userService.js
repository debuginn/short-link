import request from '@/utils/request';

// 用户注册
const register = ({ name, telephone, password }) => {
  return request.post('auth/register', { name, telephone, password });
};

// 获取用户信息
const userInfo = () => {
  return request.get('auth/info');
};

// 用户登录
const login = ({ telephone, password }) => {
  return request.post('auth/login', { telephone, password });
};


export default {
  register,
  userInfo,
  login,
};
