import Mock from 'mockjs'
import homeApi from './mockData/home'

Mock.mock('/home/getTableData', homeApi.getHomeData)