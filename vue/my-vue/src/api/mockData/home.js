export default {

    getHomeData: ()=>{
        return {
            code: 200,
            data: {
                tableData: [
                    {
                        name : "name11",
                        todayBuy : "todayB",
                        monthBuy : "monthB",
                        totalBuy : "totalB"
                    }
                ]
            }
        }
    },
    

    getCountData:()=>{
        return {
            code: 200,
            data: {
                countData: [
                    {
                        name : "今日支付订单",
                        value : 1234,
                        icon : "CircleCheck",
                        color : "#2ec7c9"
                    },
                    {
                        name : "今日收藏订单",
                        value : 210,
                        icon : "Star",
                        color : "#ffb980"
                    },
                    {
                        name : "今日未支付订单",
                        value : 1234,
                        icon : "Goods",
                        color : "#5ab1ef"
                    },
                    {
                        name : "今日支付订单",
                        value : 1234,
                        icon : "CircleCheck",
                        color : "#2ec7c9"
                    },
                    {
                        name : "今日收藏订单",
                        value : 210,
                        icon : "Star",
                        color : "#ffb980"
                    },
                    {
                        name : "今日未支付订单",
                        value : 1234,
                        icon : "Goods",
                        color : "#5ab1ef"
                    },
                ]
            }
        }
    }
}