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
            url:'/user/getUser',
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

    editUser(params){
        return request({
            url:'/user/edit',
            method: 'post',
            data: params,
            mock: true
        })
    },

    deleteUser(params){
        return request({
            url:'/user/delete',
            method: 'get',
            data: params,
            mock: true
        })
    },

}