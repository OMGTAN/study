import request  from "./request";

export default {

    getTableData(params){
        return request({
            url:'/home/getTableData',
            method: 'get',
            data: params,
            mock: true
        })
    },

    getCountData(params){
        return request({
            url:'/home/getCountData',
            method: 'get',
            data: params,
            mock: true
        })
    },

    getUserData(params){
        return request({
            url:'/user/getUserData',
            method: 'get',
            data: params,
            mock: true
        })
    },
}