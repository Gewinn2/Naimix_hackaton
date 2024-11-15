import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: 'WelcomePage',
      path: '/',
      redirect: {name: "LoginPage"},
    },
    {
      name: 'LoginPage',
      path: '/login',
      component: () => import('../pages/LoginPage.vue'),
      meta: {requiresAuth: false},
    },
    {
      name: 'SignUpPage',
      path: '/signup',
      redirect: {name: 'SignUpPage1'},
      meta: {requiresAuth: false},
      children: [
        {
          name: 'SignUpPage1',
          path: 'step-1',
          component: () => import('../pages/SignUpPage1.vue'),
          meta: {requiresAuth: false},
        },
        {
          name: 'SignUpPage2',
          path: 'step-2',
          component: () => import('../pages/SignUpPage2.vue'),
          meta: {requiresAuth: false},
        },
      ],
    },
    {
      name: 'MainPage',
      path: '/main',
      component: () => import('../pages/MainPage.vue'),
      meta: {requiresAuth: true},
    },
  ],
});