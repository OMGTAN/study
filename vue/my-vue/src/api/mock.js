import Mock from 'mockjs'
import homeApi from './mockData/home'
import userApi from './mockData/user'
import permissionApi from './mockData/permission'


Mock.mock(/home\/getTableData/, homeApi.getHomeData)
Mock.mock(/home\/getCountData/, homeApi.getCountData)



Mock.mock(/user\/getUser/, 'get', userApi.getUserList)
Mock.mock(/user\/add/, 'post', userApi.createUser)
Mock.mock(/user\/edit/, 'post', userApi.updateUser)
Mock.mock(/user\/delete/, 'get', userApi.deleteUser)


Mock.mock(/permission\/getMenu/, 'post', permissionApi.getMenu)