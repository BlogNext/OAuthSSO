/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-05-12 14:05:23
 * @LastEditros: 
 * @LastEditTime: 2021-05-13 23:02:00
 */
import { history } from 'umi'
import qs from 'qs'
import { loginCode, createCode } from '../api/index'
import { message } from 'antd'


class OAuthSSO {

  constructor(client_id, redirect_url) {
    this.client_id = client_id;
    this.redirect_url = redirect_url;
  }

  // 获取pre_auth_code
  async createCode (params) {
    console.log( 'createCode')
    const reqData = {
      ...params,
      client_id: this.client_id,
      redirect_url: this.redirect_url
    }

    let res = await createCode (reqData)
    if(res.code === 0) {
      // success
      message.success(res.msg, 2, () => {
        console.log('获取成功', res)
        let referrer = document.referrer;
        let prefix = referrer.indexOf('?') > -1 ? '&' : '?'

        location.href = `${referrer}${prefix}pre_auth_code=${res.data.pre_auth_code}`
      })
    } else {
      message.error(res.msg, 2)
    }
  };

  // 博客登录
  async ready () {
    if(window) {
      console.log(this)
      console.log('---------浏览器---------')
    
      // 
      const searchQuery = qs.parse(location.search, {ignoreQueryPrefix: true })
      
      if(searchQuery.pre_auth_code) {
        let res = await loginCode({pre_code: searchQuery.pre_auth_code})
  
        if(res.code === 0) {
          console.log('登录成功')
          let referrer = document.referrer.split('pre_auth_code')[0];
          let prefix = referrer.indexOf('?') > -1 ? '&' : '?'
          location.href = `${referrer}${prefix}token=${res.data}`
        } else {
          // login error
          console.log('失败')
        }
      }
  
  
      
    } else {
      
      console.log('---------非浏览器---------')
      return false;
    }
  }



  
  
  
}



export default OAuthSSO;