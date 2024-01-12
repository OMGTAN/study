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
        
    </el-row>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
    setup() {
        const tableData = [
            {
                name : "name1",
                todayBuy : "todayB",
                monthBuy : "monthB",
                totalBuy : "totalB"
            }
        ]
        const tableLable =  {
            name : "姓名",
            todayBuy : "今日购买",
            monthBuy : "本月购买",
            totalBuy : "总购买"
        }

        return{
            tableLable,
            tableData
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

</style>