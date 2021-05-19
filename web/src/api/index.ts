/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-04-27 16:45:02
 * @LastEditros: 
 * @LastEditTime: 2021-05-19 14:41:41
 */

import request from "@/utils/request"

// 获取pre_auth_code
export const create = async(data: any) => {
  return request(`https://login.laughingzhu.cn/api/oauth/create_pre_auth_code`, {
    method: 'POST',
    data
  })
}