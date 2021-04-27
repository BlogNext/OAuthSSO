/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-04-22 14:55:06
 * @LastEditros: 
 * @LastEditTime: 2021-04-27 16:26:02
 */
import { defineConfig } from 'umi';

export default defineConfig({
  title: 'OAuthSSO',
  antd: {},
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    { path: '/', component: '@/pages/index' },
  ],
  publicPath: './',
  base: '/',
  runtimePublicPath: true,

});
