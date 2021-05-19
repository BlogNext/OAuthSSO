/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-04-22 14:55:06
 * @LastEditros: 
 * @LastEditTime: 2021-05-19 13:54:42
 */
import React, { useState, useEffect } from 'react';

import { Form, Input, Button, message } from 'antd'
import Success from '../component/success'
import Error from '../component/error'
import { history } from 'umi'

import './style.less'
export default () => {
  const [status, setStatus] = useState(false)

  useEffect(() => {
    _checkParam()
  }, [''])


  const _checkParam = () => {
    const { client_id, redirect_url } = history.location.query
    if(client_id && redirect_url) {
      // 都存在
    } else {
    }
  }


  
  return (
    <div className="login flex">
      {status ? <Success /> : <Error />}
    </div>
  );
}


