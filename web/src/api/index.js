/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-04-27 16:45:02
 * @LastEditros: 
 * @LastEditTime: 2021-05-13 17:59:33
 */

import request from "../utils/request.js";

export const createCode = async (data) => {
  return request(`http://154.8.142.48:8085/api/oauth/create_pre_auth_code`, {
    method: 'POST',
    data
  })
}



export const loginCode = async (data) => {
  return request(`https://blog.laughingzhu.cn/front/login/login_blog_next_pre_code`, {
    method: 'POST',
    data
  })
}
