/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-04-22 14:55:06
 * @LastEditros: 
 * @LastEditTime: 2021-05-14 17:54:11
 */
import { defineConfig } from 'umi';

export default defineConfig({
  title: 'OAuthSSO',
  antd: {},
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    { path: '/error', component: '@/pages/error/index'},
    { path: '/', component: '@/pages/index' },
  ],
  publicPath: './',
  base: '/',
  runtimePublicPath: true,

});
