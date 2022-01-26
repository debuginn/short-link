import storageService from '@/service/storageService';
import userService from '@/service/userService';

const userModule = {
  namespaced: true,

  state: {
    token: storageService.get(storageService.USER_TOKEN),
    userInfo: storageService.get(storageService.USER_INFO) ? JSON.parse(storageService.get(storageService.USER_INFO)) : null,
  },

  mutations: {
    SET_TOKEN(state, token) {
      // 更新本地缓存
      storageService.set(storageService.USER_TOKEN, token);
      // 更新 state
      state.token = token;
    },
    SET_USERINFO(state, userInfo) {
      // 更新本地缓存
      storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
      // 更新 state
      state.userInfo = userInfo;
    },
  },

  actions: {
    // 用户注册
    register(context, { name, telephone, password }) {
      return new Promise((resolve, reject) => {
        userService.register({ name, telephone, password }).then((res) => {
          // 保存 token
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.userInfo();
        }).then((response) => {
          // 保存用户信息
          context.commit('SET_USERINFO', response.data.data.user);
          resolve(response);
        }).catch((err) => {
          reject(err);
        });
      });
    },

    // 用户登录
    login(context, { telephone, password }) {
      return new Promise((resolve, reject) => {
        userService.login({ telephone, password }).then((res) => {
          // 保存 token
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.userInfo();
        }).then((response) => {
          // 保存用户信息
          context.commit('SET_USERINFO', response.data.data.user);
          resolve(response);
        }).catch((err) => {
          reject(err);
        });
      });
    },

    // 用户登出
    logout({ commit }) {
      // 清除 Token
      commit('SET_TOKEN', '');
      storageService.set(storageService.USER_TOKEN, '');
      // 清除用户信息
      commit('SET_USERINFO', '');
      storageService.set(storageService.USER_INFO, '');
    },

  },

};

export default userModule;
