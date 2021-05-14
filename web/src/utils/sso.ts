/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-05-12 14:05:23
 * @LastEditros: 
 * @LastEditTime: 2021-05-14 10:05:36
 */
import qs from 'qs'
import request from "./request";



class OAuthSSO {
  private client_id: string;
  private redirect_url: string;

  constructor(client_id: string, redirect_url: string) {
    this.client_id = client_id;
    this.redirect_url = redirect_url;
  }

  /**
   * @desc 获取pre_auth_code
   * @param {*} params 
   */
  async create (params: any, error?: ((arg0: any) => any) | undefined) {
    const reqData = {
      ...params,
      client_id: this.client_id,
      redirect_url: this.redirect_url
    }

    let res = await this.createRequest (reqData)
    if(res.code === 0) {
      // success
      console.log('获取成功')
      let referrer = document.referrer;
      let prefix = referrer.indexOf('?') > -1 ? '&' : '?'
      location.href = `${referrer}${prefix}pre_auth_code=${res.data.pre_auth_code}`
    } else {
      console.log('失败')

      error && error(res.msg)
    }
  };

  // 博客登录
  async login () {
    if(window) {
      /**
       * 浏览器环境
       */
    
      const searchQuery = qs.parse(location.search, {ignoreQueryPrefix: true })
      if(searchQuery.pre_auth_code) {

        let res = await this.loginRequest(this.redirect_url, {pre_code: searchQuery.pre_auth_code})
        if(res.code === 0) {
            console.log('登录成功', res)
            let referrer = document.referrer.split('pre_auth_code')[0];
            let prefix = referrer.indexOf('?') > -1 ? '&' : '?'
            location.href = `${referrer}${prefix}token=${res.data}`

        } else {
          // login error
          console.log('登录失败')
        }
      }
  
  
      
    } else {
      /**
       * 非浏览器环境，node
       */
      return false;
    }
  }


  /**
   * @description 获取预授权码接口
   * @param data 
   * @returns 
   */
  async createRequest (data: any) {
    return request(`http://154.8.142.48:8085/api/oauth/create_pre_auth_code`, {
      method: 'POST',
      data
    })
  }

/**
   * @description 颁发令牌
   * @param data 
   * @returns 
   */
  async loginRequest (url: string, data: any) {
    return request(`${url}`, {
      method: 'POST',
      data
    })
  }

}



export default OAuthSSO;