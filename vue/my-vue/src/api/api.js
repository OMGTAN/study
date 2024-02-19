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
            url:'/home/getUser',
            method: 'get',
            data: params,
            mock: true
        })
    },

    addUser(params){
        return request({
            url:'/user/add',
            method: 'post',
            data: params,
            mock: true
        })
    },

}