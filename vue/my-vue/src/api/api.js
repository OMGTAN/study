import request  from "./request";

export default {
    getTableData(params){
        return request({
            url:'/home/getTableData',
            method: 'get',
            data: params,
            mock: true
        })
    }
}