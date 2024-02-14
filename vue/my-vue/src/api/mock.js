import Mock from 'mockjs'
import homeApi from './mockData/home'
import userApi from './mockData/user'


Mock.mock('/home/getTableData', homeApi.getHomeData)
Mock.mock('/home/getCountData', homeApi.getCountData)
Mock.mock("/user/getUserList", 'get', userApi.getUserList)