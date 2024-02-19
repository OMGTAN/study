import axios from 'axios'
import config from '../config' 
import { ElMessage, ElSubMenu } from 'element-plus'
const NETWORK_ERR = '网络异常...，请重试'

const service = axios.create({
    baseURL: config.baseApi
})

service.interceptors.request.use((req)=>{
    return req;
})

service.interceptors.response.use((res)=>{
    const {code, data, msg} = res.data
    if(code == 200){
        return data
    }else{
        ElMessage.error(msg|| NETWORK_ERR)
        return Promise.reject(msg|| NETWORK_ERR)
    }
})

function request(options){
    options.method = options.method||'get'
    if(options.method.toLowerCase() == 'get'){
        options.params = options.data
    }
    let isMock = options.mock || config.mock

    if(config.env = 'prod'){
        service.defaults.baseURL = config.baseApi
    }else{
        service.defaults.baseURL = isMock ? config.mockApi : config.baseApi
    }

    return service(options)
}

export default request