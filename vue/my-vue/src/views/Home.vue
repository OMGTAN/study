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
            <el-card>
                <div ref="echarts" style="height : 280px">

                </div>
            </el-card>
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

        const xData = [
            {苹果: 3995, 小米: 3541, 华为: 2409, oppo: 1110, vivo: 3894, 一加: 1212},
            {苹果: 4995, 小米: 4541, 华为: 3409, oppo: 2110, vivo: 4894, 一加: 2212 },
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
                data: ["aaa"],
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
            color: ["#2ec7e9", "#b6a2de", "#5ab1ef", "#ffb980", "#d87a80", "#8d98b3", ],
            series: [
                {
                data: [10, 22, 28, 43, 49],
                type: 'line',
                stack: 'x'
                },
                {
                data: [5, 4, 3, 5, 10],
                type: 'line',
                stack: 'x'
                }
            ]
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

        const getChatData = ()=>{
            const keys = Object.keys(xData[0]);
            const series =[];
            keys.forEach((key)=>{
                series.push({
                    name: key,
                    data: xData.map((item)=>{}),
                    type: "line",
                })
            });
            orderData.xData = xData;
            orderData.series = series;
            xOptions.xAxis.data = xData;
            xOptions.series = series;
            let hEcharts = echarts.init(proxy.$refs["echarts"])
            hEcharts.setOption(xOptions)
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

    .el-card{
        margin-top: 20px;
    }

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
</style>