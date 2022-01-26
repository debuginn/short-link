<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3" sm="6" offset-sm="3" xl="6" offset-xl="3">
        <b-card title="登录">
          <b-form>
            <b-form-group id="input-group-2" label="手机号码:" label-for="input-2">
              <b-form-input
                id="input-2"
                v-model="$v.user.telephone.$model"
                type="number"
                required
                placeholder="请输入你的电话号码（必须）"
                :state="validateState('telephone')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">手机号不符合要求</b-form-invalid-feedback>
            </b-form-group>
            <b-form-group id="input-group-3" label="用户密码:" label-for="input-3">
              <b-form-input
                id="input-3"
                v-model="$v.user.password.$model"
                type="password"
                required
                placeholder="请输入你的用户密码（必须）"
                :state="validateState('password')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">密码必须大于或等于8位</b-form-invalid-feedback>
            </b-form-group>
            <b-button block variant="info" @click="login">登录</b-button>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
// 引入 vuelidate 插件
import { required, minLength } from 'vuelidate/lib/validators';
import customValidator from '@/helper/validator';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        telephone: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        validator: customValidator.telephoneValidator,
      },
      password: {
        minLength: minLength(8),
      },
    },
  },
  methods: {
    ...mapActions('userModule', { userLogin: 'login' }),
    validateState(name) {
      // ES6 赋值结构
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      // 验证数据
      this.$v.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      this.userLogin(this.user).then(() => {
        // 跳转主页
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        if (typeof (err.response) !== 'undefined') {
          this.$bvToast.toast(err.response.data.msg, {
            title: '数据验证错误',
            variant: 'danger',
            solid: true,
          });
        }
      });
    },
  },
};
</script>
