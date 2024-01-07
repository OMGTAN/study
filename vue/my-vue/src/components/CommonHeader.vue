<template>
    <el-header>
        <div class="l-container">
            <el-button size="small">
            <el-icon :size="20">
                <Menu />
            </el-icon>
        </el-button>
        </div>
        <div class="r-container">
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

<style lang="less" scoped>
    header{
        display: flex;
        justify-content: space-between;
        align-content: center;
        width: 100%;
    }
    .r-container{
        .user{
            width: 40px;
            height: 40px;
            border-radius: 50%;
        }
    }

</style>