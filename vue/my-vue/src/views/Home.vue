<template>
    <el-row class="home" :gutter="20">
        <el-col :span="8" >
            <el-card shadow = "hover">
                <div class="user">
                    <img src="../assets/user.jpg" alt=""/>
                    <div style="margin-left: 0px;">
                        <p class="name">Admin</p>
                        <p class="role">超级管理员</p>
                    </div>
                </div>
                <div class="loginInfo">
                    <p>上次登录时间：<span>2024-01-01</span></p>
                    <p>上次登录的地点：<span>北京</span></p>
                </div>
            </el-card>
            <el-card shadow = "hover" style="margin-top: 20px;">
                <el-table :data="tableData">
                    <el-table-column v-for="(key, val) in tableLable " :key="val" :prop = "val" :label="key"></el-table-column>

                </el-table>
            </el-card>
        </el-col>

        <el-col :span = "16" >
            <div class="num">
                <el-card :body-style="{display:'flex', padding: 0}" v-for="item in countData" :key="item.name">
                    <component class="icons" :is="item.icon" :style="{background:item.color}"></component>
                    <div class="details">
                        <p class="num">￥{{ item.value }}</p>
                        <p class="txt">{{ item.name }}</p>
                    </div>
                </el-card>
                
            </div>
            <el-card style="margin-top: 20px;">
                <div ref="echarts" style="height : 280px">
                </div>
                
            </el-card>
            <div class="graph">
                    <el-card style="height: 260px">
                        <div ref="userecharts" style="height: 240px"></div>
                    </el-card>
                    <el-card style="height: 260px">
                        <div ref="videoecharts" style="height: 240px"></div>
                    </el-card>
                </div>
        </el-col>
        
    </el-row>
</template>

<script>
import { defineComponent, getCurrentInstance, onMounted, reactive, ref } from 'vue'
import axios from 'axios'
import * as echarts from 'echarts'

