/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-04-27 16:45:02
 * @LastEditros: 
 * @LastEditTime: 2021-04-27 18:17:29
 */

import request from "../utils/request.js";

export const login = async (data: any) => {
  return request(`http://154.8.142.48:8085/api/oauth/create_pre_auth_code`, {
    method: 'POST',
    data
  })
}