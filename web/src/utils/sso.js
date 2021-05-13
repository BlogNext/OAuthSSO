/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-05-12 14:05:23
 * @LastEditros: 
 * @LastEditTime: 2021-05-13 18:42:27
 */
import { history } from 'umi'
import qs from 'qs'
import { loginCode, createCode } from '../api/index'

class OAuthSSO {

  constructor(client_id, redirect_url) {
    this.client_id = client_id;
    this.redirect_url = redirect_url;
  }

  async createCode (nickname, passowrd) {
    console.log(nickname, passowrd, 'createCode')
  };

  async ready () {
    if(window) {
      console.log(this)
      console.log('---------浏览器---------')
    
      // 
      const searchQuery = qs.parse(location.search, {ignoreQueryPrefix: true })
      
      if(searchQuery.pre_auth_code) {
        let res = await loginCode({pre_code: searchQuery.pre_auth_code})
  
        if(res.code === 0) {
          // login success
          // location.hjref = `${}`
        } else {
          // login error
        }
      }
  
  
      
    } else {
      
      console.log('---------非浏览器---------')
      return false;
    }
  }



  
  
  
}



export default OAuthSSO;