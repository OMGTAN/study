<template>
    <el-header>
        <div class="l-content">
            <el-button size="small" @click="handleIsCollapse">
                <el-icon :size="20">
                    <Menu />
                </el-icon>
            </el-button>
            <h3>首页</h3>
        </div>
        <div class="r-content">
            <el-dropdown>
                <span class="el-dropdown-link">
                    <img class="user" src="../assets/user.png" />
                </span>
                <template #dropdown>
                <el-dropdown-menu>
                    <el-dropdown-item>个人中心</el-dropdown-item>
                    <el-dropdown-item>退出</el-dropdown-item>
                </el-dropdown-menu>
                </template>
            </el-dropdown>
        </div>
        
    </el-header>
</template>

<script>
import { defineComponent } from 'vue'
import {useStore} from 'vuex'

export default defineComponent({
    setup() {
        let store = useStore();
        let handleIsCollapse = ()=>{
            store.commit("updateIsCollapse")
        }

        return {
            handleIsCollapse
        }
    },
})
</script>


<style lang="less" scoped>
    header{
        display: flex;
        justify-content: space-between;
        align-content: center;
        width: 100%;
        background-color: #333;
    }
    .r-content{
        .user{
            width: 40px;
            height: 40px;
            border-radius: 50%;
        }
    }
    .l-content{
        display: flex;
        align-items: center;
        .el-button{
            margin-right: 20px;
        }
        h3{
            color: #fff;
        }

    }

</style>