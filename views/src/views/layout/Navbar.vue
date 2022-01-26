<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-container>
        <b-navbar-brand @click="$router.replace({name:'Home'})">ShortLink</b-navbar-brand>
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item-dropdown right v-if="userInfo">
              <template v-slot:button-content>
                <strong>{{ userInfo.name }}</strong>
              </template>
              <b-dropdown-item href="#">个人中心</b-dropdown-item>
              <b-dropdown-item @click="logout">登出</b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
          <b-navbar-nav v-if="!userInfo">
            <b-nav-item v-if="$route.name != 'Login'" @click="$router.replace({name:'Login'})">登录</b-nav-item>
            <b-nav-item
              v-if="$route.name != 'Register'"
              @click="$router.replace({name:'Register'})"
            >注册</b-nav-item>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  methods: mapActions('userModule', ['logout']),
};
</script>
