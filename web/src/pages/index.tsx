/*
 * @Description: 
 * @Author: LaughingZhu
 * @Date: 2021-04-22 14:55:06
 * @LastEditros: 
 * @LastEditTime: 2021-04-28 15:31:31
 */
import React, { useState } from 'react';
import { history } from 'umi'
import { Form, Input, Button, message } from 'antd'
import { CloseOutlined } from '@ant-design/icons';
import { login } from '../api'
import './style.less'
export default () => {
  const [type, setType] = useState(0)
  console.log(history)


  const onFinish = async(values: any) => {
    console.log('Success:', values);
    let res = await login ({...values, ...history.location.query})
    if(res.code === 0) {
      // success
      console.log('登录成功', res)
    } else {
      message.error(res.msg, 2)
    }
  };

  const closeLogin = () => {
  }

  const onBlur = (type: number) => {
    console.log(111)
    setType(type)
  }

  const onFocus = (type: number) => {
    setType(type)
  }
  return (
    <div className="login flex">
      <div className="login-wrapper flex">
        <img src={require('../assets/login/normal.png')} alt="" className={`login-wrapper-carton login-wrapper-carton--normal ${type !== 0 && 'hidden'}`}/>
        <img src={require('../assets/login/greeting.png')} alt="" className={`login-wrapper-carton login-wrapper-carton--greeting ${type !== 1 && 'hidden'}`}/>
        <img  src={require('../assets/login/blindfold.png')} alt="" className={`login-wrapper-carton login-wrapper-carton--blindfold ${type !== 2 && 'hidden'}`}/>
        <div className="login-wrapper-top flex">
          <div className="login-wrapper-top--title">账密登录</div>
          <CloseOutlined onClick={closeLogin} className='login-wrapper-top--close' />
        </div>

        <Form
          name="basic"
          style={{ width: '100%'}}
          initialValues={{ remember: true }}
          onFinish={onFinish}
          // onFinishFailed={onFinishFailed}
        >
          <Form.Item
            className="login-wrapper--item flex"
            name="nickname"
            
            rules={[{ required: true, message: 'Please input your username!' }]}
          >
            <Input onFocus={() => onFocus(1)} onBlur={() => onBlur(0)} placeholder='username' />
          </Form.Item>

          <Form.Item
            className="login-wrapper--item flex"
            name="password"
            rules={[{ required: true, message: 'Please input your password!' }]}
          >
            <Input.Password onFocus={() => onFocus(2)} onBlur={() => onBlur(0)} placeholder='password' />
          </Form.Item>

          <Form.Item className="login-wrapper--item flex">
            <Button type='primary' htmlType="submit" block>
              Submit
            </Button>
          </Form.Item>


        </Form>


        <div className="login-wrapper-footer flex">
          注册登录即表示同意 <span> 用户协议</span> 、 <span>隐私政策</span>
        </div>
      </div>
    </div>
  );
}