export default defineComponent({
    setup() {
        let tableData = ref([]);
        let countData = ref([]);
        const {proxy}  = getCurrentInstance();
        const tableLable =  {
            name : "姓名",
            todayBuy : "今日购买",
            monthBuy : "本月购买",
            totalBuy : "总购买"
        }

        let getTableData = async ()=>{
            // axios.get('/home/getTableData').then((res)=>{
            //     // console.log(res);
            //     tableData.value = res.data.data.tableData
            // }
            // )

            let res = await proxy.$api.getTableData();
            tableData.value = res.tableData
        };

        let getCountData = async ()=>{
            let res = await proxy.$api.getCountData();
            countData.value = res.countData
        };

        const xData = {
            date:['20241001', '20241002', '20241003', '20241004', '20241005', '20241006', '20241007', ],
            data:
            [
                {苹果: 3995, 小米: 3541, 华为: 2409, oppo: 1110, vivo: 3894, 一加: 1212},
                {苹果: 2995, 小米: 4541, 华为: 3409, oppo: 2110, vivo: 4894, 一加: 2212 },
                {苹果: 1995, 小米: 4541, 华为: 3409, oppo: 2110, vivo: 4894, 一加: 2212 },
                {苹果: 4295, 小米: 4541, 华为: 3409, oppo: 2110, vivo: 4894, 一加: 2212 },
                {苹果: 4395, 小米: 4541, 华为: 3409, oppo: 2110, vivo: 4894, 一加: 2212 },
                {苹果: 1995, 小米: 4541, 华为: 3409, oppo: 2110, vivo: 4894, 一加: 2212 },
                {苹果: 4095, 小米: 4541, 华为: 3409, oppo: 2110, vivo: 4894, 一加: 2212 },
            ]
        }
        const uData = [
            {
                date:"周一",
                new: 5,
                active: 200
            },
            {
                date:"周二",
                new: 10,
                active: 500
            },
            {
                date:"周三",
                new: 12,
                active: 550
            },
            {
                date:"周四",
                new: 60,
                active: 800
            },
            {
                date:"周五",
                new: 65,
                active: 550
            },
        ]
        const vData = [
            {
                name: "小米",
                value: 2999
            },
            {
                name: "苹果",
                value: 5999
            },
            {
                name: "vivo",
                value: 2999
            },
            {
                name: "oppo",
                value: 2999
            },
            {
                name: "魅族",
                value: 2999
            },
            {
                name: "三星",
                value: 2999
            },
        ] 

        let xOptions = reactive({
            textStyle:{
                color: "#333"
            },
            grid:{
                left: "20%"
            },
            tooltip:{
                trigger: "axis"
            },
            xAxis: {
                type:"category",
                data: [],
                axisLine:{
                    lineStyle:{
                        color: "#17b3a3"
                    }
                },
                axisLabel:{
                    interval: 0,
                    color: "#333",
                },
            },
            yAxis: [
                {
                    type: "value",
                    axisLine:{
                        lineStyle:{
                            color: "#17b3a3"
                        }
                    },
                }
            ],
            color: ["#2ec7e9", "#b6a2de", "#5ab1ef", "#ffb980", "#d87a80", "#8d98b3"],
            series: []
        })
        let pieOptions = reactive({
            tooltip:{
                trigger:"item",
            },
            color:[
                "#0f78f4",
                "#dd536b",
                "#9462e5",
                "#a6a6a6",
                "#e1bb22",
                "#39c362",
                "#3ed1cf",
            ]

        })
        let orderData = reactive({
            xData:[],
            series:[]
        })
        let userData = reactive({
            xData:[],
            series:[]
        })
        let videoData = reactive({
            series:[]
        })

        const getChatData = ()=>{
            const keys = Object.keys(xData.data[0]);
            const series =[];
            keys.forEach((key)=>{
                series.push({
                    name: key,
                    data: xData.data.map((item)=>item[key]),
                    type: "line",
                })
            });
            orderData.xData = xData.date;
            orderData.series = series;
            xOptions.xAxis.data = orderData.xData;
            xOptions.series = orderData.series;
            let hEcharts = echarts.init(proxy.$refs["echarts"])
            hEcharts.setOption(xOptions)

            //柱状图
            userData.xData = uData.map((item)=> item.date)
            userData.series = [
                {
                    name: "新增用户",
                    data: uData.map((item)=> item.new),
                    type: "bar",
                },
                {
                    name: "活跃用户",
                    data: uData.map((item)=> item.active),
                    type: "bar",
                },
            ];
            xOptions.xAxis.data = userData.xData;
            xOptions.series = userData.series;
            let uEcharts = echarts.init(proxy.$refs["userecharts"])
            uEcharts.setOption(xOptions)

            //饼状图
            videoData.series=[
                {
                    data: vData,
                    type: "pie"
                }
            ]
            pieOptions.series = videoData.series
            let vEcharts = echarts.init(proxy.$refs["videoecharts"])
            vEcharts.setOption(pieOptions)
        }

        onMounted(()=>{
            getTableData();
            getCountData();
            getChatData();
        })
        return{
            tableLable,
            tableData,
            countData,
        }
    },
})
</script>


<style lang="less" scoped>

    .home{
        .user{
            display: flex;
            align-items: center;
            padding-bottom: 20px;
            border-bottom: 1px solid #ccc;
            margin-bottom: 20px;
            margin-top: 0px;
            img{
                height: 150px;
                width: 150px;
                border-radius: 50%;
                margin-right: 40px;
            }
        }
        .loginInfo{
            line-height: 30px;
            font-size: 14;
            color: #999;
            
            span{
                color: #666;
                margin-left: 70px;
            }
        }
    }

    // .el-card{
    //     // margin-top: 20px;
    // }

    .num{
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
        .el-card{
            width: 32%;
            margin-bottom: 10px;
        }
        .icons{
            width: 80px;
            height: 80px;
            font-size: 30px;
            text-align: center;
            line-height: 80px;
            color:#fff;
        }
        .details{
            margin-left: 15px;
            display: flex;
            flex-direction: column;
            justify-content: center;
            .num{
                font-size: 30px;
                margin-bottom: 10px;
            }
            .txt{
                font-size: 14px;
                text-align: center;
                color: #999;
            }
        }
    }
    .graph{
        margin-top: 20px;
        display: flex;
        justify-content: space-between;
        .el-card{
            width: 48%;
        }
    }
</style>