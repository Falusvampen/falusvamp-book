import { useAuthUser } from '/store/authUser';

export default defineNuxtRouteMiddleware((to, from) => {
  if (to.path === '/login' && useAuthUser().isLoggedin) return navigateTo('/');
  if (to.path === '/login' || useAuthUser().isLoggedin) return;
  return navigateTo('/login');
});
