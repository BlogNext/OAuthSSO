/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-05-14 17:43:41
 * @LastEditros: 
 * @LastEditTime: 2021-05-19 14:43:22
 */
// 错误页面
import React, { useEffect, useState } from 'react'
import { history } from 'umi';
import { Result, Button } from 'antd';
import { FrownOutlined } from '@ant-design/icons';



export default () => {
  const [msg, setMsg] = useState('')
  useEffect(() => {
    _checkParam()
    return () => {}
  }, [])

  const _checkParam = () => {
    const { client_id, redirect_url } = history.location.query
    if(!client_id) {
      // client_id 不存在
      console.log(2222)
      setMsg('client_id参数错误')
      return false
    }
    if(!redirect_url) {
      setMsg('redirect_url参数错误')
      return false
    }

    return false
  }

  return (
    <Result
      style={{width: '100vw', height: '100vh', color: '#fff'}}
      icon={<FrownOutlined />}
      title={msg}
      extra={''}
    />
  )
}