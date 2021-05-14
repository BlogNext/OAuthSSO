/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-04-27 16:45:02
 * @LastEditros: 
 * @LastEditTime: 2021-05-14 17:20:08
 */

import request from "@/utils/request"

// 获取pre_auth_code
export const create = async(data: any) => {
  return request(`https://log.laughingzhu.cn/api/oauth/create_pre_auth_code`, {
    method: 'POST',
    data
  })
}