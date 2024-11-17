import { createRouter, createWebHistory } from "vue-router";
import { useHeaderStore } from "@/stores/headerStore";
import { useLoginStore } from "@/stores/loginStore";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: 'WelcomePage',
      path: '/',
      component: () => import('../pages/WelcomePage.vue'),
      meta: {requiresAuth: false},
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
    {
      name: 'AccountPage',
      path: '/account',
      component: () => import('../pages/AccountPage.vue'),
      meta: {requiresAuth: true},
    },
    {
      name: 'CosmogrammPage',
      path: '/cosmogramm',
      component: () => import('../pages/CosmogrammPage.vue'),
      meta: {requiresAuth: true},
    },
    {
      name: 'TarotPage',
      path: '/tarot',
      component: () => import('../pages/TarotPage.vue'),
      meta: {requiresAuth: true},
    },
  ],
});

router.beforeEach((to, from, next) => {
  const headerStore = useHeaderStore();
  headerStore.currentPage = to.name === undefined ? '' : to.name.toString();

  if(to.name === undefined) next({name: 'WelcomePage'});

  if(to.name === 'SignUpPage2'){
    const loginStore = useLoginStore();
    if(loginStore.SignupEmail === '' || loginStore.SignupPassword === '' || loginStore.SignupPhoneNumber === ''){
      next({name: 'SignUpPage1'});
      return;
    }
  }
  
  next();
});